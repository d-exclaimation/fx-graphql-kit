package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Thought holds the schema definition for the Thought entity.
type Thought struct {
	ent.Schema
}

// Fields of the Thought.
func (Thought) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.Text("body").
			NotEmpty().
			Default(""),
		field.String("imageURL"),
		field.Int64("userId"),
	}
}

// Edges of the Thought.
func (Thought) Edges() []ent.Edge {
	return nil
}
