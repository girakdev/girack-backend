package main

import (
    "database/sql"
    "fmt"
    "log"

    // postgres driver
    _ "github.com/lib/pq"
)

type TestUser struct {
    UserID   int
    Password string
}

func main() {

    var Db *sql.DB
    Db, err := sql.Open("postgres", "host=postgres user=app_user password=password dbname=app_db sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    sql := "SELECT user_id, user_password FROM TEST_USER WHERE user_id=$1;"

    pstatement, err := Db.Prepare(sql)
    if err != nil {
        log.Fatal(err)
    }

    queryID := 1
    var testUser TestUser

    err = pstatement.QueryRow(queryID).Scan(&testUser.UserID, &testUser.Password)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(testUser.UserID, testUser.Password)
}
