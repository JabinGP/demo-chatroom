package database

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/JabinGP/demo-chatroom/config"
	"github.com/JabinGP/demo-chatroom/model"
	"github.com/JabinGP/demo-chatroom/model/pojo"
	"github.com/jinzhu/gorm"
)

var once sync.Once

// DB database connect
var DB *gorm.DB

func init() {
	once.Do(func() {
		dbType := config.Viper.GetString("database.driver")
		switch dbType {
		case "mysql":
			initMysql()
		default:
			panic(errors.New("only support mysql"))
		}

		initTable()
	})
}

// init when use mysql
func initMysql() {
	dbConf := model.DbConf{
		DbType:   config.Viper.GetString("database.driver"),
		DbHost:   config.Viper.GetString("mysql.dbHost"),
		DbPort:   config.Viper.GetString("mysql.dbPort"),
		DbName:   config.Viper.GetString("mysql.dbName"),
		DbParams: config.Viper.GetString("mysql.dbParams"),
		DbUser:   config.Viper.GetString("mysql.dbUser"),
		DbPasswd: config.Viper.GetString("mysql.dbPasswd"),
	}

	dbType := dbConf.DbType
	dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbConf.DbUser, dbConf.DbPasswd, dbConf.DbHost, dbConf.DbPort, dbConf.DbName, dbConf.DbParams)

	var err error
	DB, err = gorm.Open(dbType, dbURL)
	if err != nil {
		log.Printf("Open mysql failed,err:%v\n", err)
		panic(err)
	}
}

// auto init table if not exist
func initTable() {
	// auto create table
	DB.AutoMigrate(&pojo.User{})
	DB.AutoMigrate(&pojo.Message{})

	// Enable Logger, show detailed log
	DB.LogMode(true)
}
