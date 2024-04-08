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

package setting

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the setting type in the database.
	Label = "setting"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDefault holds the string denoting the default field in the database.
	FieldDefault = "default"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldIsActive holds the string denoting the isactive field in the database.
	FieldIsActive = "is_active"
	// FieldReadOnly holds the string denoting the readonly field in the database.
	FieldReadOnly = "read_only"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the setting in the database.
	Table = "settings"
)

// Columns holds all SQL columns for setting fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDefault,
	FieldValue,
	FieldIsActive,
	FieldReadOnly,
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
	// DefaultIsActive holds the default value on creation for the "isActive" field.
	DefaultIsActive bool
	// DefaultReadOnly holds the default value on creation for the "readOnly" field.
	DefaultReadOnly bool
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Setting queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDefault orders the results by the default field.
func ByDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefault, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByIsActive orders the results by the isActive field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByReadOnly orders the results by the readOnly field.
func ByReadOnly(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReadOnly, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
