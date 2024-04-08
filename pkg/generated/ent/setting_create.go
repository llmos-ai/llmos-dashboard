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
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/setting"
)

// SettingCreate is the builder for creating a Setting entity.
type SettingCreate struct {
	config
	mutation *SettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (sc *SettingCreate) SetName(s string) *SettingCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDefault sets the "default" field.
func (sc *SettingCreate) SetDefault(s string) *SettingCreate {
	sc.mutation.SetDefault(s)
	return sc
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (sc *SettingCreate) SetNillableDefault(s *string) *SettingCreate {
	if s != nil {
		sc.SetDefault(*s)
	}
	return sc
}

// SetValue sets the "value" field.
func (sc *SettingCreate) SetValue(s string) *SettingCreate {
	sc.mutation.SetValue(s)
	return sc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (sc *SettingCreate) SetNillableValue(s *string) *SettingCreate {
	if s != nil {
		sc.SetValue(*s)
	}
	return sc
}

// SetIsActive sets the "isActive" field.
func (sc *SettingCreate) SetIsActive(b bool) *SettingCreate {
	sc.mutation.SetIsActive(b)
	return sc
}

// SetNillableIsActive sets the "isActive" field if the given value is not nil.
func (sc *SettingCreate) SetNillableIsActive(b *bool) *SettingCreate {
	if b != nil {
		sc.SetIsActive(*b)
	}
	return sc
}

// SetReadOnly sets the "readOnly" field.
func (sc *SettingCreate) SetReadOnly(b bool) *SettingCreate {
	sc.mutation.SetReadOnly(b)
	return sc
}

// SetNillableReadOnly sets the "readOnly" field if the given value is not nil.
func (sc *SettingCreate) SetNillableReadOnly(b *bool) *SettingCreate {
	if b != nil {
		sc.SetReadOnly(*b)
	}
	return sc
}

// SetCreatedAt sets the "createdAt" field.
func (sc *SettingCreate) SetCreatedAt(t time.Time) *SettingCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (sc *SettingCreate) SetNillableCreatedAt(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// Mutation returns the SettingMutation object of the builder.
func (sc *SettingCreate) Mutation() *SettingMutation {
	return sc.mutation
}

// Save creates the Setting in the database.
func (sc *SettingCreate) Save(ctx context.Context) (*Setting, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingCreate) SaveX(ctx context.Context) *Setting {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SettingCreate) defaults() {
	if _, ok := sc.mutation.Default(); !ok {
		v := setting.DefaultDefault
		sc.mutation.SetDefault(v)
	}
	if _, ok := sc.mutation.IsActive(); !ok {
		v := setting.DefaultIsActive
		sc.mutation.SetIsActive(v)
	}
	if _, ok := sc.mutation.ReadOnly(); !ok {
		v := setting.DefaultReadOnly
		sc.mutation.SetReadOnly(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := setting.DefaultCreatedAt
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Setting.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := setting.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Setting.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Default(); !ok {
		return &ValidationError{Name: "default", err: errors.New(`ent: missing required field "Setting.default"`)}
	}
	if v, ok := sc.mutation.Default(); ok {
		if err := setting.DefaultValidator(v); err != nil {
			return &ValidationError{Name: "default", err: fmt.Errorf(`ent: validator failed for field "Setting.default": %w`, err)}
		}
	}
	if _, ok := sc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "isActive", err: errors.New(`ent: missing required field "Setting.isActive"`)}
	}
	if _, ok := sc.mutation.ReadOnly(); !ok {
		return &ValidationError{Name: "readOnly", err: errors.New(`ent: missing required field "Setting.readOnly"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Setting.createdAt"`)}
	}
	return nil
}

func (sc *SettingCreate) sqlSave(ctx context.Context) (*Setting, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettingCreate) createSpec() (*Setting, *sqlgraph.CreateSpec) {
	var (
		_node = &Setting{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(setting.Table, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(setting.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Default(); ok {
		_spec.SetField(setting.FieldDefault, field.TypeString, value)
		_node.Default = value
	}
	if value, ok := sc.mutation.Value(); ok {
		_spec.SetField(setting.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := sc.mutation.IsActive(); ok {
		_spec.SetField(setting.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := sc.mutation.ReadOnly(); ok {
		_spec.SetField(setting.FieldReadOnly, field.TypeBool, value)
		_node.ReadOnly = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(setting.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sc *SettingCreate) OnConflict(opts ...sql.ConflictOption) *SettingUpsertOne {
	sc.conflict = opts
	return &SettingUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SettingCreate) OnConflictColumns(columns ...string) *SettingUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertOne{
		create: sc,
	}
}

type (
	// SettingUpsertOne is the builder for "upsert"-ing
	//  one Setting node.
	SettingUpsertOne struct {
		create *SettingCreate
	}

	// SettingUpsert is the "OnConflict" setter.
	SettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *SettingUpsert) SetName(v string) *SettingUpsert {
	u.Set(setting.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SettingUpsert) UpdateName() *SettingUpsert {
	u.SetExcluded(setting.FieldName)
	return u
}

// SetDefault sets the "default" field.
func (u *SettingUpsert) SetDefault(v string) *SettingUpsert {
	u.Set(setting.FieldDefault, v)
	return u
}

// UpdateDefault sets the "default" field to the value that was provided on create.
func (u *SettingUpsert) UpdateDefault() *SettingUpsert {
	u.SetExcluded(setting.FieldDefault)
	return u
}

// SetValue sets the "value" field.
func (u *SettingUpsert) SetValue(v string) *SettingUpsert {
	u.Set(setting.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SettingUpsert) UpdateValue() *SettingUpsert {
	u.SetExcluded(setting.FieldValue)
	return u
}

// ClearValue clears the value of the "value" field.
func (u *SettingUpsert) ClearValue() *SettingUpsert {
	u.SetNull(setting.FieldValue)
	return u
}

// SetIsActive sets the "isActive" field.
func (u *SettingUpsert) SetIsActive(v bool) *SettingUpsert {
	u.Set(setting.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "isActive" field to the value that was provided on create.
func (u *SettingUpsert) UpdateIsActive() *SettingUpsert {
	u.SetExcluded(setting.FieldIsActive)
	return u
}

// SetReadOnly sets the "readOnly" field.
func (u *SettingUpsert) SetReadOnly(v bool) *SettingUpsert {
	u.Set(setting.FieldReadOnly, v)
	return u
}

// UpdateReadOnly sets the "readOnly" field to the value that was provided on create.
func (u *SettingUpsert) UpdateReadOnly() *SettingUpsert {
	u.SetExcluded(setting.FieldReadOnly)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SettingUpsertOne) UpdateNewValues() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(setting.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SettingUpsertOne) Ignore() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertOne) DoNothing() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreate.OnConflict
// documentation for more info.
func (u *SettingUpsertOne) Update(set func(*SettingUpsert)) *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *SettingUpsertOne) SetName(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateName() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateName()
	})
}

// SetDefault sets the "default" field.
func (u *SettingUpsertOne) SetDefault(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetDefault(v)
	})
}

// UpdateDefault sets the "default" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateDefault() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateDefault()
	})
}

// SetValue sets the "value" field.
func (u *SettingUpsertOne) SetValue(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateValue() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateValue()
	})
}

// ClearValue clears the value of the "value" field.
func (u *SettingUpsertOne) ClearValue() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearValue()
	})
}

// SetIsActive sets the "isActive" field.
func (u *SettingUpsertOne) SetIsActive(v bool) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "isActive" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateIsActive() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateIsActive()
	})
}

