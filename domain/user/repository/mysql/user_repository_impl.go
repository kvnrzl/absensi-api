package mysql

import (
	"context"

	"absensi-api.com/domain/user/repository"
	"absensi-api.com/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByUsername(ctx context.Context, db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	if err := db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uint) (*model.User, error) {
	var user model.User
	if err := db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
