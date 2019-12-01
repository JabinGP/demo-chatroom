package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	"github.com/spf13/viper"

	"github.com/JabinGP/demo-chatroom/database"
	"github.com/JabinGP/demo-chatroom/models"
	"github.com/JabinGP/demo-chatroom/routes"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// 数据库连接实例
	dbCon *gorm.DB
)

func main() {
	// read config
	initConfig()
	// connect database
	initDb()

	app := iris.New()

	// set logger level
	app.Logger().SetLevel(viper.GetString("server.logger.level"))
	// add recover to recover from any http-relative panics
	app.Use(recover.New())
	// add logger to log the requests to the terminal
	app.Use(logger.New())

	// router
	routes.Route(app, dbCon)

	// listen in 8888 port
	app.Run(iris.Addr(":8888"), iris.WithoutServerError(iris.ErrServerClosed))
}

func initConfig() {
	// scan the file named config in the root directory
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	// read config, if failed, configure by default
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Read config successfully: ", viper.ConfigFileUsed())
	} else {
		log.Printf("Read failed: %s \n", err)
		panic(err)
	}
}

// initDb
func initDb() {
	dbConf := models.DbConf{
		DbType:   viper.GetString("database.driver"),
		DbHost:   viper.GetString("mysql.dbHost"),
		DbPort:   viper.GetString("mysql.dbPort"),
		DbName:   viper.GetString("mysql.dbName"),
		DbParams: viper.GetString("mysql.dbParams"),
		DbUser:   viper.GetString("mysql.dbUser"),
		DbPasswd: viper.GetString("mysql.dbPasswd"),
	}

	dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbConf.DbUser, dbConf.DbPasswd, dbConf.DbHost, dbConf.DbPort, dbConf.DbName, dbConf.DbParams)

	var err error
	dbCon, err = database.Con(dbConf.DbType, dbURL)
	if err != nil {
		log.Printf("Open mysql failed,err:%v\n", err)
		panic(err)
	}
}
