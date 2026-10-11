package models

import (
	"time"
	"uuid"
)

type Session struct {
	BaseModel
	UserID    uuid.UUID `gorm:"type:uuid"`
	User      User      `gorm:"foreignKey:UserID"`
	TokenHash string    `gorm:"unique"`
	ExpiresAt time.Time
}
