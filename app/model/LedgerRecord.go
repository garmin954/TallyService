package model

import "time"

// 账本记录
var LedgerRecordTable string = "ledger_records"

type LedgerRecord struct {
	BaseModel
	Type       int8      `gorm:"type:tinyint(4);not null;default:0;comment:类型"  json:"type" form:"type"" binding:"required"`
	Amount     float32   `gorm:"type:decimal(10,2);not null;default:0.00;comment:金额" json:"amount" form:"amount" binding:"required"`
	Remark     string    `gorm:"type:tinytext;not null;comment:描述" json:"remark" form:"remark"`
	Vouchers   string    `gorm:"type:tinytext;not null;comment:收据凭证" json:"vouchers" form:"vouchers"`
	Date       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:消费日期" json:"date" form:"date"`
	ClassifyId int       `gorm:"index;type:int(8);not null;comment:分类id" json:"classify_id" form:"classify_id" binding:"required"`
	LedgerId   int       `gorm:"index;type:int(8);not null;comment:账本id" json:"ledger_id" form:"ledger_id"`
	Uid        int       `gorm:"index;type:int(8);not null;comment:用户ID" json:"uid" form:"-"`
}
