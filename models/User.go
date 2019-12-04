package models

import "github.com/jinzhu/gorm"

// User table users model
type User struct {
	gorm.Model
	Username string
	Passwd   string
	// 1 -> girl, 2 -> boy
	Gender   int64
	Age      int64
	Interest string
}
