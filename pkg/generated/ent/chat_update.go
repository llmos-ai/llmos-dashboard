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
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/chat"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
)

// ChatUpdate is the builder for updating Chat entities.
type ChatUpdate struct {
	config
	hooks    []Hook
	mutation *ChatMutation
}

// Where appends a list predicates to the ChatUpdate builder.
func (cu *ChatUpdate) Where(ps ...predicate.Chat) *ChatUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ChatUpdate) SetTitle(s string) *ChatUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableTitle(s *string) *ChatUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *ChatUpdate) SetUserID(u uuid.UUID) *ChatUpdate {
	cu.mutation.SetUserID(u)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableUserID(u *uuid.UUID) *ChatUpdate {
	if u != nil {
		cu.SetUserID(*u)
	}
	return cu
}

// SetChat sets the "chat" field.
func (cu *ChatUpdate) SetChat(s string) *ChatUpdate {
	cu.mutation.SetChat(s)
	return cu
}

// SetNillableChat sets the "chat" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableChat(s *string) *ChatUpdate {
	if s != nil {
		cu.SetChat(*s)
	}
	return cu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cu *ChatUpdate) SetOwnerID(id uuid.UUID) *ChatUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *ChatUpdate) SetOwner(u *User) *ChatUpdate {
	return cu.SetOwnerID(u.ID)
}

// Mutation returns the ChatMutation object of the builder.
func (cu *ChatUpdate) Mutation() *ChatMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *ChatUpdate) ClearOwner() *ChatUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChatUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChatUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChatUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChatUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ChatUpdate) check() error {
	if v, ok := cu.mutation.Title(); ok {
		if err := chat.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Chat.title": %w`, err)}
		}
	}
	if _, ok := cu.mutation.OwnerID(); cu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Chat.owner"`)
	}
	return nil
}

func (cu *ChatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(chat.Table, chat.Columns, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(chat.FieldTitle, field.TypeString, value)
	}
	if value, ok := cu.mutation.Chat(); ok {
		_spec.SetField(chat.FieldChat, field.TypeString, value)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.OwnerTable,
			Columns: []string{chat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.OwnerTable,
			Columns: []string{chat.OwnerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChatUpdateOne is the builder for updating a single Chat entity.
type ChatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatMutation
}

// SetTitle sets the "title" field.
func (cuo *ChatUpdateOne) SetTitle(s string) *ChatUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableTitle(s *string) *ChatUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *ChatUpdateOne) SetUserID(u uuid.UUID) *ChatUpdateOne {
	cuo.mutation.SetUserID(u)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableUserID(u *uuid.UUID) *ChatUpdateOne {
	if u != nil {
		cuo.SetUserID(*u)
	}
	return cuo
}

// SetChat sets the "chat" field.
func (cuo *ChatUpdateOne) SetChat(s string) *ChatUpdateOne {
	cuo.mutation.SetChat(s)
	return cuo
}

// SetNillableChat sets the "chat" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableChat(s *string) *ChatUpdateOne {
	if s != nil {
		cuo.SetChat(*s)
	}
	return cuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cuo *ChatUpdateOne) SetOwnerID(id uuid.UUID) *ChatUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *ChatUpdateOne) SetOwner(u *User) *ChatUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// Mutation returns the ChatMutation object of the builder.
func (cuo *ChatUpdateOne) Mutation() *ChatMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *ChatUpdateOne) ClearOwner() *ChatUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// Where appends a list predicates to the ChatUpdate builder.
func (cuo *ChatUpdateOne) Where(ps ...predicate.Chat) *ChatUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChatUpdateOne) Select(field string, fields ...string) *ChatUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chat entity.
func (cuo *ChatUpdateOne) Save(ctx context.Context) (*Chat, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChatUpdateOne) SaveX(ctx context.Context) *Chat {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChatUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChatUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ChatUpdateOne) check() error {
	if v, ok := cuo.mutation.Title(); ok {
		if err := chat.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Chat.title": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.OwnerID(); cuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Chat.owner"`)
	}
	return nil
}

func (cuo *ChatUpdateOne) sqlSave(ctx context.Context) (_node *Chat, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(chat.Table, chat.Columns, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chat.FieldID)
		for _, f := range fields {
			if !chat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(chat.FieldTitle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Chat(); ok {
		_spec.SetField(chat.FieldChat, field.TypeString, value)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.OwnerTable,
			Columns: []string{chat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.OwnerTable,
			Columns: []string{chat.OwnerColumn},
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
	_node = &Chat{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
