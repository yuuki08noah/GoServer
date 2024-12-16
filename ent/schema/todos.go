package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Todos holds the schema definition for the Todos entity.
type Todos struct {
	ent.Schema
}

// Fields of the Todos.
func (Todos) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("title").NotEmpty(),
		field.String("description"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Todos.
func (Todos) Edges() []ent.Edge {
	return nil
}
