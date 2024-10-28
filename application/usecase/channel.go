//go:generate mockgen -source=channel.go -destination=./mock/channel_mock.go -package=mock
package usecase

import (
	"context"
	"errors"

	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/domain/repository"
	"github.com/girakdev/girack-backend/internal/pulid"
)

type channelUsecase struct {
	channelRepository repository.ChannelRepository
}

func NewChannnelUsecase(
	channelRepository repository.ChannelRepository,
) *channelUsecase {
	return &channelUsecase{
		channelRepository: channelRepository,
	}
}

type ChannelUsecase interface {
	ChannelLister
	ChannelCreator
	ChannelDeleter
	ChannelGetter
}

var _ ChannelGetter = (*channelUsecase)(nil)
var _ ChannelLister = (*channelUsecase)(nil)
var _ ChannelCreator = (*channelUsecase)(nil)
var _ ChannelDeleter = (*channelUsecase)(nil)

type (
	ChannelGetter interface {
		GetChannel(ctx context.Context, input *GetChannelInput) (output *GetChannelOutput, err error)
	}
	GetChannelInput struct {
		ID pulid.ID
	}
	GetChannelOutput struct {
		Channel *model.Channel
	}
)

type (
	ChannelLister interface {
		GetChannelList(ctx context.Context, input *GetChannelListInput) (output *GetChannelListOutput, err error)
	}
	GetChannelListInput struct {
	}
	GetChannelListOutput struct {
		Channels []*model.Channel
	}
)

type (
	ChannelCreator interface {
		CreateChannel(ctx context.Context, input *CreateChannelInput) (output *CreateChannelOutput, err error)
	}
	CreateChannelInput struct {
		Name string
	}
	CreateChannelOutput struct {
		Channel *model.Channel
	}
)

type (
	ChannelDeleter interface {
		DeleteChannel(ctx context.Context, input *DeleteChannelInput) (output *DeleteChannelOutput, err error)
	}
	DeleteChannelInput struct {
		ID pulid.ID
	}
	DeleteChannelOutput struct {
	}
)

func (u *channelUsecase) GetChannel(ctx context.Context, input *GetChannelInput) (output *GetChannelOutput, err error) {
	gcOut, err := u.channelRepository.GetChannel(ctx, &repository.GetChannelInput{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &GetChannelOutput{
		Channel: gcOut.Channel,
	}, nil
}

func (u *channelUsecase) GetChannelList(ctx context.Context, input *GetChannelListInput) (output *GetChannelListOutput, err error) {
	gcOutm, err := u.channelRepository.GetChannelList(ctx, &repository.GetChannelListInput{})
	if err != nil {
		return nil, err
	}

	return &GetChannelListOutput{
		Channels: gcOutm.Channels,
	}, nil
}

func (u *channelUsecase) CreateChannel(ctx context.Context, input *CreateChannelInput) (output *CreateChannelOutput, err error) {
	ccOut, err := u.channelRepository.CreateChannel(ctx, &repository.CreateChannelInput{
		ID:   newPULIDFunc(model.ULIDChannelPrefix),
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return &CreateChannelOutput{
		Channel: ccOut.Channel,
	}, nil
}

func (u *channelUsecase) DeleteChannel(ctx context.Context, input *DeleteChannelInput) (output *DeleteChannelOutput, err error) {
	_, err = u.channelRepository.DeleteChannel(ctx, &repository.DeleteChannelInput{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &DeleteChannelOutput{}, nil
}
