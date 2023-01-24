// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/generation"
	"github.com/stablecog/go-apps/database/ent/generationmodel"
	"github.com/stablecog/go-apps/database/ent/predicate"
)

// GenerationModelUpdate is the builder for updating GenerationModel entities.
type GenerationModelUpdate struct {
	config
	hooks    []Hook
	mutation *GenerationModelMutation
}

// Where appends a list predicates to the GenerationModelUpdate builder.
func (gmu *GenerationModelUpdate) Where(ps ...predicate.GenerationModel) *GenerationModelUpdate {
	gmu.mutation.Where(ps...)
	return gmu
}

// SetName sets the "name" field.
func (gmu *GenerationModelUpdate) SetName(s string) *GenerationModelUpdate {
	gmu.mutation.SetName(s)
	return gmu
}

// SetIsFree sets the "is_free" field.
func (gmu *GenerationModelUpdate) SetIsFree(b bool) *GenerationModelUpdate {
	gmu.mutation.SetIsFree(b)
	return gmu
}

// SetNillableIsFree sets the "is_free" field if the given value is not nil.
func (gmu *GenerationModelUpdate) SetNillableIsFree(b *bool) *GenerationModelUpdate {
	if b != nil {
		gmu.SetIsFree(*b)
	}
	return gmu
}

// SetUpdatedAt sets the "updated_at" field.
func (gmu *GenerationModelUpdate) SetUpdatedAt(t time.Time) *GenerationModelUpdate {
	gmu.mutation.SetUpdatedAt(t)
	return gmu
}

// AddGenerationIDs adds the "generations" edge to the Generation entity by IDs.
func (gmu *GenerationModelUpdate) AddGenerationIDs(ids ...uuid.UUID) *GenerationModelUpdate {
	gmu.mutation.AddGenerationIDs(ids...)
	return gmu
}

// AddGenerations adds the "generations" edges to the Generation entity.
func (gmu *GenerationModelUpdate) AddGenerations(g ...*Generation) *GenerationModelUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gmu.AddGenerationIDs(ids...)
}

// Mutation returns the GenerationModelMutation object of the builder.
func (gmu *GenerationModelUpdate) Mutation() *GenerationModelMutation {
	return gmu.mutation
}

// ClearGenerations clears all "generations" edges to the Generation entity.
func (gmu *GenerationModelUpdate) ClearGenerations() *GenerationModelUpdate {
	gmu.mutation.ClearGenerations()
	return gmu
}

// RemoveGenerationIDs removes the "generations" edge to Generation entities by IDs.
func (gmu *GenerationModelUpdate) RemoveGenerationIDs(ids ...uuid.UUID) *GenerationModelUpdate {
	gmu.mutation.RemoveGenerationIDs(ids...)
	return gmu
}

