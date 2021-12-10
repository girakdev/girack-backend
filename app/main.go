package main

import (
  //"github.com/gin-gonic/gin"
  "database/sql"
  _ "github.com/lib/pq"
  //"app/controller"
  //"app/db"
  "app/entity"
  "fmt"
)

func main() {
  //db.InitDB()
  conf := "host=girack_db user=root password=password dbname=postgres sslmode=disable"
  db, err := sql.Open("postgres", conf)
  defer db.Close()
  if err != nil {
    panic(err.Error())
  }

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }

  rows, err := db.Query("SELECT * FROM users")
  if err != nil {
    panic(err.Error())
  }

  u := entity.User{}

  for rows.Next() {
    err := rows.Scan(&u.Email, &u.Name)
    if err != nil {
      panic(err)
    }
    fmt.Println(u.Name, u.Email)

  }
}

/*
func main() {
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

*/
