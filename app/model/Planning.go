package model

// 账本
var PlanningTable string = "plannings"

type Planning struct {
	BaseModel
	ID       int     `gorm:"primary_key"`
	Name     string  `gorm:"type:varchar(50);not null;comment:计划名称"`
	Type     int8    `gorm:"type:tinyint(4);not null;default:0;comment:类型"`
	Amount   float32 `gorm:"type:decimal(10,2);not null;default:0.00;comment:额度"`
	Desc     string  `gorm:"type:tinytext;not null;comment:描述"`
	Cover    string  `gorm:"type:varchar(250);not null;default:'';comment:封面"`
	Uid      int     `gorm:"index;type:int(8);comment:创建人用户ID"`
	LedgerId int     `gorm:"index;type:int(8);comment:账本id"`
}
