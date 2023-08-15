package moduleA

import "gorm.io/gorm"

type User struct {
	gorm.Model `gorm:"embedded"`
	Name       string
}

type UserAccessor interface {
	GetUserByID(userID uint) *User
}

type UserServiceImpl struct {
	// Dependencies or data store
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (u *UserServiceImpl) GetUserByID(userID uint) *User {
	// Simulate getting user data from a database
	return &User{Model: gorm.Model{ID: uint(userID)}, Name: "John Doe"}
}
