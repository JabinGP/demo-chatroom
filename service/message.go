package service

import (
	"errors"
	"time"

	"github.com/JabinGP/demo-chatroom/model/pojo"
	"xorm.io/xorm"
)

// MessageService message service
type MessageService struct {
	db *xorm.Engine
}

// Query query message by ID, senderID, receiverID, beginTime, endTime
func (messageService *MessageService) Query(beginID int64, beginTime int64, endTime int64, receiverID int64) ([]pojo.MessageWithUser, error) {
	var msgList []pojo.MessageWithUser

	// Query received message and sended message
	tmpDB := messageService.db.Where("message.receiver_id in (?,?) or message.sender_id = ?", 0, receiverID, receiverID)

	// // Get sender message
	// // if senderID != 0 {
	// tmpDB = tmpDB.Or("sender_id = ?", senderID)
	// // }

	// Limit begin time
	tmpDB = tmpDB.Where("message.send_time >= ?", beginTime)
	// Limit end time
	if endTime != 0 {
		tmpDB = tmpDB.Where("message.send_time <= ?", endTime)
	}

	// Limit message id
	tmpDB = tmpDB.Where("message.id > ?", beginID)

	tmpDB = tmpDB.Join("LEFT", []string{"user", "sender"}, "message.sender_id = sender.id")

	tmpDB = tmpDB.Join("LEFT", []string{"user", "receiver"}, "message.receiver_id = receiver.id")
	// Execute query
	if err := tmpDB.Find(&msgList); err != nil {
		return nil, err
	}

	return msgList, nil
}

// Insert insert message and return insert id
func (messageService *MessageService) Insert(senderID int64, receiverID int64, content string) (int64, error) {
	msg := pojo.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		SendTime:   time.Now().Unix(),
	}

	// If no sender id
	if msg.SenderID == 0 {
		return 0, errors.New("Can't insert message with SenderID(0)")
	}

	// Execute insert
	if _, err := messageService.db.Insert(&msg); err != nil {
		return 0, err
	}

	return msg.ID, nil
}
