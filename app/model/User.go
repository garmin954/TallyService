package model

import (
	"gorm.io/gorm"
)

var UserTable string = "users"

type User struct {
	BaseModel
	Username string `gorm:"index:idx_username,unique;type:varchar(20);not null;comment:昵称" json:"username"`
	Nickname string `gorm:"type:varchar(20);not null;comment:昵称" json:"nickname"`
	Avatar   string `gorm:"type:varchar(250);not null;default:'';comment:头像" json:"avatar"`
	Mobile   string `gorm:"type:varchar(11);not null;default:'';comment:手机" json:"mobile"`
	Password string `gorm:"type:varchar(36);not null;default:'';comment:密码" json:"-"`
	Salt     string `gorm:"type:varchar(20);not null;comment:盐" json:"-"`
	Unionid  string `gorm:"type:varchar(36);not null;default:'';comment:union_id" json:"-"`
	Openid   string `gorm:"index:idx_openid,unique;type:varchar(36);not null;default:'';comment:openid" json:"-"`
}

type MUser struct {
}

// 初始化用户的
func (user *MUser) InitUserLedger(tx *gorm.DB, uid int) error {
	// 创建账本
	ledger := Ledger{Name: "我的账本", Desc: "记录血汗钱的流失😅", Uid: uid, Cover: ""}

	tx.Table(LedgerTable).Create(&ledger)

	// 创建账本用户关联
	userLedger := LedgerUser{LedgerId: ledger.ID, Uid: uid, Status: 1}
	rs := tx.Table(LedgerUserTable).Create(&userLedger)

	if rs.Error != nil {
		return rs.Error
	}

	return nil
}
