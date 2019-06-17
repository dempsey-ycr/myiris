package web

import (
	"context"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"myiris/library/middleware"
	"myiris/routes"
)

func Iris() {
	configurator := iris.WithConfiguration(iris.YAML("./config/conf.d/iris.yaml"))

	app := iris.New()
	app.Logger().SetLevel("debug")
	log, closed := middleware.UseLogFile(false)
	app.Use(log)
	app.Use(recover.New())

	//app.Use(pprof.New()) // 启用pprof

	// listen CONTROL + C(COMMAND + C) 或者当发送的kill命令是ENABLED BY-DEFAULT
	iris.RegisterOnInterrupt(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = closed()
		_ = app.Shutdown(ctx)
	})

	routes.RegisterErrors(app)
	routes.RegisterUsers(app)

	// 监听 tcp 0.0.0.0:9090
	go func() {
		//_ = app.NewHost(&http.Server{Addr: ":9090"}).ListenAndServe()
		_ = http.ListenAndServe(":9090", nil) // http://localhost:9090/debug/pprof
	}()

	// 主监听 0.0.0.0:8080
	_ = app.Run(iris.Addr(":8080"), configurator)
}
