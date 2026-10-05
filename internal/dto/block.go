package dto

type BlockResponse struct {
	BaseModel
	Title    string  `json:"title"`
	Text     *string `json:"text"`
	Image    *string `json:"image"`
	URL      *string `json:"url"`
	Type     string  `json:"type"`
	IsActive *bool   `json:"is_active"`
}

type BlockRequest struct {
	Title    string  `json:"title"`
	Text     *string `json:"text"`
	Image    *string `json:"image"`
	URL      *string `json:"url"`
	Type     string  `json:"type"`
	IsActive *bool   `json:"is_active"`
}
