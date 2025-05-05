package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/application/usecase"
	mockusecase "github.com/girakdev/girack-backend/application/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestChannelController_GetChannel(t *testing.T) {
	t.Parallel()
	t.Run("OK", func(t *testing.T) {
		t.Parallel()
		var (
			id      = "1"
			channel = model.Channel{
				ID: model.ID(id),

				Name: "channel",
			}
		)

		ctrl := gomock.NewController(t)
		channelUsecase := mockusecase.NewMockChannelUsecase(ctrl)
		channelUsecase.EXPECT().GetChannel(
			gomock.Any(),
			&usecase.GetChannelInput{
				ID: channel.ID,
			},
		).Return(
			&usecase.GetChannelOutput{
				Channel: &model.Channel{
					ID:   channel.ID,
					Name: channel.Name,
				},
			},
			nil,
		)

		c := NewChannelHandler(channelUsecase)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		ctx.AddParam("id", id)
		c.GetChannel(ctx)
	})
}

func TestChannelController_ListChannel(t *testing.T) {
	t.Parallel()
	t.Run("OK", func(t *testing.T) {
		t.Parallel()
		var (
			channels = []*model.Channel{
				{
					ID:   model.ID("1"),
					Name: "channel1",
				},
				{
					ID:   model.ID("2"),
					Name: "channel2",
				},
			}
		)

		ctrl := gomock.NewController(t)
		channelUsecase := mockusecase.NewMockChannelUsecase(ctrl)
		channelUsecase.EXPECT().GetChannelList(
			gomock.Any(),
			&usecase.GetChannelListInput{},
		).Return(
			&usecase.GetChannelListOutput{
				Channels: channels,
			},
			nil,
		)

		c := NewChannelHandler(channelUsecase)
		resp := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(resp)
		c.ListChannel(ctx)

		if http.StatusOK != resp.Code {
			t.Errorf("expected: %d, got: %d", http.StatusOK, resp.Code)
		}

		var data []*model.Channel
		if err := json.Unmarshal(resp.Body.Bytes(), &data); err != nil {
			t.Errorf("JSON Unmarshal error: %v", err)
			return
		}
		if diff := cmp.Diff(channels, data); diff != "" {
			t.Errorf("-want +got %v", diff)
		}
	})
}

func TestChannelController_CreateChannel(t *testing.T) {
	t.Parallel()
	t.Run("OK", func(t *testing.T) {
		t.Parallel()
		var (
			id      = "1"
			name    = "channel"
			channel = model.Channel{
				ID:   model.ID(id),
				Name: name,
			}
		)

		ctrl := gomock.NewController(t)
		channelUsecase := mockusecase.NewMockChannelUsecase(ctrl)
		channelUsecase.EXPECT().CreateChannel(
			gomock.Any(),
			&usecase.CreateChannelInput{
				Name: name,
			},
		).Return(
			&usecase.CreateChannelOutput{
				Channel: &model.Channel{
					ID:   channel.ID,
					Name: channel.Name,
				},
			},
			nil,
		)

		var err error
		body, err := json.Marshal(channel)
		if err != nil {
			t.Error(err)
		}

		c := NewChannelHandler(channelUsecase)
		resp := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(resp)
		ctx.Request, err = http.NewRequest(http.MethodPost, "/channels", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
			return
		}

		c.CreateChannel(ctx)
		if http.StatusOK != resp.Code {
			t.Errorf("expected: %d, got: %d", http.StatusOK, ctx.Writer.Status())
		}

		var data model.Channel
		if err := json.Unmarshal(resp.Body.Bytes(), &data); err != nil {
			t.Errorf("JSON Unmarshal error: %v", err)
			return
		}
		if diff := cmp.Diff(channel, data); diff != "" {
			t.Errorf("-want +got %v", diff)
		}
	})
}

func TestChannelController_DeleteChannel(t *testing.T) {
	t.Parallel()
	t.Run("OK", func(t *testing.T) {
		t.Parallel()
		var (
			id = "1"
		)

		ctrl := gomock.NewController(t)
		channelUsecase := mockusecase.NewMockChannelUsecase(ctrl)
		channelUsecase.EXPECT().DeleteChannel(
			gomock.Any(),
			&usecase.DeleteChannelInput{
				ID: model.ID(id),
			},
		).Return(
			nil,
			nil,
		)

		c := NewChannelHandler(channelUsecase)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		ctx.AddParam("id", id)
		c.DeleteChannel(ctx)
	})
}
