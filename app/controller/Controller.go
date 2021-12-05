package controller

import (
  "github.com/gin-gonic/gin"
  "database/sql"

  _ "github.com/lib/pq"
)

const (
  conf = "host=postgres port=5555 user=girak password=password dbname=girack sslmode=disable"
)

func CreateNewController() (controller *Controller){
  conn, err := sql.Open("postgres", conf)
  if err != nil {
    panic(err.Error)
  }
  defer conn.Close()

  controller = new(Controller)
  controller.db = conn

  return controller
}

type Controller struct {
  db *sql.DB
}

func (conn *Controller) CreateUser(c *gin.Context) {
  // shori
}

func (conn *Controller) Index(c *gin.Context){
  // shori
}

func (conn *Controller) Show(c *gin.Context){
  // shori
}
