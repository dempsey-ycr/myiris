package routes

import "github.com/kataras/iris"

const notFoundHTML = "<h1> custom http error page </h1>"

func RegisterErrors(app *iris.Application) {
	// set a custom 404 handler
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML(notFoundHTML)
	})
}
