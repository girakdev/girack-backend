package main

import (
  "github.com/gin-gonic/gin"
  "app/controller"
)

func main() {
  router := gin.Default()

  conn := controller.CreateNewController()

  router.POST("/users", conn.CreateUser)
  router.GET("/users", conn.Index)
  router.GET("/users:id", conn.Show)

  router.Run()
}

