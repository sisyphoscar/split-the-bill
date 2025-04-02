package member

type MemberRepository interface {
	Create(member *Member) error
}
