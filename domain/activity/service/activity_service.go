package service

import (
	"context"

	"absensi-api.com/model"
)

type ActivityService interface {
	Save(ctx context.Context, user *model.Activity) (*model.Activity, error)
	FindAll(ctx context.Context) ([]*model.Activity, error)
	FindByID(ctx context.Context, id uint) (*model.Activity, error)
	Update(ctx context.Context, activity *model.Activity) (*model.Activity, error)
	Delete(ctx context.Context, id uint) error
}
