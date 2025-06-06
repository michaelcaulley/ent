// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/builder"
)

// BuilderCreate is the builder for creating a Builder entity.
type BuilderCreate struct {
	config
	mutation *BuilderMutation
	hooks    []Hook
}

// Mutation returns the BuilderMutation object of the builder.
func (_c *BuilderCreate) Mutation() *BuilderMutation {
	return _c.mutation
}

// Save creates the Builder in the database.
func (_c *BuilderCreate) Save(ctx context.Context) (*Builder, error) {
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *BuilderCreate) SaveX(ctx context.Context) *Builder {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *BuilderCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *BuilderCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *BuilderCreate) check() error {
	return nil
}

func (_c *BuilderCreate) gremlinSave(ctx context.Context) (*Builder, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := _c.gremlin().Query()
	if err := _c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Builder{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *BuilderCreate) gremlin() *dsl.Traversal {
	v := g.AddV(builder.Label)
	return v.ValueMap(true)
}

// BuilderCreateBulk is the builder for creating many Builder entities in bulk.
type BuilderCreateBulk struct {
	config
	err      error
	builders []*BuilderCreate
}
