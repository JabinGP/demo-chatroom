package controllers

import (
	"log"

	"github.com/JabinGP/demo-chatroom/database"
	"github.com/JabinGP/demo-chatroom/models"
	"github.com/JabinGP/demo-chatroom/tools"
	"github.com/kataras/iris/v12"
)

var db = database.DB

// Login user login
func Login(ctx iris.Context) {
	type Req struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
	}

	type Res struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}

	req := Req{}
	ctx.ReadJSON(&req)

	user := models.User{}
	// if username unexisted
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		log.Println(err)
		ctx.JSON(new(models.ResModel).WithError("unexisted username"))
		return
	}

	// if passwd are inconsistent
	if user.Passwd != req.Passwd {
		ctx.JSON(new(models.ResModel).WithError("username or passwd error"))
		return
	}

	// ok
	// get token
	token, err := tools.GetJWTString(user.Username, user.ID)
	if err != nil {
		ctx.JSON(new(models.ResModel).WithError(err.Error()))
	}

	res := Res{
		Username: user.Username,
		Token:    token,
	}
	ctx.JSON(new(models.ResModel).WithData(res))
}

// Register user register
func Register(ctx iris.Context) {
	type Req struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
	}
	type Res struct {
		Username string `json:"username"`
	}
	req := Req{}
	ctx.ReadJSON(&req)

	// username and passwd can't be blank
	if req.Username == "" || req.Passwd == "" {
		ctx.JSON(new(models.ResModel).WithError("username or passwd cann't be blank"))
		return
	}

	exist := models.User{}
	db.Select("username").First(&exist)

	// can't be same username
	if exist.Username != "" {
		ctx.JSON(new(models.ResModel).WithError("existed username"))
		return
	}

	// new user and insert into database
	newUser := models.User{
		Username: req.Username,
		Passwd:   req.Passwd,
	}
	if err := db.Create(&newUser).Error; err != nil {
		ctx.JSON(new(models.ResModel).WithError("insert into database error"))
		return
	}

	res := Res{
		Username: newUser.Username,
	}
	ctx.JSON(new(models.ResModel).WithData(res))
}

// GetUserList return user list
func GetUserList(ctx iris.Context) {
	type Req struct {
		Username string `json:"username"`
	}
	type Res struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	}

	req := Req{}
	ctx.ReadQuery(&req)
	res := []Res{}

	db.Table("users").Select("id,username").Where("username like ?", "%"+req.Username+"%").Find(&res)

	ctx.JSON(new(models.ResModel).WithData(res))
}
