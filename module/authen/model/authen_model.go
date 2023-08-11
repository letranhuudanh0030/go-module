package model

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserSSO struct {
	IsAdmin  bool
	UserFile string
	Birthday string
	Gender   string
}
