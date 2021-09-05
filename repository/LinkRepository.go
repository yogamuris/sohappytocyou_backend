package repository

import (
	"context"
	"database/sql"
	"github.com/yogamuris/sohappytocyou/entity"
)

type LinkRepository interface {
	Show(ctx context.Context, tx *sql.Tx, id int) (entity.Link, error)
	List(ctx context.Context, db *sql.DB, username string) ([]entity.Link, error)
	Save(ctx context.Context, tx *sql.Tx, link entity.Link) (entity.Link, error)
	Update(ctx context.Context, tx *sql.Tx, link entity.Link) (entity.Link, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) bool
}
