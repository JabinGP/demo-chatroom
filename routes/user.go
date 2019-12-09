package routes

import (
	"github.com/JabinGP/demo-chatroom/controllers"
	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeUser(party router.Party) {
	party.Post("/login", controllers.Login)

	party.Post("/user", controllers.Register)
	party.Get("/user", controllers.GetUserList)
	party.Put("/user", middleware.JWT.Serve, controllers.UpdateUser)
}
