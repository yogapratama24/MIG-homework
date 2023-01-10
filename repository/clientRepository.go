package repository

import (
	"homework_mitramas/model"
	"log"

	"gorm.io/gorm"
)

type ClientRepository interface {
	GetClient() (user []model.User, err error)
	CreateClient(request model.User) error
	UpdateClient(request model.User) error
	DeleteClient(Id int) error
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *clientRepository {
	return &clientRepository{db}
}

func (r *clientRepository) GetClient() (user []model.User, err error) {
	db := r.db
	if err := db.Table("users").
		Where("users.role_id = ?", 2).
		Select("users.id as id, users.user_name as user_name, users.email as email").
		Scan(&user).Error; err != nil {
		log.Printf("Error get data user with err: %s", err)
		return nil, err
	}
	return user, nil
}

func (r *clientRepository) CreateClient(user model.User) error {
	db := r.db
	if err := db.Create(&user); err.Error != nil {
		log.Printf("Error create data client with err: %s\n", err.Error)
		return err.Error
	}
	return nil
}

func (r *clientRepository) UpdateClient(user model.User) error {
	db := r.db
	err := db.Model(user).Updates(user)
	if err.Error != nil {
		log.Printf("Error update data client with err: %v", err)
		return err.Error
	}

	return nil
}

func (r *clientRepository) DeleteClient(Id int) error {
	db := r.db
	err := db.Delete(model.User{}, Id)
	if err.Error != nil {
		log.Printf("Error delete data client with err: %v", err)
		return err.Error
	}

	return nil
}
