package controller

import (
	"github.com/JabinGP/demo-chatroom/model"
	"github.com/JabinGP/demo-chatroom/model/pojo"
	"github.com/JabinGP/demo-chatroom/model/reqo"
	"github.com/JabinGP/demo-chatroom/model/reso"
	"github.com/JabinGP/demo-chatroom/tool"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"log"
)

// PostLogin user login
func PostLogin(ctx iris.Context) {
	req := reqo.PostLogin{}
	ctx.ReadJSON(&req)

	// Query user by username
	user, err := userService.QueryByUsername(req.Username)
	if err != nil {
		ctx.JSON(new(model.ResModel).WithError(err.Error()))
		return
	}

	log.Println(user, req)
	// If passwd are inconsistent
	if user.Passwd != req.Passwd {
		ctx.JSON(new(model.ResModel).WithError("username or passwd error"))
		return
	}

	// Login Ok
	// Get token
	token, err := tool.GetJWTString(user.Username, user.ID)
	if err != nil {
		ctx.JSON(new(model.ResModel).WithError(err.Error()))
	}

	res := reso.PostLogin{
		Username: user.Username,
		ID:       user.ID,
		Token:    token,
	}
	ctx.JSON(new(model.ResModel).WithData(res))
}

// PostUser user register
func PostUser(ctx iris.Context) {
	req := reqo.PostUser{}
	ctx.ReadJSON(&req)

	// Username and passwd can't be blank
	if req.Username == "" || req.Passwd == "" {
		ctx.JSON(new(model.ResModel).WithError("username or passwd cann't be blank"))
		return
	}

	// Query user with same username
	exist, _ := userService.QueryByUsername(req.Username)

	// Can't be same username
	if exist.Username != "" {
		ctx.JSON(new(model.ResModel).WithError("existed username"))
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
		ctx.JSON(new(model.ResModel).WithError(err.Error()))
		return
	}

	res := reso.PostUser{
		Username: newUser.Username,
		ID:       userID,
	}
	ctx.JSON(new(model.ResModel).WithData(res))
}

// GetUser return user list
func GetUser(ctx iris.Context) {
	req := reqo.GetUser{}
	ctx.ReadQuery(&req)
	resList := []reso.GetUser{}

	userList, err := userService.Query(req.Username, req.ID)
	if err != nil {
		ctx.JSON(new(model.ResModel).WithError(err.Error()))
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
	ctx.JSON(new(model.ResModel).WithData(resList))
}

// PutUser update user information
func PutUser(ctx iris.Context) {
	req := reqo.PutUser{}
	ctx.ReadJSON(&req)
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(jwtInfo["userId"].(float64))

	// // Query user by userID
	// user, err := userService.QueryByID(userID)
	// if err != nil {
	// 	ctx.JSON(new(model.ResModel).WithError(err.Error()))
	// 	return
	// }

	user := pojo.User{}
	user.ID = userID
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
		ctx.JSON(new(model.ResModel).WithError(err.Error()))
		return
	}

	// Get updated user
	updatedUser, err := userService.QueryByID(user.ID)
	if err != nil {
		ctx.JSON(new(model.ResModel).WithError(err.Error()))
		return
	}

	res := reso.PutUser{
		ID:       updatedUser.ID,
		Username: updatedUser.Username,
		Gender:   updatedUser.Gender,
		Age:      updatedUser.Age,
		Interest: updatedUser.Interest,
	}
	ctx.JSON(new(model.ResModel).WithData(res))
}
