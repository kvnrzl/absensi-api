package mysql

import (
	"context"

	"absensi-api.com/domain/activity/repository"
	"absensi-api.com/model"
	"gorm.io/gorm"
)

type ActivityRepositoryImpl struct{}

func NewActivityRepositoryImpl() repository.ActivityRepository {
	return &ActivityRepositoryImpl{}
}

func (r *ActivityRepositoryImpl) Save(ctx context.Context, db *gorm.DB, activity *model.Activity) (*model.Activity, error) {
	err := db.WithContext(ctx).Create(activity).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (r *ActivityRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]*model.Activity, error) {
	var activities []*model.Activity
	err := db.WithContext(ctx).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *ActivityRepositoryImpl) FindByDate(ctx context.Context, db *gorm.DB, userID uint, date string) ([]*model.Activity, error) {
	var activities []*model.Activity
	err := db.WithContext(ctx).Where("user_id = ? AND DATE(start_time) = ?", userID, date).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *ActivityRepositoryImpl) FindByUserID(ctx context.Context, db *gorm.DB, userID uint) ([]*model.Activity, error) {
	var activities []*model.Activity
	err := db.WithContext(ctx).Where("user_id = ?", userID).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *ActivityRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uint, userID uint) (*model.Activity, error) {
	var activity model.Activity
	err := db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&activity).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *ActivityRepositoryImpl) Update(ctx context.Context, db *gorm.DB, activity *model.Activity) (*model.Activity, error) {
	err := db.WithContext(ctx).Where("id = ?", activity.ID).Updates(&activity).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (r *ActivityRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uint, userID uint) error {
	err := db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&model.Activity{}).Error
	if err != nil {
		return err
	}
	return nil
}
