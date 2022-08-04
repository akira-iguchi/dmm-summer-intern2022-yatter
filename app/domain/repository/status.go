package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error)
	FindByStatusId(ctx context.Context, statusId int64) (*object.Status, error)
}
