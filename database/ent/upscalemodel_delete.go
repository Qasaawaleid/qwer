// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/upscalemodel"
)

// UpscaleModelDelete is the builder for deleting a UpscaleModel entity.
type UpscaleModelDelete struct {
	config
	hooks    []Hook
	mutation *UpscaleModelMutation
}

// Where appends a list predicates to the UpscaleModelDelete builder.
func (umd *UpscaleModelDelete) Where(ps ...predicate.UpscaleModel) *UpscaleModelDelete {
	umd.mutation.Where(ps...)
	return umd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (umd *UpscaleModelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UpscaleModelMutation](ctx, umd.sqlExec, umd.mutation, umd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (umd *UpscaleModelDelete) ExecX(ctx context.Context) int {
	n, err := umd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (umd *UpscaleModelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: upscalemodel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscalemodel.FieldID,
			},
		},
	}
	if ps := umd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, umd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	umd.mutation.done = true
	return affected, err
}

// UpscaleModelDeleteOne is the builder for deleting a single UpscaleModel entity.
type UpscaleModelDeleteOne struct {
	umd *UpscaleModelDelete
}

// Exec executes the deletion query.
func (umdo *UpscaleModelDeleteOne) Exec(ctx context.Context) error {
	n, err := umdo.umd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{upscalemodel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (umdo *UpscaleModelDeleteOne) ExecX(ctx context.Context) {
	umdo.umd.ExecX(ctx)
}
