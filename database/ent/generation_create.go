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
	"github.com/stablecog/go-apps/database/ent/generation"
	"github.com/stablecog/go-apps/database/ent/model"
	"github.com/stablecog/go-apps/database/ent/negativeprompt"
	"github.com/stablecog/go-apps/database/ent/prompt"
	"github.com/stablecog/go-apps/database/ent/scheduler"
	"github.com/stablecog/go-apps/database/ent/user"
	"github.com/stablecog/go-apps/database/enttypes"
)

// GenerationCreate is the builder for creating a Generation entity.
type GenerationCreate struct {
	config
	mutation *GenerationMutation
	hooks    []Hook
}

// SetPromptID sets the "prompt_id" field.
func (gc *GenerationCreate) SetPromptID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetPromptID(u)
	return gc
}

// SetNegativePromptID sets the "negative_prompt_id" field.
func (gc *GenerationCreate) SetNegativePromptID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetNegativePromptID(u)
	return gc
}

// SetModelID sets the "model_id" field.
func (gc *GenerationCreate) SetModelID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetModelID(u)
	return gc
}

// SetImageID sets the "image_id" field.
func (gc *GenerationCreate) SetImageID(s string) *GenerationCreate {
	gc.mutation.SetImageID(s)
	return gc
}

// SetWidth sets the "width" field.
func (gc *GenerationCreate) SetWidth(i int) *GenerationCreate {
	gc.mutation.SetWidth(i)
	return gc
}

// SetHeight sets the "height" field.
func (gc *GenerationCreate) SetHeight(i int) *GenerationCreate {
	gc.mutation.SetHeight(i)
	return gc
}

// SetSeed sets the "seed" field.
func (gc *GenerationCreate) SetSeed(ei enttypes.BigInt) *GenerationCreate {
	gc.mutation.SetSeed(ei)
	return gc
}

// SetNillableSeed sets the "seed" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableSeed(ei *enttypes.BigInt) *GenerationCreate {
	if ei != nil {
		gc.SetSeed(*ei)
	}
	return gc
}

// SetNumInferenceSteps sets the "num_inference_steps" field.
func (gc *GenerationCreate) SetNumInferenceSteps(i int) *GenerationCreate {
	gc.mutation.SetNumInferenceSteps(i)
	return gc
}

// SetGuidanceScale sets the "guidance_scale" field.
func (gc *GenerationCreate) SetGuidanceScale(f float64) *GenerationCreate {
	gc.mutation.SetGuidanceScale(f)
	return gc
}

// SetHidden sets the "hidden" field.
func (gc *GenerationCreate) SetHidden(b bool) *GenerationCreate {
	gc.mutation.SetHidden(b)
	return gc
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableHidden(b *bool) *GenerationCreate {
	if b != nil {
		gc.SetHidden(*b)
	}
	return gc
}

// SetSchedulerID sets the "scheduler_id" field.
func (gc *GenerationCreate) SetSchedulerID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetSchedulerID(u)
	return gc
}

// SetUserID sets the "user_id" field.
func (gc *GenerationCreate) SetUserID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetUserID(u)
	return gc
}

// SetUserTier sets the "user_tier" field.
func (gc *GenerationCreate) SetUserTier(gt generation.UserTier) *GenerationCreate {
	gc.mutation.SetUserTier(gt)
	return gc
}

// SetNillableUserTier sets the "user_tier" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableUserTier(gt *generation.UserTier) *GenerationCreate {
	if gt != nil {
		gc.SetUserTier(*gt)
	}
	return gc
}

// SetServerURL sets the "server_url" field.
func (gc *GenerationCreate) SetServerURL(s string) *GenerationCreate {
	gc.mutation.SetServerURL(s)
	return gc
}

// SetCountryCode sets the "country_code" field.
func (gc *GenerationCreate) SetCountryCode(s string) *GenerationCreate {
	gc.mutation.SetCountryCode(s)
	return gc
}

// SetDeviceType sets the "device_type" field.
func (gc *GenerationCreate) SetDeviceType(s string) *GenerationCreate {
	gc.mutation.SetDeviceType(s)
	return gc
}

// SetDeviceOs sets the "device_os" field.
func (gc *GenerationCreate) SetDeviceOs(s string) *GenerationCreate {
	gc.mutation.SetDeviceOs(s)
	return gc
}

