//go:generate mockgen -source=user.go -destination=./mock/user_mock.go -package=mock
package repository

import (
	"context"

	"github.com/girakdev/girack-backend/application/model"
)

type UserRepositry interface {
	UserGetter
	UserLister
	UserCreator
	UserDeleter
}

type (
	UserGetter interface {
		GetUser(ctx context.Context, input *GetUserInput) (output *GetUserOutput, err error)
	}
	GetUserInput struct {
		ID model.ID
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
		ID model.ID
	}
	DeleteUserOutput struct {
	}
)