// RemoveGenerations removes "generations" edges to Generation entities.
func (gmu *GenerationModelUpdate) RemoveGenerations(g ...*Generation) *GenerationModelUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gmu.RemoveGenerationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmu *GenerationModelUpdate) Save(ctx context.Context) (int, error) {
	gmu.defaults()
	return withHooks[int, GenerationModelMutation](ctx, gmu.sqlSave, gmu.mutation, gmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmu *GenerationModelUpdate) SaveX(ctx context.Context) int {
	affected, err := gmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmu *GenerationModelUpdate) Exec(ctx context.Context) error {
	_, err := gmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmu *GenerationModelUpdate) ExecX(ctx context.Context) {
	if err := gmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmu *GenerationModelUpdate) defaults() {
	if _, ok := gmu.mutation.UpdatedAt(); !ok {
		v := generationmodel.UpdateDefaultUpdatedAt()
		gmu.mutation.SetUpdatedAt(v)
	}
}

func (gmu *GenerationModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   generationmodel.Table,
			Columns: generationmodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generationmodel.FieldID,
			},
		},
	}
	if ps := gmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmu.mutation.Name(); ok {
		_spec.SetField(generationmodel.FieldName, field.TypeString, value)
	}
	if value, ok := gmu.mutation.IsFree(); ok {
		_spec.SetField(generationmodel.FieldIsFree, field.TypeBool, value)
	}
	if value, ok := gmu.mutation.UpdatedAt(); ok {
		_spec.SetField(generationmodel.FieldUpdatedAt, field.TypeTime, value)
	}
	if gmu.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   generationmodel.GenerationsTable,
			Columns: []string{generationmodel.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmu.mutation.RemovedGenerationsIDs(); len(nodes) > 0 && !gmu.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   generationmodel.GenerationsTable,
			Columns: []string{generationmodel.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmu.mutation.GenerationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   generationmodel.GenerationsTable,
			Columns: []string{generationmodel.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{generationmodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gmu.mutation.done = true
	return n, nil
}

// GenerationModelUpdateOne is the builder for updating a single GenerationModel entity.
type GenerationModelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GenerationModelMutation
}

// SetName sets the "name" field.
func (gmuo *GenerationModelUpdateOne) SetName(s string) *GenerationModelUpdateOne {
	gmuo.mutation.SetName(s)
	return gmuo
}

// SetIsFree sets the "is_free" field.
func (gmuo *GenerationModelUpdateOne) SetIsFree(b bool) *GenerationModelUpdateOne {
	gmuo.mutation.SetIsFree(b)
	return gmuo
}

// SetNillableIsFree sets the "is_free" field if the given value is not nil.
func (gmuo *GenerationModelUpdateOne) SetNillableIsFree(b *bool) *GenerationModelUpdateOne {
	if b != nil {
		gmuo.SetIsFree(*b)
	}
	return gmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (gmuo *GenerationModelUpdateOne) SetUpdatedAt(t time.Time) *GenerationModelUpdateOne {
	gmuo.mutation.SetUpdatedAt(t)
	return gmuo
}

// AddGenerationIDs adds the "generations" edge to the Generation entity by IDs.
func (gmuo *GenerationModelUpdateOne) AddGenerationIDs(ids ...uuid.UUID) *GenerationModelUpdateOne {
	gmuo.mutation.AddGenerationIDs(ids...)
	return gmuo
}

// AddGenerations adds the "generations" edges to the Generation entity.
func (gmuo *GenerationModelUpdateOne) AddGenerations(g ...*Generation) *GenerationModelUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gmuo.AddGenerationIDs(ids...)
}

// Mutation returns the GenerationModelMutation object of the builder.
func (gmuo *GenerationModelUpdateOne) Mutation() *GenerationModelMutation {
	return gmuo.mutation
}

// ClearGenerations clears all "generations" edges to the Generation entity.
func (gmuo *GenerationModelUpdateOne) ClearGenerations() *GenerationModelUpdateOne {
	gmuo.mutation.ClearGenerations()
	return gmuo
}

// RemoveGenerationIDs removes the "generations" edge to Generation entities by IDs.
func (gmuo *GenerationModelUpdateOne) RemoveGenerationIDs(ids ...uuid.UUID) *GenerationModelUpdateOne {
	gmuo.mutation.RemoveGenerationIDs(ids...)
	return gmuo
}

// RemoveGenerations removes "generations" edges to Generation entities.
func (gmuo *GenerationModelUpdateOne) RemoveGenerations(g ...*Generation) *GenerationModelUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gmuo.RemoveGenerationIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmuo *GenerationModelUpdateOne) Select(field string, fields ...string) *GenerationModelUpdateOne {
	gmuo.fields = append([]string{field}, fields...)
	return gmuo
}

// Save executes the query and returns the updated GenerationModel entity.
func (gmuo *GenerationModelUpdateOne) Save(ctx context.Context) (*GenerationModel, error) {
	gmuo.defaults()
	return withHooks[*GenerationModel, GenerationModelMutation](ctx, gmuo.sqlSave, gmuo.mutation, gmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmuo *GenerationModelUpdateOne) SaveX(ctx context.Context) *GenerationModel {
	node, err := gmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmuo *GenerationModelUpdateOne) Exec(ctx context.Context) error {
	_, err := gmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmuo *GenerationModelUpdateOne) ExecX(ctx context.Context) {
	if err := gmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmuo *GenerationModelUpdateOne) defaults() {
	if _, ok := gmuo.mutation.UpdatedAt(); !ok {
		v := generationmodel.UpdateDefaultUpdatedAt()
		gmuo.mutation.SetUpdatedAt(v)
	}
}

func (gmuo *GenerationModelUpdateOne) sqlSave(ctx context.Context) (_node *GenerationModel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   generationmodel.Table,
			Columns: generationmodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generationmodel.FieldID,
			},
		},
	}
	id, ok := gmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GenerationModel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, generationmodel.FieldID)
		for _, f := range fields {
			if !generationmodel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != generationmodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmuo.mutation.Name(); ok {
		_spec.SetField(generationmodel.FieldName, field.TypeString, value)
	}
	if value, ok := gmuo.mutation.IsFree(); ok {
		_spec.SetField(generationmodel.FieldIsFree, field.TypeBool, value)
	}
	if value, ok := gmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(generationmodel.FieldUpdatedAt, field.TypeTime, value)
	}
	if gmuo.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   generationmodel.GenerationsTable,
			Columns: []string{generationmodel.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmuo.mutation.RemovedGenerationsIDs(); len(nodes) > 0 && !gmuo.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   generationmodel.GenerationsTable,
			Columns: []string{generationmodel.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmuo.mutation.GenerationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   generationmodel.GenerationsTable,
			Columns: []string{generationmodel.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GenerationModel{config: gmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{generationmodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gmuo.mutation.done = true
	return _node, nil
}
