package session

import (
	"github.com/go-xorm/xorm"
	"go.mongodb.org/mongo-driver/mongo"
	"myiris/library/redis"
)

var (
	Conf  *Storage
	Redis *redis.Cache
	Mongo *mongo.Client
	Mysql *xorm.Engine
)

type Storage struct {
	LocalAddr     string
	RedisUrl      string `yaml:"redisUrl"`
	RedisPoolSize int    `yaml:"redisPoolSize"`
	RedisDb       int    `yaml:"redisDB"`
	RedisPass     string `yaml:"redisPassword"`
	MongoDB       string `yaml:"mongoDB"`       // mongo db
	MongoProc     string `yaml:"mongoHostProc"` // mongo host
	MysqlProc     string `yaml:"mysqlHostProc"` // mysql source
	QueueName     string `yaml:"queueName"`
	QueueType     string `yaml:"queueType"`
	LogName       string `yaml:"logName"`  // 输出文件名
	LogDir        string `yaml:"logDir"`   // 输出路径
	LogLevel      string `yaml:"logLevel"` // 输出层级
	LogMod        string `yaml:"logMod"`   // 输出方式
}
