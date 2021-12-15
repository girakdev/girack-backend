package router

import (
  "github.com/gin-gonic/gin"
  "app/controller"
)

func CreateRouter() *gin.Engine {
  router := gin.Default()

  girackRouter := router.Group("/girack")
  {
    v1 := girackRouter.Group("/v1")
    {
      users := v1.Group("/users")
      {
        users.POST("", controller.CreateUser)
        users.PUT(":id", controller.UpdateUser)
        users.GET(":id", controller.GetUser)
        users.GET("", controller.GetAllUser)
        users.DELETE(":id", controller.DeleteUser)
      }
 /*     channels := v1.Group("/channels")
      {
        channels.POST("", controller.CreatChannel)
        channels.POST("", controller.UpdateChannel)
        channels.GET("", controller.GetChannel)
        channels.GET("", controller.GetAllChannell)
        channels.DELETE("", controller.DeleteChannel)
      }
  */
    }
  }

  router.Run()

  return router
}
