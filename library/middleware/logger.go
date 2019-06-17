package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"myiris/library/tool"
)

var (
	logConf = logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	}
	excludeExtensions = [...]string{
		".js",
		".css",
		".jpg",
		".png",
		".ico",
		".svg",
	}

	logsDir = "./logs/"

	deleteFileOnExit = true
)

func UseLogFile(file bool) (h iris.Handler, close func() error) {
	close = func() error { return nil }
	switch file {
	case true:
		h, close = newRequestLoggerRelease()
	default:
		deleteFileOnExit = false
		h = newRequestLoggerDebug()
	}
	return
}

func newRequestLoggerDebug() iris.Handler {
	return logger.New(logConf)
}

func todayFileName() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFileName()
	//打开一个输出文件，如果重新启动服务器，它将追加到今天的文件中

	f, err := os.OpenFile(tool.CreatePath(logsDir+filename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func newRequestLoggerRelease() (h iris.Handler, close func() error) {
	close = func() error { return nil }

	logFile := newLogFile()

	close = func() error {
		err := logFile.Close()
		if deleteFileOnExit {
			err = os.Remove(logFile.Name())
		}
		return err
	}

	logConf.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		_, _ = logFile.Write([]byte(output))
	}

	//我们不想使用记录器，一些静态请求等
	logConf.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	h = logger.New(logConf)
	return
}
