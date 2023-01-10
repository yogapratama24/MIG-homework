package repository

import (
	"homework_mitramas/model"
	"log"

	"gorm.io/gorm"
)

type MemberRepository interface {
	GetMember(userId int) (member []model.Member, err error)
	CreateMember(request model.Member) error
	UpdateMember(request model.Member) error
	DeleteMember(Id int) error
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) GetMember(userId int) (member []model.Member, err error) {
	db := r.db
	if err := db.Table("members").
		Where("members.user_id = ?", userId).
		Select("members.id as id, members.member_name, members.address, members.phone_number, members.user_id").
		Scan(&member).Error; err != nil {
		log.Printf("Error get data member with err: %s", err)
		return nil, err
	}
	return member, nil
}

func (r *memberRepository) CreateMember(member model.Member) error {
	db := r.db
	if err := db.Create(&member); err.Error != nil {
		log.Printf("Error create data member with err: %s\n", err.Error)
		return err.Error
	}
	return nil
}

func (r *memberRepository) UpdateMember(member model.Member) error {
	db := r.db
	err := db.Model(member).Updates(member)
	if err.Error != nil {
		log.Printf("Error update data member with err: %v", err)
		return err.Error
	}

	return nil
}

func (r *memberRepository) DeleteMember(Id int) error {
	db := r.db
	err := db.Delete(model.Member{}, Id)
	if err.Error != nil {
		log.Printf("Error delete data member with err: %v", err)
		return err.Error
	}

	return nil
}
