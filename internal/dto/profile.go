package dto

type ProfileResponse struct {
	BaseModel
	Slug       string  `json:"username"`
	Name       string  `json:"name"`
	Bio        *string `json:"bio"`
	AvatarURL  *string `json:"avatar_url"`
	Background string  `json:"background"`
}

type ProfileRequest struct {
	Slug       string  `json:"slug" binding:"required,min=3"`
	Name       string  `json:"name" binding:"required,min=3"`
	Bio        *string `json:"bio"`
	AvatarURL  *string `json:"avatar_url"`
	Background string  `json:"background" binding:"required"`
}
