package controller

import (
	"time"

	"github.com/JabinGP/demo-chatroom/model"
	"github.com/JabinGP/demo-chatroom/model/reqo"
	"github.com/JabinGP/demo-chatroom/model/reso"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// PostMessage send message
func PostMessage(ctx iris.Context) {
	req := reqo.PostMessage{}
	ctx.ReadJSON(&req)
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(jwtInfo["userId"].(float64))

	insertID, err := messageService.Insert(userID, req.ReceiverID, req.Content)
	if err != nil {
		ctx.JSON(new(model.ResModel).WithError(err.Error()))
		return
	}

	res := reso.PostMessage{
		ID: insertID,
	}

	ctx.JSON(new(model.ResModel).WithData(res))
}

// GetMessage get all message
func GetMessage(ctx iris.Context) {
	req := reqo.GetMessage{}
	ctx.ReadQuery(&req)
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(jwtInfo["userId"].(float64))

	msgList, err := messageService.Query(
		req.BeginID,
		time.Unix(req.BeginTime, 0),
		time.Unix(req.EndTime, 0),
		0,
		userID,
	)
	if err != nil {
		ctx.JSON(new(model.ResModel).WithError(err.Error()))
		return
	}

	// Build response object
	var resList []reso.GetMessage

	for _, msg := range msgList {
		// Set private according to ReceiverID
		private := false
		if msg.ReceiverID != 0 {
			private = true
		}

		// Get single res
		res := reso.GetMessage{
			ID:         msg.ID,
			SenderID:   msg.SenderID,
			SenderName: msg.Sender.Username,
			Content:    msg.Content,
			SendTime:   msg.SendTime.Unix(),
			Private:    private,
		}

		// Add into resList
		resList = append(resList, res)
	}
	ctx.JSON(new(model.ResModel).WithData(resList))
}
