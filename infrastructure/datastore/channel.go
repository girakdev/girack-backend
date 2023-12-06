package datastore

import (
	"context"

	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/domain/repository"
	"github.com/girakdev/girack-backend/ent"
)

type channelRepository struct {
	client *ent.Client
}

var _ repository.ChannelRepository = (*channelRepository)(nil)

func NewChannelRepository() *channelRepository {
	return &channelRepository{}
}

func (r channelRepository) GetChannels(ctx context.Context, input *repository.GetChannelsInput) (*repository.GetChannelsOutput, error) {
	channels, err := r.client.Channel.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*model.Channel, len(channels))
	for _, v := range channels {
		var c *model.Channel
		c.ID = v.ID
		c.Name = v.Name

		res = append(res, c)
	}

	return &repository.GetChannelsOutput{
		Channels: res,
	}, nil
}
