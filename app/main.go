package main

import (
  "github.com/gin-gonic/gin"
  _ "github.com/lib/pq"
  "app/controller"
  "app/middleware"
)


func init() {
  router := gin.Default()

  router.Use(middleware.unchara)

  girackRouter := router.Group("/girack")
  {
    v1 := router.Group("/v1")
    {
      router.POST("/users", )
      router.GET("/users", )
      router.GET("/users/:id", )
    }
  }

  router.Run()
}

