package service

import (
	"context"

	"absensi-api.com/model"
)

type UserService interface {
	Register(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, user *model.User) (string, error)
}
