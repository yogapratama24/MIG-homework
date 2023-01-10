package service

import (
	"homework_mitramas/model"
	"homework_mitramas/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type ClientService interface {
	GetClient() (user []model.UserResponse, err error)
	CreateClient(request model.UserCreateRequest) error
	UpdateClient(request model.UserUpdateRequest) error
	DeleteClient(Id int) error
}

type clientService struct {
	clientRepository repository.ClientRepository
}

func NewClientService(repository repository.ClientRepository) *clientService {
	return &clientService{repository}
}

func (s *clientService) GetClient() (user []model.UserResponse, err error) {
	data, err := s.clientRepository.GetClient()
	if err != nil {
		return nil, err
	}

	for i := range data {
		var userRow model.UserResponse

		userRow.Id = data[i].Id
		userRow.UserName = data[i].UserName
		userRow.Email = data[i].Email

		user = append(user, userRow)
	}

	return user, nil
}

func (s *clientService) CreateClient(request model.UserCreateRequest) error {
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 4)
	if err != nil {
		log.Printf("Error hashing password with err: %s", err)
		return err
	}
	request.Password = string(password)

	clientCreate := model.User{
		UserName: request.UserName,
		Password: request.Password,
		Email:    request.Email,
		RoleId:   2,
	}
	if err := s.clientRepository.CreateClient(clientCreate); err != nil {
		return err
	}

	return nil
}

func (s *clientService) UpdateClient(request model.UserUpdateRequest) error {
	clientCreate := model.User{
		Id:       request.Id,
		UserName: request.UserName,
		Email:    request.Email,
		RoleId:   2,
	}
	if err := s.clientRepository.UpdateClient(clientCreate); err != nil {
		return err
	}

	return nil
}

func (s *clientService) DeleteClient(Id int) error {
	if err := s.clientRepository.DeleteClient(Id); err != nil {
		return err
	}

	return nil
}
