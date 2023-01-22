// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/upscale"
	"github.com/stablecog/go-apps/database/ent/upscaleoutput"
)

// UpscaleOutputCreate is the builder for creating a UpscaleOutput entity.
type UpscaleOutputCreate struct {
	config
	mutation *UpscaleOutputMutation
	hooks    []Hook
}

// SetImageURL sets the "image_url" field.
func (uoc *UpscaleOutputCreate) SetImageURL(s string) *UpscaleOutputCreate {
	uoc.mutation.SetImageURL(s)
	return uoc
}

// SetUpscaleID sets the "upscale_id" field.
func (uoc *UpscaleOutputCreate) SetUpscaleID(u uuid.UUID) *UpscaleOutputCreate {
	uoc.mutation.SetUpscaleID(u)
	return uoc
}

// SetCreatedAt sets the "created_at" field.
func (uoc *UpscaleOutputCreate) SetCreatedAt(t time.Time) *UpscaleOutputCreate {
	uoc.mutation.SetCreatedAt(t)
	return uoc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableCreatedAt(t *time.Time) *UpscaleOutputCreate {
	if t != nil {
		uoc.SetCreatedAt(*t)
	}
	return uoc
}

// SetUpdatedAt sets the "updated_at" field.
func (uoc *UpscaleOutputCreate) SetUpdatedAt(t time.Time) *UpscaleOutputCreate {
	uoc.mutation.SetUpdatedAt(t)
	return uoc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableUpdatedAt(t *time.Time) *UpscaleOutputCreate {
	if t != nil {
		uoc.SetUpdatedAt(*t)
	}
	return uoc
}

// SetID sets the "id" field.
func (uoc *UpscaleOutputCreate) SetID(u uuid.UUID) *UpscaleOutputCreate {
	uoc.mutation.SetID(u)
	return uoc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableID(u *uuid.UUID) *UpscaleOutputCreate {
	if u != nil {
		uoc.SetID(*u)
	}
	return uoc
}

// SetUpscalesID sets the "upscales" edge to the Upscale entity by ID.
func (uoc *UpscaleOutputCreate) SetUpscalesID(id uuid.UUID) *UpscaleOutputCreate {
	uoc.mutation.SetUpscalesID(id)
	return uoc
}

// SetUpscales sets the "upscales" edge to the Upscale entity.
func (uoc *UpscaleOutputCreate) SetUpscales(u *Upscale) *UpscaleOutputCreate {
	return uoc.SetUpscalesID(u.ID)
}

// Mutation returns the UpscaleOutputMutation object of the builder.
func (uoc *UpscaleOutputCreate) Mutation() *UpscaleOutputMutation {
	return uoc.mutation
}

// Save creates the UpscaleOutput in the database.
func (uoc *UpscaleOutputCreate) Save(ctx context.Context) (*UpscaleOutput, error) {
	uoc.defaults()
	return withHooks[*UpscaleOutput, UpscaleOutputMutation](ctx, uoc.sqlSave, uoc.mutation, uoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uoc *UpscaleOutputCreate) SaveX(ctx context.Context) *UpscaleOutput {
	v, err := uoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uoc *UpscaleOutputCreate) Exec(ctx context.Context) error {
	_, err := uoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uoc *UpscaleOutputCreate) ExecX(ctx context.Context) {
	if err := uoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uoc *UpscaleOutputCreate) defaults() {
	if _, ok := uoc.mutation.CreatedAt(); !ok {
		v := upscaleoutput.DefaultCreatedAt()
		uoc.mutation.SetCreatedAt(v)
	}
	if _, ok := uoc.mutation.UpdatedAt(); !ok {
		v := upscaleoutput.DefaultUpdatedAt()
		uoc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uoc.mutation.ID(); !ok {
		v := upscaleoutput.DefaultID()
		uoc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uoc *UpscaleOutputCreate) check() error {
	if _, ok := uoc.mutation.ImageURL(); !ok {
		return &ValidationError{Name: "image_url", err: errors.New(`ent: missing required field "UpscaleOutput.image_url"`)}
	}
	if _, ok := uoc.mutation.UpscaleID(); !ok {
		return &ValidationError{Name: "upscale_id", err: errors.New(`ent: missing required field "UpscaleOutput.upscale_id"`)}
	}
	if _, ok := uoc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UpscaleOutput.created_at"`)}
	}
	if _, ok := uoc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UpscaleOutput.updated_at"`)}
	}
	if _, ok := uoc.mutation.UpscalesID(); !ok {
		return &ValidationError{Name: "upscales", err: errors.New(`ent: missing required edge "UpscaleOutput.upscales"`)}
	}
	return nil
}

func (uoc *UpscaleOutputCreate) sqlSave(ctx context.Context) (*UpscaleOutput, error) {
	if err := uoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	uoc.mutation.id = &_node.ID
	uoc.mutation.done = true
	return _node, nil
}

func (uoc *UpscaleOutputCreate) createSpec() (*UpscaleOutput, *sqlgraph.CreateSpec) {
	var (
		_node = &UpscaleOutput{config: uoc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: upscaleoutput.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscaleoutput.FieldID,
			},
		}
	)
	if id, ok := uoc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uoc.mutation.ImageURL(); ok {
		_spec.SetField(upscaleoutput.FieldImageURL, field.TypeString, value)
		_node.ImageURL = value
	}
	if value, ok := uoc.mutation.CreatedAt(); ok {
		_spec.SetField(upscaleoutput.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uoc.mutation.UpdatedAt(); ok {
		_spec.SetField(upscaleoutput.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := uoc.mutation.UpscalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscaleoutput.UpscalesTable,
			Columns: []string{upscaleoutput.UpscalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UpscaleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UpscaleOutputCreateBulk is the builder for creating many UpscaleOutput entities in bulk.
type UpscaleOutputCreateBulk struct {
	config
	builders []*UpscaleOutputCreate
}

// Save creates the UpscaleOutput entities in the database.
func (uocb *UpscaleOutputCreateBulk) Save(ctx context.Context) ([]*UpscaleOutput, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uocb.builders))
	nodes := make([]*UpscaleOutput, len(uocb.builders))
	mutators := make([]Mutator, len(uocb.builders))
	for i := range uocb.builders {
		func(i int, root context.Context) {
			builder := uocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UpscaleOutputMutation)
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
					_, err = mutators[i+1].Mutate(root, uocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, uocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uocb *UpscaleOutputCreateBulk) SaveX(ctx context.Context) []*UpscaleOutput {
	v, err := uocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uocb *UpscaleOutputCreateBulk) Exec(ctx context.Context) error {
	_, err := uocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uocb *UpscaleOutputCreateBulk) ExecX(ctx context.Context) {
	if err := uocb.Exec(ctx); err != nil {
		panic(err)
	}
}