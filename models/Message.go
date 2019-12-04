package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Message table messages model
type Message struct {
	gorm.Model
	SenderID   uint
	ReceiverID uint
	Content    string
	SendTime   time.Time
}
