package pojo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Message message object model
type Message struct {
	gorm.Model
	SenderID   uint
	Sender     User `gorm:"ForeignKey:SenderID"`
	ReceiverID uint
	Receiver   User `gorm:"ForeignKey:ReceiverID"`
	Content    string
	SendTime   time.Time
}
