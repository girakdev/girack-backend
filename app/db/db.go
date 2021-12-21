package db

import (
  "fmt"
  "os"

  "database/sql"
  _ "github.com/lib/pq"

  "log"
)

var (
  Db *sql.DB
  err error
)

var (
  schema = "host=%s user=%s password=%s dbname=%s sslmode=disable"
  host     = "girack_db"
  userName = os.Getenv("POSTGRES_USER")
  password = os.Getenv("POSTGRES_PASSWORD")
  dbName   = os.Getenv("POSTGRES_DB")
)

func InitDB() {
  conf := fmt.Sprintf(schema, host, userName, password, dbName)

  Db, err = sql.Open("postgres", conf)
  if err != nil {
    log.Fatal(err)
  }
  err = Db.Ping()

  if err != nil {
    log.Fatal(err)
  }

  log.Println("database successfully configured")
}

func CloseDB() {
  err = Db.Close()
  if err != nil {
    log.Fatal(err)
  }
}
