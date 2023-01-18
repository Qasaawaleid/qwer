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
	"github.com/stablecog/go-apps/database/ent/generationg"
	"github.com/stablecog/go-apps/database/ent/predicate"
	"github.com/stablecog/go-apps/database/ent/upscale"
	"github.com/stablecog/go-apps/database/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (uu *UserUpdate) SetStripeCustomerID(s string) *UserUpdate {
	uu.mutation.SetStripeCustomerID(s)
	return uu
}

// SetSubscriptionTier sets the "subscription_tier" field.
func (uu *UserUpdate) SetSubscriptionTier(ut user.SubscriptionTier) *UserUpdate {
	uu.mutation.SetSubscriptionTier(ut)
	return uu
}

// SetNillableSubscriptionTier sets the "subscription_tier" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSubscriptionTier(ut *user.SubscriptionTier) *UserUpdate {
	if ut != nil {
		uu.SetSubscriptionTier(*ut)
	}
	return uu
}

// SetSubscriptionCategory sets the "subscription_category" field.
func (uu *UserUpdate) SetSubscriptionCategory(uc user.SubscriptionCategory) *UserUpdate {
	uu.mutation.SetSubscriptionCategory(uc)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetConfirmedAt sets the "confirmed_at" field.
func (uu *UserUpdate) SetConfirmedAt(t time.Time) *UserUpdate {
	uu.mutation.SetConfirmedAt(t)
	return uu
}

// AddUpscaleIDs adds the "upscale" edge to the Upscale entity by IDs.
func (uu *UserUpdate) AddUpscaleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddUpscaleIDs(ids...)
	return uu
}

// AddUpscale adds the "upscale" edges to the Upscale entity.
func (uu *UserUpdate) AddUpscale(u ...*Upscale) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUpscaleIDs(ids...)
}

// AddGenerationIDs adds the "generation" edge to the Generation entity by IDs.
func (uu *UserUpdate) AddGenerationIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddGenerationIDs(ids...)
	return uu
}

// AddGeneration adds the "generation" edges to the Generation entity.
func (uu *UserUpdate) AddGeneration(g ...*Generation) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGenerationIDs(ids...)
}

// AddGenerationGIDs adds the "generation_g" edge to the GenerationG entity by IDs.
func (uu *UserUpdate) AddGenerationGIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddGenerationGIDs(ids...)
	return uu
}

// AddGenerationG adds the "generation_g" edges to the GenerationG entity.
func (uu *UserUpdate) AddGenerationG(g ...*GenerationG) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGenerationGIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearUpscale clears all "upscale" edges to the Upscale entity.
func (uu *UserUpdate) ClearUpscale() *UserUpdate {
	uu.mutation.ClearUpscale()
	return uu
}

// RemoveUpscaleIDs removes the "upscale" edge to Upscale entities by IDs.
func (uu *UserUpdate) RemoveUpscaleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveUpscaleIDs(ids...)
	return uu
}

// RemoveUpscale removes "upscale" edges to Upscale entities.
func (uu *UserUpdate) RemoveUpscale(u ...*Upscale) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUpscaleIDs(ids...)
}

// ClearGeneration clears all "generation" edges to the Generation entity.
func (uu *UserUpdate) ClearGeneration() *UserUpdate {
	uu.mutation.ClearGeneration()
	return uu
}

// RemoveGenerationIDs removes the "generation" edge to Generation entities by IDs.
func (uu *UserUpdate) RemoveGenerationIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveGenerationIDs(ids...)
	return uu
}

// RemoveGeneration removes "generation" edges to Generation entities.
func (uu *UserUpdate) RemoveGeneration(g ...*Generation) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGenerationIDs(ids...)
}

// ClearGenerationG clears all "generation_g" edges to the GenerationG entity.
func (uu *UserUpdate) ClearGenerationG() *UserUpdate {
	uu.mutation.ClearGenerationG()
	return uu
}

// RemoveGenerationGIDs removes the "generation_g" edge to GenerationG entities by IDs.
func (uu *UserUpdate) RemoveGenerationGIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveGenerationGIDs(ids...)
	return uu
}

