package main

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	_ "github.com/go-sql-driver/mysql"

	"github.com/JabinGP/demo-chatroom/config"
	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/JabinGP/demo-chatroom/routes"
)

func main() {
	// // read config
	// initConfig()
	// // connect database
	// initDb()

	app := iris.New()
	// set logger level
	app.Logger().SetLevel(config.Viper.GetString("server.logger.level"))
	// add recover to recover from any http-relative panics
	app.Use(recover.New())
	// add logger to log the requests to the terminal
	app.Use(logger.New())

	// globally allow options method to enable CORS
	app.AllowMethods(iris.MethodOptions)
	// add global CORS handler
	app.Use(middleware.CORS)

	// router
	routes.Route(app)

	// listen in 8888 port
	app.Run(iris.Addr(":8888"), iris.WithoutServerError(iris.ErrServerClosed))
}

// func initConfig() {
// 	// scan the file named config in the root directory
// 	viper.AddConfigPath("./")
// 	viper.SetConfigName("config")

// 	// read config, if failed, configure by default
// 	if err := viper.ReadInConfig(); err == nil {
// 		log.Println("Read config successfully: ", viper.ConfigFileUsed())
// 	} else {
// 		log.Printf("Read failed: %s \n", err)
// 		panic(err)
// 	}
// }
