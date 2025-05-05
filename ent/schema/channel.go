package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/girakdev/girack-backend/application/model"
)

// User holds the schema definition for the User entity.
type Channel struct {
	ent.Schema
}

// Fields of the User.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(model.ID("")).
			DefaultFunc(func() model.ID { return model.MustNew("Channel") }),
		field.String("name"),
	}
}

// Edges of the User.
func (Channel) Edges() []ent.Edge {
	return nil
}
