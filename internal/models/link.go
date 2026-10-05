package models

import (
	"uuid"
)

type Link struct {
	BaseModel
	UserID    uuid.UUID `gorm:"type:uuid"`
	User      User      `gorm:"foreignKey:UserID"`
	ProfileID uuid.UUID `gorm:"type:uuid"`
	Profile   Profile   `gorm:"foreignKey:ProfileID"`
	URL       string
	Image     string
	IsActive  *bool `gorm:"default:false"`
}
