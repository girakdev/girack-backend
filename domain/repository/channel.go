//go:generate mockgen -source=channel.go -destination=./mock/channel_mock.go -package=mock
package repository

import (
	"context"

	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/internal/pulid"
)

type ChannelRepository interface {
	ChannelsGetter
	ChannelCreator
}

type (
	ChannelsGetter interface {
		GetChannels(ctx context.Context, input *GetChannelsInput) (output *GetChannelsOutput, err error)
	}
	GetChannelsInput struct {
	}
	GetChannelsOutput struct {
		Channels []*model.Channel
	}
)

type (
	ChannelCreator interface {
		CreateChannel(ctx context.Context, input *CreateChannelInput) (output *CreateChannelOutput, err error)
	}
	CreateChannelInput struct {
		ID   pulid.ID
		Name string
	}
	CreateChannelOutput struct {
		Channel *model.Channel
	}
)
