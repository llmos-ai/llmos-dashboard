package v1

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Modelfile holds the schema definition for the Modelfile entity.
type Modelfile struct {
	ent.Schema
}

// Fields of the Modelfile.
func (Modelfile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique(),
		field.String("tagName").StorageKey("tag_name").NotEmpty().Unique(),
		field.String("modelfile").Default("").NotEmpty(),
		field.UUID("userId", uuid.UUID{}).StorageKey("user_id"),
		field.Time("createdAt").StorageKey("created_at").Default(time.Now()).Immutable(),
	}
}

// Edges of the Modelfile.
func (Modelfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("modelfiles").
			Field("userId").
			Unique().
			Required(),
	}
}

func (Modelfile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("userId", "tagName"),
	}
}
