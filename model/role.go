package model

type Role struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name" gorm:"unique"`
}