// SetReadOnly sets the "readOnly" field.
func (u *SettingUpsertOne) SetReadOnly(v bool) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetReadOnly(v)
	})
}

// UpdateReadOnly sets the "readOnly" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateReadOnly() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateReadOnly()
	})
}

// Exec executes the query.
func (u *SettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SettingUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SettingUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SettingCreateBulk is the builder for creating many Setting entities in bulk.
type SettingCreateBulk struct {
	config
	err      error
	builders []*SettingCreate
	conflict []sql.ConflictOption
}

// Save creates the Setting entities in the database.
func (scb *SettingCreateBulk) Save(ctx context.Context) ([]*Setting, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Setting, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettingCreateBulk) SaveX(ctx context.Context) []*Setting {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (scb *SettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *SettingUpsertBulk {
	scb.conflict = opts
	return &SettingUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SettingCreateBulk) OnConflictColumns(columns ...string) *SettingUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertBulk{
		create: scb,
	}
}

// SettingUpsertBulk is the builder for "upsert"-ing
// a bulk of Setting nodes.
type SettingUpsertBulk struct {
	create *SettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SettingUpsertBulk) UpdateNewValues() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(setting.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SettingUpsertBulk) Ignore() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertBulk) DoNothing() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreateBulk.OnConflict
// documentation for more info.
func (u *SettingUpsertBulk) Update(set func(*SettingUpsert)) *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *SettingUpsertBulk) SetName(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateName() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateName()
	})
}

// SetDefault sets the "default" field.
func (u *SettingUpsertBulk) SetDefault(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetDefault(v)
	})
}

// UpdateDefault sets the "default" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateDefault() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateDefault()
	})
}

// SetValue sets the "value" field.
func (u *SettingUpsertBulk) SetValue(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateValue() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateValue()
	})
}

// ClearValue clears the value of the "value" field.
func (u *SettingUpsertBulk) ClearValue() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearValue()
	})
}

// SetIsActive sets the "isActive" field.
func (u *SettingUpsertBulk) SetIsActive(v bool) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "isActive" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateIsActive() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateIsActive()
	})
}

// SetReadOnly sets the "readOnly" field.
func (u *SettingUpsertBulk) SetReadOnly(v bool) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetReadOnly(v)
	})
}

// UpdateReadOnly sets the "readOnly" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateReadOnly() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateReadOnly()
	})
}

// Exec executes the query.
func (u *SettingUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
