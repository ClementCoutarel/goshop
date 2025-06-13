package models

type Role int64

const (
	Guest Role = iota
	Admin
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type UserCreateDTO struct {
	Name     string `json:"name" validate:"required,max=50,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,alphanum"`
	Role     Role   `json:"role"`
}

type UserGetDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
