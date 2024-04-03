package v1

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Fields of the User.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique(),
		field.String("title").NotEmpty(),
		field.UUID("user_id", uuid.UUID{}),
		field.String("chat"),
		field.Time("created_at").Default(time.Now()).Immutable(),
	}
}

// Edges of the User.
func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "chats" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("owner", User.Type).
			Ref("chats").
			Field("user_id").
			// setting the edge to unique, ensure
			// that a chat can have only one owner user.
			Unique().
			Required(),
	}
}

func (Chat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
