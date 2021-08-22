package repository

import (
	"context"
	"database/sql"
	"github.com/yogamuris/sohappytocyou/entity"
)

type UserRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error)
	Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	ChangePassword(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
