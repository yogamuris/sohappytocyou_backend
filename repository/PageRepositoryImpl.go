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
	query := "select id, username, background, photo, description from page where username = ?"
	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		return entity.Page{}, err
	}

	defer rows.Close()

	page := entity.Page{}
	if rows.Next() {
		err := rows.Scan(&page.Id, &page.Username, &page.Background, &page.Photo, &page.Description)
		if err != nil {
			return page, err
		}
	} else {
		return page, errors.New("page not found")
	}

	var links []entity.Link

	query = "select id, url, visited from link where id_page = ?;"
	linkRows, err := tx.QueryContext(ctx, query, page.Id)
	if err != nil {
		return page, err
	}

	defer linkRows.Close()

	link := entity.Link{}
	for linkRows.Next() {
		err := linkRows.Scan(&link.Id, &link.Url, &link.Visited)
		if err != nil {
			return page, err
		}

		links = append(links, link)
	}

	page.Links = links

	return page, nil
}

func (p PageRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, page entity.Page) (entity.Page, error) {
	query := "insert into page(id_user, username, background, photo, description, created_at) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, query, page.IdUser, page.Username, page.Background, page.Photo, page.Description, page.CreatedAt)
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
	query := "update page set background = ?, photo = ?, description = ?, modified_at = ? where username = ?"
	_, err := tx.ExecContext(ctx, query, page.Background, page.Photo, page.Description, page.ModifiedAt, page.Username)

	if err != nil {
		return entity.Page{}, err
	}

	return page, nil
}

func (p PageRepositoryImpl) GetUsernameId(ctx context.Context, tx *sql.Tx, username string) (int, error) {
	query := "select id from user where username = ?;"
	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id)
		if err != nil {
			return -1, err
		}

		return user.Id, err
	} else {
		return -1, errors.New("user not found")
	}
}
