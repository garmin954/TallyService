package model

import "time"

type BaseModel struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_at"`
}
