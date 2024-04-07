/*
Copyright YEAR 1block.ai.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldProfileImageURL holds the string denoting the profile_image_url field in the database.
	FieldProfileImageURL = "profile_image_url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeChats holds the string denoting the chats edge name in mutations.
	EdgeChats = "chats"
	// EdgeModelfiles holds the string denoting the modelfiles edge name in mutations.
	EdgeModelfiles = "modelfiles"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ChatsTable is the table that holds the chats relation/edge.
	ChatsTable = "chats"
	// ChatsInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	ChatsInverseTable = "chats"
	// ChatsColumn is the table column denoting the chats relation/edge.
	ChatsColumn = "user_id"
	// ModelfilesTable is the table that holds the modelfiles relation/edge.
	ModelfilesTable = "modelfiles"
	// ModelfilesInverseTable is the table name for the Modelfile entity.
	// It exists in this package in order to avoid circular dependency with the "modelfile" package.
	ModelfilesInverseTable = "modelfiles"
	// ModelfilesColumn is the table column denoting the modelfiles relation/edge.
	ModelfilesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldRole,
	FieldProfileImageURL,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultProfileImageURL holds the default value on creation for the "profile_image_url" field.
	DefaultProfileImageURL string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Role defines the type for the "role" enum field.
type Role string

// RolePending is the default value of the Role enum.
const DefaultRole = RolePending

// Role values.
const (
	RoleAdmin   Role = "admin"
	RoleUser    Role = "user"
	RolePending Role = "pending"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleAdmin, RoleUser, RolePending:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByProfileImageURL orders the results by the profile_image_url field.
func ByProfileImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfileImageURL, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByChatsCount orders the results by chats count.
func ByChatsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChatsStep(), opts...)
	}
}

// ByChats orders the results by chats terms.
func ByChats(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChatsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByModelfilesCount orders the results by modelfiles count.
func ByModelfilesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newModelfilesStep(), opts...)
	}
}

// ByModelfiles orders the results by modelfiles terms.
func ByModelfiles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModelfilesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newChatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChatsTable, ChatsColumn),
	)
}
func newModelfilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ModelfilesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ModelfilesTable, ModelfilesColumn),
	)
}
