package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/yogamuris/sohappytocyou/entity"
	"github.com/yogamuris/sohappytocyou/repository"
)

type PageServiceImpl struct {
	PageRepository repository.PageRepository
	Db	*sql.DB
	Validate *validator.Validate
}

func NewPageService(pageRepository repository.PageRepository, db *sql.DB, validate *validator.Validate) PageService{
	return &PageServiceImpl{
		PageRepository: pageRepository,
		Db:             db,
		Validate:       validate,
	}
}

func (p PageServiceImpl) Show(ctx context.Context, tx *sql.Tx, username string) (entity.Page, error) {
	panic("implement me")
}

func (p PageServiceImpl) Save(ctx context.Context, tx *sql.Tx, page entity.Page) (entity.Page, error) {
	panic("implement me")
}

func (p PageServiceImpl) Update(ctx context.Context, tx *sql.Tx, page entity.Page) (entity.Page, error) {
	panic("implement me")
}


