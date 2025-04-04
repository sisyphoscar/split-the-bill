package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"member-service/internal/domain/member"
	"time"
)

type MemberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

func (r *MemberRepository) Create(member *member.Member) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO members (name, email, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, member.Name, member.Email, member.CreatedAt, member.UpdatedAt).Scan(&member.ID)

	if err != nil {
		return fmt.Errorf("failed to create member: %w", err)
	}

	return nil
}
