package member

type MemberService struct {
	repo MemberRepository
}

func NewMemberService(repo MemberRepository) *MemberService {
	return &MemberService{repo: repo}
}

func (s *MemberService) Create(name string, email string) (*Member, error) {
	member, err := newMember(name, email)
	if err != nil {
		return nil, err
	}
	err = s.repo.Create(member)
	if err != nil {
		return nil, err
	}
	return member, nil
}
