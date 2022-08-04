package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"

	"log"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &status{db: db}
}

func (r *status) AllStatuses(ctx context.Context) (*object.Timelines, error) {
	timelines := new(object.Timelines)
	log.Printf("Serve on huuuuuuuuuuuuuuuuuuuuuuttp://%s", timelines)
	rows, err := r.db.QueryxContext(ctx, "select s.id, s.content, s.create_at, a.username as 'account.username', a.create_at as 'account.create_at' from status as s inner join account as a on s.account_id = a.id")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	var status object.Status
	var statuses []object.Status
	for rows.Next() {
		rows.Scan(&status)
		statuses = append(statuses, status)
	}

	timelines = statuses

	return timelines, nil
}
