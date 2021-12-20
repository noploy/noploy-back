package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	ID         int64 `gorm:"primaryKey"`
	Name       string
	Hash       string
	LastLogin  *time.Time
	LastAccess *time.Time
}
