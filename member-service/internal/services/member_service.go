package services

import (
	"member-service/internal/entities"
	repository "member-service/internal/repositories"
)

type MemberService struct {
	repo *repository.MemberRepository
}

func NewMemberService(repo *repository.MemberRepository) *MemberService {
	return &MemberService{repo: repo}
}

func (s *MemberService) CreateMember(name, email string) (entities.Member, error) {
	member := &entities.Member{Name: name, Email: email}

	return s.repo.Create(member)
}
