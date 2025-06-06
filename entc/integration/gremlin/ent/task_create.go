// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/ent/schema/task"
	enttask "entgo.io/ent/entc/integration/gremlin/ent/task"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetPriority sets the "priority" field.
func (_c *TaskCreate) SetPriority(v task.Priority) *TaskCreate {
	_c.mutation.SetPriority(v)
	return _c
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (_c *TaskCreate) SetNillablePriority(v *task.Priority) *TaskCreate {
	if v != nil {
		_c.SetPriority(*v)
	}
	return _c
}

// SetPriorities sets the "priorities" field.
func (_c *TaskCreate) SetPriorities(v map[string]task.Priority) *TaskCreate {
	_c.mutation.SetPriorities(v)
	return _c
}

// SetCreatedAt sets the "created_at" field.
func (_c *TaskCreate) SetCreatedAt(v time.Time) *TaskCreate {
	_c.mutation.SetCreatedAt(v)
	return _c
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (_c *TaskCreate) SetNillableCreatedAt(v *time.Time) *TaskCreate {
	if v != nil {
		_c.SetCreatedAt(*v)
	}
	return _c
}

// SetName sets the "name" field.
func (_c *TaskCreate) SetName(v string) *TaskCreate {
	_c.mutation.SetName(v)
	return _c
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_c *TaskCreate) SetNillableName(v *string) *TaskCreate {
	if v != nil {
		_c.SetName(*v)
	}
	return _c
}

// SetOwner sets the "owner" field.
func (_c *TaskCreate) SetOwner(v string) *TaskCreate {
	_c.mutation.SetOwner(v)
	return _c
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (_c *TaskCreate) SetNillableOwner(v *string) *TaskCreate {
	if v != nil {
		_c.SetOwner(*v)
	}
	return _c
}

// SetOrder sets the "order" field.
func (_c *TaskCreate) SetOrder(v int) *TaskCreate {
	_c.mutation.SetOrder(v)
	return _c
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (_c *TaskCreate) SetNillableOrder(v *int) *TaskCreate {
	if v != nil {
		_c.SetOrder(*v)
	}
	return _c
}

// SetOrderOption sets the "order_option" field.
func (_c *TaskCreate) SetOrderOption(v int) *TaskCreate {
	_c.mutation.SetOrderOption(v)
	return _c
}

// SetNillableOrderOption sets the "order_option" field if the given value is not nil.
func (_c *TaskCreate) SetNillableOrderOption(v *int) *TaskCreate {
	if v != nil {
		_c.SetOrderOption(*v)
	}
	return _c
}

// SetOp sets the "op" field.
func (_c *TaskCreate) SetOp(v string) *TaskCreate {
	_c.mutation.SetOpField(v)
	return _c
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (_c *TaskCreate) SetNillableOp(v *string) *TaskCreate {
	if v != nil {
		_c.SetOp(*v)
	}
	return _c
}

// Mutation returns the TaskMutation object of the builder.
func (_c *TaskCreate) Mutation() *TaskMutation {
	return _c.mutation
}

// Save creates the Task in the database.
func (_c *TaskCreate) Save(ctx context.Context) (*Task, error) {
	_c.defaults()
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *TaskCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *TaskCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *TaskCreate) defaults() {
	if _, ok := _c.mutation.Priority(); !ok {
		v := enttask.DefaultPriority
		_c.mutation.SetPriority(v)
	}
	if _, ok := _c.mutation.CreatedAt(); !ok {
		v := enttask.DefaultCreatedAt()
		_c.mutation.SetCreatedAt(v)
	}
	if _, ok := _c.mutation.GetOp(); !ok {
		v := enttask.DefaultOp
		_c.mutation.SetOpField(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *TaskCreate) check() error {
	if _, ok := _c.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Task.priority"`)}
	}
	if v, ok := _c.mutation.Priority(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Task.priority": %w`, err)}
		}
	}
	if _, ok := _c.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	if _, ok := _c.mutation.GetOp(); !ok {
		return &ValidationError{Name: "op", err: errors.New(`ent: missing required field "Task.op"`)}
	}
	if v, ok := _c.mutation.GetOp(); ok {
		if err := enttask.OpValidator(v); err != nil {
			return &ValidationError{Name: "op", err: fmt.Errorf(`ent: validator failed for field "Task.op": %w`, err)}
		}
	}
	return nil
}

func (_c *TaskCreate) gremlinSave(ctx context.Context) (*Task, error) {
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
	rnode := &Task{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *TaskCreate) gremlin() *dsl.Traversal {
	v := g.AddV(enttask.Label)
	if value, ok := _c.mutation.Priority(); ok {
		v.Property(dsl.Single, enttask.FieldPriority, value)
	}
	if value, ok := _c.mutation.Priorities(); ok {
		v.Property(dsl.Single, enttask.FieldPriorities, value)
	}
	if value, ok := _c.mutation.CreatedAt(); ok {
		v.Property(dsl.Single, enttask.FieldCreatedAt, value)
	}
	if value, ok := _c.mutation.Name(); ok {
		v.Property(dsl.Single, enttask.FieldName, value)
	}
	if value, ok := _c.mutation.Owner(); ok {
		v.Property(dsl.Single, enttask.FieldOwner, value)
	}
	if value, ok := _c.mutation.Order(); ok {
		v.Property(dsl.Single, enttask.FieldOrder, value)
	}
	if value, ok := _c.mutation.OrderOption(); ok {
		v.Property(dsl.Single, enttask.FieldOrderOption, value)
	}
	if value, ok := _c.mutation.GetOp(); ok {
		v.Property(dsl.Single, enttask.FieldOp, value)
	}
	return v.ValueMap(true)
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	err      error
	builders []*TaskCreate
}
