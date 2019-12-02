package models

import "github.com/jinzhu/gorm"

// User table users model
type User struct {
	gorm.Model
	Username string
	Passwd   string
}
