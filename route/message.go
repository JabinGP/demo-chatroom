package route

import (
	"github.com/JabinGP/demo-chatroom/controller"
	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeMessage(party router.Party) {
	party.Post("/message", middleware.JWT.Serve, controller.PostMessage)
	// party.Get("/message", middleware.JWT.Serve, controller.GetMessageList)
	party.Get("/message", middleware.JWT.Serve, controller.GetMessage)
}
