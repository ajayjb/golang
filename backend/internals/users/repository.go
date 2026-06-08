package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (
		name,
		email
	)
	VALUES ($1, $2)
	RETURNING id
`

	return r.db.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
	).Scan(&user.ID)
}
