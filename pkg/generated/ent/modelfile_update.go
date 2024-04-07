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
	"github.com/google/uuid"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
)

// ModelfileUpdate is the builder for updating Modelfile entities.
type ModelfileUpdate struct {
	config
	hooks    []Hook
	mutation *ModelfileMutation
}

// Where appends a list predicates to the ModelfileUpdate builder.
func (mu *ModelfileUpdate) Where(ps ...predicate.Modelfile) *ModelfileUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetTagName sets the "tagName" field.
func (mu *ModelfileUpdate) SetTagName(s string) *ModelfileUpdate {
	mu.mutation.SetTagName(s)
	return mu
}

// SetNillableTagName sets the "tagName" field if the given value is not nil.
func (mu *ModelfileUpdate) SetNillableTagName(s *string) *ModelfileUpdate {
	if s != nil {
		mu.SetTagName(*s)
	}
	return mu
}

// SetModelfile sets the "modelfile" field.
func (mu *ModelfileUpdate) SetModelfile(s string) *ModelfileUpdate {
	mu.mutation.SetModelfile(s)
	return mu
}

// SetNillableModelfile sets the "modelfile" field if the given value is not nil.
func (mu *ModelfileUpdate) SetNillableModelfile(s *string) *ModelfileUpdate {
	if s != nil {
		mu.SetModelfile(*s)
	}
	return mu
}

// SetUserId sets the "userId" field.
func (mu *ModelfileUpdate) SetUserId(u uuid.UUID) *ModelfileUpdate {
	mu.mutation.SetUserId(u)
	return mu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (mu *ModelfileUpdate) SetNillableUserId(u *uuid.UUID) *ModelfileUpdate {
	if u != nil {
		mu.SetUserId(*u)
	}
	return mu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (mu *ModelfileUpdate) SetOwnerID(id uuid.UUID) *ModelfileUpdate {
	mu.mutation.SetOwnerID(id)
	return mu
}

// SetOwner sets the "owner" edge to the User entity.
func (mu *ModelfileUpdate) SetOwner(u *User) *ModelfileUpdate {
	return mu.SetOwnerID(u.ID)
}

// Mutation returns the ModelfileMutation object of the builder.
func (mu *ModelfileUpdate) Mutation() *ModelfileMutation {
	return mu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (mu *ModelfileUpdate) ClearOwner() *ModelfileUpdate {
	mu.mutation.ClearOwner()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *ModelfileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ModelfileUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ModelfileUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ModelfileUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *ModelfileUpdate) check() error {
	if v, ok := mu.mutation.TagName(); ok {
		if err := modelfile.TagNameValidator(v); err != nil {
			return &ValidationError{Name: "tagName", err: fmt.Errorf(`ent: validator failed for field "Modelfile.tagName": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Modelfile(); ok {
		if err := modelfile.ModelfileValidator(v); err != nil {
			return &ValidationError{Name: "modelfile", err: fmt.Errorf(`ent: validator failed for field "Modelfile.modelfile": %w`, err)}
		}
	}
	if _, ok := mu.mutation.OwnerID(); mu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Modelfile.owner"`)
	}
	return nil
}

func (mu *ModelfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(modelfile.Table, modelfile.Columns, sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.TagName(); ok {
		_spec.SetField(modelfile.FieldTagName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Modelfile(); ok {
		_spec.SetField(modelfile.FieldModelfile, field.TypeString, value)
	}
	if mu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelfile.OwnerTable,
			Columns: []string{modelfile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelfile.OwnerTable,
			Columns: []string{modelfile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{modelfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// ModelfileUpdateOne is the builder for updating a single Modelfile entity.
type ModelfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ModelfileMutation
}

// SetTagName sets the "tagName" field.
func (muo *ModelfileUpdateOne) SetTagName(s string) *ModelfileUpdateOne {
	muo.mutation.SetTagName(s)
	return muo
}

// SetNillableTagName sets the "tagName" field if the given value is not nil.
func (muo *ModelfileUpdateOne) SetNillableTagName(s *string) *ModelfileUpdateOne {
	if s != nil {
		muo.SetTagName(*s)
	}
	return muo
}

// SetModelfile sets the "modelfile" field.
func (muo *ModelfileUpdateOne) SetModelfile(s string) *ModelfileUpdateOne {
	muo.mutation.SetModelfile(s)
	return muo
}

// SetNillableModelfile sets the "modelfile" field if the given value is not nil.
func (muo *ModelfileUpdateOne) SetNillableModelfile(s *string) *ModelfileUpdateOne {
	if s != nil {
		muo.SetModelfile(*s)
	}
	return muo
}

// SetUserId sets the "userId" field.
func (muo *ModelfileUpdateOne) SetUserId(u uuid.UUID) *ModelfileUpdateOne {
	muo.mutation.SetUserId(u)
	return muo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (muo *ModelfileUpdateOne) SetNillableUserId(u *uuid.UUID) *ModelfileUpdateOne {
	if u != nil {
		muo.SetUserId(*u)
	}
	return muo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (muo *ModelfileUpdateOne) SetOwnerID(id uuid.UUID) *ModelfileUpdateOne {
	muo.mutation.SetOwnerID(id)
	return muo
}

// SetOwner sets the "owner" edge to the User entity.
func (muo *ModelfileUpdateOne) SetOwner(u *User) *ModelfileUpdateOne {
	return muo.SetOwnerID(u.ID)
}

// Mutation returns the ModelfileMutation object of the builder.
func (muo *ModelfileUpdateOne) Mutation() *ModelfileMutation {
	return muo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (muo *ModelfileUpdateOne) ClearOwner() *ModelfileUpdateOne {
	muo.mutation.ClearOwner()
	return muo
}

// Where appends a list predicates to the ModelfileUpdate builder.
func (muo *ModelfileUpdateOne) Where(ps ...predicate.Modelfile) *ModelfileUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *ModelfileUpdateOne) Select(field string, fields ...string) *ModelfileUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Modelfile entity.
func (muo *ModelfileUpdateOne) Save(ctx context.Context) (*Modelfile, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ModelfileUpdateOne) SaveX(ctx context.Context) *Modelfile {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *ModelfileUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ModelfileUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *ModelfileUpdateOne) check() error {
	if v, ok := muo.mutation.TagName(); ok {
		if err := modelfile.TagNameValidator(v); err != nil {
			return &ValidationError{Name: "tagName", err: fmt.Errorf(`ent: validator failed for field "Modelfile.tagName": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Modelfile(); ok {
		if err := modelfile.ModelfileValidator(v); err != nil {
			return &ValidationError{Name: "modelfile", err: fmt.Errorf(`ent: validator failed for field "Modelfile.modelfile": %w`, err)}
		}
	}
	if _, ok := muo.mutation.OwnerID(); muo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Modelfile.owner"`)
	}
	return nil
}

func (muo *ModelfileUpdateOne) sqlSave(ctx context.Context) (_node *Modelfile, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(modelfile.Table, modelfile.Columns, sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Modelfile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, modelfile.FieldID)
		for _, f := range fields {
			if !modelfile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != modelfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.TagName(); ok {
		_spec.SetField(modelfile.FieldTagName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Modelfile(); ok {
		_spec.SetField(modelfile.FieldModelfile, field.TypeString, value)
	}
	if muo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelfile.OwnerTable,
			Columns: []string{modelfile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelfile.OwnerTable,
			Columns: []string{modelfile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Modelfile{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{modelfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
