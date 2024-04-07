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
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
)

// ModelfileQuery is the builder for querying Modelfile entities.
type ModelfileQuery struct {
	config
	ctx        *QueryContext
	order      []modelfile.OrderOption
	inters     []Interceptor
	predicates []predicate.Modelfile
	withOwner  *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ModelfileQuery builder.
func (mq *ModelfileQuery) Where(ps ...predicate.Modelfile) *ModelfileQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *ModelfileQuery) Limit(limit int) *ModelfileQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *ModelfileQuery) Offset(offset int) *ModelfileQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *ModelfileQuery) Unique(unique bool) *ModelfileQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *ModelfileQuery) Order(o ...modelfile.OrderOption) *ModelfileQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryOwner chains the current query on the "owner" edge.
func (mq *ModelfileQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(modelfile.Table, modelfile.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, modelfile.OwnerTable, modelfile.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Modelfile entity from the query.
// Returns a *NotFoundError when no Modelfile was found.
func (mq *ModelfileQuery) First(ctx context.Context) (*Modelfile, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{modelfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *ModelfileQuery) FirstX(ctx context.Context) *Modelfile {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Modelfile ID from the query.
// Returns a *NotFoundError when no Modelfile ID was found.
func (mq *ModelfileQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{modelfile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *ModelfileQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Modelfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Modelfile entity is found.
// Returns a *NotFoundError when no Modelfile entities are found.
func (mq *ModelfileQuery) Only(ctx context.Context) (*Modelfile, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{modelfile.Label}
	default:
		return nil, &NotSingularError{modelfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *ModelfileQuery) OnlyX(ctx context.Context) *Modelfile {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Modelfile ID in the query.
// Returns a *NotSingularError when more than one Modelfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *ModelfileQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{modelfile.Label}
	default:
		err = &NotSingularError{modelfile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *ModelfileQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Modelfiles.
func (mq *ModelfileQuery) All(ctx context.Context) ([]*Modelfile, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Modelfile, *ModelfileQuery]()
	return withInterceptors[[]*Modelfile](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *ModelfileQuery) AllX(ctx context.Context) []*Modelfile {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Modelfile IDs.
func (mq *ModelfileQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(modelfile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *ModelfileQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *ModelfileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*ModelfileQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *ModelfileQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *ModelfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *ModelfileQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ModelfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *ModelfileQuery) Clone() *ModelfileQuery {
	if mq == nil {
		return nil
	}
	return &ModelfileQuery{
		config:     mq.config,
		ctx:        mq.ctx.Clone(),
		order:      append([]modelfile.OrderOption{}, mq.order...),
		inters:     append([]Interceptor{}, mq.inters...),
		predicates: append([]predicate.Modelfile{}, mq.predicates...),
		withOwner:  mq.withOwner.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *ModelfileQuery) WithOwner(opts ...func(*UserQuery)) *ModelfileQuery {
	query := (&UserClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withOwner = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TagName string `json:"tagName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Modelfile.Query().
//		GroupBy(modelfile.FieldTagName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *ModelfileQuery) GroupBy(field string, fields ...string) *ModelfileGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ModelfileGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = modelfile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TagName string `json:"tagName,omitempty"`
//	}
//
//	client.Modelfile.Query().
//		Select(modelfile.FieldTagName).
//		Scan(ctx, &v)
func (mq *ModelfileQuery) Select(fields ...string) *ModelfileSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &ModelfileSelect{ModelfileQuery: mq}
	sbuild.label = modelfile.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ModelfileSelect configured with the given aggregations.
func (mq *ModelfileQuery) Aggregate(fns ...AggregateFunc) *ModelfileSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *ModelfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !modelfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *ModelfileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Modelfile, error) {
	var (
		nodes       = []*Modelfile{}
		_spec       = mq.querySpec()
		loadedTypes = [1]bool{
			mq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Modelfile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Modelfile{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withOwner; query != nil {
		if err := mq.loadOwner(ctx, query, nodes, nil,
			func(n *Modelfile, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *ModelfileQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Modelfile, init func(*Modelfile), assign func(*Modelfile, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Modelfile)
	for i := range nodes {
		fk := nodes[i].UserId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "userId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mq *ModelfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *ModelfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(modelfile.Table, modelfile.Columns, sqlgraph.NewFieldSpec(modelfile.FieldID, field.TypeUUID))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, modelfile.FieldID)
		for i := range fields {
			if fields[i] != modelfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mq.withOwner != nil {
			_spec.Node.AddColumnOnce(modelfile.FieldUserId)
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *ModelfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(modelfile.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = modelfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ModelfileGroupBy is the group-by builder for Modelfile entities.
type ModelfileGroupBy struct {
	selector
	build *ModelfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *ModelfileGroupBy) Aggregate(fns ...AggregateFunc) *ModelfileGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *ModelfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ModelfileQuery, *ModelfileGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *ModelfileGroupBy) sqlScan(ctx context.Context, root *ModelfileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ModelfileSelect is the builder for selecting fields of Modelfile entities.
type ModelfileSelect struct {
	*ModelfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *ModelfileSelect) Aggregate(fns ...AggregateFunc) *ModelfileSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *ModelfileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ModelfileQuery, *ModelfileSelect](ctx, ms.ModelfileQuery, ms, ms.inters, v)
}

func (ms *ModelfileSelect) sqlScan(ctx context.Context, root *ModelfileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
