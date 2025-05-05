package usecase

import (
	"context"
	"testing"

	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/domain/repository"
	mockrepository "github.com/girakdev/girack-backend/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestChannelUsecase_GetChannel(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		ctrl := gomock.NewController(t)
		channelRepository := mockrepository.NewMockChannelRepository(ctrl)

		channelUsecase := NewChannnelUsecase(channelRepository)

		channelRepository.EXPECT().GetChannel(
			gomock.Any(),
			&repository.GetChannelInput{
				ID: "1",
			},
		).Return(
			&repository.GetChannelOutput{
				Channel: &model.Channel{
					ID:   "1",
					Name: "channel",
				},
			}, nil)

		out, err := channelUsecase.GetChannel(ctx, &GetChannelInput{ID: "1"})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		want := &GetChannelOutput{
			Channel: &model.Channel{
				ID:   "1",
				Name: "channel",
			},
		}

		if diff := cmp.Diff(out, want); diff != "" {
			t.Errorf("-want +got %v", diff)
		}
	})
}

func TestChannelUsecase_ListChannel(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		ctrl := gomock.NewController(t)
		channelRepository := mockrepository.NewMockChannelRepository(ctrl)

		channelUsecase := NewChannnelUsecase(channelRepository)

		channelRepository.EXPECT().GetChannelList(
			gomock.Any(),
			&repository.GetChannelListInput{},
		).Return(
			&repository.GetChannelListOutput{
				Channels: []*model.Channel{
					{
						ID:   "1",
						Name: "channel1",
					},
					{
						ID:   "2",
						Name: "channel2",
					},
				},
			},
			nil,
		)

		out, err := channelUsecase.GetChannelList(ctx, &GetChannelListInput{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		want := &GetChannelListOutput{
			Channels: []*model.Channel{
				{
					ID:   "1",
					Name: "channel1",
				},
				{
					ID:   "2",
					Name: "channel2",
				},
			},
		}

		if diff := cmp.Diff(out, want); diff != "" {
			t.Errorf("-want +got %v", diff)
		}
	})
}

func TestChannelUsecase_CreateChannel(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		channelRepository := mockrepository.NewMockChannelRepository(ctrl)

		channelUsecase := NewChannnelUsecase(channelRepository)

		newIDFunc = func(_ string) model.ID {
			return "1"
		}

		channelRepository.EXPECT().CreateChannel(
			gomock.Any(),
			&repository.CreateChannelInput{
				ID:   "1",
				Name: "channel",
			},
		).Return(
			&repository.CreateChannelOutput{
				Channel: &model.Channel{
					ID:   "1",
					Name: "channel",
				},
			},
			nil,
		)

		out, err := channelUsecase.CreateChannel(ctx, &CreateChannelInput{Name: "channel"})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		want := &CreateChannelOutput{
			Channel: &model.Channel{
				ID:   "1",
				Name: "channel",
			},
		}

		if diff := cmp.Diff(out, want); diff != "" {
			t.Errorf("-want +got %v", diff)
		}
	})
}

func TestChannelController_DeleteChannel(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		channelRepository := mockrepository.NewMockChannelRepository(ctrl)

		channelUsecase := NewChannnelUsecase(channelRepository)

		channelRepository.EXPECT().DeleteChannel(
			gomock.Any(),
			&repository.DeleteChannelInput{
				ID: "1",
			},
		).Return(
			&repository.DeleteChannelOutput{},
			nil,
		)

		_, err := channelUsecase.DeleteChannel(ctx, &DeleteChannelInput{ID: "1"})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
