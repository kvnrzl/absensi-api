package repository

import (
	"context"

	"absensi-api.com/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, db *gorm.DB, user *model.User) (*model.User, error)
	FindByUsername(ctx context.Context, db *gorm.DB, username string) (*model.User, error)
}
