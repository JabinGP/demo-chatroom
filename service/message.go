package service

import (
	"errors"
	"time"

	"github.com/JabinGP/demo-chatroom/model/pojo"
	"github.com/jinzhu/gorm"
)

// MessageService message service
type MessageService struct {
	db *gorm.DB
}

// Query query message by ID, senderID, receiverID, beginTime, endTime
func (messageService *MessageService) Query(beginID uint, beginTime time.Time, endTime time.Time, senderID uint, receiverID uint) ([]pojo.Message, error) {
	var msgList []pojo.Message

	// Fill nested struct User
	tmpDB := messageService.db.Preload("Sender").Preload("Receiver")

	// Limit sender
	if senderID != 0 {
		tmpDB = tmpDB.Where("sender_id = ?", senderID)
	}

	// Limit receiver
	tmpDB = tmpDB.Where("receiver_id in (?,?)", 0, receiverID)

	// Limit begin time
	tmpDB = tmpDB.Where("send_time >= ?", beginTime)
	// Limit end time
	if endTime != time.Unix(0, 0) && (endTime != time.Time{}) {
		tmpDB = tmpDB.Where("send_time <= ?", endTime)
	}

	// Limit message id
	tmpDB = tmpDB.Where("id >= ?", beginID)

	// Execute query
	if err := tmpDB.Find(&msgList).Error; err != nil {
		return nil, err
	}

	return msgList, nil
}

// Insert insert message and return insert id
func (messageService *MessageService) Insert(senderID uint, receiverID uint, content string) (uint, error) {
	msg := pojo.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		SendTime:   time.Now(),
	}

	// If no sender id
	if msg.SenderID == 0 {
		return 0, errors.New("Can't insert message with SenderID(0)")
	}

	// Execute insert
	if err := messageService.db.Table("messages").Create(&msg).Error; err != nil {
		return 0, err
	}

	return msg.ID, nil
}
