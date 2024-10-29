package model

import "github.com/girakdev/girack-backend/internal/pulid"

type User struct {
	ID   pulid.ID
	Name string
}

const ULIDUserPrefix = "user"
