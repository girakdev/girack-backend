package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/girakdev/girack-backend/application/usecase"
	"github.com/girakdev/girack-backend/domain/model"
	"github.com/girakdev/girack-backend/internal/pulid"
)

type channelController struct {
	channnelUsecase usecase.ChannelUsecase
}

func NewChannelHandler(channelUsecase usecase.ChannelUsecase) *channelController {
	return &channelController{
		channnelUsecase: channelUsecase,
	}
}

// @Summary	List Channel
// @Schemes
// @Description	Get List Channel
// @Tags			channels
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]model.Channel
// @Router			/channels [get]
func (c *channelController) ListChannel(ctx *gin.Context) {
	gcOut, err := c.channnelUsecase.GetChannelList(ctx, &usecase.GetChannelListInput{})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var channels []model.Channel
	channels = make([]model.Channel, 0)

	for _, v := range gcOut.Channels {
		channels = append(channels, model.Channel{
			ID:   string(v.ID),
			Name: v.Name,
		})
	}

	ctx.JSON(http.StatusOK, channels)
}

// @Summary	List Channel
// @Schemes
// @Description	Get List Channel
// @Tags			channels
// @Accept			json
// @Produce		json
// @Param name body string true "チャンネル名"
// @Success		200	{object}	model.Channel
// @Router			/channels [post]
func (c *channelController) CreateChannel(ctx *gin.Context) {
	p := &model.PostChannelPayload{}
	if err := ctx.BindJSON(p); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ccOut, err := c.channnelUsecase.CreateChannel(ctx, &usecase.CreateChannelInput{
		Name: p.Name,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, model.Channel{
		ID:   string(ccOut.Channel.ID),
		Name: p.Name,
	})
}

// @Summary	Delete Channel
// @Schemes
// @Description	Delete Channel
// @Tags			channels
// @Accept			json
// @Produce		json
// @Success		200	{}
// @Router			/channels [delete]
func (c *channelController) DeleteChannel(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := c.channnelUsecase.DeleteChannel(ctx, &usecase.DeleteChannelInput{
		ID: pulid.ID(id),
	})

	if err != nil {
		if err == usecase.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "ok"})
}
