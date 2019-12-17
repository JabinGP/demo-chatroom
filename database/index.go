package database

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/JabinGP/demo-chatroom/config"
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

// Init when use mysql
func initMysql() {

	dbType := config.Viper.GetString("database.driver")
	dbHost := config.Viper.GetString("mysql.dbHost")
	dbPort := config.Viper.GetString("mysql.dbPort")
	dbName := config.Viper.GetString("mysql.dbName")
	dbParams := config.Viper.GetString("mysql.dbParams")
	dbUser := config.Viper.GetString("mysql.dbUser")
	dbPasswd := config.Viper.GetString("mysql.dbPasswd")

	dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbUser, dbPasswd, dbHost, dbPort, dbName, dbParams)

	var err error
	DB, err = gorm.Open(dbType, dbURL)
	if err != nil {
		log.Printf("Open mysql failed,err:%v\n", err)
		panic(err)
	}
}

// Auto init table if not exist
func initTable() {
	// Auto create table
	DB.AutoMigrate(&pojo.User{})
	DB.AutoMigrate(&pojo.Message{})

	// Enable Logger, show detailed log
	DB.LogMode(true)
}
