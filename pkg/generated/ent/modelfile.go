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
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
)

// Modelfile is the model entity for the Modelfile schema.
type Modelfile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// TagName holds the value of the "tag_name" field.
	TagName string `json:"tag_name,omitempty"`
	// Modelfile holds the value of the "modelfile" field.
	Modelfile string `json:"modelfile,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Modelfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case modelfile.FieldID, modelfile.FieldUserID:
			values[i] = new(sql.NullInt64)
		case modelfile.FieldTagName, modelfile.FieldModelfile:
			values[i] = new(sql.NullString)
		case modelfile.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Modelfile fields.
func (m *Modelfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case modelfile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case modelfile.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				m.UserID = int(value.Int64)
			}
		case modelfile.FieldTagName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag_name", values[i])
			} else if value.Valid {
				m.TagName = value.String
			}
		case modelfile.FieldModelfile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field modelfile", values[i])
			} else if value.Valid {
				m.Modelfile = value.String
			}
		case modelfile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Modelfile.
// This includes values selected through modifiers, order, etc.
func (m *Modelfile) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Modelfile.
// Note that you need to call Modelfile.Unwrap() before calling this method if this Modelfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Modelfile) Update() *ModelfileUpdateOne {
	return NewModelfileClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Modelfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Modelfile) Unwrap() *Modelfile {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Modelfile is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Modelfile) String() string {
	var builder strings.Builder
	builder.WriteString("Modelfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", m.UserID))
	builder.WriteString(", ")
	builder.WriteString("tag_name=")
	builder.WriteString(m.TagName)
	builder.WriteString(", ")
	builder.WriteString("modelfile=")
	builder.WriteString(m.Modelfile)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Modelfiles is a parsable slice of Modelfile.
type Modelfiles []*Modelfile
