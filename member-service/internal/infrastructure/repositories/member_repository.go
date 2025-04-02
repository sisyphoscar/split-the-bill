package repositories

import (
	"database/sql"
	"member-service/internal/domain/member"
)

type MemberRepository struct {
	db *sql.DB
}

// TODO: query timeout

func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

func (r *MemberRepository) Create(member *member.Member) error {
	query := "INSERT INTO members (name, email, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, member.Name, member.Email, member.CreatedAt, member.UpdatedAt).Scan(&member.ID)

	return err
}
