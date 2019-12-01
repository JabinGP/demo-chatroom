package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Con connect and return *gorm.DB
func Con(dbType string, conURL string) (*gorm.DB, error) {
	dbCon, err := gorm.Open(dbType, conURL)
	if err != nil {
		log.Printf("Open mysql failed,err:%v\n", err)
		return nil, err
	}
	return dbCon, nil
}
