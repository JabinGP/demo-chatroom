package routes

import (
	"github.com/JabinGP/demo-chatroom/controllers"
	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeToken(party router.Party) {
	party.Get("/tokeninfo", middleware.JWT.Serve, controllers.VerifyToken)
}
