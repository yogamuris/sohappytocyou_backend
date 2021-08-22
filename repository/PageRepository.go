package repository

import "github.com/yogamuris/sohappytocyou/entity"

type PageRepository interface {
	FindByUsername(username string) (*entity.Link, error)
	Save(page entity.Page) error
	AddLink(link entity.Link) error
	Update(page entity.Page) error
	Delete(page entity.Page) error
}
