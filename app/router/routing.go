package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

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
      v1.GET("/logout", controller.Logout)

      authUserGroup := v1.Group("/auth")
      authUserGroup.Use(AuthRequired)
      {
        users := authUserGroup.Group("/users")
        {
          users.PUT(":id", controller.UpdateUser)
          users.GET(":id", controller.GetUser)
          users.GET("", controller.GetAllUser)
          users.DELETE(":id", controller.DeleteUser)
        }
/*
        channels := authUserGroup.Group("/channels")
        {
          channels.POST("", controller.CreatChannel)
          channels.PUT("", controller.UpdateChannel)
          channels.GET(":id", controller.GetChannel)
          channels.GET("", controller.GetAllChannell)
          channels.DELETE("", controller.DeleteChannel)
        }
        message := authUserGroup.Group("/messages")
        {
          message.POST("", controller.CreateMessage)
          message.PUT("", controller.UpdateMessage)
          message.GET("", controller.GetMessage)
          message.DELETE("", controller.DeleteMessage)
        }
*/
      }
    }
  }

  return router
}

func AuthRequired(c *gin.Context) {
  session := sessions.Default(c)
  user := session.Get("userid")
  if user == nil {
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
    return
  }
  c.Next()
}

