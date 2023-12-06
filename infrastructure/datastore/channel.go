package datastore

import (
	"context"

	"github.com/girakdev/girack-backend/domain/repository"
)

type channelRepository struct {
}

var _ repository.ChannelRepository = (*channelRepository)(nil)

func NewChannelRepository() *channelRepository {
	return &channelRepository{}
}

func (r channelRepository) GetChannels(ctx context.Context, input *repository.GetChannelsInput) (output *repository.GetChannelsOutput, err error) {
	return &repository.GetChannelsOutput{}, nil
}
