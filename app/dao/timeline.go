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
	timeline struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) PublicTimelines(ctx context.Context, maxId int, sinceId int, limit int) ([]*object.Status, error) {
	rows, err := r.db.QueryxContext(ctx, "select s.id, s.content, s.create_at, a.username as 'account.username', a.create_at as 'account.create_at' from status as s inner join account as a on s.account_id = a.id where s.id < ? and s.id > ? limit ?", maxId, sinceId, limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}
	var timelines []*object.Status
	for rows.Next() {
		status := new(object.Status)
		err = rows.StructScan(&status)
		if err != nil {
			return nil, err
		}
		timelines = append(timelines, status)

	}
	return timelines, nil
}
