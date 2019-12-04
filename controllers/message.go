package controllers

import (
	"time"

	"github.com/JabinGP/demo-chatroom/models"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// SendMessage send message
func SendMessage(ctx iris.Context) {
	type Req struct {
		ReceiverID uint   `json:"receiverId"`
		Content    string `json:"content"`
	}

	req := Req{}
	ctx.ReadJSON(&req)
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(jwtInfo["userId"].(float64))

	newMsg := models.Message{
		SenderID:   userID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
		SendTime:   time.Now(),
	}

	if err := db.Table("messages").Create(&newMsg).Error; err != nil {
		ctx.JSON(new(models.ResModel).WithError(err.Error()))
		return
	}

	ctx.JSON(new(models.ResModel).WithData(nil))
}

// GetMessageList get message that send to me
func GetMessageList(ctx iris.Context) {
	type Req struct {
		Time      int64 `json:"time"`
		BeginTime int64 `json:"beginTime"`
		EndTime   int64 `json:"endTime"`
	}
	type Res struct {
		ID         uint   `json:"id"`
		SenderID   uint   `json:"senderId"`
		SenderName string `json:"senderName" gorm:"column:username"`
		Content    string `json:"content"`
		SendTime   int64  `json:"sendTime"`
	}

	req := Req{}
	ctx.ReadQuery(&req)
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(jwtInfo["userId"].(float64))

	msgList := []Res{}

	tmpDB := db.Table("messages").
		Joins("left join users on messages.sender_id = users.id").
		Select(`messages.id,
						unix_timestamp(messages.send_time) as send_time, 
						users.username,
						messages.sender_id,
						messages.content`)

	if req.EndTime == 0 { // no limited end time
		tmpDB = tmpDB.Where(`(messages.receiver_id = ? 
													or messages.receiver_id = 0)
													and unix_timestamp(messages.send_time) >= ? `,
			userID, req.BeginTime)
	} else { // has limited end time
		tmpDB = tmpDB.Where(`(messages.receiver_id = ? 
													or messages.receiver_id = 0)
													and unix_timestamp(messages.send_time) >= ? 
													and unix_timestamp(messages.send_time) <= ? `,
			userID, req.BeginTime, req.EndTime)
	}

	if err := tmpDB.Find(&msgList).Error; err != nil {
		ctx.JSON(new(models.ResModel).WithError(err.Error()))
	}
	ctx.JSON(new(models.ResModel).WithData(msgList))
}
