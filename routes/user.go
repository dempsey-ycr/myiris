package routes

import (
	"github.com/kataras/iris"
	"myiris/controllers/user"
)

func RegisterUsers(app *iris.Application) {
	usersMiddleware := func(ctx iris.Context) {
		ctx.Next()
	}

	users := app.Party("/user", usersMiddleware)
	{
		users.Get("/get", user.NewUserInstance().GetInfo)
		users.Get("/token", user.NewUserInstance().GetToken)
		users.Post("/info",  user.NewUserInstance().PostUser)
	}

	setting := users.Party("/setting")
	{
		setting.Post("/change", user.NewUserInstance().Change)
		setting.Get("/info", user.NewUserInstance().GetInfo)
	}
}
