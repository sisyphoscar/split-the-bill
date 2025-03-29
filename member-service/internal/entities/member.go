package entities

import "time"

type Member struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type MemberRepository interface {
	Create(member *Member) error
}
