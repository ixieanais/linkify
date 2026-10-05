package dto

type UserResponse struct {
	BaseModel
	Email string `json:"email"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
