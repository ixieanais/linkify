package models

import (
	"uuid"
)

type Block struct {
	BaseModel
	UserID    uuid.UUID `gorm:"type:uuid"`
	User      User      `gorm:"foreignKey:UserID"`
	ProfileID uuid.UUID `gorm:"type:uuid"`
	Profile   Profile   `gorm:"foreignKey:ProfileID"`
	Title     string
	Text      *string
	Image     *string
	URL       *string
	Type      string `gorm:"type:text;default:'link'"`
	IsActive  *bool  `gorm:"default:false"`
}
