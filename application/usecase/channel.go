package usecase

import (
	"context"

	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/domain/repository"
)

type channelUsecase struct {
	channelRepository repository.ChannelRepository
}

var _ ChannelLister = (*channelUsecase)(nil)

func NewChannnelUsecase(
	channelRepository repository.ChannelRepository,
) *channelUsecase {
	return &channelUsecase{
		channelRepository: channelRepository,
	}
}

type ChannelUsecase interface {
	ChannelLister
}

type (
	ChannelLister interface {
		GetChannelList(ctx context.Context, input *GetChannelListInput) (output *GetChannelListOut, err error)
	}
	GetChannelListInput struct {
	}
	GetChannelListOut struct {
		Channels []model.Channel
	}
)

func (u *channelUsecase) GetChannelList(ctx context.Context, input *GetChannelListInput) (output *GetChannelListOut, err error) {
	return &GetChannelListOut{}, nil
}
