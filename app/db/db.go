package db

import (
  "database/sql"
  _ "github.com/lib/pq"
  "log"
)

var db *sql.DB

func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func Init() {
  // security hole
  conf := "host=girack_db user=root password=password dbname=postgres sslmode=disable"

  db, err := sql.Open("postgres", conf)
  logFatal(err)
  defer db.Close()
  err = db.Ping()
  logFatal(err)

  log.Println("database successfully configured")
}

func GetDB() *sql.DB {
  return db
}
