package route

import (
	"github.com/JabinGP/demo-chatroom/controller"
	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeUser(party router.Party) {
	party.Post("/login", controller.PostLogin)

	party.Post("/user", controller.PostUser)
	party.Get("/user", controller.GetUser)
	party.Put("/user", middleware.JWT.Serve, middleware.Logined, controller.PutUser)
}
