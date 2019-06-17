package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
	"myiris/cache/session"
	"myiris/library/logger"
	"myiris/library/redis"
)

// 初始化配置文件
func loadConfig(absPath string, cfg interface{}) error {
	f, err := ioutil.ReadFile(filepath.Join(absPath+"/config", "conf.d/storage.yaml"))
	if err != nil {
		return fmt.Errorf("reading config file path error: %s", err.Error())
	}

	if err = yaml.Unmarshal(f, cfg); err != nil {
		return fmt.Errorf("解析.yaml 文件出错：%s", err.Error())
	}
	return nil
}

// 初始化logger
func initLogger(storage *session.Storage) {
	logger.Init(storage.LogName, storage.LogMod, storage.LogDir, storage.LogLevel)
}

// 初始化redis
func initRedis(storage *session.Storage) {
	session.Redis = redis.NewRedisCache(storage.RedisDb, storage.RedisUrl, storage.RedisPass, 2000)
}

// 初始化mongo
func initMongo(storage *session.Storage) {
	var err error
	session.Mongo, err = mongo.NewClient(options.Client().ApplyURI(storage.MongoProc))
	if err != nil {
		logger.Error("mongo newClient: ", err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err = session.Mongo.Connect(ctx); err != nil {
		logger.Error("mongo client connect: ", err)
		return
	}
	if err = session.Mongo.Ping(ctx, readpref.Primary()); err != nil {
		logger.Error("mongo client ping: ", err)
		return
	}
	logger.Debug("mongo client init success...")
	//c := session.Mongo.Database(storage.MongoDB).Collection("")
}

// 初始化mysql
func initMysql(storage *session.Storage) {
	engine, err := xorm.NewEngine("mysql", storage.MysqlProc)
	if err != nil {
		logger.Error("mysql dial: ", err)
		return
	}

	if err = engine.Ping(); err != nil {
		logger.Error("mysql ping: ", err)
		return
	}
	//engine.SetTableMapper(core.SnakeMapper{})    // struct 与table、fields 之间名称映射

	// engine内部支持连接池
	//engine.SetMaxOpenConns()
	//engine.SetMaxIdleConns()
	//engine.SetConnMaxLifetime()

	// LRU 缓冲
	cache := xorm.NewLRUCacher(xorm.NewMemoryStore(), 100)
	engine.SetDefaultCacher(cache)

	// 打印mysql 日志
	engine.ShowSQL(true)

	session.Mysql = engine
	_, _ = session.Mysql.Clone()
	logger.Debug("mysql session init success...")
}

func Initialize() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	adsPath := strings.Replace(dir, "\\", "/", -1)

	session.Conf = &session.Storage{}
	if err = loadConfig(adsPath, session.Conf); err != nil {
		panic(err)
	}

	initLogger(session.Conf)
	initRedis(session.Conf)
	initMongo(session.Conf)
	initMysql(session.Conf)
}
