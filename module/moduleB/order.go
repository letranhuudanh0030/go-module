package moduleB

import (
	"todo/module/moduleA"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model `gorm:"embedded"`
	UserID     uint `gorm:"foreignKey:UserID"`
}

type OrderService interface {
	GetOrderByID(orderID int) *OrderWithUser
}

type OrderServiceImpl struct {
	// Dependencies
	userAccessor moduleA.UserAccessor
}

func NewOrderServiceImpl(userAccessor moduleA.UserAccessor) *OrderServiceImpl {
	return &OrderServiceImpl{userAccessor: userAccessor}
}

func (o *OrderServiceImpl) GetOrderByID(orderID int) *OrderWithUser {
	order := &Order{
		Model:  gorm.Model{ID: uint(orderID)},
		UserID: 456,
	}
	user := o.userAccessor.GetUserByID(order.UserID)

	return &OrderWithUser{Order: order, User: user}
}

type OrderWithUser struct {
	*Order
	User *moduleA.User
}
