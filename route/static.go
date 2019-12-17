package route

import "github.com/kataras/iris/v12"

func routeStatic(app *iris.Application) {
	// static source
	app.HandleDir("/", "./assets")
}
