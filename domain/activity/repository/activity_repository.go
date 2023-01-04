package repository

import (
	"context"

	"absensi-api.com/model"
	"gorm.io/gorm"
)

type ActivityRepository interface {
	Save(ctx context.Context, db *gorm.DB, user *model.Activity) (*model.Activity, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]*model.Activity, error)
	FindByID(ctx context.Context, db *gorm.DB, id uint) (*model.Activity, error)
	Update(ctx context.Context, db *gorm.DB, user *model.Activity) (*model.Activity, error)
	Delete(ctx context.Context, db *gorm.DB, id uint) error
}
