package main

import (
  "github.com/gin-gonic/gin"
  "app/controller"
  "app/db"
)
func main() {
  db.InitDB()
  router := gin.Default()

  girackRouter := router.Group("/girack")
  {
    v1 := girackRouter.Group("/v1")
    {
      users := v1.Group("/users")
      {
        users.POST("", controller.CreateUser)
        users.POST(":id", controller.UpdateUser)
        users.GET(":id", controller.GetUser)
        users.GET("", controller.GetAllUser)
        users.DELETE(":id", controller.DeleteUser)
      }
      channels := v1.Group("/channels")
      {
        channels.POST("", controller.CreateUser)
        channels.POST("", controller.CreateUser)
        channels.GET("", controller.CreateUser)
        channels.GET("", controller.CreateUser)
        channels.DELETE("", controller.DeleteUser)
      }
    }
  }

  router.Run()
  db.CloseDB()
}
