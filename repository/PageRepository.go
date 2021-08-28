package repository

import (
	"context"
	"database/sql"
	"github.com/yogamuris/sohappytocyou/entity"
)

type PageRepository interface {
	Show(ctx context.Context, tx *sql.Tx, username string) (entity.Page, error)
	Save(ctx context.Context, tx *sql.Tx, page entity.Page) (entity.Page, error)
	Update(ctx context.Context, tx *sql.Tx, page entity.Page) (entity.Page, error)
	GetUsernameId(ctx context.Context, tx *sql.Tx, username string) (int, error)
}
