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
	status struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) FindByStatusId(ctx context.Context, statusId int64) (*object.Status, error) {
	status := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select s.id, s.content, s.create_at, a.username as 'account.username', a.create_at as 'account.create_at' from status as s inner join account as a on s.account_id = a.id where s.id = ?", statusId).StructScan(status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return status, nil
}

// CreateStatus : Statusを作成
func (r status) CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error) {
	result, err := r.db.ExecContext(ctx, "insert into status (account_id, content) value (?, ?)", status.AccountID, status.Content)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	status_id, err := result.LastInsertId()
	if err != nil {
		return nil, nil
	}

	if err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", status_id).StructScan(status); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return status, nil
}
