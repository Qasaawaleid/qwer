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
	"github.com/yekta/stablecog/go-apps/database/ent/generationrealtime"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
)

// GenerationRealtimeUpdate is the builder for updating GenerationRealtime entities.
type GenerationRealtimeUpdate struct {
	config
	hooks    []Hook
	mutation *GenerationRealtimeMutation
}

// Where appends a list predicates to the GenerationRealtimeUpdate builder.
func (gru *GenerationRealtimeUpdate) Where(ps ...predicate.GenerationRealtime) *GenerationRealtimeUpdate {
	gru.mutation.Where(ps...)
	return gru
}

// SetCountryCode sets the "country_code" field.
func (gru *GenerationRealtimeUpdate) SetCountryCode(s string) *GenerationRealtimeUpdate {
	gru.mutation.SetCountryCode(s)
	return gru
}

// SetDurationMs sets the "duration_ms" field.
func (gru *GenerationRealtimeUpdate) SetDurationMs(i int) *GenerationRealtimeUpdate {
	gru.mutation.ResetDurationMs()
	gru.mutation.SetDurationMs(i)
	return gru
}

// AddDurationMs adds i to the "duration_ms" field.
func (gru *GenerationRealtimeUpdate) AddDurationMs(i int) *GenerationRealtimeUpdate {
	gru.mutation.AddDurationMs(i)
	return gru
}

// SetStatus sets the "status" field.
func (gru *GenerationRealtimeUpdate) SetStatus(ge generationrealtime.Status) *GenerationRealtimeUpdate {
	gru.mutation.SetStatus(ge)
	return gru
}

// SetUsesDefaultServer sets the "uses_default_server" field.
func (gru *GenerationRealtimeUpdate) SetUsesDefaultServer(b bool) *GenerationRealtimeUpdate {
	gru.mutation.SetUsesDefaultServer(b)
	return gru
}

// SetWidth sets the "width" field.
func (gru *GenerationRealtimeUpdate) SetWidth(i int) *GenerationRealtimeUpdate {
	gru.mutation.ResetWidth()
	gru.mutation.SetWidth(i)
	return gru
}

// AddWidth adds i to the "width" field.
func (gru *GenerationRealtimeUpdate) AddWidth(i int) *GenerationRealtimeUpdate {
	gru.mutation.AddWidth(i)
	return gru
}

// SetHeight sets the "height" field.
func (gru *GenerationRealtimeUpdate) SetHeight(i int) *GenerationRealtimeUpdate {
	gru.mutation.ResetHeight()
	gru.mutation.SetHeight(i)
	return gru
}

// AddHeight adds i to the "height" field.
func (gru *GenerationRealtimeUpdate) AddHeight(i int) *GenerationRealtimeUpdate {
	gru.mutation.AddHeight(i)
	return gru
}

// SetNumInterferenceSteps sets the "num_interference_steps" field.
func (gru *GenerationRealtimeUpdate) SetNumInterferenceSteps(i int) *GenerationRealtimeUpdate {
	gru.mutation.ResetNumInterferenceSteps()
	gru.mutation.SetNumInterferenceSteps(i)
	return gru
}

// AddNumInterferenceSteps adds i to the "num_interference_steps" field.
func (gru *GenerationRealtimeUpdate) AddNumInterferenceSteps(i int) *GenerationRealtimeUpdate {
	gru.mutation.AddNumInterferenceSteps(i)
	return gru
}

// SetUpdatedAt sets the "updated_at" field.
func (gru *GenerationRealtimeUpdate) SetUpdatedAt(t time.Time) *GenerationRealtimeUpdate {
	gru.mutation.SetUpdatedAt(t)
	return gru
}

// SetUserTier sets the "user_tier" field.
func (gru *GenerationRealtimeUpdate) SetUserTier(gt generationrealtime.UserTier) *GenerationRealtimeUpdate {
	gru.mutation.SetUserTier(gt)
	return gru
}

// SetNillableUserTier sets the "user_tier" field if the given value is not nil.
func (gru *GenerationRealtimeUpdate) SetNillableUserTier(gt *generationrealtime.UserTier) *GenerationRealtimeUpdate {
	if gt != nil {
		gru.SetUserTier(*gt)
	}
	return gru
}

