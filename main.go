package main

import (
  "context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	conn = "host=postgres port=5432 user=test_user password=password dbname=girack sslmode=disable"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func slowQuery(ctx context.Context) error {
	_, err := db.ExecContext(ctx, "SELECT pg_sleep(5)")
	return err
}

func slowHandler(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	err := slowQuery(req.Context())
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Fprintln(w, "OK")
	fmt.Printf("slowHandler took: %v\n", time.Since(start))
}

func main() {
	var err error

	db, err = sql.Open("postgres", conn)
	logFatal(err)

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

	err = db.PingContext(ctx)
	logFatal(err)

	srv := http.Server{
		Addr:         ":8080",
		WriteTimeout: 2 * time.Second,
    Handler:      http.TimeoutHandler(
      http.HandlerFunc(slowHandler),
      1*time.Second,
      "TIMEOUT!!",
    ),
	}

	log.Println("Start Http Server...")
	log.Fatal(srv.ListenAndServe())
}
