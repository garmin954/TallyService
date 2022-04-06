package model

import "time"

// 账本记录
var LedgerRecordTable string = "ledger_records"

type LedgerRecord struct {
	BaseModel
	Title    string    `gorm:"type:varchar(50);not null;default:'';comment:标题"`
	Type     int8      `gorm:"type:tinyint(4);not null;default:0;comment:类型"`
	Amount   float32   `gorm:"type:decimal(10,2);not null;default:0.00;comment:金额"`
	Desc     string    `gorm:"type:tinytext;not null;comment:描述"`
	Vouchers string    `gorm:"type:tinytext;not null;comment:收据凭证"`
	Date     time.Time `gorm:"not null;comment:消费日期"`
	LedgerId int       `gorm:"index;type:int(8);comment:账本id"`
	Uid      int       `gorm:"index;type:int(8);comment:用户ID"`
}
