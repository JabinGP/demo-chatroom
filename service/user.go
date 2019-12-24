package service

import (
	"github.com/JabinGP/demo-chatroom/model/pojo"
	"xorm.io/xorm"
)

// UserService user service
type UserService struct {
	db *xorm.Engine
}

// Query query user by username and id
func (userService *UserService) Query(username string, id uint) ([]pojo.User, error) {
	var userList []pojo.User

	// Limit username
	tmpDB := userService.db.Where("username like ?", "%"+username+"%")

	// Limit id
	if id != 0 {
		tmpDB.Where("id = ?", id)
	}

	// Execute query
	if err := tmpDB.Find(&userList); err != nil {
		return nil, err
	}

	return userList, nil
}

// QueryByUsername return one user
func (userService *UserService) QueryByUsername(username string) (pojo.User, error) {
	var user = pojo.User{
		Username: username,
	}
	has, err := userService.db.Get(&user)
	if err != nil {
		return pojo.User{}, err
	}
	if !has {
		return pojo.User{}, nil
	}
	return user, nil
}

// QueryByID return one user
func (userService *UserService) QueryByID(id int64) (pojo.User, error) {
	var user = pojo.User{
		ID: id,
	}

	if _, err := userService.db.Get(&user); err != nil {
		return pojo.User{}, err
	}

	return user, nil
}

// Insert insert a new user and return id
func (userService *UserService) Insert(user pojo.User) (int64, error) {
	if _, err := userService.db.Insert(&user); err != nil {
		return 0, err
	}
	return user.ID, nil
}

// Update update user and return current user infomation
func (userService *UserService) Update(user pojo.User) error {
	if _, err := userService.db.Update(&user); err != nil {
		return err
	}
	return nil
}
