package repository

import (
	"homework_mitramas/model"
	"log"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(request model.User) (data model.User, err error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) Login(request model.User) (data model.User, err error) {
	db := r.db
	if err = db.Where("email = ?", request.Email).First(&data).Error; err != nil {
		log.Printf("Error get data user login with err: %s", err)
		return data, err
	}
	return data, nil
}
