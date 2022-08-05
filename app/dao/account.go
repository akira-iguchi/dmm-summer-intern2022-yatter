package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// CreateUser : ユーザを作成
func (r account) CreateUser(ctx context.Context, user *object.Account) (*object.Account, error) {
	if _, err := r.db.ExecContext(ctx, "insert into account (username, password_hash) value (?, ?)", user.Username, user.PasswordHash); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", user.Username).StructScan(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return user, nil
}

func (r *account) UpdateUser(ctx context.Context, accountId int64, displayName *string, note *string, avatar *string, header *string) (*object.Account, error) {
	if _, err := r.db.ExecContext(ctx, "update account set display_name = ?, note = ?, avatar = ?, header = ? where id = ?", displayName, note, avatar, header, accountId); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	user := new(object.Account)
	if err := r.db.QueryRowxContext(ctx, "select * from account where id = ?", accountId).StructScan(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}
	return user, nil
}
