package model

var UserTable string = "users"

type User struct {
	BaseModel
	ID       int    `gorm:"primary_key"`
	Nickname string `gorm:"vachar(20)";not null`
	Unionid  string `gorm:"vachar(100)";not null`
}
