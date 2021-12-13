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
      v1.POST("/users", controller.CreateUser)
      v1.GET("/users", controller.DeleteUser)
      v1.GET("/users/:id", controller.GetUser)
    }
  }

  router.Run()
}
