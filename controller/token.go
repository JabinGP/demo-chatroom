package controller

import (
	"github.com/JabinGP/demo-chatroom/model/reso"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// GetTokenInfo Verify that the token is valid and return token info
func GetTokenInfo(ctx iris.Context) {
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := int64(jwtInfo["userId"].(float64))
	userName := jwtInfo["userName"].(string)

	res := reso.GetTokenInfo{
		ID:       userID,
		Username: userName,
	}

	ctx.JSON((res))
}
