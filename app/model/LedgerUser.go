package model

// 账本关联用户
var LedgerUserTable string = "ledger_users"

type LedgerUser struct {
	BaseModel
	Status   int8 `gorm:"type:tinyint(4);comment:状态"`
	LedgerId int  `gorm:"index;type:int(8);comment:账本id"`
	Uid      int  `gorm:"index;type:int(8);comment:用户ID"`
}