// RemoveGenerationG removes "generation_g" edges to GenerationG entities.
func (uu *UserUpdate) RemoveGenerationG(g ...*GenerationG) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGenerationGIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.SubscriptionTier(); ok {
		if err := user.SubscriptionTierValidator(v); err != nil {
			return &ValidationError{Name: "subscription_tier", err: fmt.Errorf(`ent: validator failed for field "User.subscription_tier": %w`, err)}
		}
	}
	if v, ok := uu.mutation.SubscriptionCategory(); ok {
		if err := user.SubscriptionCategoryValidator(v); err != nil {
			return &ValidationError{Name: "subscription_category", err: fmt.Errorf(`ent: validator failed for field "User.subscription_category": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.StripeCustomerID(); ok {
		_spec.SetField(user.FieldStripeCustomerID, field.TypeString, value)
	}
	if value, ok := uu.mutation.SubscriptionTier(); ok {
		_spec.SetField(user.FieldSubscriptionTier, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.SubscriptionCategory(); ok {
		_spec.SetField(user.FieldSubscriptionCategory, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.ConfirmedAt(); ok {
		_spec.SetField(user.FieldConfirmedAt, field.TypeTime, value)
	}
	if uu.mutation.UpscaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UpscaleTable,
			Columns: []string{user.UpscaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUpscaleIDs(); len(nodes) > 0 && !uu.mutation.UpscaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UpscaleTable,
			Columns: []string{user.UpscaleColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UpscaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UpscaleTable,
			Columns: []string{user.UpscaleColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GenerationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationTable,
			Columns: []string{user.GenerationColumn},
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
	if nodes := uu.mutation.RemovedGenerationIDs(); len(nodes) > 0 && !uu.mutation.GenerationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationTable,
			Columns: []string{user.GenerationColumn},
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
	if nodes := uu.mutation.GenerationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationTable,
			Columns: []string{user.GenerationColumn},
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
	if uu.mutation.GenerationGCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationGTable,
			Columns: []string{user.GenerationGColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationg.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGenerationGIDs(); len(nodes) > 0 && !uu.mutation.GenerationGCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationGTable,
			Columns: []string{user.GenerationGColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GenerationGIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationGTable,
			Columns: []string{user.GenerationGColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (uuo *UserUpdateOne) SetStripeCustomerID(s string) *UserUpdateOne {
	uuo.mutation.SetStripeCustomerID(s)
	return uuo
}

// SetSubscriptionTier sets the "subscription_tier" field.
func (uuo *UserUpdateOne) SetSubscriptionTier(ut user.SubscriptionTier) *UserUpdateOne {
	uuo.mutation.SetSubscriptionTier(ut)
	return uuo
}

// SetNillableSubscriptionTier sets the "subscription_tier" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSubscriptionTier(ut *user.SubscriptionTier) *UserUpdateOne {
	if ut != nil {
		uuo.SetSubscriptionTier(*ut)
	}
	return uuo
}

// SetSubscriptionCategory sets the "subscription_category" field.
func (uuo *UserUpdateOne) SetSubscriptionCategory(uc user.SubscriptionCategory) *UserUpdateOne {
	uuo.mutation.SetSubscriptionCategory(uc)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetConfirmedAt sets the "confirmed_at" field.
func (uuo *UserUpdateOne) SetConfirmedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetConfirmedAt(t)
	return uuo
}

// AddUpscaleIDs adds the "upscale" edge to the Upscale entity by IDs.
func (uuo *UserUpdateOne) AddUpscaleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddUpscaleIDs(ids...)
	return uuo
}

// AddUpscale adds the "upscale" edges to the Upscale entity.
func (uuo *UserUpdateOne) AddUpscale(u ...*Upscale) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUpscaleIDs(ids...)
}

// AddGenerationIDs adds the "generation" edge to the Generation entity by IDs.
func (uuo *UserUpdateOne) AddGenerationIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddGenerationIDs(ids...)
	return uuo
}

// AddGeneration adds the "generation" edges to the Generation entity.
func (uuo *UserUpdateOne) AddGeneration(g ...*Generation) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGenerationIDs(ids...)
}

// AddGenerationGIDs adds the "generation_g" edge to the GenerationG entity by IDs.
func (uuo *UserUpdateOne) AddGenerationGIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddGenerationGIDs(ids...)
	return uuo
}

// AddGenerationG adds the "generation_g" edges to the GenerationG entity.
func (uuo *UserUpdateOne) AddGenerationG(g ...*GenerationG) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGenerationGIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearUpscale clears all "upscale" edges to the Upscale entity.
func (uuo *UserUpdateOne) ClearUpscale() *UserUpdateOne {
	uuo.mutation.ClearUpscale()
	return uuo
}

// RemoveUpscaleIDs removes the "upscale" edge to Upscale entities by IDs.
func (uuo *UserUpdateOne) RemoveUpscaleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveUpscaleIDs(ids...)
	return uuo
}

// RemoveUpscale removes "upscale" edges to Upscale entities.
func (uuo *UserUpdateOne) RemoveUpscale(u ...*Upscale) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUpscaleIDs(ids...)
}

// ClearGeneration clears all "generation" edges to the Generation entity.
func (uuo *UserUpdateOne) ClearGeneration() *UserUpdateOne {
	uuo.mutation.ClearGeneration()
	return uuo
}

// RemoveGenerationIDs removes the "generation" edge to Generation entities by IDs.
func (uuo *UserUpdateOne) RemoveGenerationIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveGenerationIDs(ids...)
	return uuo
}

// RemoveGeneration removes "generation" edges to Generation entities.
func (uuo *UserUpdateOne) RemoveGeneration(g ...*Generation) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGenerationIDs(ids...)
}

// ClearGenerationG clears all "generation_g" edges to the GenerationG entity.
func (uuo *UserUpdateOne) ClearGenerationG() *UserUpdateOne {
	uuo.mutation.ClearGenerationG()
	return uuo
}

// RemoveGenerationGIDs removes the "generation_g" edge to GenerationG entities by IDs.
func (uuo *UserUpdateOne) RemoveGenerationGIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveGenerationGIDs(ids...)
	return uuo
}

// RemoveGenerationG removes "generation_g" edges to GenerationG entities.
func (uuo *UserUpdateOne) RemoveGenerationG(g ...*GenerationG) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGenerationGIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.SubscriptionTier(); ok {
		if err := user.SubscriptionTierValidator(v); err != nil {
			return &ValidationError{Name: "subscription_tier", err: fmt.Errorf(`ent: validator failed for field "User.subscription_tier": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.SubscriptionCategory(); ok {
		if err := user.SubscriptionCategoryValidator(v); err != nil {
			return &ValidationError{Name: "subscription_category", err: fmt.Errorf(`ent: validator failed for field "User.subscription_category": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.StripeCustomerID(); ok {
		_spec.SetField(user.FieldStripeCustomerID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.SubscriptionTier(); ok {
		_spec.SetField(user.FieldSubscriptionTier, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.SubscriptionCategory(); ok {
		_spec.SetField(user.FieldSubscriptionCategory, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.ConfirmedAt(); ok {
		_spec.SetField(user.FieldConfirmedAt, field.TypeTime, value)
	}
	if uuo.mutation.UpscaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UpscaleTable,
			Columns: []string{user.UpscaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUpscaleIDs(); len(nodes) > 0 && !uuo.mutation.UpscaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UpscaleTable,
			Columns: []string{user.UpscaleColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UpscaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UpscaleTable,
			Columns: []string{user.UpscaleColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GenerationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationTable,
			Columns: []string{user.GenerationColumn},
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
	if nodes := uuo.mutation.RemovedGenerationIDs(); len(nodes) > 0 && !uuo.mutation.GenerationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationTable,
			Columns: []string{user.GenerationColumn},
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
	if nodes := uuo.mutation.GenerationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationTable,
			Columns: []string{user.GenerationColumn},
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
	if uuo.mutation.GenerationGCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationGTable,
			Columns: []string{user.GenerationGColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationg.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGenerationGIDs(); len(nodes) > 0 && !uuo.mutation.GenerationGCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationGTable,
			Columns: []string{user.GenerationGColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GenerationGIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GenerationGTable,
			Columns: []string{user.GenerationGColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}