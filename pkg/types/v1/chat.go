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

// note:run go generate if the filed changes
type Histroy struct {
	CurrentID string             `json:"currentId"`
	Messages  map[string]Message `json:"messages"`
}

type Message struct {
	ChildrenIds []string    `json:"childrenIds"`
	Content     string      `json:"content"`
	Context     string      `json:"context,omitempty"`
	ID          string      `json:"id"`
	ParentId    string      `json:"parentId"`
	Role        string      `json:"role"`
	Timestamp   int64       `json:"timestamp"`
	Done        bool        `json:"done,omitempty"`
	Info        interface{} `json:"info,omitempty"`
}

// Fields of the User.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique(),
		field.String("title").NotEmpty(),
		field.UUID("userId", uuid.UUID{}).StorageKey("user_id"),
		//field.Text("chat").NotEmpty(),
		field.JSON("models", []string{}),
		//field.String("options").NotEmpty(),
		field.JSON("tags", []string{}),
		field.JSON("history", Histroy{}),
		field.JSON("messages", []Message{}),
		field.Time("createdAt").StorageKey("created_at").Default(time.Now()).Immutable(),
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
			Field("userId").
			// setting the edge to unique, ensure
			// that a chat can have only one owner user.
			Unique().
			Required(),
	}
}

func (Chat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("userId"),
	}
}
