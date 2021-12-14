package db

import (
  "database/sql"
  _ "github.com/lib/pq"
  "log"
)

var Db *sql.DB
var err error

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


  Db, err = sql.Open("postgres", conf)
  logFatal(err)
  err = Db.Ping()
  logFatal(err)

  log.Println("database successfully configured")
}

func CloseDB() {
  err = Db.Close()
  logFatal(err)
}
