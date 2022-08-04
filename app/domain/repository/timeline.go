package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	AllStatuses(ctx context.Context) (*object.Timelines, error)
}
