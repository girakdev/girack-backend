package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/girakdev/girack-backend/internal/pulid"
)

// User holds the schema definition for the User entity.
type Channel struct {
	ent.Schema
}

// Fields of the User.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(pulid.ID("")).
			DefaultFunc(func() pulid.ID { return pulid.MustNew("US") }),
		field.Int("name"),
	}
}

// Edges of the User.
func (Channel) Edges() []ent.Edge {
	return nil
}
