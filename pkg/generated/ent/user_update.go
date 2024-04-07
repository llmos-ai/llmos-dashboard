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
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(u user.Role) *UserUpdate {
	uu.mutation.SetRole(u)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(u *user.Role) *UserUpdate {
	if u != nil {
		uu.SetRole(*u)
	}
	return uu
}

// SetProfileImageUrl sets the "profileImageUrl" field.
func (uu *UserUpdate) SetProfileImageUrl(s string) *UserUpdate {
	uu.mutation.SetProfileImageUrl(s)
	return uu
}

// SetNillableProfileImageUrl sets the "profileImageUrl" field if the given value is not nil.
func (uu *UserUpdate) SetNillableProfileImageUrl(s *string) *UserUpdate {
	if s != nil {
		uu.SetProfileImageUrl(*s)
	}
	return uu
}

// AddChatIDs adds the "chats" edge to the Chat entity by IDs.
func (uu *UserUpdate) AddChatIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddChatIDs(ids...)
	return uu
}

// AddChats adds the "chats" edges to the Chat entity.
func (uu *UserUpdate) AddChats(c ...*Chat) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddChatIDs(ids...)
}

// AddModelfileIDs adds the "modelfiles" edge to the Modelfile entity by IDs.
func (uu *UserUpdate) AddModelfileIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddModelfileIDs(ids...)
	return uu
}

// AddModelfiles adds the "modelfiles" edges to the Modelfile entity.
func (uu *UserUpdate) AddModelfiles(m ...*Modelfile) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.AddModelfileIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearChats clears all "chats" edges to the Chat entity.
func (uu *UserUpdate) ClearChats() *UserUpdate {
	uu.mutation.ClearChats()
	return uu
}

// RemoveChatIDs removes the "chats" edge to Chat entities by IDs.
func (uu *UserUpdate) RemoveChatIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveChatIDs(ids...)
	return uu
}

// RemoveChats removes "chats" edges to Chat entities.
func (uu *UserUpdate) RemoveChats(c ...*Chat) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveChatIDs(ids...)
}

// ClearModelfiles clears all "modelfiles" edges to the Modelfile entity.
func (uu *UserUpdate) ClearModelfiles() *UserUpdate {
	uu.mutation.ClearModelfiles()
	return uu
}

// RemoveModelfileIDs removes the "modelfiles" edge to Modelfile entities by IDs.
func (uu *UserUpdate) RemoveModelfileIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveModelfileIDs(ids...)
	return uu
}

// RemoveModelfiles removes "modelfiles" edges to Modelfile entities.
func (uu *UserUpdate) RemoveModelfiles(m ...*Modelfile) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.RemoveModelfileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.ProfileImageUrl(); ok {
		_spec.SetField(user.FieldProfileImageUrl, field.TypeString, value)
	}
	if uu.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedChatsIDs(); len(nodes) > 0 && !uu.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ModelfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ModelfilesTable,
			Columns: []string{user.ModelfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedModelfilesIDs(); len(nodes) > 0 && !uu.mutation.ModelfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ModelfilesTable,
			Columns: []string{user.ModelfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ModelfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ModelfilesTable,
			Columns: []string{user.ModelfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(u user.Role) *UserUpdateOne {
	uuo.mutation.SetRole(u)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(u *user.Role) *UserUpdateOne {
	if u != nil {
		uuo.SetRole(*u)
	}
	return uuo
}

// SetProfileImageUrl sets the "profileImageUrl" field.
func (uuo *UserUpdateOne) SetProfileImageUrl(s string) *UserUpdateOne {
	uuo.mutation.SetProfileImageUrl(s)
	return uuo
}

// SetNillableProfileImageUrl sets the "profileImageUrl" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProfileImageUrl(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetProfileImageUrl(*s)
	}
	return uuo
}

// AddChatIDs adds the "chats" edge to the Chat entity by IDs.
func (uuo *UserUpdateOne) AddChatIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddChatIDs(ids...)
	return uuo
}

// AddChats adds the "chats" edges to the Chat entity.
func (uuo *UserUpdateOne) AddChats(c ...*Chat) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddChatIDs(ids...)
}

// AddModelfileIDs adds the "modelfiles" edge to the Modelfile entity by IDs.
func (uuo *UserUpdateOne) AddModelfileIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddModelfileIDs(ids...)
	return uuo
}

// AddModelfiles adds the "modelfiles" edges to the Modelfile entity.
func (uuo *UserUpdateOne) AddModelfiles(m ...*Modelfile) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.AddModelfileIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearChats clears all "chats" edges to the Chat entity.
func (uuo *UserUpdateOne) ClearChats() *UserUpdateOne {
	uuo.mutation.ClearChats()
	return uuo
}

// RemoveChatIDs removes the "chats" edge to Chat entities by IDs.
func (uuo *UserUpdateOne) RemoveChatIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveChatIDs(ids...)
	return uuo
}

// RemoveChats removes "chats" edges to Chat entities.
func (uuo *UserUpdateOne) RemoveChats(c ...*Chat) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveChatIDs(ids...)
}

// ClearModelfiles clears all "modelfiles" edges to the Modelfile entity.
func (uuo *UserUpdateOne) ClearModelfiles() *UserUpdateOne {
	uuo.mutation.ClearModelfiles()
	return uuo
}

// RemoveModelfileIDs removes the "modelfiles" edge to Modelfile entities by IDs.
func (uuo *UserUpdateOne) RemoveModelfileIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveModelfileIDs(ids...)
	return uuo
}

// RemoveModelfiles removes "modelfiles" edges to Modelfile entities.
func (uuo *UserUpdateOne) RemoveModelfiles(m ...*Modelfile) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.RemoveModelfileIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.ProfileImageUrl(); ok {
		_spec.SetField(user.FieldProfileImageUrl, field.TypeString, value)
	}
	if uuo.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedChatsIDs(); len(nodes) > 0 && !uuo.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ModelfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ModelfilesTable,
			Columns: []string{user.ModelfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedModelfilesIDs(); len(nodes) > 0 && !uuo.mutation.ModelfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ModelfilesTable,
			Columns: []string{user.ModelfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ModelfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ModelfilesTable,
			Columns: []string{user.ModelfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
