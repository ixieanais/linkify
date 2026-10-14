package models

import (
	"uuid"
)

type Profile struct {
	BaseModel
	UserID     uuid.UUID `gorm:"type:uuid"`
	User       User      `gorm:"foreignKey:UserID"`
	Slug       string    `gorm:"unique"`
	Name       string
	Bio        *string
	AvatarURL  *string
	Background string
}
