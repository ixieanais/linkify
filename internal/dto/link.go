package dto

type LinkResponse struct {
	BaseModel
	URL      string `json:"url"`
	Image    string `json:"image"`
	IsActive *bool  `json:"is_active"`
}

type LinkRequest struct {
	URL      string `json:"url"`
	Image    string `json:"image"`
	IsActive *bool  `json:"is_active"`
}
