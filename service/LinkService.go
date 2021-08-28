package service

import (
	"context"
	"github.com/yogamuris/sohappytocyou/entity/web"
)

type LinkService interface {
	List(ctx context.Context, username string) (web.LinkListResponse, error)
	Show(ctx context.Context, id int) (web.LinkResponse, error)
	Save(ctx context.Context, request web.LinkSaveRequest) (web.LinkResponse, error)
	Update(ctx context.Context, request web.LinkUpdateRequest) (web.LinkResponse, error)
	Delete(ctx context.Context, request web.LinkDeleteRequest) bool
}
