package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/girakdev/girack-backend/application/usecase"
	"github.com/girakdev/girack-backend/controller/model"
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
func (c *channelController) ListChannel(g *gin.Context) {
	gcOut, err := c.channnelUsecase.GetChannelList(g, &usecase.GetChannelListInput{})
	if err != nil {
		log.Println(err)
		g.JSON(http.StatusInternalServerError, `uooo`)
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

	if err != nil {
		g.JSON(http.StatusInternalServerError, `usecase does not work`)
		return
	}

	g.JSON(http.StatusOK, channels)
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
func (c *channelController) CreateChannel(g *gin.Context) {
	type payload struct {
		Name string `json:"name"`
	}

	p := &payload{}
	if err := g.BindJSON(p); err != nil {
		g.JSON(http.StatusInternalServerError, `parse error`)
		return
	}
	log.Println(p.Name)
	ccOut, err := c.channnelUsecase.CreateChannel(g, &usecase.CreateChannelInput{
		Name: p.Name,
	})

	if err != nil {
		g.JSON(http.StatusInternalServerError, `usecase`)
		return
	}
	g.JSON(http.StatusOK, model.Channel{
		ID:   string(ccOut.Channel.ID),
		Name: p.Name,
	})
}
