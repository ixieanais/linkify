package dto

type ProfileResponse struct {
	BaseModel
	Username  string  `json:"username"`
	URL       string  `json:"url"`
	Bio       *string `json:"bio"`
	AvatarURL *string `json:"avatar_url"`
}

type ProfileRequest struct {
	Username  string  `json:"username"`
	URL       string  `json:"url"`
	Bio       *string `json:"bio"`
	AvatarURL *string `json:"avatar_url"`
}
