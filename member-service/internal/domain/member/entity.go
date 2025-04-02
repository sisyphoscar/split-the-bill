package member

import (
	"errors"
	"time"
)

type Member struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func newMember(name string, email string) (*Member, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	now := time.Now()

	return &Member{
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}, nil
}
