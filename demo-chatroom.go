package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	"github.com/spf13/viper"

	"github.com/JabinGP/demo-chatroom/models"
)

func main() {
	// init config
	initConfig()

	// init database
	dbConf := models.DbConf{
		DbHost:   "1",
		DbPort:   "2",
		DbName:   "3",
		DbParams: "4",
		DbUser:   "5",
		DbPasswd: "6",
	}

	log.Println(dbConf)
	app := iris.New()
	app.Logger().SetLevel(viper.GetString("server.logger.level"))

	// add recover to recover from any http-relative panics
	app.Use(recover.New())
	// add logger to log the requests to the terminal
	app.Use(logger.New())

	// router
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Iris Running")
	})

	app.Get("/{a:string}/{b:string}", func(ctx iris.Context) {
		a := ctx.Params().Get("a")
		b := ctx.Params().Get("b")
		ctx.HTML("a=" + a + ",b=" + b)
	})

	// listen in 8888 port
	app.Run(iris.Addr(":8888"), iris.WithoutServerError(iris.ErrServerClosed))
}

func initConfig() {
	// viper.SetConfigType("toml")

	// 扫描项目根目录下名为config的配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	// 读取
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("读取配置文件成功:", viper.ConfigFileUsed())
	} else {
		log.Printf("读取配置文件失败: %s \n", err)
		log.Println("将以系统预设值作为参数默认值")
	}
}
