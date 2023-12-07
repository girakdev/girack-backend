package model

import "github.com/girakdev/girack-backend/internal/pulid"

type Channel struct {
	ID   pulid.ID
	Name string
}

const ULIDChannelPrefix = "channel"
