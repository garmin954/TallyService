package model

import (
	"fmt"

	"gorm.io/gorm"
)

var UserTable string = "users"

type User struct {
	BaseModel
	Username string `gorm:"index:idx_username,unique;type:varchar(20);not null;comment:昵称" json:"username"`
	NickName string `gorm:"type:varchar(20);not null;comment:昵称" json:"nickname"`
	Avatar   string `gorm:"type:varchar(250);not null;default:'';comment:头像" json:"avatar"`
	Mobile   string `gorm:"type:varchar(11);not null;default:'';comment:手机" json:"mobile"`
	Password string `gorm:"type:varchar(36);not null;default:'';comment:密码" json:"-"`
	Salt     string `gorm:"type:varchar(20);not null;comment:盐" json:"-"`
	Unionid  string `gorm:"index:idx_union_id,unique;type:varchar(36);not null;default:'';comment:union_id" json:"-"`
	Openid   string `gorm:"type:varchar(36);not null;default:'';comment:openid" json:"-"`
}

type MUser struct {
}

// 初始化用户的
func (user *MUser) InitUserLedger(tx *gorm.DB, uid int) error {
	// 创建账本
	ledger := Ledger{Name: "我的账本", Desc: "记录血汗钱的流失", Uid: uid, Cover: ""}

	r := tx.Table(LedgerTable).Create(&ledger)
	fmt.Println(ledger.ID)
	if r.Error != nil {
		return r.Error
	}

	// 创建账本用户关联
	userLedger := LedgerUser{LedgerId: ledger.ID, Uid: uid, Status: 1}
	rs := tx.Table(LedgerUserTable).Create(userLedger)
	fmt.Println(rs)

	if rs.Error != nil {
		return rs.Error
	}

	return nil
}
