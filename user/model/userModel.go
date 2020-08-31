package model

// User usermodel
type User struct {
	ID       uint   `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Bio      string `json:"Bio"`
	Image    string `json:"Image"`
	Password string `json:"Password" `
}

// LoginUser  userModel
type LoginUser struct {
	Email    string `json:"Email"  example:"email" binding:"required"`
	Password string `json:"Password" example:"password" binding:"required"`
}
