package controllers

import (
	"github.com/JabinGP/demo-chatroom/models"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// VerifyToken Verify that the token is valid and return token info
func VerifyToken(ctx iris.Context) {
	type Res struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(jwtInfo["userId"].(float64))
	userName := jwtInfo["userName"].(string)

	res := Res{
		ID:       userID,
		Username: userName,
	}

	ctx.JSON(new(models.ResModel).WithData(res))
}
