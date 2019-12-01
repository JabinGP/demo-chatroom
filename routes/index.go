package routes

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"

	"github.com/JabinGP/demo-chatroom/models"
)

// Route 123
func Route(app *iris.Application, dbCon *gorm.DB) {
	type Res models.ResModel
	log.Println("1")
	v1 := app.Party("v1")
	v1.Get("/test1", func(ctx iris.Context) {
		res := models.ResModel{}
		res.WithData(iris.Map{
			"username": "JabinGP",
		})
		ctx.JSON(res)
	})

	v1.Get("/test2", func(ctx iris.Context) {
		type DbModel struct {
			Title string `json:"title"`
			ID    int64  `json:"id"`
		}
		var dbResList []DbModel
		res := models.ResModel{}
		dbCon.Select("id,title").Find(&dbResList)
		res.WithData(dbResList)
		ctx.JSON(res)
	})
}
