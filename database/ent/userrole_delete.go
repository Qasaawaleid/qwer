// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stablecog/go-apps/database/ent/predicate"
	"github.com/stablecog/go-apps/database/ent/userrole"
)

// UserRoleDelete is the builder for deleting a UserRole entity.
type UserRoleDelete struct {
	config
	hooks    []Hook
	mutation *UserRoleMutation
}

// Where appends a list predicates to the UserRoleDelete builder.
func (urd *UserRoleDelete) Where(ps ...predicate.UserRole) *UserRoleDelete {
	urd.mutation.Where(ps...)
	return urd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urd *UserRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UserRoleMutation](ctx, urd.sqlExec, urd.mutation, urd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (urd *UserRoleDelete) ExecX(ctx context.Context) int {
	n, err := urd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urd *UserRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		},
	}
	if ps := urd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, urd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	urd.mutation.done = true
	return affected, err
}

// UserRoleDeleteOne is the builder for deleting a single UserRole entity.
type UserRoleDeleteOne struct {
	urd *UserRoleDelete
}

// Exec executes the deletion query.
func (urdo *UserRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := urdo.urd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urdo *UserRoleDeleteOne) ExecX(ctx context.Context) {
	urdo.urd.ExecX(ctx)
}