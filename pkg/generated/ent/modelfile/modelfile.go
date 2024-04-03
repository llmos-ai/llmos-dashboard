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
)

const (
	// Label holds the string label denoting the modelfile type in the database.
	Label = "modelfile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTagName holds the string denoting the tag_name field in the database.
	FieldTagName = "tag_name"
	// FieldModelfile holds the string denoting the modelfile field in the database.
	FieldModelfile = "modelfile"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the modelfile in the database.
	Table = "modelfiles"
)

// Columns holds all SQL columns for modelfile fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldTagName,
	FieldModelfile,
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
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(int) error
	// TagNameValidator is a validator for the "tag_name" field. It is called by the builders before save.
	TagNameValidator func(string) error
	// DefaultModelfile holds the default value on creation for the "modelfile" field.
	DefaultModelfile string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Modelfile queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByTagName orders the results by the tag_name field.
func ByTagName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTagName, opts...).ToFunc()
}

// ByModelfile orders the results by the modelfile field.
func ByModelfile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelfile, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
