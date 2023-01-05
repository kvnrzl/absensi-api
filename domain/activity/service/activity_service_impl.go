package service

import (
	"context"

	"absensi-api.com/domain/activity/repository"
	"absensi-api.com/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ActivityServiceImpl struct {
	activityRepository repository.ActivityRepository
	db                 *gorm.DB
	validate           *validator.Validate
}

func NewActivityServiceImpl(activityRepository repository.ActivityRepository, db *gorm.DB, validate *validator.Validate) ActivityService {
	return &ActivityServiceImpl{
		activityRepository: activityRepository,
		db:                 db,
		validate:           validate,
	}
}

func (s *ActivityServiceImpl) Save(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	if err := s.validate.Struct(activity); err != nil {
		return nil, model.ErrInvalidJsonRequest
	}

	return s.activityRepository.Save(ctx, s.db, activity)
}

func (s *ActivityServiceImpl) FindAll(ctx context.Context) ([]*model.Activity, error) {
	return s.activityRepository.FindAll(ctx, s.db)
}

func (s *ActivityServiceImpl) FindByDate(ctx context.Context, userID uint, date string) ([]*model.Activity, error) {
	return s.activityRepository.FindByDate(ctx, s.db, userID, date)
}

func (s *ActivityServiceImpl) FindByUserID(ctx context.Context, userID uint) ([]*model.Activity, error) {
	return s.activityRepository.FindByUserID(ctx, s.db, userID)
}

func (s *ActivityServiceImpl) Update(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	if _, err := s.activityRepository.FindByID(ctx, s.db, activity.ID, activity.UserID); err != nil {
		return nil, err
	}

	if _, err := s.activityRepository.Update(ctx, s.db, activity); err != nil {
		return nil, err
	}

	return s.activityRepository.FindByID(ctx, s.db, activity.ID, activity.UserID)
}

func (s *ActivityServiceImpl) Delete(ctx context.Context, id uint, userID uint) error {
	if _, err := s.activityRepository.FindByID(ctx, s.db, id, userID); err != nil {
		return err
	}

	return s.activityRepository.Delete(ctx, s.db, id, userID)
}
