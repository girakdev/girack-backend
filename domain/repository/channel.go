//go:generate mockgen -source=channel.go -destination=./mock/channel_mock.go -package=mock
package repository

import (
	"context"

	"github.com/girakdev/girack-backend/application/model"
)

type ChannelRepository interface {
	ChannelGetter
	ChannelListGetter
	ChannelCreator
	ChannelDeleter
}

type (
	ChannelGetter interface {
		GetChannel(ctx context.Context, input *GetChannelInput) (output *GetChannelOutput, err error)
	}
	GetChannelInput struct {
		ID model.ID
	}
	GetChannelOutput struct {
		Channel *model.Channel
	}
)
type (
	ChannelListGetter interface {
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
		ID   model.ID
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
		ID model.ID
	}
	DeleteChannelOutput struct {
	}
)
