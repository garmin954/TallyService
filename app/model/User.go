package model

var UserTable string = "users"

type User struct {
	BaseModel
	Nickname string `gorm:"type:varchar(20);not null;comment:昵称"`
	Avatar   string `gorm:"type:varchar(250);not null;comment:头像"`
	Unionid  string `gorm:"type:varchar(100);not null;comment:union_id"`
}
