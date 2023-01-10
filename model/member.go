package model

import "time"

type Member struct {
	Id          int       `json:"id"`
	MemberName  string    `json:"member_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	UserId      int       `json:"user_id"`
	User        *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MemberCreateRequest struct {
	MemberName  string `json:"member_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
	UserId      int    `json:"user_id" validate:"required"`
}

type MemberResponse struct {
	Id          int    `json:"id"`
	MemberName  string `json:"member_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
}

type MemberUpdateRequest struct {
	Id          int    `json:"id" validate:"required"`
	MemberName  string `json:"member_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
}
