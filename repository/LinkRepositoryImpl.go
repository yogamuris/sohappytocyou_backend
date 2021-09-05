package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yogamuris/sohappytocyou/entity"
)

type LinkRepositoryImpl struct {
}

func NewLinkRepository() LinkRepository {
	return &LinkRepositoryImpl{}
}

func (l LinkRepositoryImpl) Show(ctx context.Context, tx *sql.Tx, id int) (entity.Link, error) {
	query := "select id, id_page, url, visited from link where id = ?;"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return entity.Link{}, err
	}

	var link entity.Link
	if rows.Next() {
		err := rows.Scan(&link.Id, &link.IdPage, &link.Url, &link.Visited)
		if err != nil {
			return entity.Link{}, err
		}

		return link, nil
	} else {
		return entity.Link{}, errors.New("link not found")
	}
}

func (l LinkRepositoryImpl) List(ctx context.Context, db *sql.DB, username string) ([]entity.Link, error) {
	query := "select id from page where username = ?;"
	rows, err := db.QueryContext(ctx, query, username)

	var page entity.Page
	if rows.Next() {
		rows.Scan(&page.Id)
	}

	defer rows.Close()

	query = "select id, id_page, url, visited from link where id_page = ?;"
	rows, err = db.QueryContext(ctx, query, page.Id)
	if err != nil {
		return nil, err
	}

	var links []entity.Link

	defer rows.Close()

	for rows.Next() {
		var link entity.Link
		err := rows.Scan(&link.Id, &link.IdPage, &link.Url, &link.Visited)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}

func (l LinkRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, link entity.Link) (entity.Link, error) {
	query := "insert into link(id_page, url, visited, created_at) values (?,?,?,?);"
	result, err := tx.ExecContext(ctx, query, link.IdPage, link.Url, link.Visited, link.CreatedAt)
	if err != nil {
		return entity.Link{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.Link{}, err
	}

	link.Id = int(id)

	return link, nil
}

func (l LinkRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, link entity.Link) (entity.Link, error) {
	panic("implement me")
}

func (l LinkRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) bool {
	query := "delete from link where id = ?;"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return false
	}

	return true
}
