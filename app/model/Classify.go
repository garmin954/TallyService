package model

var ClassifyTable string = "classifies"

type Classify struct {
	BaseModel
	Name  string `gorm:"type:varchar(8);not null;comment:名称"`
	Cover string `gorm:"type:varchar(250);not null;default:'';comment:封面"`
	Sort  int    `gorm:"type:tinyint(4);not null;default:50;comment:排序"`
	Type  int8   `gorm:"type:tinyint(4);not null;default:0;comment:类型:1支出，2入账"`
	Uid   int    `gorm:"index;type:int(8);comment:用户ID"`
}
