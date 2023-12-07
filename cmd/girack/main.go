package main

import (
	"log"

	"github.com/girakdev/girack-backend/controller"
	"github.com/girakdev/girack-backend/ent"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=girack password=isataku sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	controller.Router(client).Run()
}
