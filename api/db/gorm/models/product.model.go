package models

import (
	"database/sql"
)

type Product struct {
	ID        int64        `gorm:"column:id;primaryKey"`
	Code      string       `gorm:"column:code"`
	Name      string       `gorm:"column:name"`
	Price     int64        `gorm:"column:price"`
	CreatedAt sql.NullTime `gorm:"column:created_at"`
}

// TableName sets the insert table name for this struct type
func (s *Product) TableName() string {
	return "products"
}
