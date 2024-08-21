package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty(),
		field.Time("due_at").Nillable().Optional().Default(nil),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
