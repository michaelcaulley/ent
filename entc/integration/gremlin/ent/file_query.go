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
	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/file"
	"entgo.io/ent/entc/integration/gremlin/ent/filetype"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// FileQuery is the builder for querying File entities.
type FileQuery struct {
	config
	ctx        *QueryContext
	order      []file.OrderOption
	inters     []Interceptor
	predicates []predicate.File
	withOwner  *UserQuery
	withType   *FileTypeQuery
	withField  *FieldTypeQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the FileQuery builder.
func (_q *FileQuery) Where(ps ...predicate.File) *FileQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *FileQuery) Limit(limit int) *FileQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *FileQuery) Offset(offset int) *FileQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *FileQuery) Unique(unique bool) *FileQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *FileQuery) Order(o ...file.OrderOption) *FileQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryOwner chains the current query on the "owner" edge.
func (_q *FileQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := _q.gremlinQuery(ctx)
		fromU = gremlin.InE(user.FilesLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryType chains the current query on the "type" edge.
func (_q *FileQuery) QueryType() *FileTypeQuery {
	query := (&FileTypeClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := _q.gremlinQuery(ctx)
		fromU = gremlin.InE(filetype.FilesLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryField chains the current query on the "field" edge.
func (_q *FileQuery) QueryField() *FieldTypeQuery {
	query := (&FieldTypeClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := _q.gremlinQuery(ctx)
		fromU = gremlin.OutE(file.FieldLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first File entity from the query.
// Returns a *NotFoundError when no File was found.
func (_q *FileQuery) First(ctx context.Context) (*File, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{file.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *FileQuery) FirstX(ctx context.Context) *File {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first File ID from the query.
// Returns a *NotFoundError when no File ID was found.
func (_q *FileQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{file.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *FileQuery) FirstIDX(ctx context.Context) string {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single File entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one File entity is found.
// Returns a *NotFoundError when no File entities are found.
func (_q *FileQuery) Only(ctx context.Context) (*File, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{file.Label}
	default:
		return nil, &NotSingularError{file.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *FileQuery) OnlyX(ctx context.Context) *File {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only File ID in the query.
// Returns a *NotSingularError when more than one File ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *FileQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{file.Label}
	default:
		err = &NotSingularError{file.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *FileQuery) OnlyIDX(ctx context.Context) string {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Files.
func (_q *FileQuery) All(ctx context.Context) ([]*File, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*File, *FileQuery]()
	return withInterceptors[[]*File](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *FileQuery) AllX(ctx context.Context) []*File {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of File IDs.
func (_q *FileQuery) IDs(ctx context.Context) (ids []string, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(file.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *FileQuery) IDsX(ctx context.Context) []string {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *FileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*FileQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *FileQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *FileQuery) Exist(ctx context.Context) (bool, error) {
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
func (_q *FileQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *FileQuery) Clone() *FileQuery {
	if _q == nil {
		return nil
	}
	return &FileQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]file.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.File{}, _q.predicates...),
		withOwner:  _q.withOwner.Clone(),
		withType:   _q.withType.Clone(),
		withField:  _q.withField.Clone(),
		// clone intermediate query.
		gremlin: _q.gremlin.Clone(),
		path:    _q.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *FileQuery) WithOwner(opts ...func(*UserQuery)) *FileQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withOwner = query
	return _q
}

// WithType tells the query-builder to eager-load the nodes that are connected to
// the "type" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *FileQuery) WithType(opts ...func(*FileTypeQuery)) *FileQuery {
	query := (&FileTypeClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withType = query
	return _q
}

// WithField tells the query-builder to eager-load the nodes that are connected to
// the "field" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *FileQuery) WithField(opts ...func(*FieldTypeQuery)) *FileQuery {
	query := (&FieldTypeClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withField = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SetID int `json:"set_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.File.Query().
//		GroupBy(file.FieldSetID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *FileQuery) GroupBy(field string, fields ...string) *FileGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FileGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = file.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SetID int `json:"set_id,omitempty"`
//	}
//
//	client.File.Query().
//		Select(file.FieldSetID).
//		Scan(ctx, &v)
func (_q *FileQuery) Select(fields ...string) *FileSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &FileSelect{FileQuery: _q}
	sbuild.label = file.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FileSelect configured with the given aggregations.
func (_q *FileQuery) Aggregate(fns ...AggregateFunc) *FileSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *FileQuery) prepareQuery(ctx context.Context) error {
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
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.gremlin = prev
	}
	return nil
}

func (_q *FileQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*File, error) {
	res := &gremlin.Response{}
	traversal := _q.gremlinQuery(ctx)
	if len(_q.ctx.Fields) > 0 {
		fields := make([]any, len(_q.ctx.Fields))
		for i, f := range _q.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var _ms Files
	if err := _ms.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range _ms {
		_ms[i].config = _q.config
	}
	return _ms, nil
}

func (_q *FileQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _q.gremlinQuery(ctx).Count().Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (_q *FileQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(file.Label)
	if _q.gremlin != nil {
		v = _q.gremlin.Clone()
	}
	for _, p := range _q.predicates {
		p(v)
	}
	if len(_q.order) > 0 {
		v.Order()
		for _, p := range _q.order {
			p(v)
		}
	}
	switch limit, offset := _q.ctx.Limit, _q.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := _q.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// FileGroupBy is the group-by builder for File entities.
type FileGroupBy struct {
	selector
	build *FileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FileGroupBy) Aggregate(fns ...AggregateFunc) *FileGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, ent.OpQueryGroupBy)
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileQuery, *FileGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FileGroupBy) gremlinScan(ctx context.Context, root *FileQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range fgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *fgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*fgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := fgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*fgb.flds)+len(fgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// FileSelect is the builder for selecting fields of File entities.
type FileSelect struct {
	*FileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FileSelect) Aggregate(fns ...AggregateFunc) *FileSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, ent.OpQuerySelect)
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileQuery, *FileSelect](ctx, fs.FileQuery, fs, fs.inters, v)
}

func (fs *FileSelect) gremlinScan(ctx context.Context, root *FileQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := fs.ctx.Fields; len(fields) == 1 {
		if fields[0] != file.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(fs.ctx.Fields))
		for i, f := range fs.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := fs.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
