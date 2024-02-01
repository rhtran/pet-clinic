package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int            `json:"id",gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-",gorm:"index"`
}
