package router

import (
  "github.com/gin-gonic/gin"
  "app/controller"
)

var Router *gin.Engine

func init() {
  router := gin.Default()

  conn := controller.CreateNewController()

  router.POST("/users", conn.CreateUser)
  router.GET("/users", conn.Index)
  router.GET("/users/:id", conn.Show)

  Router = router
}

