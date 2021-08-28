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

type LinkServiceImpl struct {
	LinkRepository repository.LinkRepository
	Db             *sql.DB
	Validate       *validator.Validate
}

func NewLinkService(repository repository.LinkRepository, DB *sql.DB, validate *validator.Validate) LinkService {
	return &LinkServiceImpl{
		LinkRepository: repository,
		Db:             DB,
		Validate:       validate,
	}
}

func (service LinkServiceImpl) List(ctx context.Context, username string) (web.LinkListResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return web.LinkListResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	links, err := service.LinkRepository.List(ctx, tx, username)
	if err != nil {
		return web.LinkListResponse{}, err
	}

	var linkResponses []web.LinkResponse

	for _, link := range links {
		linkResponses = append(linkResponses, web.LinkResponse{
			Id:      link.Id,
			IdPage:  link.IdPage,
			Url:     link.Url,
			Visited: link.Visited,
		})
	}

	return web.LinkListResponse{Links: linkResponses}, nil
}

func (service LinkServiceImpl) Show(ctx context.Context, id int) (web.LinkResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return web.LinkResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	link, err := service.LinkRepository.Show(ctx, tx, id)
	if err != nil {
		return web.LinkResponse{}, err
	}

	return web.LinkResponse{
		Id:      link.Id,
		IdPage:  link.IdPage,
		Url:     link.Url,
		Visited: link.Visited,
	}, nil
}

func (service LinkServiceImpl) Save(ctx context.Context, request web.LinkSaveRequest) (web.LinkResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return web.LinkResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	link := entity.Link{
		IdPage:    request.IdPage,
		Url:       request.Url,
		Visited:   0,
		CreatedAt: time.Now(),
	}

	link, err = service.LinkRepository.Save(ctx, tx, link)
	if err != nil {
		return web.LinkResponse{}, err
	}

	return web.LinkResponse{
		Id:      link.Id,
		IdPage:  link.IdPage,
		Url:     link.Url,
		Visited: link.Visited,
	}, nil
}

func (service LinkServiceImpl) Update(ctx context.Context, request web.LinkUpdateRequest) (web.LinkResponse, error) {
	panic("implement me")
}

func (service LinkServiceImpl) Delete(ctx context.Context, request web.LinkDeleteRequest) bool {
	tx, err := service.Db.Begin()
	if err != nil {
		return false
	}
	defer helper.CommitOrRollback(tx)

	ok := service.LinkRepository.Delete(ctx, tx, request.Id)
	if !ok {
		return false
	}

	return true
}
