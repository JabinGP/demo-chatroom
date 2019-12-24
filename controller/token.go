package controller

import (
	"github.com/JabinGP/demo-chatroom/model"
	"github.com/JabinGP/demo-chatroom/model/reso"
	"github.com/kataras/iris/v12"
)

// GetTokenInfo 验证token是否有效，如果有效则返回token携带的信息
func GetTokenInfo(ctx iris.Context) {
	logined := ctx.Values().Get("logined").(model.Logined)

	res := reso.GetTokenInfo{
		ID:       logined.ID,
		Username: logined.Username,
	}
	ctx.JSON(res)
}
