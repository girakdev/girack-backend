package controller

import (
  "log"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gin-contrib/sessions"
	"golang.org/x/crypto/bcrypt"

	"app/db"
)

type AuthUser struct {
  Email string `json:"email"`
  Name string `json:"name"`
  Password string `json:"password"`
}

func Register(c *gin.Context) {
  db := db.Db
  user := AuthUser{}
  query := "INSERT INTO users (email, name, password) VALUES($1, $2, $3)"
  c.BindJSON(&user)

  stmt, err := db.Prepare(query)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": ""})
    return
  }
  defer stmt.Close()
  user.Password, err = passwordHash(user.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "inviled password format"})
    return
  }

  _, err = stmt.Exec(user.Email, user.Name, user.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "inviled password format"})
    return
  }
  message := "Create " + user.Name
  c.JSON(http.StatusOK, gin.H{"message": message})
}

func Login(c *gin.Context) {
  session := sessions.Default(c)
  user := AuthUser{}
  c.BindJSON(&user)

  dbpassword, err := GetPasswordByEmail(user.Email)
  if err == sql.ErrNoRows {
    c.JSON(http.StatusUnauthorized, gin.H{"error1": "incorrect Email or Password"})
    log.Println(err)
    return
  } else if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "ServerError"})
    log.Println(err)
    return
  }

  err = bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(user.Password))
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect Email or Password"})
    log.Println(err)
    return
  }

  session.Set("user", user.Email)
  if err = session.Save(); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"errror": "Failed to save session"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func Logout(c *gin.Context) {
  session := sessions.Default(c)
  user := session.Get("user")
  if user == nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session toekn"})
    return
  }
  session.Delete("user")
  err := session.Save()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}


func passwordHash(pw string) (string, error) {
  hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  }
  return string(hash), err
}

func GetPasswordByEmail(email string) (password string, err error) {
  db := db.Db
  query := "SELECT password FROM users WHERE email = $1"

  stmt, err := db.Prepare(query)
  if err != nil {
    return "", err
  }
  err = stmt.QueryRow(email).Scan(&password)
  if err != nil {
    return "", err
  }
  return
}
