package controller

import (
  "github.com/gin-gonic/gin"
  "database/sql"

  _ "github.com/lib/pq"
)


type Controller struct {
  db *sql.DB
}

func CreateNewController() (controller *Controller){
  conn, err := sql.Open("postgres", "host=postgres port=5555 user=girak password=password dbname=girack sslmode=disable")
  if err != nil {
    panic(err.Error)
  }
  defer conn.Close()

  controller = new(Controller)
  controller.db = conn

  return controller
}

func (conn *Controller) CreateUser(c *gin.Context) {
  c.JSON( 200, gin.H{ "message": "CreateUser", })
  // shori
}

func (conn *Controller) Index(c *gin.Context){
  c.JSON( 200, gin.H{ "message": "Index", })
  // shori
}

func (conn *Controller) Show(c *gin.Context){
  c.JSON( 200, gin.H{ "message": "Show", })
  // shori
}
