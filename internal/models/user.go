package models

type User struct {
	BaseModel
	Email    string
	Password string
}
