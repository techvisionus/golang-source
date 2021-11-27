package models

import (
	"time"
)

type Order struct {
	ID         int64     `gorm:"column:id;primaryKey"`
	Code       string    `gorm:"column:code"`
	Product    Product   `gorm:"foreignKey:ID"`
	TotalPrice int64     `gorm:"column:total_price"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

// TableName sets the insert table name for this struct type
func (s *Order) TableName() string {
	return "orders"
}
