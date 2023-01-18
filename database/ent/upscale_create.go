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
	"github.com/stablecog/go-apps/database/ent/user"
	"github.com/stablecog/go-apps/database/enttypes"
)

// UpscaleCreate is the builder for creating a Upscale entity.
type UpscaleCreate struct {
	config
	mutation *UpscaleMutation
	hooks    []Hook
}

// SetWidth sets the "width" field.
func (uc *UpscaleCreate) SetWidth(i int) *UpscaleCreate {
	uc.mutation.SetWidth(i)
	return uc
}

// SetHeight sets the "height" field.
func (uc *UpscaleCreate) SetHeight(i int) *UpscaleCreate {
	uc.mutation.SetHeight(i)
	return uc
}

// SetScale sets the "scale" field.
func (uc *UpscaleCreate) SetScale(i int) *UpscaleCreate {
	uc.mutation.SetScale(i)
	return uc
}

// SetStatus sets the "status" field.
func (uc *UpscaleCreate) SetStatus(u upscale.Status) *UpscaleCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetServerURL sets the "server_url" field.
func (uc *UpscaleCreate) SetServerURL(s string) *UpscaleCreate {
	uc.mutation.SetServerURL(s)
	return uc
}

// SetDurationMsg sets the "duration_msg" field.
func (uc *UpscaleCreate) SetDurationMsg(i int) *UpscaleCreate {
	uc.mutation.SetDurationMsg(i)
	return uc
}

// SetType sets the "type" field.
func (uc *UpscaleCreate) SetType(s string) *UpscaleCreate {
	uc.mutation.SetType(s)
	return uc
}

// SetPrompt sets the "prompt" field.
func (uc *UpscaleCreate) SetPrompt(s string) *UpscaleCreate {
	uc.mutation.SetPrompt(s)
	return uc
}

// SetNegativePrompt sets the "negative_prompt" field.
func (uc *UpscaleCreate) SetNegativePrompt(s string) *UpscaleCreate {
	uc.mutation.SetNegativePrompt(s)
	return uc
}

// SetSeed sets the "seed" field.
func (uc *UpscaleCreate) SetSeed(ei enttypes.BigInt) *UpscaleCreate {
	uc.mutation.SetSeed(ei)
	return uc
}

// SetNillableSeed sets the "seed" field if the given value is not nil.
func (uc *UpscaleCreate) SetNillableSeed(ei *enttypes.BigInt) *UpscaleCreate {
	if ei != nil {
		uc.SetSeed(*ei)
	}
	return uc
}

// SetNumInferenceSteps sets the "num_inference_steps" field.
func (uc *UpscaleCreate) SetNumInferenceSteps(i int) *UpscaleCreate {
	uc.mutation.SetNumInferenceSteps(i)
	return uc
}

// SetGuidanceScale sets the "guidance_scale" field.
func (uc *UpscaleCreate) SetGuidanceScale(f float64) *UpscaleCreate {
	uc.mutation.SetGuidanceScale(f)
	return uc
}

// SetCountryCode sets the "country_code" field.
func (uc *UpscaleCreate) SetCountryCode(s string) *UpscaleCreate {
	uc.mutation.SetCountryCode(s)
	return uc
}

// SetDeviceType sets the "device_type" field.
func (uc *UpscaleCreate) SetDeviceType(s string) *UpscaleCreate {
	uc.mutation.SetDeviceType(s)
	return uc
}

// SetDeviceOs sets the "device_os" field.
func (uc *UpscaleCreate) SetDeviceOs(s string) *UpscaleCreate {
	uc.mutation.SetDeviceOs(s)
	return uc
}

// SetDeviceBrowser sets the "device_browser" field.
func (uc *UpscaleCreate) SetDeviceBrowser(s string) *UpscaleCreate {
	uc.mutation.SetDeviceBrowser(s)
	return uc
}

// SetUserAgent sets the "user_agent" field.
func (uc *UpscaleCreate) SetUserAgent(s string) *UpscaleCreate {
	uc.mutation.SetUserAgent(s)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UpscaleCreate) SetCreatedAt(t time.Time) *UpscaleCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UpscaleCreate) SetNillableCreatedAt(t *time.Time) *UpscaleCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UpscaleCreate) SetUpdatedAt(t time.Time) *UpscaleCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UpscaleCreate) SetNillableUpdatedAt(t *time.Time) *UpscaleCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetUserID sets the "user_id" field.
