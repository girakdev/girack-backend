package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/girakdev/girack-backend/application/usecase"
)

type channelController struct {
	channnelUsecase usecase.ChannelUsecase
}

func NewChannelHandler(channelUsecase usecase.ChannelUsecase) *channelController {
	return &channelController{
		channnelUsecase: channelUsecase,
	}
}

func (a *channelController) ListChannel(c *gin.Context) {
	c.JSON(http.StatusOK, []byte("{}"))
}
