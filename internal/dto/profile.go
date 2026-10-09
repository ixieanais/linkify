package dto

type ProfileResponse struct {
	BaseModel
	Username   string  `json:"username"`
	Name       string  `json:"name"`
	Bio        *string `json:"bio"`
	AvatarURL  *string `json:"avatar_url"`
	Background string  `json:"background"`
}

type ProfileRequest struct {
	Username   string  `json:"username"`
	Name       string  `json:"name"`
	Bio        *string `json:"bio"`
	AvatarURL  *string `json:"avatar_url"`
	Background string  `json:"background"`
}
