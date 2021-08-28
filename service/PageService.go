package service

import (
	"context"
	"github.com/yogamuris/sohappytocyou/entity/web"
)

type PageService interface {
	Show(ctx context.Context, request web.PageRequest) (web.PageResponse, error)
	Save(ctx context.Context, request web.PageSaveRequest) (web.PageResponse, error)
	Update(ctx context.Context, request web.PageUpdateRequest) (web.PageResponse, error)
}
