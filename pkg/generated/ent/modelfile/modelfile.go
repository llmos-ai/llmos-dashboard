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

package modelfile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the modelfile type in the database.
	Label = "modelfile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTagName holds the string denoting the tagname field in the database.
	FieldTagName = "tag_name"
	// FieldModelfile holds the string denoting the modelfile field in the database.
	FieldModelfile = "modelfile"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the modelfile in the database.
	Table = "modelfiles"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "modelfiles"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_id"
)

// Columns holds all SQL columns for modelfile fields.
var Columns = []string{
	FieldID,
	FieldTagName,
	FieldModelfile,
	FieldUserId,
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
	// TagNameValidator is a validator for the "tagName" field. It is called by the builders before save.
	TagNameValidator func(string) error
	// DefaultModelfile holds the default value on creation for the "modelfile" field.
	DefaultModelfile string
	// ModelfileValidator is a validator for the "modelfile" field. It is called by the builders before save.
	ModelfileValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Modelfile queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTagName orders the results by the tagName field.
func ByTagName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTagName, opts...).ToFunc()
}

// ByModelfile orders the results by the modelfile field.
func ByModelfile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelfile, opts...).ToFunc()
}

// ByUserId orders the results by the userId field.
func ByUserId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserId, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
