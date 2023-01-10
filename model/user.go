package model

import "time"

type User struct {
	Id        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email" gorm:"unique"`
	RoleId    int       `json:"role_id"`
	Role      *Role     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type UserCreateRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required"`
}

type UserUpdateRequest struct {
	Id       int    `json:"id" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required"`
}
