package model

// 账本关联用户
var LedgerUserTable string = "ledger_users"

type LedgerUser struct {
	BaseModel
	Status   int `gorm:"type:tinyint(4);comment:状态" form:"status"  json:"status"`
	LedgerId int `gorm:"index;type:int(8);comment:账本id" form:"Ledger_id"  json:"Ledger_id"`
	Uid      int `gorm:"index;type:int(8);comment:用户ID" form:"uid"  json:"uid"`
}
