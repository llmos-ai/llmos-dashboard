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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/llmos/llmos-dashboard/pkg/generated/ent/modelfile"
)

// ModelfileCreate is the builder for creating a Modelfile entity.
type ModelfileCreate struct {
	config
	mutation *ModelfileMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (mc *ModelfileCreate) SetUserID(i int) *ModelfileCreate {
	mc.mutation.SetUserID(i)
	return mc
}

// SetTagName sets the "tag_name" field.
func (mc *ModelfileCreate) SetTagName(s string) *ModelfileCreate {
	mc.mutation.SetTagName(s)
	return mc
}

// SetModelfile sets the "modelfile" field.
func (mc *ModelfileCreate) SetModelfile(s string) *ModelfileCreate {
	mc.mutation.SetModelfile(s)
	return mc
}

// SetNillableModelfile sets the "modelfile" field if the given value is not nil.
func (mc *ModelfileCreate) SetNillableModelfile(s *string) *ModelfileCreate {
	if s != nil {
		mc.SetModelfile(*s)
	}
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *ModelfileCreate) SetCreatedAt(t time.Time) *ModelfileCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *ModelfileCreate) SetNillableCreatedAt(t *time.Time) *ModelfileCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// Mutation returns the ModelfileMutation object of the builder.
func (mc *ModelfileCreate) Mutation() *ModelfileMutation {
	return mc.mutation
}

// Save creates the Modelfile in the database.
func (mc *ModelfileCreate) Save(ctx context.Context) (*Modelfile, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *ModelfileCreate) SaveX(ctx context.Context) *Modelfile {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *ModelfileCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *ModelfileCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *ModelfileCreate) defaults() {
	if _, ok := mc.mutation.Modelfile(); !ok {
		v := modelfile.DefaultModelfile
		mc.mutation.SetModelfile(v)
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := modelfile.DefaultCreatedAt
		mc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *ModelfileCreate) check() error {
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Modelfile.user_id"`)}
	}
	if v, ok := mc.mutation.UserID(); ok {
		if err := modelfile.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Modelfile.user_id": %w`, err)}
		}
	}
	if _, ok := mc.mutation.TagName(); !ok {
		return &ValidationError{Name: "tag_name", err: errors.New(`ent: missing required field "Modelfile.tag_name"`)}
	}
	if v, ok := mc.mutation.TagName(); ok {
		if err := modelfile.TagNameValidator(v); err != nil {
			return &ValidationError{Name: "tag_name", err: fmt.Errorf(`ent: validator failed for field "Modelfile.tag_name": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Modelfile(); !ok {
		return &ValidationError{Name: "modelfile", err: errors.New(`ent: missing required field "Modelfile.modelfile"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Modelfile.created_at"`)}
	}
	return nil
}

func (mc *ModelfileCreate) sqlSave(ctx context.Context) (*Modelfile, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *ModelfileCreate) createSpec() (*Modelfile, *sqlgraph.CreateSpec) {
	var (
		_node = &Modelfile{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(modelfile.Table, sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.UserID(); ok {
		_spec.SetField(modelfile.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := mc.mutation.TagName(); ok {
		_spec.SetField(modelfile.FieldTagName, field.TypeString, value)
		_node.TagName = value
	}
	if value, ok := mc.mutation.Modelfile(); ok {
		_spec.SetField(modelfile.FieldModelfile, field.TypeString, value)
		_node.Modelfile = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(modelfile.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// ModelfileCreateBulk is the builder for creating many Modelfile entities in bulk.
type ModelfileCreateBulk struct {
	config
	err      error
	builders []*ModelfileCreate
}

// Save creates the Modelfile entities in the database.
func (mcb *ModelfileCreateBulk) Save(ctx context.Context) ([]*Modelfile, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Modelfile, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ModelfileMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *ModelfileCreateBulk) SaveX(ctx context.Context) []*Modelfile {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *ModelfileCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *ModelfileCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
