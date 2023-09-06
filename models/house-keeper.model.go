package models

import (
	"time"
)

type HouseKeeper struct {
	ID        int64     `gorm:"primaryKey" json:"id" `
	FirstName string    `gorm:"size:128" json:"first_name"`
	LastName  string    `gorm:"size:64" json:"last_name"`
	Phone     string    `gorm:"size:64" json:"phone"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (HouseKeeper) TableName() string {
	return "house_keepers"
}