// SetDeviceBrowser sets the "device_browser" field.
func (gc *GenerationCreate) SetDeviceBrowser(s string) *GenerationCreate {
	gc.mutation.SetDeviceBrowser(s)
	return gc
}

// SetUserAgent sets the "user_agent" field.
func (gc *GenerationCreate) SetUserAgent(s string) *GenerationCreate {
	gc.mutation.SetUserAgent(s)
	return gc
}

// SetDurationMs sets the "duration_ms" field.
func (gc *GenerationCreate) SetDurationMs(i int) *GenerationCreate {
	gc.mutation.SetDurationMs(i)
	return gc
}

// SetStatus sets the "status" field.
func (gc *GenerationCreate) SetStatus(ge generation.Status) *GenerationCreate {
	gc.mutation.SetStatus(ge)
	return gc
}

// SetFailureReason sets the "failure_reason" field.
func (gc *GenerationCreate) SetFailureReason(s string) *GenerationCreate {
	gc.mutation.SetFailureReason(s)
	return gc
}

// SetImageObjectName sets the "image_object_name" field.
func (gc *GenerationCreate) SetImageObjectName(s string) *GenerationCreate {
	gc.mutation.SetImageObjectName(s)
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GenerationCreate) SetCreatedAt(t time.Time) *GenerationCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableCreatedAt(t *time.Time) *GenerationCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GenerationCreate) SetUpdatedAt(t time.Time) *GenerationCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableUpdatedAt(t *time.Time) *GenerationCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GenerationCreate) SetID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetID(u)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableID(u *uuid.UUID) *GenerationCreate {
	if u != nil {
		gc.SetID(*u)
	}
	return gc
}

// SetUser sets the "user" edge to the User entity.
func (gc *GenerationCreate) SetUser(u *User) *GenerationCreate {
	return gc.SetUserID(u.ID)
}

// SetModel sets the "model" edge to the Model entity.
func (gc *GenerationCreate) SetModel(m *Model) *GenerationCreate {
	return gc.SetModelID(m.ID)
}

// SetPrompt sets the "prompt" edge to the Prompt entity.
func (gc *GenerationCreate) SetPrompt(p *Prompt) *GenerationCreate {
	return gc.SetPromptID(p.ID)
}

// SetNegativePrompt sets the "negative_prompt" edge to the NegativePrompt entity.
func (gc *GenerationCreate) SetNegativePrompt(n *NegativePrompt) *GenerationCreate {
	return gc.SetNegativePromptID(n.ID)
}

// SetScheduler sets the "scheduler" edge to the Scheduler entity.
func (gc *GenerationCreate) SetScheduler(s *Scheduler) *GenerationCreate {
	return gc.SetSchedulerID(s.ID)
}

// Mutation returns the GenerationMutation object of the builder.
func (gc *GenerationCreate) Mutation() *GenerationMutation {
	return gc.mutation
}

