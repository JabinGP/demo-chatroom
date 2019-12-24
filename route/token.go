package route

import (
	"github.com/JabinGP/demo-chatroom/controller"
	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeToken(party router.Party) {
	party.Get("/token/info", middleware.JWT.Serve, middleware.Logined, controller.GetTokenInfo)
}
