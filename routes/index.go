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
		v1.Get("/user", middleware.JWT.Serve, controllers.GetUserList)
	}

}
