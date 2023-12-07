//go:generate mockgen -source=channel.go -destination=./mock/channel_mock.go -package=mock
package usecase

import (
	"context"

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
}

var _ ChannelLister = (*channelUsecase)(nil)
var _ ChannelCreator = (*channelUsecase)(nil)

type (
	ChannelLister interface {
		GetChannelList(ctx context.Context, input *GetChannelListInput) (output *GetChannelListOut, err error)
	}
	GetChannelListInput struct {
	}
	GetChannelListOut struct {
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

func (u *channelUsecase) GetChannelList(ctx context.Context, input *GetChannelListInput) (output *GetChannelListOut, err error) {
	gcOutm, err := u.channelRepository.GetChannels(ctx, &repository.GetChannelsInput{})
	if err != nil {
		return nil, err
	}

	return &GetChannelListOut{
		Channels: gcOutm.Channels,
	}, nil
}

func (u *channelUsecase) CreateChannel(ctx context.Context, input *CreateChannelInput) (output *CreateChannelOutput, err error) {
	ccOut, err := u.channelRepository.CreateChannel(ctx, &repository.CreateChannelInput{
		ID:   pulid.MustNew(model.ULIDChannelPrefix),
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return &CreateChannelOutput{
		Channel: ccOut.Channel,
	}, nil
}
