package controllers

import (
	"log"

	"github.com/JabinGP/demo-chatroom/database"
	"github.com/JabinGP/demo-chatroom/models"
	"github.com/kataras/iris/v12"
)

var db = database.DB

// Login user login
func Login(ctx iris.Context) {
	type loginReq struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
	}

	type loginRes struct {
		Username string `json:"username"`
	}

	req := loginReq{}
	ctx.ReadJSON(&req)

	user := models.User{}
	// if username unexisted
	if err := db.Select("passwd").Where("username = ?", req.Username).First(&user).Error; err != nil {
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
	res := loginRes{
		Username: req.Username,
	}
	ctx.JSON(new(models.ResModel).WithData(res))
}

// Register user register
func Register(ctx iris.Context) {
	type registerReq struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
	}
	type registerRes struct {
		Username string `json:"username"`
	}
	req := registerReq{}
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

	res := registerRes{
		Username: req.Username,
	}
	ctx.JSON(new(models.ResModel).WithData(res))
}
