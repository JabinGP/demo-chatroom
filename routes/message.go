package routes

import (
	"github.com/JabinGP/demo-chatroom/controllers"
	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeMessage(party router.Party) {
	party.Post("/message", middleware.JWT.Serve, controllers.SendMessage)
	party.Get("/message", middleware.JWT.Serve, controllers.GetMessageList)
}
