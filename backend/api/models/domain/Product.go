package domain

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id            int            `gorm:"primaryKey; autoIncrement"`
	Name          string         `gorm:"type:varchar(255);not null"`
	PurchasePrice float64        `gorm:"type:decimal(10,2);not null"`
	SellingPrice  float64        `gorm:"type:decimal(10,2);not null"`
	Stock         int            `gorm:"type:INT"`
	CreatedAt     time.Time      `gorm:"autoCreateTime;not null"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime;not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
