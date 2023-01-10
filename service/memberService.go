package service

import (
	"homework_mitramas/model"
	"homework_mitramas/repository"
)

type MemberService interface {
	GetMember(userID int) (member []model.MemberResponse, err error)
	CreateMember(request model.MemberCreateRequest) error
	UpdateMember(request model.MemberUpdateRequest) error
	DeleteMember(Id int) error
}

type memberService struct {
	memberRepository repository.MemberRepository
}

func NewMemberService(repository repository.MemberRepository) *memberService {
	return &memberService{repository}
}

func (s *memberService) GetMember(userID int) (member []model.MemberResponse, err error) {
	data, err := s.memberRepository.GetMember(userID)
	if err != nil {
		return nil, err
	}

	for i := range data {
		var memberRow model.MemberResponse

		memberRow.Id = data[i].Id
		memberRow.MemberName = data[i].MemberName
		memberRow.PhoneNumber = data[i].PhoneNumber
		memberRow.Address = data[i].Address
		memberRow.UserId = data[i].UserId

		member = append(member, memberRow)
	}

	return member, nil
}

func (s *memberService) CreateMember(request model.MemberCreateRequest) error {
	memberCreate := model.Member{
		MemberName:  request.MemberName,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		UserId:      request.UserId,
	}
	if err := s.memberRepository.CreateMember(memberCreate); err != nil {
		return err
	}

	return nil
}

func (s *memberService) UpdateMember(request model.MemberUpdateRequest) error {
	memberUpdate := model.Member{
		Id:          request.Id,
		MemberName:  request.MemberName,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}
	if err := s.memberRepository.UpdateMember(memberUpdate); err != nil {
		return err
	}

	return nil
}

func (s *memberService) DeleteMember(Id int) error {
	if err := s.memberRepository.DeleteMember(Id); err != nil {
		return err
	}

	return nil
}
