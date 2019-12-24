package middleware

import (
	"github.com/JabinGP/demo-chatroom/model"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
)

// Logined 获取JWT中的用户数据，封装成实体存放在ctx中供请求调用
var Logined context.Handler

func initUserInfo() {
	Logined = func(ctx context.Context) {
		jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
		id := int64(jwtInfo["userId"].(float64))
		username := jwtInfo["userName"].(string)
		logined := model.Logined{
			ID:       id,
			Username: username,
		}
		ctx.Values().Set("logined", logined)
		ctx.Next()
	}
}
