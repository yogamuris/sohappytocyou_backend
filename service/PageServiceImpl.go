package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/yogamuris/sohappytocyou/entity"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/helper"
	"github.com/yogamuris/sohappytocyou/repository"
	"time"
)

type PageServiceImpl struct {
	PageRepository repository.PageRepository
	Db             *sql.DB
	Validate       *validator.Validate
}

func NewPageService(pageRepository repository.PageRepository, db *sql.DB, validate *validator.Validate) PageService {
	return &PageServiceImpl{
		PageRepository: pageRepository,
		Db:             db,
		Validate:       validate,
	}
}

func (service PageServiceImpl) Show(ctx context.Context, username string) (web.PageResponse, error) {
	page, err := service.PageRepository.Show(ctx, service.Db, username)
	if err != nil {
		return web.PageResponse{}, err
	}

	return web.PageResponse{
		Id:          page.Id,
		Username:    page.Username,
		Background:  page.Background,
		Photo:       page.Photo,
		Description: page.Description,
		Links:       page.Links,
	}, nil
}

func (service PageServiceImpl) Save(ctx context.Context, request web.PageSaveRequest) (web.PageResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return web.PageResponse{}, err
	}

	defer helper.CommitOrRollback(tx)

	userId, err := service.PageRepository.GetUsernameId(ctx, tx, request.Username)
	if err != nil {
		return web.PageResponse{}, err
	}

	page := entity.Page{
		IdUser:      userId,
		Username:    request.Username,
		Background:  request.Background,
		Photo:       request.Photo,
		Description: request.Description,
		CreatedAt:   time.Now(),
	}

	page, err = service.PageRepository.Save(ctx, tx, page)
	if err != nil {
		return web.PageResponse{}, err
	}

	return web.PageResponse{
		Id:          page.Id,
		Username:    page.Username,
		Background:  page.Background,
		Photo:       page.Photo,
		Description: page.Description,
		Links:       nil,
	}, nil
}

func (service PageServiceImpl) Update(ctx context.Context, request web.PageUpdateRequest) (web.PageResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return web.PageResponse{}, err
	}

	defer helper.CommitOrRollback(tx)

	page := entity.Page{
		Username:    request.Username,
		Background:  request.Background,
		Photo:       request.Photo,
		Description: request.Description,
		ModifiedAt:  time.Now(),
	}

	page, err = service.PageRepository.Update(ctx, tx, page)
	if err != nil {
		return web.PageResponse{}, err
	}

	return web.PageResponse{
		Id:          page.Id,
		Username:    page.Username,
		Background:  page.Background,
		Photo:       page.Photo,
		Description: page.Description,
		Links:       nil,
	}, nil
}
