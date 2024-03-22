package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wildanfaz/cinema-ticket/internal/models"
)

type ImplementUsers struct {
	db *pgxpool.Conn
}

type Users interface {
	Register(ctx context.Context, payload *models.User) error
	Profile(ctx context.Context, email string) (*models.User, error)
}

func NewUsersRepository(db *pgxpool.Conn) Users {
	return &ImplementUsers{
		db: db,
	}
}

func (r *ImplementUsers) Register(ctx context.Context, payload *models.User) error {
	q := `
	INSERT INTO users
	(full_name, email, password)
	VALUES
	($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, q, payload.FullName, payload.Email, payload.Password)

	return err
}

func (r *ImplementUsers) Profile(ctx context.Context, email string) (*models.User, error) {
	var user = new(models.User)

	q := `
	SELECT
		id, full_name, email, password, role, balance, created_at, updated_at
	FROM
		users
	WHERE
		email = $1
	`

	err := r.db.QueryRow(ctx, q, email).Scan(
		&user.Id, &user.FullName, &user.Email,
		&user.Password, &user.Role, &user.Balance,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}

	user.ToLocal()

	return user, err
}
