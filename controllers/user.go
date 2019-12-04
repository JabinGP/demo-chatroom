package controllers

import (
	"log"

	"github.com/JabinGP/demo-chatroom/models"
	"github.com/JabinGP/demo-chatroom/tools"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// Login user login
func Login(ctx iris.Context) {
	type Req struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
	}

	type Res struct {
		Username string `json:"username"`
		ID       uint   `json:"id"`
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
		ID:       user.ID,
		Token:    token,
	}
	ctx.JSON(new(models.ResModel).WithData(res))
}

// Register user register
func Register(ctx iris.Context) {
	type Req struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
		Gender   int64  `json:"gender"`
		Age      int64  `json:"age"`
		Interest string `json:"interest"`
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
	db.Where("username = ?", req.Username).Select("username").First(&exist)

	// can't be same username
	if exist.Username != "" {
		ctx.JSON(new(models.ResModel).WithError("existed username"))
		return
	}

	// new user and insert into database
	newUser := models.User{
		Username: req.Username,
		Passwd:   req.Passwd,
		Gender:   req.Gender,
		Age:      req.Age,
		Interest: req.Interest,
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
		ID       uint   `json:"id"`
	}
	type Res struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Gender   int64  `json:"gender"`
		Age      int64  `json:"age"`
		Interest string `json:"interest"`
	}

	req := Req{}
	ctx.ReadQuery(&req)
	res := []Res{}

	if req.ID != 0 {
		db.Table("users").Where("username like ? and id = ?", "%"+req.Username+"%", req.ID).Find(&res)
	} else {
		db.Table("users").Where("username like ?", "%"+req.Username+"%").Find(&res)
	}

	ctx.JSON(new(models.ResModel).WithData(res))
}

// UpdateUser update user information
func UpdateUser(ctx iris.Context) {
	type Req struct {
		Gender   int64  `json:"gender"`
		Age      int64  `json:"age"`
		Interest string `json:"interest"`
	}

	type Res struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Gender   int64  `json:"gender"`
		Age      int64  `json:"age"`
		Interest string `json:"interest"`
	}

	req := Req{}
	ctx.ReadJSON(&req)
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(jwtInfo["userId"].(float64))

	user := models.User{}
	if err := db.Table("users").Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(new(models.ResModel).WithError(err.Error()))
		return
	}

	if req.Gender != 0 {
		user.Gender = req.Gender
	}

	if req.Age != 0 {
		user.Age = req.Age
	}

	if req.Interest != "" {
		user.Interest = req.Interest
	}

	if err := db.Table("users").Where("id = ?", userID).Update(&user).Error; err != nil {
		ctx.JSON(new(models.ResModel).WithError(err.Error()))
		return
	}

	res := Res{
		ID:       userID,
		Username: user.Username,
		Gender:   user.Gender,
		Age:      user.Age,
		Interest: user.Interest,
	}
	ctx.JSON(new(models.ResModel).WithData(res))
}
