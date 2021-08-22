package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yogamuris/sohappytocyou/entity"
)

type UserRepositoryImpl struct {

}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.User, error) {
	query := "select id, username, email, created_at, verified_at from user where id = ?;"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.VerifiedAt)
		if err != nil {
			panic(err)
		}

		return user, nil
	} else {
		return user, errors.New("User not found")
	}
}

func (u UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error) {
	query := "select id, username, email, created_at, verified_at from user where username = ?;"
	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.VerifiedAt)
		if err != nil {
			return user, err
		}

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	query := "insert into user(username, email, password, created_at) values(?, ?, ?, ?);"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.CreatedAt)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.User{}
	}

	user.Id = int(id)
	return user
}

func (u UserRepositoryImpl) ChangePassword(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	query := "update user set password = ? where username = ?;"
	_, err := tx.ExecContext(ctx, query, user.Password, user.Username)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "delete from user where id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

