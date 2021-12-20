package router

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"

  "app/controller"
)

func CreateRouter() *gin.Engine {
  router := gin.Default()

  store := cookie.NewStore([]byte("secret"))
  router.Use(sessions.Sessions("mysession", store))

  girackRouter := router.Group("/girack")
  {
    v1 := girackRouter.Group("/v1")
    {
      v1.POST("/register", controller.Register)
      v1.POST("/login", controller.Login)
      v1.POST("/logout", controller.Logout)

      authUserGroup := v1.Group("auth")
      authUserGroup.Use(sessionCheck())
      {
        users := authUserGroup.Group("/users")
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
  }

  return router
}

func sessionCheck() gin.HandlerFunc {
  return func(c *gin.Context) {
  }
}
