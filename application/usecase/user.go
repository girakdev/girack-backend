//go:generate mockgen -source=user.go -destination=./mock/user_mock.go -package=mock
package usecase

import (
	"context"

	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/domain/repository"
)

type userUsecase struct {
	userRepository repository.UserRepositry
}

func NewUserUsecase(userRepository repository.UserRepositry) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
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
	guOut, err := u.userRepository.GetUser(ctx, &repository.GetUserInput{
		ID: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{
		User: guOut.User,
	}, nil
}

func (u *userUsecase) ListUser(ctx context.Context, input *ListUserInput) (output *ListUserOutput, err error) {
	luOut, err := u.userRepository.ListUser(ctx, &repository.ListUserInput{})
	if err != nil {
		return nil, err
	}

	return &ListUserOutput{
		Users: luOut.Users,
	}, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, input *CreateUserInput) (output *CreateUserOutput, err error) {
	cuOut, err := u.userRepository.CreateUser(ctx, &repository.CreateUserInput{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		User: cuOut.User,
	}, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, input *DeleteUserInput) (output *DeleteUserOutput, err error) {
	_, err = u.userRepository.DeleteUser(ctx, &repository.DeleteUserInput{
		ID: input.ID,
	})
	if err != nil {

		return nil, err
	}

	return &DeleteUserOutput{}, nil
}
