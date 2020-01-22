package controller

import (
	"errors"
	"log"

	"github.com/JabinGP/demo-chatroom/model"
	"github.com/JabinGP/demo-chatroom/model/pojo"
	"github.com/JabinGP/demo-chatroom/model/reqo"
	"github.com/JabinGP/demo-chatroom/model/reso"
	"github.com/JabinGP/demo-chatroom/tool"
	"github.com/kataras/iris/v12"
)

// PostLogin user login
func PostLogin(ctx iris.Context) {
	req := reqo.PostLogin{}
	ctx.ReadJSON(&req)

	// Query user by username
	user, err := userService.QueryByUsername(req.Username)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	log.Println(user, req)
	// If passwd are inconsistent
	if user.Passwd != req.Passwd {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(errors.New("用户名或密码错误")))
		return
	}

	// Login Ok
	// Get token
	token, err := tool.GetJWTString(user.Username, user.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorBuildJWT(err))
	}

	res := reso.PostLogin{
		Username: user.Username,
		ID:       user.ID,
		Token:    token,
	}
	ctx.JSON(res)
}

// PostUser user register
func PostUser(ctx iris.Context) {
	req := reqo.PostUser{}
	ctx.ReadJSON(&req)

	// Username and passwd can't be blank
	if req.Username == "" || req.Passwd == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(errors.New("用户名和密码不能为空")))
		return
	}

	// Query user with same username
	exist, _ := userService.QueryByUsername(req.Username)

	// Can't be same username
	if exist.Username != "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(errors.New("用户名已存在")))
		return
	}

	// New user and insert into database
	newUser := pojo.User{
		Username: req.Username,
		Passwd:   req.Passwd,
		Gender:   req.Gender,
		Age:      req.Age,
		Interest: req.Interest,
	}
	userID, err := userService.Insert(newUser)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorInsertDatabase(err))
		return
	}

	res := reso.PostUser{
		Username: newUser.Username,
		ID:       userID,
	}
	ctx.JSON(res)
}

// GetUser return user list
func GetUser(ctx iris.Context) {
	req := reqo.GetUser{}
	ctx.ReadQuery(&req)
	resList := []reso.GetUser{}

	userList, err := userService.Query(req.Username, req.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	for _, user := range userList {
		res := reso.GetUser{
			ID:       user.ID,
			Username: user.Username,
			Gender:   user.Gender,
			Age:      user.Age,
			Interest: user.Interest,
		}

		resList = append(resList, res)
	}
	ctx.JSON(resList)
}

// PutUser update user information
func PutUser(ctx iris.Context) {
	req := reqo.PutUser{}
	ctx.ReadJSON(&req)
	logined := ctx.Values().Get("logined").(model.Logined)

	// // Query user by userID
	// user, err := userService.QueryByID(userID)
	// if err != nil {
	// 	ctx.JSON(new(model.ResModel).WithError(err.Error()))
	// 	return
	// }

	user := pojo.User{
		ID: logined.ID,
	}
	// Replace if set
	if req.Gender != 0 {
		user.Gender = req.Gender
	}
	if req.Age != 0 {
		user.Age = req.Age
	}
	if req.Interest != "" {
		user.Interest = req.Interest
	}

	// Update user
	err := userService.Update(user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	// Get updated user
	updatedUser, err := userService.QueryByID(user.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	res := reso.PutUser{
		ID:       updatedUser.ID,
		Username: updatedUser.Username,
		Gender:   updatedUser.Gender,
		Age:      updatedUser.Age,
		Interest: updatedUser.Interest,
	}
	ctx.JSON(res)
}