// Mutation returns the GenerationRealtimeMutation object of the builder.
func (gru *GenerationRealtimeUpdate) Mutation() *GenerationRealtimeMutation {
	return gru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gru *GenerationRealtimeUpdate) Save(ctx context.Context) (int, error) {
	gru.defaults()
	return withHooks[int, GenerationRealtimeMutation](ctx, gru.sqlSave, gru.mutation, gru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gru *GenerationRealtimeUpdate) SaveX(ctx context.Context) int {
	affected, err := gru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gru *GenerationRealtimeUpdate) Exec(ctx context.Context) error {
	_, err := gru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gru *GenerationRealtimeUpdate) ExecX(ctx context.Context) {
	if err := gru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gru *GenerationRealtimeUpdate) defaults() {
	if _, ok := gru.mutation.UpdatedAt(); !ok {
		v := generationrealtime.UpdateDefaultUpdatedAt()
		gru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gru *GenerationRealtimeUpdate) check() error {
	if v, ok := gru.mutation.Status(); ok {
		if err := generationrealtime.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "GenerationRealtime.status": %w`, err)}
		}
	}
	if v, ok := gru.mutation.UserTier(); ok {
		if err := generationrealtime.UserTierValidator(v); err != nil {
			return &ValidationError{Name: "user_tier", err: fmt.Errorf(`ent: validator failed for field "GenerationRealtime.user_tier": %w`, err)}
		}
	}
	return nil
}

func (gru *GenerationRealtimeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gru.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   generationrealtime.Table,
			Columns: generationrealtime.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generationrealtime.FieldID,
			},
		},
	}
	if ps := gru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gru.mutation.CountryCode(); ok {
		_spec.SetField(generationrealtime.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := gru.mutation.DurationMs(); ok {
		_spec.SetField(generationrealtime.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := gru.mutation.AddedDurationMs(); ok {
		_spec.AddField(generationrealtime.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := gru.mutation.Status(); ok {
		_spec.SetField(generationrealtime.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := gru.mutation.UsesDefaultServer(); ok {
		_spec.SetField(generationrealtime.FieldUsesDefaultServer, field.TypeBool, value)
	}
	if value, ok := gru.mutation.Width(); ok {
		_spec.SetField(generationrealtime.FieldWidth, field.TypeInt, value)
	}
	if value, ok := gru.mutation.AddedWidth(); ok {
		_spec.AddField(generationrealtime.FieldWidth, field.TypeInt, value)
	}
	if value, ok := gru.mutation.Height(); ok {
		_spec.SetField(generationrealtime.FieldHeight, field.TypeInt, value)
	}
	if value, ok := gru.mutation.AddedHeight(); ok {
		_spec.AddField(generationrealtime.FieldHeight, field.TypeInt, value)
	}
	if value, ok := gru.mutation.NumInterferenceSteps(); ok {
		_spec.SetField(generationrealtime.FieldNumInterferenceSteps, field.TypeInt, value)
	}
	if value, ok := gru.mutation.AddedNumInterferenceSteps(); ok {
		_spec.AddField(generationrealtime.FieldNumInterferenceSteps, field.TypeInt, value)
	}
	if value, ok := gru.mutation.UpdatedAt(); ok {
		_spec.SetField(generationrealtime.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gru.mutation.UserTier(); ok {
		_spec.SetField(generationrealtime.FieldUserTier, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{generationrealtime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gru.mutation.done = true
	return n, nil
}

// GenerationRealtimeUpdateOne is the builder for updating a single GenerationRealtime entity.
type GenerationRealtimeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GenerationRealtimeMutation
}

// SetCountryCode sets the "country_code" field.
func (gruo *GenerationRealtimeUpdateOne) SetCountryCode(s string) *GenerationRealtimeUpdateOne {
	gruo.mutation.SetCountryCode(s)
	return gruo
}

// SetDurationMs sets the "duration_ms" field.
func (gruo *GenerationRealtimeUpdateOne) SetDurationMs(i int) *GenerationRealtimeUpdateOne {
	gruo.mutation.ResetDurationMs()
	gruo.mutation.SetDurationMs(i)
	return gruo
}

// AddDurationMs adds i to the "duration_ms" field.
func (gruo *GenerationRealtimeUpdateOne) AddDurationMs(i int) *GenerationRealtimeUpdateOne {
	gruo.mutation.AddDurationMs(i)
	return gruo
}

// SetStatus sets the "status" field.
func (gruo *GenerationRealtimeUpdateOne) SetStatus(ge generationrealtime.Status) *GenerationRealtimeUpdateOne {
	gruo.mutation.SetStatus(ge)
	return gruo
}

// SetUsesDefaultServer sets the "uses_default_server" field.
func (gruo *GenerationRealtimeUpdateOne) SetUsesDefaultServer(b bool) *GenerationRealtimeUpdateOne {
	gruo.mutation.SetUsesDefaultServer(b)
	return gruo
}

// SetWidth sets the "width" field.
func (gruo *GenerationRealtimeUpdateOne) SetWidth(i int) *GenerationRealtimeUpdateOne {
	gruo.mutation.ResetWidth()
	gruo.mutation.SetWidth(i)
	return gruo
}

// AddWidth adds i to the "width" field.
func (gruo *GenerationRealtimeUpdateOne) AddWidth(i int) *GenerationRealtimeUpdateOne {
	gruo.mutation.AddWidth(i)
	return gruo
}

// SetHeight sets the "height" field.
func (gruo *GenerationRealtimeUpdateOne) SetHeight(i int) *GenerationRealtimeUpdateOne {
	gruo.mutation.ResetHeight()
	gruo.mutation.SetHeight(i)
	return gruo
}

// AddHeight adds i to the "height" field.
func (gruo *GenerationRealtimeUpdateOne) AddHeight(i int) *GenerationRealtimeUpdateOne {
	gruo.mutation.AddHeight(i)
	return gruo
}

// SetNumInterferenceSteps sets the "num_interference_steps" field.
func (gruo *GenerationRealtimeUpdateOne) SetNumInterferenceSteps(i int) *GenerationRealtimeUpdateOne {
	gruo.mutation.ResetNumInterferenceSteps()
	gruo.mutation.SetNumInterferenceSteps(i)
	return gruo
}

// AddNumInterferenceSteps adds i to the "num_interference_steps" field.
func (gruo *GenerationRealtimeUpdateOne) AddNumInterferenceSteps(i int) *GenerationRealtimeUpdateOne {
	gruo.mutation.AddNumInterferenceSteps(i)
	return gruo
}

// SetUpdatedAt sets the "updated_at" field.
func (gruo *GenerationRealtimeUpdateOne) SetUpdatedAt(t time.Time) *GenerationRealtimeUpdateOne {
	gruo.mutation.SetUpdatedAt(t)
	return gruo
}

// SetUserTier sets the "user_tier" field.
func (gruo *GenerationRealtimeUpdateOne) SetUserTier(gt generationrealtime.UserTier) *GenerationRealtimeUpdateOne {
	gruo.mutation.SetUserTier(gt)
	return gruo
}

// SetNillableUserTier sets the "user_tier" field if the given value is not nil.
func (gruo *GenerationRealtimeUpdateOne) SetNillableUserTier(gt *generationrealtime.UserTier) *GenerationRealtimeUpdateOne {
	if gt != nil {
		gruo.SetUserTier(*gt)
	}
	return gruo
}

// Mutation returns the GenerationRealtimeMutation object of the builder.
func (gruo *GenerationRealtimeUpdateOne) Mutation() *GenerationRealtimeMutation {
	return gruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gruo *GenerationRealtimeUpdateOne) Select(field string, fields ...string) *GenerationRealtimeUpdateOne {
	gruo.fields = append([]string{field}, fields...)
	return gruo
}

// Save executes the query and returns the updated GenerationRealtime entity.
func (gruo *GenerationRealtimeUpdateOne) Save(ctx context.Context) (*GenerationRealtime, error) {
	gruo.defaults()
	return withHooks[*GenerationRealtime, GenerationRealtimeMutation](ctx, gruo.sqlSave, gruo.mutation, gruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gruo *GenerationRealtimeUpdateOne) SaveX(ctx context.Context) *GenerationRealtime {
	node, err := gruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gruo *GenerationRealtimeUpdateOne) Exec(ctx context.Context) error {
	_, err := gruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gruo *GenerationRealtimeUpdateOne) ExecX(ctx context.Context) {
	if err := gruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gruo *GenerationRealtimeUpdateOne) defaults() {
	if _, ok := gruo.mutation.UpdatedAt(); !ok {
		v := generationrealtime.UpdateDefaultUpdatedAt()
		gruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gruo *GenerationRealtimeUpdateOne) check() error {
	if v, ok := gruo.mutation.Status(); ok {
		if err := generationrealtime.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "GenerationRealtime.status": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.UserTier(); ok {
		if err := generationrealtime.UserTierValidator(v); err != nil {
			return &ValidationError{Name: "user_tier", err: fmt.Errorf(`ent: validator failed for field "GenerationRealtime.user_tier": %w`, err)}
		}
	}
	return nil
}

func (gruo *GenerationRealtimeUpdateOne) sqlSave(ctx context.Context) (_node *GenerationRealtime, err error) {
	if err := gruo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   generationrealtime.Table,
			Columns: generationrealtime.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generationrealtime.FieldID,
			},
		},
	}
	id, ok := gruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GenerationRealtime.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, generationrealtime.FieldID)
		for _, f := range fields {
			if !generationrealtime.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != generationrealtime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gruo.mutation.CountryCode(); ok {
		_spec.SetField(generationrealtime.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := gruo.mutation.DurationMs(); ok {
		_spec.SetField(generationrealtime.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.AddedDurationMs(); ok {
		_spec.AddField(generationrealtime.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.Status(); ok {
		_spec.SetField(generationrealtime.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := gruo.mutation.UsesDefaultServer(); ok {
		_spec.SetField(generationrealtime.FieldUsesDefaultServer, field.TypeBool, value)
	}
	if value, ok := gruo.mutation.Width(); ok {
		_spec.SetField(generationrealtime.FieldWidth, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.AddedWidth(); ok {
		_spec.AddField(generationrealtime.FieldWidth, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.Height(); ok {
		_spec.SetField(generationrealtime.FieldHeight, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.AddedHeight(); ok {
		_spec.AddField(generationrealtime.FieldHeight, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.NumInterferenceSteps(); ok {
		_spec.SetField(generationrealtime.FieldNumInterferenceSteps, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.AddedNumInterferenceSteps(); ok {
		_spec.AddField(generationrealtime.FieldNumInterferenceSteps, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.UpdatedAt(); ok {
		_spec.SetField(generationrealtime.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gruo.mutation.UserTier(); ok {
		_spec.SetField(generationrealtime.FieldUserTier, field.TypeEnum, value)
	}
	_node = &GenerationRealtime{config: gruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{generationrealtime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gruo.mutation.done = true
	return _node, nil
}
