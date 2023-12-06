//go:generate mockgen -source=channel.go -destination=./mock/channel_mock.go -package=mock
package repository

import (
	"context"

	"github.com/girakdev/girack-backend/application/model"
)

type ChannelRepository interface {
	ChannelsGetter
}

type (
	ChannelsGetter interface {
		GetChannels(ctx context.Context, input *GetChannelsInput) (output *GetChannelsOutput, err error)
	}
	GetChannelsInput struct {
	}
	GetChannelsOutput struct {
		Channels []model.Channel
	}
)