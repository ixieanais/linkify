package dto

import (
	"uuid"
)

type BlockResponse struct {
	BaseModel
	ProfileID uuid.UUID        `json:"profile_id"`
	Profile   *ProfileResponse `json:"profile,omitempty"`
	Title     string           `json:"title"`
	Text      *string          `json:"text"`
	Image     *string          `json:"image"`
	URL       *string          `json:"url"`
	Type      string           `json:"type"`
	IsActive  *bool            `json:"is_active"`
}

type BlockRequest struct {
	ProfileID uuid.UUID `json:"profile_id"`
	Title     string    `json:"title"`
	Text      *string   `json:"text"`
	Image     *string   `json:"image"`
	URL       *string   `json:"url"`
	Type      string    `json:"type"`
	IsActive  *bool     `json:"is_active"`
}
