package service

import (
	"homework_mitramas/model"
	"homework_mitramas/repository"
)

type AuthService interface {
	Login(request model.UserLogin) (model.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository}
}

func (s *authService) Login(request model.UserLogin) (data model.User, err error) {
	adminLogin := model.User{
		Email:    request.Email,
		Password: request.Password,
	}
	data, err = s.authRepository.Login(adminLogin)
	if err != nil {
		return data, err
	}

	return data, nil
}
