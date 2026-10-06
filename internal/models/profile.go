package models

import (
	"uuid"
)

type Profile struct {
	BaseModel
	UserID    uuid.UUID `gorm:"type:uuid"`
	User      User      `gorm:"foreignKey:UserID"`
	Username  string
	URL       string `gorm:"unique"`
	Bio       *string
	AvatarURL *string
}
