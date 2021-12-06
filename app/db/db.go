package db

import (
  "databaes/sql"
  _ "github.com/lib/pq"
)

var (
  db *sql.DB
  err error
)

func Init() {
  var conf = "host=postgres port=5555 user=girak password=password dbname=girack sslmode=disable"
  db, err = sql.Open("postgres", conf)
  if err != nil {
    panic(err)
  }

}

func GetDB() *sql.DB {
  return db
}

func Close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}
