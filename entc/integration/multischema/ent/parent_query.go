// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/multischema/ent/internal"
	"entgo.io/ent/entc/integration/multischema/ent/parent"
	"entgo.io/ent/entc/integration/multischema/ent/predicate"
	"entgo.io/ent/entc/integration/multischema/ent/user"
	"entgo.io/ent/schema/field"
)

// ParentQuery is the builder for querying Parent entities.
type ParentQuery struct {
	config
	ctx        *QueryContext
	order      []parent.OrderOption
	inters     []Interceptor
	predicates []predicate.Parent
	withChild  *UserQuery
	withParent *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ParentQuery builder.
func (_q *ParentQuery) Where(ps ...predicate.Parent) *ParentQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *ParentQuery) Limit(limit int) *ParentQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *ParentQuery) Offset(offset int) *ParentQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *ParentQuery) Unique(unique bool) *ParentQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *ParentQuery) Order(o ...parent.OrderOption) *ParentQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryChild chains the current query on the "child" edge.
func (_q *ParentQuery) QueryChild() *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(parent.Table, parent.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, parent.ChildTable, parent.ChildColumn),
		)
		schemaConfig := _q.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Parent
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (_q *ParentQuery) QueryParent() *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(parent.Table, parent.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, parent.ParentTable, parent.ParentColumn),
		)
		schemaConfig := _q.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Parent
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Parent entity from the query.
// Returns a *NotFoundError when no Parent was found.
func (_q *ParentQuery) First(ctx context.Context) (*Parent, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{parent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *ParentQuery) FirstX(ctx context.Context) *Parent {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Parent ID from the query.
// Returns a *NotFoundError when no Parent ID was found.
func (_q *ParentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{parent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *ParentQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Parent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Parent entity is found.
// Returns a *NotFoundError when no Parent entities are found.
func (_q *ParentQuery) Only(ctx context.Context) (*Parent, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{parent.Label}
	default:
		return nil, &NotSingularError{parent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *ParentQuery) OnlyX(ctx context.Context) *Parent {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Parent ID in the query.
// Returns a *NotSingularError when more than one Parent ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *ParentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{parent.Label}
	default:
		err = &NotSingularError{parent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *ParentQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Parents.
func (_q *ParentQuery) All(ctx context.Context) ([]*Parent, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Parent, *ParentQuery]()
	return withInterceptors[[]*Parent](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *ParentQuery) AllX(ctx context.Context) []*Parent {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Parent IDs.
func (_q *ParentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(parent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *ParentQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *ParentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*ParentQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *ParentQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *ParentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *ParentQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ParentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *ParentQuery) Clone() *ParentQuery {
	if _q == nil {
		return nil
	}
	return &ParentQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]parent.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Parent{}, _q.predicates...),
		withChild:  _q.withChild.Clone(),
		withParent: _q.withParent.Clone(),
		// clone intermediate query.
		sql:       _q.sql.Clone(),
		path:      _q.path,
		modifiers: append([]func(*sql.Selector){}, _q.modifiers...),
	}
}

// WithChild tells the query-builder to eager-load the nodes that are connected to
// the "child" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *ParentQuery) WithChild(opts ...func(*UserQuery)) *ParentQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withChild = query
	return _q
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *ParentQuery) WithParent(opts ...func(*UserQuery)) *ParentQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withParent = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ByAdoption bool `json:"by_adoption,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Parent.Query().
//		GroupBy(parent.FieldByAdoption).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *ParentQuery) GroupBy(field string, fields ...string) *ParentGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ParentGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = parent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ByAdoption bool `json:"by_adoption,omitempty"`
//	}
//
//	client.Parent.Query().
//		Select(parent.FieldByAdoption).
//		Scan(ctx, &v)
func (_q *ParentQuery) Select(fields ...string) *ParentSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &ParentSelect{ParentQuery: _q}
	sbuild.label = parent.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ParentSelect configured with the given aggregations.
func (_q *ParentQuery) Aggregate(fns ...AggregateFunc) *ParentSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *ParentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	for _, f := range _q.ctx.Fields {
		if !parent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.sql = prev
	}
	return nil
}

func (_q *ParentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Parent, error) {
	var (
		nodes       = []*Parent{}
		_spec       = _q.querySpec()
		loadedTypes = [2]bool{
			_q.withChild != nil,
			_q.withParent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Parent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Parent{config: _q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = _q.schemaConfig.Parent
	ctx = internal.NewSchemaConfigContext(ctx, _q.schemaConfig)
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, _q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := _q.withChild; query != nil {
		if err := _q.loadChild(ctx, query, nodes, nil,
			func(n *Parent, e *User) { n.Edges.Child = e }); err != nil {
			return nil, err
		}
	}
	if query := _q.withParent; query != nil {
		if err := _q.loadParent(ctx, query, nodes, nil,
			func(n *Parent, e *User) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *ParentQuery) loadChild(ctx context.Context, query *UserQuery, nodes []*Parent, init func(*Parent), assign func(*Parent, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Parent)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (_q *ParentQuery) loadParent(ctx context.Context, query *UserQuery, nodes []*Parent, init func(*Parent), assign func(*Parent, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Parent)
	for i := range nodes {
		fk := nodes[i].ParentID
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
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (_q *ParentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Node.Schema = _q.schemaConfig.Parent
	ctx = internal.NewSchemaConfigContext(ctx, _q.schemaConfig)
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *ParentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(parent.Table, parent.Columns, sqlgraph.NewFieldSpec(parent.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, parent.FieldID)
		for i := range fields {
			if fields[i] != parent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if _q.withChild != nil {
			_spec.Node.AddColumnOnce(parent.FieldUserID)
		}
		if _q.withParent != nil {
			_spec.Node.AddColumnOnce(parent.FieldParentID)
		}
	}
	if ps := _q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := _q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := _q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := _q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (_q *ParentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(parent.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = parent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(_q.schemaConfig.Parent)
	ctx = internal.NewSchemaConfigContext(ctx, _q.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range _q.modifiers {
		m(selector)
	}
	for _, p := range _q.predicates {
		p(selector)
	}
	for _, p := range _q.order {
		p(selector)
	}
	if offset := _q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := _q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (_q *ParentQuery) Modify(modifiers ...func(s *sql.Selector)) *ParentSelect {
	_q.modifiers = append(_q.modifiers, modifiers...)
	return _q.Select()
}

// ParentGroupBy is the group-by builder for Parent entities.
type ParentGroupBy struct {
	selector
	build *ParentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ParentGroupBy) Aggregate(fns ...AggregateFunc) *ParentGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ParentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ParentQuery, *ParentGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ParentGroupBy) sqlScan(ctx context.Context, root *ParentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ParentSelect is the builder for selecting fields of Parent entities.
type ParentSelect struct {
	*ParentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ParentSelect) Aggregate(fns ...AggregateFunc) *ParentSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ParentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ParentQuery, *ParentSelect](ctx, ps.ParentQuery, ps, ps.inters, v)
}

func (ps *ParentSelect) sqlScan(ctx context.Context, root *ParentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *ParentSelect) Modify(modifiers ...func(s *sql.Selector)) *ParentSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
