package model

import (
	"time"

	"gorm.io/gorm"
)

// Base is a struct that represents the base model for all other models in the application.
// It includes fields for ID, creation time, update time, and deletion time.
type Base struct {
	ID        int            `json:"id",gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-",gorm:"index"`
}
