package repository

import (
	"context"

	"absensi-api.com/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	// FindAll() ([]*model.User, error)
	Save(ctx context.Context, db *gorm.DB, user *model.User) (*model.User, error)
	FindByUsername(ctx context.Context, db *gorm.DB, username string) (*model.User, error)
	FindByID(ctx context.Context, db *gorm.DB, id uint) (*model.User, error)
	// Update(user *model.User) (*model.User, error)
	// Delete(user *model.User) error
}
