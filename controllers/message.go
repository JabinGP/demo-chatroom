package controllers

import (
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
	}

	if err := db.Table("messages").Create(&newMsg).Error; err != nil {
		ctx.JSON(new(models.ResModel).WithError(err.Error()))
		return
	}

	ctx.JSON(new(models.ResModel).WithData(nil))
}

// GetMessageList get message that send to me
func GetMessageList(ctx iris.Context) {
	type Res struct {
		SenderID   uint   `json:"senderId"`
		SenderName string `json:"senderName" gorm:"column:username"`
		Content    string `json:"content"`
	}

	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(jwtInfo["userId"].(float64))

	msgList := []Res{}
	if err := db.Table("messages").
		Select("*,users.username").
		Where("receiver_id = ?", userID).
		Joins("left join users on messages.sender_id = users.id").
		Find(&msgList).Error; err != nil {
		ctx.JSON(new(models.ResModel).WithError(err.Error()))
		return
	}

	ctx.JSON(new(models.ResModel).WithData(msgList))
}
