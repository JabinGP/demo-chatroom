package routes

import (
	"log"

	"github.com/kataras/iris/v12"

	"github.com/JabinGP/demo-chatroom/controllers"
	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/JabinGP/demo-chatroom/models"
)

// Route 123
func Route(app *iris.Application) {
	type Res models.ResModel
	log.Println("1")
	v1 := app.Party("v1")
	{
		v1.Post("/login", controllers.Login)

		v1.Post("/user", controllers.Register)
		v1.Get("/user", controllers.GetUserList)
		v1.Put("/user", middleware.JWT.Serve, controllers.UpdateUser)

		v1.Post("/message", middleware.JWT.Serve, controllers.SendMessage)
		v1.Get("/message", middleware.JWT.Serve, controllers.GetMessageList)
	}

}
