package route

import (
	"github.com/kataras/iris/v12"
)

// Route ...
func Route(app *iris.Application) {
	routeStatic(app)
	routeRedirect(app)
	v1 := app.Party("/v1")
	{
		routeToken(v1)
		routeUser(v1)
		routeMessage(v1)
	}
}
