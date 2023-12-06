package repository

import (
	"testing"

	"github.com/girakdev/girack-backend/ent/enttest"
	_ "github.com/lib/pq"
)

func TestChannelRepository_GetChannels(t *testing.T) {
	client := enttest.Open(t, "postgres", "host=db port=5432 user=postgres dbname=girack password=isataku sslmode=disable")
	defer client.Close()
}
