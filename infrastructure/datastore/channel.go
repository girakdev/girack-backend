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

func (r *channelRepository) GetChannel(ctx context.Context, input *repository.GetChannelInput) (*repository.GetChannelOutput, error) {
	channel, err := r.client.Channel.Get(ctx, input.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &repository.GetChannelOutput{
		Channel: &model.Channel{
			ID:   channel.ID,
			Name: channel.Name,
		},
	}, nil
}

func (r *channelRepository) GetChannelList(ctx context.Context, input *repository.GetChannelListInput) (*repository.GetChannelListOutput, error) {
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

	return &repository.GetChannelListOutput{
		Channels: res,
	}, nil
}

func (r *channelRepository) CreateChannel(ctx context.Context, input *repository.CreateChannelInput) (*repository.CreateChannelOutput, error) {
	channel, err := r.client.Channel.Create().SetID(input.ID).SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &repository.CreateChannelOutput{
		Channel: &model.Channel{
			ID:   model.ID(channel.ID),
			Name: channel.Name,
		},
	}, nil
}

func (r *channelRepository) DeleteChannel(ctx context.Context, input *repository.DeleteChannelInput) (*repository.DeleteChannelOutput, error) {
	err := r.client.Channel.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &repository.DeleteChannelOutput{}, nil
}
