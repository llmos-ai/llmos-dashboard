package v1

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	entv1User "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique(),
		field.String("name").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty().Sensitive(),
		field.Enum("role").Default("pending").
			Values("admin", "user", "pending").Default("pending"),
		field.String("profileImageUrl").Default("").StorageKey("profile_image_url"),
		field.Time("createdAt").StorageKey("created_at").Default(time.Now()).Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chats", Chat.Type),
		edge.To("modelfiles", Modelfile.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "email").Unique(),
		index.Fields("role"),
	}
}

func GetUserRole(role string) entv1User.Role {
	switch role {
	case "admin":
		return entv1User.RoleAdmin
	case "user":
		return entv1User.RoleUser
	default:
		return entv1User.RolePending
	}
}
