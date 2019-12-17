package service

import (
	"github.com/JabinGP/demo-chatroom/model/pojo"
	"github.com/jinzhu/gorm"
)

// UserService user service
type UserService struct {
	db *gorm.DB
}

// Query query user by username and id
func (userService *UserService) Query(username string, id uint) ([]pojo.User, error) {
	var userList []pojo.User

	// Limit username
	tmpDB := userService.db.Table("users").Where("username like ?", "%"+username+"%")

	// Limit id
	if id != 0 {
		tmpDB.Where("id = ?", id)
	}

	// Execute query
	if err := tmpDB.Find(&userList).Error; err != nil {
		return nil, err
	}

	return userList, nil
}

// QueryByUsername return one user
func (userService *UserService) QueryByUsername(username string) (pojo.User, error) {
	var user = pojo.User{}
	user.Username = username
	if err := userService.db.Model(&user).Where("username = ?", user.Username).First(&user).Error; err != nil {
		return pojo.User{}, err
	}

	return user, nil
}

// QueryByID return one user
func (userService *UserService) QueryByID(id uint) (pojo.User, error) {
	var user = pojo.User{}
	user.ID = id
	if err := userService.db.Model(&user).First(&user).Error; err != nil {
		return pojo.User{}, err
	}

	return user, nil
}

// Insert insert a new user and return id
func (userService *UserService) Insert(user pojo.User) (uint, error) {
	if err := userService.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// Update update user and return current user infomation
func (userService *UserService) Update(user pojo.User) error {
	if err := userService.db.Model(&user).Update(&user).Error; err != nil {
		return err
	}
	return nil
}
