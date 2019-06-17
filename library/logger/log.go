package logger

import (
	"os"
	"strings"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//初始值
const (
	maxAge       = 7 * 24 * time.Hour
	rotationTime = time.Hour * 2
)

var logger *logrus.Logger
var _log *logrus.Entry

func init() {

	SetConfig(SWITCH) //打开错误日志堆栈

	logger = logrus.New()
	_log = logrus.NewEntry(logger)

}

//Init 初始化
func Init(serverName, model, path string, level string) {

	var lt logrus.Level
	//# 0 Panic、1 fatal、2 error、3 warn 、4 info 、5 debug
	switch strings.ToLower(level) {
	case "debug":
		lt = 5
	case "info":
		lt = 4
	case "warn":
		lt = 3
	case "error":
		lt = 2
	case "fatal":
		lt = 1
	case "panic":
		lt = 0
	default:
		lt = 4
	}

	_log = logger.WithFields(logrus.Fields{
		"server": serverName,
		"mode":   model,
	})
	switch model {
	case "test":
		initTest(lt)
	default:
		initRelease(serverName, path, lt)
	}
}

func initTest(level logrus.Level) {
	// Log as JSON instead of the default ASCII formatter.
	logger.Formatter = &logrus.TextFormatter{}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.Out = os.Stdout

	// Only log the debug severity or above.
	logger.Level = level

	filenameHook := NewLineHook()
	filenameHook.Field = "line"
	logger.AddHook(filenameHook)

	_log.Debug("logger initialization is complete...")
}

func initRelease(serverName, logRootPath string, level logrus.Level) {
	logPath := logRootPath + serverName + "/"
	mkdirLogRootPath(logPath)

	writer, err := rotatelogs.New(
		logPath+"%Y-%m-%d_%H:%M.txt",
		rotatelogs.WithLinkName(logPath),          // WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)

	if err != nil {
		logrus.Fatalf("config local file system logger error. %+v", err)
	}

	filenameHook := NewLineHook()
	filenameHook.Field = "line"
	logger.AddHook(filenameHook)

	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		},
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	logger.AddHook(lfHook)
	logger.Level = level

	_log.Debug("logger initialization is complete...")
}

// pathExists 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkdirLogRootPath(path string) {
	exist, err := pathExists(path)
	if err != nil {
		logrus.Fatalf("get dir error![%v]\n", err)
		return
	}

	if !exist {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logrus.Fatalf("mkdir %s failed![%v]\n", path, err)
		}
	}
}
