package models

type User struct {
	User_ID    string  `json:"user_id"`
	Name       *string `json:"name" validate:"required,min=2,max=30"`
	Email      *string `json:"email" validate:"required,email"`
	Password   *string `json:"password" validate:"required,min=6"`
}
