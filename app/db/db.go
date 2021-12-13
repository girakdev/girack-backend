package db

import (
  "database/sql"
  _ "github.com/lib/pq"
  "log"
  "app/entity"
)

var db *sql.DB

const (
  queryGetUser    = "SELECT email, name FROM users WHERE id=$1"
)
func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func InitDB() {
  // security hole
  conf := "host=girack_db user=girack_user password=password dbname=girack_db sslmode=disable"

  db, err := sql.Open("postgres", conf)
  logFatal(err)
  defer db.Close()
  err = db.Ping()
  logFatal(err)

  log.Println("database successfully configured")

  /*
  user := entity.User{}

  stmt, err := db.Prepare(queryGetUser)
  defer stmt.Close()
  logFatal(err)

  err = stmt.QueryRow(1).Scan(&user.Email, &user.Name)
  log.Println(user)
  */
}

func GetDB() *sql.DB {
  return db
}
