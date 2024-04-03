package v1

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Modelfile holds the schema definition for the Modelfile entity.
type Modelfile struct {
	ent.Schema
}

// Fields of the User.
func (Modelfile) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Positive().Unique(),
		field.String("tag_name").NotEmpty().Unique(),
		field.String("modelfile").Default(""),
		field.Time("created_at").Default(time.Now()).Immutable(),
	}
}

// Edges of the User.
func (Modelfile) Edges() []ent.Edge {
	return nil
}
