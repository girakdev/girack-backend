//go:generate mockgen -source=user.go -destination=./mock/user_mock.go -package=mock
package usecase

import (
	"context"

	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/internal/pulid"
)

type userUsecase struct {
}

func NewUserUsecase() *userUsecase {
	return &userUsecase{}
}

type UserUsecase interface {
	UserGetter
	UserLister
	UserDeleter
	UserDeleter
}

var _ UserGetter = (*userUsecase)(nil)
var _ UserLister = (*userUsecase)(nil)
var _ UserDeleter = (*userUsecase)(nil)
var _ UserDeleter = (*userUsecase)(nil)

type (
	UserGetter interface {
		GetUser(ctx context.Context, input *GetUserInput) (output *GetUserOutput, err error)
	}
	GetUserInput struct {
		ID pulid.ID
	}
	GetUserOutput struct {
		User *model.User
	}
)

type (
	UserLister interface {
		ListUser(ctx context.Context, input *ListUserInput) (output *ListUserOutput, err error)
	}
	ListUserInput struct {
	}

	ListUserOutput struct {
		Users []*model.User
	}
)

type (
	UserCreator interface {
		CreateUser(ctx context.Context, input *CreateUserInput) (output *CreateUserOutput, err error)
	}
	CreateUserInput struct {
		Name string
	}
	CreateUserOutput struct {
		User *model.User
	}
)

type (
	UserDeleter interface {
		DeleteUser(ctx context.Context, input *DeleteUserInput) (output *DeleteUserOutput, err error)
	}
	DeleteUserInput struct {
		ID pulid.ID
	}
	DeleteUserOutput struct {
	}
)

type (
	UserUpdater interface {
		UpdateUser(ctx context.Context, input *UpdateUserInput) (output *UpdateUserOutput, err error)
	}
	UpdateUserInput struct {
		User *model.User
	}
	UpdateUserOutput struct {
		User *model.User
	}
)

func (u *userUsecase) GetUser(ctx context.Context, input *GetUserInput) (output *GetUserOutput, err error) {
	// TODO: Implement
	return nil, nil
}

func (u *userUsecase) ListUser(ctx context.Context, input *ListUserInput) (output *ListUserOutput, err error) {
	// TODO: Implement
	return nil, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, input *CreateUserInput) (output *CreateUserOutput, err error) {
	// TODO: Implement
	return nil, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, input *DeleteUserInput) (output *DeleteUserOutput, err error) {
	// TODO: Implement
	return nil, nil
}
