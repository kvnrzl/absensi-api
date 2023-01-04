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
		return nil, err
	}

	return s.activityRepository.Save(ctx, s.db, activity)
}

func (s *ActivityServiceImpl) FindAll(ctx context.Context) ([]*model.Activity, error) {
	return s.activityRepository.FindAll(ctx, s.db)
}

func (s *ActivityServiceImpl) FindByID(ctx context.Context, id uint) (*model.Activity, error) {
	return s.activityRepository.FindByID(ctx, s.db, id)
}

func (s *ActivityServiceImpl) Update(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	if _, err := s.FindByID(ctx, activity.ID); err != nil {
		return nil, err
	}

	if _, err := s.activityRepository.Update(ctx, s.db, activity); err != nil {
		return nil, err
	}

	return s.FindByID(ctx, activity.ID)
}

func (s *ActivityServiceImpl) Delete(ctx context.Context, id uint) error {
	if _, err := s.FindByID(ctx, id); err != nil {
		return err
	}

	return s.activityRepository.Delete(ctx, s.db, id)
}
