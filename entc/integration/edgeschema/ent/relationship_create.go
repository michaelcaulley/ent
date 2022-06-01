// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationship"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/schema/field"
)

// RelationshipCreate is the builder for creating a Relationship entity.
type RelationshipCreate struct {
	config
	mutation *RelationshipMutation
	hooks    []Hook
}

// SetWeight sets the "weight" field.
func (rc *RelationshipCreate) SetWeight(i int) *RelationshipCreate {
	rc.mutation.SetWeight(i)
	return rc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (rc *RelationshipCreate) SetNillableWeight(i *int) *RelationshipCreate {
	if i != nil {
		rc.SetWeight(*i)
	}
	return rc
}

// SetUserID sets the "user_id" field.
func (rc *RelationshipCreate) SetUserID(i int) *RelationshipCreate {
	rc.mutation.SetUserID(i)
	return rc
}

// SetRelativeID sets the "relative_id" field.
func (rc *RelationshipCreate) SetRelativeID(i int) *RelationshipCreate {
	rc.mutation.SetRelativeID(i)
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *RelationshipCreate) SetUser(u *User) *RelationshipCreate {
	return rc.SetUserID(u.ID)
}

// SetRelative sets the "relative" edge to the User entity.
func (rc *RelationshipCreate) SetRelative(u *User) *RelationshipCreate {
	return rc.SetRelativeID(u.ID)
}

// Mutation returns the RelationshipMutation object of the builder.
func (rc *RelationshipCreate) Mutation() *RelationshipMutation {
	return rc.mutation
}

// Save creates the Relationship in the database.
func (rc *RelationshipCreate) Save(ctx context.Context) (*Relationship, error) {
	var (
		err  error
		node *Relationship
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RelationshipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Relationship)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RelationshipMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RelationshipCreate) SaveX(ctx context.Context) *Relationship {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RelationshipCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RelationshipCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RelationshipCreate) defaults() {
	if _, ok := rc.mutation.Weight(); !ok {
		v := relationship.DefaultWeight
		rc.mutation.SetWeight(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RelationshipCreate) check() error {
	if _, ok := rc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Relationship.weight"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Relationship.user_id"`)}
	}
	if _, ok := rc.mutation.RelativeID(); !ok {
		return &ValidationError{Name: "relative_id", err: errors.New(`ent: missing required field "Relationship.relative_id"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Relationship.user"`)}
	}
	if _, ok := rc.mutation.RelativeID(); !ok {
		return &ValidationError{Name: "relative", err: errors.New(`ent: missing required edge "Relationship.relative"`)}
	}
	return nil
}

func (rc *RelationshipCreate) sqlSave(ctx context.Context) (*Relationship, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (rc *RelationshipCreate) createSpec() (*Relationship, *sqlgraph.CreateSpec) {
	var (
		_node = &Relationship{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: relationship.Table,
		}
	)
	if value, ok := rc.mutation.Weight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: relationship.FieldWeight,
		})
		_node.Weight = value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.UserTable,
			Columns: []string{relationship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RelativeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.RelativeTable,
			Columns: []string{relationship.RelativeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelativeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RelationshipCreateBulk is the builder for creating many Relationship entities in bulk.
type RelationshipCreateBulk struct {
	config
	builders []*RelationshipCreate
}

// Save creates the Relationship entities in the database.
func (rcb *RelationshipCreateBulk) Save(ctx context.Context) ([]*Relationship, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Relationship, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RelationshipMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RelationshipCreateBulk) SaveX(ctx context.Context) []*Relationship {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RelationshipCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RelationshipCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}