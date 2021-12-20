package controller

import (
//	"net/http"
 // "log"

	"github.com/gin-gonic/gin"
 // "database/sql"
  _ "github.com/lib/pq"
  //"strconv"
  "golang.org/x/crypto/bcrypt"
  "github.com/gin-contrib/sessions"

//	"app/db"
//	"app/entity"
)

type authUser struct {
  Email string `json:"email"`
  Name string `json:"name"`
  Password string `json:"password"`
}

func Login(c *gin.Context) {
 // user := authUser{}

  //TODO ログイン処理を書く
  session := sessions.Default(c)
  //session.Set("UserID", userid)
  session.Save()
}

func Register(c *gin.Context) {
  //email, _ := c.GetPostForm("id")
  //password, _ := c.GetPostForm("password")

  //TODO 登録処理を格

}

func Logout(c *gin.Context) {
  session := sessions.Default(c)
  session.Clear()
  session.Save()
}


func passwordHash(pw string) (string, error) {
  hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  }
  return string(hash), err
}

func passwordVerify(hash, pw string) error {
  return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

