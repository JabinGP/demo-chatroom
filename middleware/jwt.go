package middleware

import (
	"github.com/JabinGP/demo-chatroom/model"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var (
	// JWT JWT Middleware
	JWT *jwt.Middleware
)

func initJWT() {
	JWT = jwt.New(jwt.Config{
		ErrorHandler: func(ctx context.Context, err error) {
			if err == nil {
				return
			}
			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(model.ErrorUnauthorized(err))
		},

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})
}
