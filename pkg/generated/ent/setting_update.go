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
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/setting"
)

// SettingUpdate is the builder for updating Setting entities.
type SettingUpdate struct {
	config
	hooks    []Hook
	mutation *SettingMutation
}

// Where appends a list predicates to the SettingUpdate builder.
func (su *SettingUpdate) Where(ps ...predicate.Setting) *SettingUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SettingUpdate) SetName(s string) *SettingUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SettingUpdate) SetNillableName(s *string) *SettingUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetDefault sets the "default" field.
func (su *SettingUpdate) SetDefault(s string) *SettingUpdate {
	su.mutation.SetDefault(s)
	return su
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (su *SettingUpdate) SetNillableDefault(s *string) *SettingUpdate {
	if s != nil {
		su.SetDefault(*s)
	}
	return su
}

// ClearDefault clears the value of the "default" field.
func (su *SettingUpdate) ClearDefault() *SettingUpdate {
	su.mutation.ClearDefault()
	return su
}

// SetValue sets the "value" field.
func (su *SettingUpdate) SetValue(s string) *SettingUpdate {
	su.mutation.SetValue(s)
	return su
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (su *SettingUpdate) SetNillableValue(s *string) *SettingUpdate {
	if s != nil {
		su.SetValue(*s)
	}
	return su
}

// ClearValue clears the value of the "value" field.
func (su *SettingUpdate) ClearValue() *SettingUpdate {
	su.mutation.ClearValue()
	return su
}

// SetIsActive sets the "isActive" field.
func (su *SettingUpdate) SetIsActive(b bool) *SettingUpdate {
	su.mutation.SetIsActive(b)
	return su
}

// SetNillableIsActive sets the "isActive" field if the given value is not nil.
func (su *SettingUpdate) SetNillableIsActive(b *bool) *SettingUpdate {
	if b != nil {
		su.SetIsActive(*b)
	}
	return su
}

// SetReadOnly sets the "readOnly" field.
func (su *SettingUpdate) SetReadOnly(b bool) *SettingUpdate {
	su.mutation.SetReadOnly(b)
	return su
}

// SetNillableReadOnly sets the "readOnly" field if the given value is not nil.
func (su *SettingUpdate) SetNillableReadOnly(b *bool) *SettingUpdate {
	if b != nil {
		su.SetReadOnly(*b)
	}
	return su
}

// Mutation returns the SettingMutation object of the builder.
func (su *SettingUpdate) Mutation() *SettingMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SettingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SettingUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SettingUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SettingUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SettingUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := setting.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Setting.name": %w`, err)}
		}
	}
	return nil
}

func (su *SettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(setting.Table, setting.Columns, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(setting.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Default(); ok {
		_spec.SetField(setting.FieldDefault, field.TypeString, value)
	}
	if su.mutation.DefaultCleared() {
		_spec.ClearField(setting.FieldDefault, field.TypeString)
	}
	if value, ok := su.mutation.Value(); ok {
		_spec.SetField(setting.FieldValue, field.TypeString, value)
	}
	if su.mutation.ValueCleared() {
		_spec.ClearField(setting.FieldValue, field.TypeString)
	}
	if value, ok := su.mutation.IsActive(); ok {
		_spec.SetField(setting.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := su.mutation.ReadOnly(); ok {
		_spec.SetField(setting.FieldReadOnly, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SettingUpdateOne is the builder for updating a single Setting entity.
type SettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SettingMutation
}

// SetName sets the "name" field.
func (suo *SettingUpdateOne) SetName(s string) *SettingUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableName(s *string) *SettingUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetDefault sets the "default" field.
func (suo *SettingUpdateOne) SetDefault(s string) *SettingUpdateOne {
	suo.mutation.SetDefault(s)
	return suo
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableDefault(s *string) *SettingUpdateOne {
	if s != nil {
		suo.SetDefault(*s)
	}
	return suo
}

// ClearDefault clears the value of the "default" field.
func (suo *SettingUpdateOne) ClearDefault() *SettingUpdateOne {
	suo.mutation.ClearDefault()
	return suo
}

// SetValue sets the "value" field.
func (suo *SettingUpdateOne) SetValue(s string) *SettingUpdateOne {
	suo.mutation.SetValue(s)
	return suo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableValue(s *string) *SettingUpdateOne {
	if s != nil {
		suo.SetValue(*s)
	}
	return suo
}

// ClearValue clears the value of the "value" field.
func (suo *SettingUpdateOne) ClearValue() *SettingUpdateOne {
	suo.mutation.ClearValue()
	return suo
}

// SetIsActive sets the "isActive" field.
func (suo *SettingUpdateOne) SetIsActive(b bool) *SettingUpdateOne {
	suo.mutation.SetIsActive(b)
	return suo
}

// SetNillableIsActive sets the "isActive" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableIsActive(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetIsActive(*b)
	}
	return suo
}

// SetReadOnly sets the "readOnly" field.
func (suo *SettingUpdateOne) SetReadOnly(b bool) *SettingUpdateOne {
	suo.mutation.SetReadOnly(b)
	return suo
}

// SetNillableReadOnly sets the "readOnly" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableReadOnly(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetReadOnly(*b)
	}
	return suo
}

// Mutation returns the SettingMutation object of the builder.
func (suo *SettingUpdateOne) Mutation() *SettingMutation {
	return suo.mutation
}

// Where appends a list predicates to the SettingUpdate builder.
func (suo *SettingUpdateOne) Where(ps ...predicate.Setting) *SettingUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SettingUpdateOne) Select(field string, fields ...string) *SettingUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Setting entity.
func (suo *SettingUpdateOne) Save(ctx context.Context) (*Setting, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SettingUpdateOne) SaveX(ctx context.Context) *Setting {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SettingUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SettingUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SettingUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := setting.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Setting.name": %w`, err)}
		}
	}
	return nil
}

func (suo *SettingUpdateOne) sqlSave(ctx context.Context) (_node *Setting, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(setting.Table, setting.Columns, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Setting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, setting.FieldID)
		for _, f := range fields {
			if !setting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != setting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(setting.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Default(); ok {
		_spec.SetField(setting.FieldDefault, field.TypeString, value)
	}
	if suo.mutation.DefaultCleared() {
		_spec.ClearField(setting.FieldDefault, field.TypeString)
	}
	if value, ok := suo.mutation.Value(); ok {
		_spec.SetField(setting.FieldValue, field.TypeString, value)
	}
	if suo.mutation.ValueCleared() {
		_spec.ClearField(setting.FieldValue, field.TypeString)
	}
	if value, ok := suo.mutation.IsActive(); ok {
		_spec.SetField(setting.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := suo.mutation.ReadOnly(); ok {
		_spec.SetField(setting.FieldReadOnly, field.TypeBool, value)
	}
	_node = &Setting{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
