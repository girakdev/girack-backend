package controller

import (
	"database/sql"
  "net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

  "app/db"
  "app/entity"
)

type GirackController interface {
  DeleteUser()
}
type girackController struct {
  Db *sql.DB
}

func CreategirackController() (gc *girackController) {
  var conf = "host=postgres port=5555 user=girak password=password dbname=girack sslmode=disable"

  conn, err := sql.Open("postgres", conf)
  defer conn.Close()

  if err != nil {
    panic(err)
  }

  gc = new(girackController)
  gc.Db = conn

  return
}
func (gc *girackController) DeleteUser(c *gin.Context){

  c.JSON(http.StatusOK, gin.H{ "message": "CreateUser", })
  //Todo
}
func (gc *girackController)CreateUser(c *gin.Context) {

  c.JSON(http.StatusOK, gin.H{ "message": "CreateUser", })
  // Todo
}
func (gc *girackController) Index(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{ "message": "Index", })
  // Todo
}
func (gc *girackController) Show(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{ "message": "Show", })
  // Todo
}