func (uc *UpscaleCreate) SetUserID(u uuid.UUID) *UpscaleCreate {
	uc.mutation.SetUserID(u)
	return uc
}

// SetUserTier sets the "user_tier" field.
func (uc *UpscaleCreate) SetUserTier(ut upscale.UserTier) *UpscaleCreate {
	uc.mutation.SetUserTier(ut)
	return uc
}

// SetNillableUserTier sets the "user_tier" field if the given value is not nil.
func (uc *UpscaleCreate) SetNillableUserTier(ut *upscale.UserTier) *UpscaleCreate {
	if ut != nil {
		uc.SetUserTier(*ut)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UpscaleCreate) SetID(u uuid.UUID) *UpscaleCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UpscaleCreate) SetNillableID(u *uuid.UUID) *UpscaleCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// SetUser sets the "user" edge to the User entity.
func (uc *UpscaleCreate) SetUser(u *User) *UpscaleCreate {
	return uc.SetUserID(u.ID)
}

// Mutation returns the UpscaleMutation object of the builder.
func (uc *UpscaleCreate) Mutation() *UpscaleMutation {
	return uc.mutation
}

// Save creates the Upscale in the database.
func (uc *UpscaleCreate) Save(ctx context.Context) (*Upscale, error) {
	uc.defaults()
	return withHooks[*Upscale, UpscaleMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UpscaleCreate) SaveX(ctx context.Context) *Upscale {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UpscaleCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UpscaleCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UpscaleCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := upscale.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := upscale.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.UserTier(); !ok {
		v := upscale.DefaultUserTier
		uc.mutation.SetUserTier(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := upscale.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UpscaleCreate) check() error {
	if _, ok := uc.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "Upscale.width"`)}
	}
	if _, ok := uc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Upscale.height"`)}
	}
	if _, ok := uc.mutation.Scale(); !ok {
		return &ValidationError{Name: "scale", err: errors.New(`ent: missing required field "Upscale.scale"`)}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Upscale.status"`)}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := upscale.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Upscale.status": %w`, err)}
		}
	}
	if _, ok := uc.mutation.ServerURL(); !ok {
		return &ValidationError{Name: "server_url", err: errors.New(`ent: missing required field "Upscale.server_url"`)}
	}
	if _, ok := uc.mutation.DurationMsg(); !ok {
		return &ValidationError{Name: "duration_msg", err: errors.New(`ent: missing required field "Upscale.duration_msg"`)}
	}
	if _, ok := uc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Upscale.type"`)}
	}
	if _, ok := uc.mutation.Prompt(); !ok {
		return &ValidationError{Name: "prompt", err: errors.New(`ent: missing required field "Upscale.prompt"`)}
	}
	if _, ok := uc.mutation.NegativePrompt(); !ok {
		return &ValidationError{Name: "negative_prompt", err: errors.New(`ent: missing required field "Upscale.negative_prompt"`)}
	}
	if _, ok := uc.mutation.NumInferenceSteps(); !ok {
		return &ValidationError{Name: "num_inference_steps", err: errors.New(`ent: missing required field "Upscale.num_inference_steps"`)}
	}
	if _, ok := uc.mutation.GuidanceScale(); !ok {
		return &ValidationError{Name: "guidance_scale", err: errors.New(`ent: missing required field "Upscale.guidance_scale"`)}
	}
	if _, ok := uc.mutation.CountryCode(); !ok {
		return &ValidationError{Name: "country_code", err: errors.New(`ent: missing required field "Upscale.country_code"`)}
	}
	if _, ok := uc.mutation.DeviceType(); !ok {
		return &ValidationError{Name: "device_type", err: errors.New(`ent: missing required field "Upscale.device_type"`)}
	}
	if _, ok := uc.mutation.DeviceOs(); !ok {
		return &ValidationError{Name: "device_os", err: errors.New(`ent: missing required field "Upscale.device_os"`)}
	}
	if _, ok := uc.mutation.DeviceBrowser(); !ok {
		return &ValidationError{Name: "device_browser", err: errors.New(`ent: missing required field "Upscale.device_browser"`)}
	}
	if _, ok := uc.mutation.UserAgent(); !ok {
		return &ValidationError{Name: "user_agent", err: errors.New(`ent: missing required field "Upscale.user_agent"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Upscale.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Upscale.updated_at"`)}
	}
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Upscale.user_id"`)}
	}
	if _, ok := uc.mutation.UserTier(); !ok {
		return &ValidationError{Name: "user_tier", err: errors.New(`ent: missing required field "Upscale.user_tier"`)}
	}
	if v, ok := uc.mutation.UserTier(); ok {
		if err := upscale.UserTierValidator(v); err != nil {
			return &ValidationError{Name: "user_tier", err: fmt.Errorf(`ent: validator failed for field "Upscale.user_tier": %w`, err)}
		}
	}
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Upscale.user"`)}
	}
	return nil
}

