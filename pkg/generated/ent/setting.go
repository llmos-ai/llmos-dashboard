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

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/setting"
)

// Setting is the model entity for the Setting schema.
type Setting struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Default holds the value of the "default" field.
	Default string `json:"default,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// IsActive holds the value of the "isActive" field.
	IsActive bool `json:"isActive,omitempty"`
	// ReadOnly holds the value of the "readOnly" field.
	ReadOnly bool `json:"readOnly,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Setting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case setting.FieldIsActive, setting.FieldReadOnly:
			values[i] = new(sql.NullBool)
		case setting.FieldID:
			values[i] = new(sql.NullInt64)
		case setting.FieldName, setting.FieldDefault, setting.FieldValue:
			values[i] = new(sql.NullString)
		case setting.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Setting fields.
func (s *Setting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case setting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case setting.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case setting.FieldDefault:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default", values[i])
			} else if value.Valid {
				s.Default = value.String
			}
		case setting.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				s.Value = value.String
			}
		case setting.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isActive", values[i])
			} else if value.Valid {
				s.IsActive = value.Bool
			}
		case setting.FieldReadOnly:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field readOnly", values[i])
			} else if value.Valid {
				s.ReadOnly = value.Bool
			}
		case setting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Setting.
// This includes values selected through modifiers, order, etc.
func (s *Setting) GetValue(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Setting.
// Note that you need to call Setting.Unwrap() before calling this method if this Setting
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Setting) Update() *SettingUpdateOne {
	return NewSettingClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Setting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Setting) Unwrap() *Setting {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Setting is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Setting) String() string {
	var builder strings.Builder
	builder.WriteString("Setting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("default=")
	builder.WriteString(s.Default)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(s.Value)
	builder.WriteString(", ")
	builder.WriteString("isActive=")
	builder.WriteString(fmt.Sprintf("%v", s.IsActive))
	builder.WriteString(", ")
	builder.WriteString("readOnly=")
	builder.WriteString(fmt.Sprintf("%v", s.ReadOnly))
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Settings is a parsable slice of Setting.
type Settings []*Setting
