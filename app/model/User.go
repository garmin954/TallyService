package model

var UserTable string = "users"

type User struct {
	BaseModel
	Username string `gorm:"index:idx_username,unique;type:varchar(20);not null;comment:昵称"`
	Avatar   string `gorm:"type:varchar(250);not null;default:'';comment:头像"`
	Mobile   string `gorm:"type:varchar(11);not null;default:'';comment:手机"`
	Password string `gorm:"type:varchar(36);not null;default:'';comment:密码"`
	Salt     string `gorm:"type:varchar(20);not null;comment:盐"`
	Unionid  string `gorm:"type:varchar(100);not null;default:'';comment:union_id"`
}