// Save creates the Generation in the database.
func (gc *GenerationCreate) Save(ctx context.Context) (*Generation, error) {
	gc.defaults()
	return withHooks[*Generation, GenerationMutation](ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GenerationCreate) SaveX(ctx context.Context) *Generation {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GenerationCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GenerationCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GenerationCreate) defaults() {
	if _, ok := gc.mutation.Hidden(); !ok {
		v := generation.DefaultHidden
		gc.mutation.SetHidden(v)
	}
	if _, ok := gc.mutation.UserTier(); !ok {
		v := generation.DefaultUserTier
		gc.mutation.SetUserTier(v)
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := generation.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := generation.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := generation.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GenerationCreate) check() error {
	if _, ok := gc.mutation.PromptID(); !ok {
		return &ValidationError{Name: "prompt_id", err: errors.New(`ent: missing required field "Generation.prompt_id"`)}
	}
	if _, ok := gc.mutation.NegativePromptID(); !ok {
		return &ValidationError{Name: "negative_prompt_id", err: errors.New(`ent: missing required field "Generation.negative_prompt_id"`)}
	}
	if _, ok := gc.mutation.ModelID(); !ok {
		return &ValidationError{Name: "model_id", err: errors.New(`ent: missing required field "Generation.model_id"`)}
	}
	if _, ok := gc.mutation.ImageID(); !ok {
		return &ValidationError{Name: "image_id", err: errors.New(`ent: missing required field "Generation.image_id"`)}
	}
	if _, ok := gc.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "Generation.width"`)}
	}
	if _, ok := gc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Generation.height"`)}
	}
	if _, ok := gc.mutation.NumInferenceSteps(); !ok {
		return &ValidationError{Name: "num_inference_steps", err: errors.New(`ent: missing required field "Generation.num_inference_steps"`)}
	}
	if _, ok := gc.mutation.GuidanceScale(); !ok {
		return &ValidationError{Name: "guidance_scale", err: errors.New(`ent: missing required field "Generation.guidance_scale"`)}
	}
	if _, ok := gc.mutation.Hidden(); !ok {
		return &ValidationError{Name: "hidden", err: errors.New(`ent: missing required field "Generation.hidden"`)}
	}
	if _, ok := gc.mutation.SchedulerID(); !ok {
		return &ValidationError{Name: "scheduler_id", err: errors.New(`ent: missing required field "Generation.scheduler_id"`)}
	}
	if _, ok := gc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Generation.user_id"`)}
	}
	if _, ok := gc.mutation.UserTier(); !ok {
		return &ValidationError{Name: "user_tier", err: errors.New(`ent: missing required field "Generation.user_tier"`)}
	}
	if v, ok := gc.mutation.UserTier(); ok {
		if err := generation.UserTierValidator(v); err != nil {
			return &ValidationError{Name: "user_tier", err: fmt.Errorf(`ent: validator failed for field "Generation.user_tier": %w`, err)}
		}
	}
	if _, ok := gc.mutation.ServerURL(); !ok {
		return &ValidationError{Name: "server_url", err: errors.New(`ent: missing required field "Generation.server_url"`)}
	}
	if _, ok := gc.mutation.CountryCode(); !ok {
		return &ValidationError{Name: "country_code", err: errors.New(`ent: missing required field "Generation.country_code"`)}
	}
	if _, ok := gc.mutation.DeviceType(); !ok {
		return &ValidationError{Name: "device_type", err: errors.New(`ent: missing required field "Generation.device_type"`)}
	}
	if _, ok := gc.mutation.DeviceOs(); !ok {
		return &ValidationError{Name: "device_os", err: errors.New(`ent: missing required field "Generation.device_os"`)}
	}
	if _, ok := gc.mutation.DeviceBrowser(); !ok {
		return &ValidationError{Name: "device_browser", err: errors.New(`ent: missing required field "Generation.device_browser"`)}
	}
	if _, ok := gc.mutation.UserAgent(); !ok {
		return &ValidationError{Name: "user_agent", err: errors.New(`ent: missing required field "Generation.user_agent"`)}
	}
	if _, ok := gc.mutation.DurationMs(); !ok {
		return &ValidationError{Name: "duration_ms", err: errors.New(`ent: missing required field "Generation.duration_ms"`)}
	}
	if _, ok := gc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Generation.status"`)}
	}
	if v, ok := gc.mutation.Status(); ok {
		if err := generation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Generation.status": %w`, err)}
		}
	}
	if _, ok := gc.mutation.FailureReason(); !ok {
		return &ValidationError{Name: "failure_reason", err: errors.New(`ent: missing required field "Generation.failure_reason"`)}
	}
	if _, ok := gc.mutation.ImageObjectName(); !ok {
		return &ValidationError{Name: "image_object_name", err: errors.New(`ent: missing required field "Generation.image_object_name"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Generation.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Generation.updated_at"`)}
	}
	if _, ok := gc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Generation.user"`)}
	}
	if _, ok := gc.mutation.ModelID(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required edge "Generation.model"`)}
	}
	if _, ok := gc.mutation.PromptID(); !ok {
		return &ValidationError{Name: "prompt", err: errors.New(`ent: missing required edge "Generation.prompt"`)}
	}
	if _, ok := gc.mutation.NegativePromptID(); !ok {
		return &ValidationError{Name: "negative_prompt", err: errors.New(`ent: missing required edge "Generation.negative_prompt"`)}
	}
	if _, ok := gc.mutation.SchedulerID(); !ok {
		return &ValidationError{Name: "scheduler", err: errors.New(`ent: missing required edge "Generation.scheduler"`)}
	}
	return nil
}

func (gc *GenerationCreate) sqlSave(ctx context.Context) (*Generation, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
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
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GenerationCreate) createSpec() (*Generation, *sqlgraph.CreateSpec) {
	var (
		_node = &Generation{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: generation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generation.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gc.mutation.ImageID(); ok {
		_spec.SetField(generation.FieldImageID, field.TypeString, value)
		_node.ImageID = value
	}
	if value, ok := gc.mutation.Width(); ok {
		_spec.SetField(generation.FieldWidth, field.TypeInt, value)
		_node.Width = value
	}
	if value, ok := gc.mutation.Height(); ok {
		_spec.SetField(generation.FieldHeight, field.TypeInt, value)
		_node.Height = value
	}
	if value, ok := gc.mutation.Seed(); ok {
		_spec.SetField(generation.FieldSeed, field.TypeInt, value)
		_node.Seed = &value
	}
	if value, ok := gc.mutation.NumInferenceSteps(); ok {
		_spec.SetField(generation.FieldNumInferenceSteps, field.TypeInt, value)
		_node.NumInferenceSteps = &value
	}
	if value, ok := gc.mutation.GuidanceScale(); ok {
		_spec.SetField(generation.FieldGuidanceScale, field.TypeFloat64, value)
		_node.GuidanceScale = value
	}
	if value, ok := gc.mutation.Hidden(); ok {
		_spec.SetField(generation.FieldHidden, field.TypeBool, value)
		_node.Hidden = value
	}
	if value, ok := gc.mutation.UserTier(); ok {
		_spec.SetField(generation.FieldUserTier, field.TypeEnum, value)
		_node.UserTier = value
	}
	if value, ok := gc.mutation.ServerURL(); ok {
		_spec.SetField(generation.FieldServerURL, field.TypeString, value)
		_node.ServerURL = value
	}
	if value, ok := gc.mutation.CountryCode(); ok {
		_spec.SetField(generation.FieldCountryCode, field.TypeString, value)
		_node.CountryCode = &value
	}
	if value, ok := gc.mutation.DeviceType(); ok {
		_spec.SetField(generation.FieldDeviceType, field.TypeString, value)
		_node.DeviceType = &value
	}
	if value, ok := gc.mutation.DeviceOs(); ok {
		_spec.SetField(generation.FieldDeviceOs, field.TypeString, value)
		_node.DeviceOs = &value
	}
	if value, ok := gc.mutation.DeviceBrowser(); ok {
		_spec.SetField(generation.FieldDeviceBrowser, field.TypeString, value)
		_node.DeviceBrowser = &value
	}
	if value, ok := gc.mutation.UserAgent(); ok {
		_spec.SetField(generation.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = &value
	}
	if value, ok := gc.mutation.DurationMs(); ok {
		_spec.SetField(generation.FieldDurationMs, field.TypeInt, value)
		_node.DurationMs = &value
	}
	if value, ok := gc.mutation.Status(); ok {
		_spec.SetField(generation.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := gc.mutation.FailureReason(); ok {
		_spec.SetField(generation.FieldFailureReason, field.TypeString, value)
		_node.FailureReason = &value
	}
	if value, ok := gc.mutation.ImageObjectName(); ok {
		_spec.SetField(generation.FieldImageObjectName, field.TypeString, value)
		_node.ImageObjectName = &value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(generation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(generation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := gc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.UserTable,
			Columns: []string{generation.UserColumn},
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
	if nodes := gc.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.ModelTable,
			Columns: []string{generation.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: model.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ModelID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.PromptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.PromptTable,
			Columns: []string{generation.PromptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prompt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PromptID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.NegativePromptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.NegativePromptTable,
			Columns: []string{generation.NegativePromptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: negativeprompt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NegativePromptID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.SchedulerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.SchedulerTable,
			Columns: []string{generation.SchedulerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scheduler.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SchedulerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GenerationCreateBulk is the builder for creating many Generation entities in bulk.
type GenerationCreateBulk struct {
	config
	builders []*GenerationCreate
}

// Save creates the Generation entities in the database.
func (gcb *GenerationCreateBulk) Save(ctx context.Context) ([]*Generation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Generation, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GenerationMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GenerationCreateBulk) SaveX(ctx context.Context) []*Generation {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GenerationCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GenerationCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}