package main

import (
	"context"
	"log"

	"github.com/girakdev/girack-backend/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=girack password=isataku sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
