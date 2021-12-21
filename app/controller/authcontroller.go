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
  "app/entity"
)

func Register(c *gin.Context) {
  db := db.Db
  user := entity.User{}
  password := c.Param("password")
  log.Println(password)
  query := "INSERT INTO users (email, name, password) VALUES($1, $2, $3)"

  c.BindJSON(&user)

  stmt, err := db.Prepare(query)
  logFatal(err)
  defer stmt.Close()
  hashpassword, err := passwordHash(password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "inviled password format"})
    return
  }

  _, err = stmt.Exec(user.Name, user.Email, hashpassword)
  logFatal(err)

  message := "Create " + user.Name
  c.JSON(http.StatusCreated, gin.H{"message": message})
}

func Login(c *gin.Context) {
  session := sessions.Default(c)
  email := c.Param("email")
  password := c.Param("password")
  log.Println(email, password)

  hashpassword, _ := passwordHash(password)
  dbpassword, err := GetPasswordByEmail(email)
  if err == sql.ErrNoRows {
    c.JSON(http.StatusUnauthorized, gin.H{"error1": "incorrect Email or Password"})
    log.Println(err)
    return
  } else if err != nil {
    log.Println(err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "EEEEEEEEEEEEEEEERROR"})
    return
  }

  err = passwordVerify(hashpassword, dbpassword)
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect Email or Password"})
    return
  }

  session.Set("user", email)
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
  session.Clear()
  if err := session.Save(); err != nil {
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

func passwordVerify(hash, pw string) error {
  return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
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
