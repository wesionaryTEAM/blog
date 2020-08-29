package model

// UserModel usermodel
type UserModel struct {
	ID       uint   `gorm:"primary_key" json:"ID"`
	Username string `gorm:"column:username" json:"Username" `
	Email    string `gorm:"column:email;unique_index" json:"Email"`
	Bio      string `gorm:"column:bio;size:1024" json:"Bio"`
	Image    string `gorm:"column:image" json:"Image"`
	Password string `gorm:"column:password;not null" json:"Password"`
}
