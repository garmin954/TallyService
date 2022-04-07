package model

import "time"

// 账本记录
var LedgerRecordTable string = "ledger_records"

type LedgerRecord struct {
	BaseModel
<<<<<<< HEAD
	Title      string    `gorm:"type:varchar(50);not null;default:'';comment:标题" json:"title" form:"title" binding:"required"`
	Type       int8      `gorm:"type:tinyint(4);not null;default:0;comment:类型"  json:"type" form:"type""`
	Amount     float32   `gorm:"type:decimal(10,2);not null;default:0.00;comment:金额" json:"amount" form:"amount" binding:"required"`
	Desc       string    `gorm:"type:tinytext;not null;comment:描述" json:"desc" form:"desc"`
	Vouchers   string    `gorm:"type:tinytext;not null;comment:收据凭证" json:"vouchers" form:"vouchers"`
	Date       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:消费日期" json:"date" form:"date"`
	ClassifyId int       `gorm:"index;type:int(8);comment:分类id" json:"classify_id" form:"classify_id" binding:"required"`
	LedgerId   int       `gorm:"index;type:int(8);comment:账本id" json:"ledger_id" form:"ledger_id"`
	Uid        int       `gorm:"index;type:int(8);comment:用户ID" json:"uid" form:"-"`
=======
	Title      string    `gorm:"type:varchar(50);not null;default:'';comment:标题"`
	Type       int8      `gorm:"type:tinyint(4);not null;default:0;comment:类型"`
	Amount     float32   `gorm:"type:decimal(10,2);not null;default:0.00;comment:金额"`
	Desc       string    `gorm:"type:tinytext;not null;comment:描述"`
	Vouchers   string    `gorm:"type:tinytext;not null;comment:收据凭证"`
	Date       time.Time `gorm:"not null;comment:消费日期"`
	ClassifyId int       `gorm:"index;type:int(8);comment:分类id"`
	LedgerId   int       `gorm:"index;type:int(8);comment:账本id"`
	Uid        int       `gorm:"index;type:int(8);comment:用户ID"`
>>>>>>> f5f42371f1c38e0ec7b08a45c1148b681c15329c
}
