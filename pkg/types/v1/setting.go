package v1

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Setting holds the schema definition for the Setting entity
type Setting struct {
	ent.Schema
}

// Fields of the Setting
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("default").Default("").NotEmpty(),
		field.String("value").Optional(),
		field.Bool("isActive").StorageKey("is_active").Default(true),
		field.Bool("readOnly").StorageKey("read_only").Default(false),
		field.Time("createdAt").StorageKey("created_at").Default(time.Now()).Immutable(),
	}
}

func (Setting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}
