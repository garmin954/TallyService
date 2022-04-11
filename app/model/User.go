package model

var UserTable string = "users"

type User struct {
	BaseModel
	Username string `gorm:"index:idx_username,unique;type:varchar(20);not null;comment:昵称" json:"username"`
	Avatar   string `gorm:"type:varchar(250);not null;default:'';comment:头像" json:"avatar"`
	Mobile   string `gorm:"type:varchar(11);not null;default:'';comment:手机" json:"mobile"`
	Password string `gorm:"type:varchar(36);not null;default:'';comment:密码" json:"-"`
	Salt     string `gorm:"type:varchar(20);not null;comment:盐" json:"-"`
	Unionid  string `gorm:"index:idx_union_id,unique;type:varchar(36);not null;default:'';comment:union_id" json:"-"`
	Openid   string `gorm:"type:varchar(36);not null;default:'';comment:openid" json:"-"`
}
