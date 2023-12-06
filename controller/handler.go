package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/girakdev/girack-backend/application/usecase"
	_ "github.com/girakdev/girack-backend/docs"
	"github.com/girakdev/girack-backend/infrastructure/datastore"
	ginfiles "github.com/swaggo/files"         // swagger embed files
	ginswagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Giarck
// @version 1.0
// @license.name isataku, whale-yay, kirota
// @description Girack webapi reference
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

	r.GET("/swagger/*any", ginswagger.WrapHandler(ginfiles.Handler))

	return r
}