func (uc *UpscaleCreate) sqlSave(ctx context.Context) (*Upscale, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
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
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UpscaleCreate) createSpec() (*Upscale, *sqlgraph.CreateSpec) {
	var (
		_node = &Upscale{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: upscale.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscale.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.Width(); ok {
		_spec.SetField(upscale.FieldWidth, field.TypeInt, value)
		_node.Width = value
	}
	if value, ok := uc.mutation.Height(); ok {
		_spec.SetField(upscale.FieldHeight, field.TypeInt, value)
		_node.Height = value
	}
	if value, ok := uc.mutation.Scale(); ok {
		_spec.SetField(upscale.FieldScale, field.TypeInt, value)
		_node.Scale = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(upscale.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.ServerURL(); ok {
		_spec.SetField(upscale.FieldServerURL, field.TypeString, value)
		_node.ServerURL = value
	}
	if value, ok := uc.mutation.DurationMsg(); ok {
		_spec.SetField(upscale.FieldDurationMsg, field.TypeInt, value)
		_node.DurationMsg = &value
	}
	if value, ok := uc.mutation.GetType(); ok {
		_spec.SetField(upscale.FieldType, field.TypeString, value)
		_node.Type = &value
	}
	if value, ok := uc.mutation.Prompt(); ok {
		_spec.SetField(upscale.FieldPrompt, field.TypeString, value)
		_node.Prompt = &value
	}
	if value, ok := uc.mutation.NegativePrompt(); ok {
		_spec.SetField(upscale.FieldNegativePrompt, field.TypeString, value)
		_node.NegativePrompt = &value
	}
	if value, ok := uc.mutation.Seed(); ok {
		_spec.SetField(upscale.FieldSeed, field.TypeInt, value)
		_node.Seed = &value
	}
	if value, ok := uc.mutation.NumInferenceSteps(); ok {
		_spec.SetField(upscale.FieldNumInferenceSteps, field.TypeInt, value)
		_node.NumInferenceSteps = &value
	}
	if value, ok := uc.mutation.GuidanceScale(); ok {
		_spec.SetField(upscale.FieldGuidanceScale, field.TypeFloat64, value)
		_node.GuidanceScale = &value
	}
	if value, ok := uc.mutation.CountryCode(); ok {
		_spec.SetField(upscale.FieldCountryCode, field.TypeString, value)
		_node.CountryCode = &value
	}
	if value, ok := uc.mutation.DeviceType(); ok {
		_spec.SetField(upscale.FieldDeviceType, field.TypeString, value)
		_node.DeviceType = &value
	}
	if value, ok := uc.mutation.DeviceOs(); ok {
		_spec.SetField(upscale.FieldDeviceOs, field.TypeString, value)
		_node.DeviceOs = &value
	}
	if value, ok := uc.mutation.DeviceBrowser(); ok {
		_spec.SetField(upscale.FieldDeviceBrowser, field.TypeString, value)
		_node.DeviceBrowser = &value
	}
	if value, ok := uc.mutation.UserAgent(); ok {
		_spec.SetField(upscale.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = &value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(upscale.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(upscale.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.UserTier(); ok {
		_spec.SetField(upscale.FieldUserTier, field.TypeEnum, value)
		_node.UserTier = value
	}
	if nodes := uc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UserTable,
			Columns: []string{upscale.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UpscaleCreateBulk is the builder for creating many Upscale entities in bulk.
type UpscaleCreateBulk struct {
	config
	builders []*UpscaleCreate
}

// Save creates the Upscale entities in the database.
func (ucb *UpscaleCreateBulk) Save(ctx context.Context) ([]*Upscale, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Upscale, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UpscaleMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UpscaleCreateBulk) SaveX(ctx context.Context) []*Upscale {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UpscaleCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UpscaleCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}