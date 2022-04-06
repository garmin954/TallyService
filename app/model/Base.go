package model

import "time"

type BaseModel struct {
	ID        int       `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
}
