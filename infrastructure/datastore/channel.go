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

func NewChannelRepository(client *ent.Client) *channelRepository {
	return &channelRepository{
		client: client,
	}
}

func (r *channelRepository) GetChannels(ctx context.Context, input *repository.GetChannelsInput) (*repository.GetChannelsOutput, error) {
	channels, err := r.client.Channel.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*model.Channel
	res = make([]*model.Channel, 0)
	for _, v := range channels {
		res = append(res, &model.Channel{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return &repository.GetChannelsOutput{
		Channels: res,
	}, nil
}
