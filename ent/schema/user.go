package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/girakdev/girack-backend/application/model"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(model.ID("")).
			DefaultFunc(func() model.ID { return model.MustNew("User") }),
		field.Int("age").Positive(),
		field.String("name").Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
