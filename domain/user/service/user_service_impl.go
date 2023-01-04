package service

import (
	"context"
	"time"

	"absensi-api.com/domain/user/repository"
	"absensi-api.com/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	db             *gorm.DB
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		db:             db,
		validate:       validate,
	}
}

func (u *UserServiceImpl) Register(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.validate.Struct(user); err != nil {
		return nil, err
	}

	exist, _ := u.userRepository.FindByUsername(ctx, u.db, user.Username)
	if exist != nil {
		return nil, model.ErrUsernameAlreadyExist
	}

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(bytePassword)

	return u.userRepository.Save(ctx, u.db, user)
}

func (u *UserServiceImpl) Login(ctx context.Context, user *model.User) (string, error) {
	if err := u.validate.Struct(user); err != nil {
		return "", err
	}

	userDB, err := u.userRepository.FindByUsername(ctx, u.db, user.Username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    "absensi-api.com",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserServiceImpl) Logout(ctx context.Context) error {
	return nil
}
