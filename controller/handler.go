package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/girakdev/girack-backend/application/usecase"
	"github.com/girakdev/girack-backend/infrastructure/datastore"
)

func Router() *gin.Engine {
	// Repository
	channelRepository := datastore.NewChannelRepository()

	// Usecase
	channelUsecase := usecase.NewChannnelUsecase(channelRepository)

	// Handler
	channelController := NewChannelHandler(channelUsecase)

	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/channels", channelController.ListChannel)

	return r
}
