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
	"github.com/google/uuid"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
)

// Modelfile is the model entity for the Modelfile schema.
type Modelfile struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// TagName holds the value of the "tagName" field.
	TagName string `json:"tagName,omitempty"`
	// Modelfile holds the value of the "modelfile" field.
	Modelfile string `json:"modelfile,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId uuid.UUID `json:"userId,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ModelfileQuery when eager-loading is set.
	Edges        ModelfileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ModelfileEdges holds the relations/edges for other nodes in the graph.
type ModelfileEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ModelfileEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Modelfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case modelfile.FieldTagName, modelfile.FieldModelfile:
			values[i] = new(sql.NullString)
		case modelfile.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case modelfile.FieldID, modelfile.FieldUserId:
			values[i] = new(uuid.UUID)
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
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case modelfile.FieldTagName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tagName", values[i])
			} else if value.Valid {
				m.TagName = value.String
			}
		case modelfile.FieldModelfile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field modelfile", values[i])
			} else if value.Valid {
				m.Modelfile = value.String
			}
		case modelfile.FieldUserId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value != nil {
				m.UserId = *value
			}
		case modelfile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
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

// QueryOwner queries the "owner" edge of the Modelfile entity.
func (m *Modelfile) QueryOwner() *UserQuery {
	return NewModelfileClient(m.config).QueryOwner(m)
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
	builder.WriteString("tagName=")
	builder.WriteString(m.TagName)
	builder.WriteString(", ")
	builder.WriteString("modelfile=")
	builder.WriteString(m.Modelfile)
	builder.WriteString(", ")
	builder.WriteString("userId=")
	builder.WriteString(fmt.Sprintf("%v", m.UserId))
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Modelfiles is a parsable slice of Modelfile.
type Modelfiles []*Modelfile
