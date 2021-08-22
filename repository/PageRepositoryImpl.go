package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yogamuris/sohappytocyou/entity"
)

type PageRepositoryImpl struct {

}

func NewPageRepository() PageRepository {
	return &PageRepositoryImpl{}
}

func (p PageRepositoryImpl) Show(ctx context.Context, tx *sql.Tx, username string) (entity.Page, error) {
	query := "select id, background, photo, description from page where username = ?"
	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		return entity.Page{}, err
	}

	defer rows.Close()

	page := entity.Page{}
	if rows.Next() {
		err := rows.Scan(&page.Id, &page.Background, &page.Photo, &page.Description)
		if err != nil {
			return page, err
		}

		return page, nil
	} else {
		return page, errors.New("page not found")
	}
}

func (p PageRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, page entity.Page) (entity.Page, error) {
	var idUser int
	query := "insert into page(id_user, username, background, photo, description, created_at) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, query, idUser, page.Username, page.Background, page.Photo, page.Description)
	if err != nil {
		return entity.Page{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.Page{}, err
	}

	page.Id = int(id)
	return page, nil
}

func (p PageRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, page entity.Page) (entity.Page, error) {
	query := "update page set background = ?, photo = ?, description = ? where username = ?"
	_, err := tx.ExecContext(ctx, query, page.Background, page.Photo, page.Description, page.Username)

	if err != nil {
		return entity.Page{}, err
	}

	return page, nil
}



