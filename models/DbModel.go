package models

import "github.com/jinzhu/gorm"

// DbModel database table model
type DbModel struct {
	gorm.Model
	Title string
}
