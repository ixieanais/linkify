package models

import (
	"uuid"
)

type Block struct {
	BaseModel
	ProfileID uuid.UUID `gorm:"type:uuid"`
	Profile   Profile   `gorm:"foreignKey:ProfileID"`
	Title     string
	Text      *string
	Image     *string
	URL       *string
	Type      string `gorm:"type:text;default:'link'"`
	IsActive  *bool  `gorm:"default:false"`
}
