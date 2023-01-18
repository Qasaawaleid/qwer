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
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
	"github.com/yekta/stablecog/go-apps/database/ent/upscale"
	"github.com/yekta/stablecog/go-apps/database/ent/user"
	"github.com/yekta/stablecog/go-apps/database/enttypes"
)

// UpscaleUpdate is the builder for updating Upscale entities.
type UpscaleUpdate struct {
	config
	hooks    []Hook
	mutation *UpscaleMutation
}

// Where appends a list predicates to the UpscaleUpdate builder.
func (uu *UpscaleUpdate) Where(ps ...predicate.Upscale) *UpscaleUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetWidth sets the "width" field.
func (uu *UpscaleUpdate) SetWidth(i int) *UpscaleUpdate {
	uu.mutation.ResetWidth()
	uu.mutation.SetWidth(i)
	return uu
}

// AddWidth adds i to the "width" field.
func (uu *UpscaleUpdate) AddWidth(i int) *UpscaleUpdate {
	uu.mutation.AddWidth(i)
	return uu
}

// SetHeight sets the "height" field.
func (uu *UpscaleUpdate) SetHeight(i int) *UpscaleUpdate {
	uu.mutation.ResetHeight()
	uu.mutation.SetHeight(i)
	return uu
}

// AddHeight adds i to the "height" field.
func (uu *UpscaleUpdate) AddHeight(i int) *UpscaleUpdate {
	uu.mutation.AddHeight(i)
	return uu
}

// SetScale sets the "scale" field.
func (uu *UpscaleUpdate) SetScale(i int) *UpscaleUpdate {
	uu.mutation.ResetScale()
	uu.mutation.SetScale(i)
	return uu
}

// AddScale adds i to the "scale" field.
func (uu *UpscaleUpdate) AddScale(i int) *UpscaleUpdate {
	uu.mutation.AddScale(i)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UpscaleUpdate) SetStatus(u upscale.Status) *UpscaleUpdate {
	uu.mutation.SetStatus(u)
	return uu
}

// SetServerURL sets the "server_url" field.
func (uu *UpscaleUpdate) SetServerURL(s string) *UpscaleUpdate {
	uu.mutation.SetServerURL(s)
	return uu
}

// SetDurationMsg sets the "duration_msg" field.
func (uu *UpscaleUpdate) SetDurationMsg(i int) *UpscaleUpdate {
	uu.mutation.ResetDurationMsg()
	uu.mutation.SetDurationMsg(i)
	return uu
}

// AddDurationMsg adds i to the "duration_msg" field.
func (uu *UpscaleUpdate) AddDurationMsg(i int) *UpscaleUpdate {
	uu.mutation.AddDurationMsg(i)
	return uu
}

// SetType sets the "type" field.
func (uu *UpscaleUpdate) SetType(s string) *UpscaleUpdate {
	uu.mutation.SetType(s)
	return uu
}

// SetPrompt sets the "prompt" field.
func (uu *UpscaleUpdate) SetPrompt(s string) *UpscaleUpdate {
	uu.mutation.SetPrompt(s)
	return uu
}

// SetNegativePrompt sets the "negative_prompt" field.
func (uu *UpscaleUpdate) SetNegativePrompt(s string) *UpscaleUpdate {
	uu.mutation.SetNegativePrompt(s)
	return uu
}

// SetSeed sets the "seed" field.
func (uu *UpscaleUpdate) SetSeed(ei enttypes.BigInt) *UpscaleUpdate {
	uu.mutation.ResetSeed()
	uu.mutation.SetSeed(ei)
	return uu
}

// SetNillableSeed sets the "seed" field if the given value is not nil.
func (uu *UpscaleUpdate) SetNillableSeed(ei *enttypes.BigInt) *UpscaleUpdate {
	if ei != nil {
		uu.SetSeed(*ei)
	}
	return uu
}

// AddSeed adds ei to the "seed" field.
func (uu *UpscaleUpdate) AddSeed(ei enttypes.BigInt) *UpscaleUpdate {
	uu.mutation.AddSeed(ei)
	return uu
}

// ClearSeed clears the value of the "seed" field.
func (uu *UpscaleUpdate) ClearSeed() *UpscaleUpdate {
	uu.mutation.ClearSeed()
	return uu
}

// SetNumInferenceSteps sets the "num_inference_steps" field.
func (uu *UpscaleUpdate) SetNumInferenceSteps(i int) *UpscaleUpdate {
	uu.mutation.ResetNumInferenceSteps()
	uu.mutation.SetNumInferenceSteps(i)
	return uu
}

// AddNumInferenceSteps adds i to the "num_inference_steps" field.
func (uu *UpscaleUpdate) AddNumInferenceSteps(i int) *UpscaleUpdate {
	uu.mutation.AddNumInferenceSteps(i)
	return uu
}

// SetGuidanceScale sets the "guidance_scale" field.
func (uu *UpscaleUpdate) SetGuidanceScale(f float64) *UpscaleUpdate {
	uu.mutation.ResetGuidanceScale()
	uu.mutation.SetGuidanceScale(f)
	return uu
}

// AddGuidanceScale adds f to the "guidance_scale" field.
func (uu *UpscaleUpdate) AddGuidanceScale(f float64) *UpscaleUpdate {
	uu.mutation.AddGuidanceScale(f)
	return uu
}

// SetCountryCode sets the "country_code" field.
func (uu *UpscaleUpdate) SetCountryCode(s string) *UpscaleUpdate {
	uu.mutation.SetCountryCode(s)
	return uu
}

// SetDeviceType sets the "device_type" field.
func (uu *UpscaleUpdate) SetDeviceType(s string) *UpscaleUpdate {
	uu.mutation.SetDeviceType(s)
	return uu
}

// SetDeviceOs sets the "device_os" field.
func (uu *UpscaleUpdate) SetDeviceOs(s string) *UpscaleUpdate {
	uu.mutation.SetDeviceOs(s)
	return uu
}

// SetDeviceBrowser sets the "device_browser" field.
func (uu *UpscaleUpdate) SetDeviceBrowser(s string) *UpscaleUpdate {
	uu.mutation.SetDeviceBrowser(s)
	return uu
}

// SetUserAgent sets the "user_agent" field.
func (uu *UpscaleUpdate) SetUserAgent(s string) *UpscaleUpdate {
	uu.mutation.SetUserAgent(s)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UpscaleUpdate) SetUpdatedAt(t time.Time) *UpscaleUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetUserID sets the "user_id" field.
func (uu *UpscaleUpdate) SetUserID(u uuid.UUID) *UpscaleUpdate {
	uu.mutation.SetUserID(u)
	return uu
}

// SetUserTier sets the "user_tier" field.
func (uu *UpscaleUpdate) SetUserTier(ut upscale.UserTier) *UpscaleUpdate {
	uu.mutation.SetUserTier(ut)
	return uu
}

// SetNillableUserTier sets the "user_tier" field if the given value is not nil.
func (uu *UpscaleUpdate) SetNillableUserTier(ut *upscale.UserTier) *UpscaleUpdate {
	if ut != nil {
		uu.SetUserTier(*ut)
	}
	return uu
}

// SetUser sets the "user" edge to the User entity.
func (uu *UpscaleUpdate) SetUser(u *User) *UpscaleUpdate {
	return uu.SetUserID(u.ID)
}

// Mutation returns the UpscaleMutation object of the builder.
func (uu *UpscaleUpdate) Mutation() *UpscaleMutation {
	return uu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uu *UpscaleUpdate) ClearUser() *UpscaleUpdate {
	uu.mutation.ClearUser()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UpscaleUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UpscaleMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UpscaleUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UpscaleUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UpscaleUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UpscaleUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := upscale.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UpscaleUpdate) check() error {
	if v, ok := uu.mutation.Status(); ok {
		if err := upscale.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Upscale.status": %w`, err)}
		}
	}
	if v, ok := uu.mutation.UserTier(); ok {
		if err := upscale.UserTierValidator(v); err != nil {
			return &ValidationError{Name: "user_tier", err: fmt.Errorf(`ent: validator failed for field "Upscale.user_tier": %w`, err)}
		}
	}
	if _, ok := uu.mutation.UserID(); uu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upscale.user"`)
	}
	return nil
}

func (uu *UpscaleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscale.Table,
			Columns: upscale.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscale.FieldID,
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
	if value, ok := uu.mutation.Width(); ok {
		_spec.SetField(upscale.FieldWidth, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedWidth(); ok {
		_spec.AddField(upscale.FieldWidth, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Height(); ok {
		_spec.SetField(upscale.FieldHeight, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedHeight(); ok {
		_spec.AddField(upscale.FieldHeight, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Scale(); ok {
		_spec.SetField(upscale.FieldScale, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedScale(); ok {
		_spec.AddField(upscale.FieldScale, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(upscale.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.ServerURL(); ok {
		_spec.SetField(upscale.FieldServerURL, field.TypeString, value)
	}
	if value, ok := uu.mutation.DurationMsg(); ok {
		_spec.SetField(upscale.FieldDurationMsg, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedDurationMsg(); ok {
		_spec.AddField(upscale.FieldDurationMsg, field.TypeInt, value)
	}
	if value, ok := uu.mutation.GetType(); ok {
		_spec.SetField(upscale.FieldType, field.TypeString, value)
	}
	if value, ok := uu.mutation.Prompt(); ok {
		_spec.SetField(upscale.FieldPrompt, field.TypeString, value)
	}
	if value, ok := uu.mutation.NegativePrompt(); ok {
		_spec.SetField(upscale.FieldNegativePrompt, field.TypeString, value)
	}
	if value, ok := uu.mutation.Seed(); ok {
		_spec.SetField(upscale.FieldSeed, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedSeed(); ok {
		_spec.AddField(upscale.FieldSeed, field.TypeInt, value)
	}
	if uu.mutation.SeedCleared() {
		_spec.ClearField(upscale.FieldSeed, field.TypeInt)
	}
	if value, ok := uu.mutation.NumInferenceSteps(); ok {
		_spec.SetField(upscale.FieldNumInferenceSteps, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedNumInferenceSteps(); ok {
		_spec.AddField(upscale.FieldNumInferenceSteps, field.TypeInt, value)
	}
	if value, ok := uu.mutation.GuidanceScale(); ok {
		_spec.SetField(upscale.FieldGuidanceScale, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.AddedGuidanceScale(); ok {
		_spec.AddField(upscale.FieldGuidanceScale, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.CountryCode(); ok {
		_spec.SetField(upscale.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := uu.mutation.DeviceType(); ok {
		_spec.SetField(upscale.FieldDeviceType, field.TypeString, value)
	}
	if value, ok := uu.mutation.DeviceOs(); ok {
		_spec.SetField(upscale.FieldDeviceOs, field.TypeString, value)
	}
	if value, ok := uu.mutation.DeviceBrowser(); ok {
		_spec.SetField(upscale.FieldDeviceBrowser, field.TypeString, value)
	}
	if value, ok := uu.mutation.UserAgent(); ok {
		_spec.SetField(upscale.FieldUserAgent, field.TypeString, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(upscale.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UserTier(); ok {
		_spec.SetField(upscale.FieldUserTier, field.TypeEnum, value)
	}
	if uu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upscale.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UpscaleUpdateOne is the builder for updating a single Upscale entity.
type UpscaleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UpscaleMutation
}

// SetWidth sets the "width" field.
func (uuo *UpscaleUpdateOne) SetWidth(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetWidth()
	uuo.mutation.SetWidth(i)
	return uuo
}

// AddWidth adds i to the "width" field.
func (uuo *UpscaleUpdateOne) AddWidth(i int) *UpscaleUpdateOne {
	uuo.mutation.AddWidth(i)
	return uuo
}

// SetHeight sets the "height" field.
func (uuo *UpscaleUpdateOne) SetHeight(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetHeight()
	uuo.mutation.SetHeight(i)
	return uuo
}

// AddHeight adds i to the "height" field.
func (uuo *UpscaleUpdateOne) AddHeight(i int) *UpscaleUpdateOne {
	uuo.mutation.AddHeight(i)
	return uuo
}

// SetScale sets the "scale" field.
func (uuo *UpscaleUpdateOne) SetScale(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetScale()
	uuo.mutation.SetScale(i)
	return uuo
}

// AddScale adds i to the "scale" field.
func (uuo *UpscaleUpdateOne) AddScale(i int) *UpscaleUpdateOne {
	uuo.mutation.AddScale(i)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UpscaleUpdateOne) SetStatus(u upscale.Status) *UpscaleUpdateOne {
	uuo.mutation.SetStatus(u)
	return uuo
}

// SetServerURL sets the "server_url" field.
func (uuo *UpscaleUpdateOne) SetServerURL(s string) *UpscaleUpdateOne {
	uuo.mutation.SetServerURL(s)
	return uuo
}

// SetDurationMsg sets the "duration_msg" field.
func (uuo *UpscaleUpdateOne) SetDurationMsg(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetDurationMsg()
	uuo.mutation.SetDurationMsg(i)
	return uuo
}

// AddDurationMsg adds i to the "duration_msg" field.
func (uuo *UpscaleUpdateOne) AddDurationMsg(i int) *UpscaleUpdateOne {
	uuo.mutation.AddDurationMsg(i)
	return uuo
}

// SetType sets the "type" field.
func (uuo *UpscaleUpdateOne) SetType(s string) *UpscaleUpdateOne {
	uuo.mutation.SetType(s)
	return uuo
}

// SetPrompt sets the "prompt" field.
func (uuo *UpscaleUpdateOne) SetPrompt(s string) *UpscaleUpdateOne {
	uuo.mutation.SetPrompt(s)
	return uuo
}

// SetNegativePrompt sets the "negative_prompt" field.
func (uuo *UpscaleUpdateOne) SetNegativePrompt(s string) *UpscaleUpdateOne {
	uuo.mutation.SetNegativePrompt(s)
	return uuo
}

// SetSeed sets the "seed" field.
func (uuo *UpscaleUpdateOne) SetSeed(ei enttypes.BigInt) *UpscaleUpdateOne {
	uuo.mutation.ResetSeed()
	uuo.mutation.SetSeed(ei)
	return uuo
}

// SetNillableSeed sets the "seed" field if the given value is not nil.
func (uuo *UpscaleUpdateOne) SetNillableSeed(ei *enttypes.BigInt) *UpscaleUpdateOne {
	if ei != nil {
		uuo.SetSeed(*ei)
	}
	return uuo
}

// AddSeed adds ei to the "seed" field.
func (uuo *UpscaleUpdateOne) AddSeed(ei enttypes.BigInt) *UpscaleUpdateOne {
	uuo.mutation.AddSeed(ei)
	return uuo
}

// ClearSeed clears the value of the "seed" field.
func (uuo *UpscaleUpdateOne) ClearSeed() *UpscaleUpdateOne {
	uuo.mutation.ClearSeed()
	return uuo
}

// SetNumInferenceSteps sets the "num_inference_steps" field.
func (uuo *UpscaleUpdateOne) SetNumInferenceSteps(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetNumInferenceSteps()
	uuo.mutation.SetNumInferenceSteps(i)
	return uuo
}

// AddNumInferenceSteps adds i to the "num_inference_steps" field.
func (uuo *UpscaleUpdateOne) AddNumInferenceSteps(i int) *UpscaleUpdateOne {
	uuo.mutation.AddNumInferenceSteps(i)
	return uuo
}

// SetGuidanceScale sets the "guidance_scale" field.
func (uuo *UpscaleUpdateOne) SetGuidanceScale(f float64) *UpscaleUpdateOne {
	uuo.mutation.ResetGuidanceScale()
	uuo.mutation.SetGuidanceScale(f)
	return uuo
}

// AddGuidanceScale adds f to the "guidance_scale" field.
func (uuo *UpscaleUpdateOne) AddGuidanceScale(f float64) *UpscaleUpdateOne {
	uuo.mutation.AddGuidanceScale(f)
	return uuo
}

// SetCountryCode sets the "country_code" field.
func (uuo *UpscaleUpdateOne) SetCountryCode(s string) *UpscaleUpdateOne {
	uuo.mutation.SetCountryCode(s)
	return uuo
}

// SetDeviceType sets the "device_type" field.
func (uuo *UpscaleUpdateOne) SetDeviceType(s string) *UpscaleUpdateOne {
	uuo.mutation.SetDeviceType(s)
	return uuo
}

// SetDeviceOs sets the "device_os" field.
func (uuo *UpscaleUpdateOne) SetDeviceOs(s string) *UpscaleUpdateOne {
	uuo.mutation.SetDeviceOs(s)
	return uuo
}

// SetDeviceBrowser sets the "device_browser" field.
func (uuo *UpscaleUpdateOne) SetDeviceBrowser(s string) *UpscaleUpdateOne {
	uuo.mutation.SetDeviceBrowser(s)
	return uuo
}

// SetUserAgent sets the "user_agent" field.
func (uuo *UpscaleUpdateOne) SetUserAgent(s string) *UpscaleUpdateOne {
	uuo.mutation.SetUserAgent(s)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UpscaleUpdateOne) SetUpdatedAt(t time.Time) *UpscaleUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetUserID sets the "user_id" field.
func (uuo *UpscaleUpdateOne) SetUserID(u uuid.UUID) *UpscaleUpdateOne {
	uuo.mutation.SetUserID(u)
	return uuo
}

// SetUserTier sets the "user_tier" field.
func (uuo *UpscaleUpdateOne) SetUserTier(ut upscale.UserTier) *UpscaleUpdateOne {
	uuo.mutation.SetUserTier(ut)
	return uuo
}

// SetNillableUserTier sets the "user_tier" field if the given value is not nil.
func (uuo *UpscaleUpdateOne) SetNillableUserTier(ut *upscale.UserTier) *UpscaleUpdateOne {
	if ut != nil {
		uuo.SetUserTier(*ut)
	}
	return uuo
}

// SetUser sets the "user" edge to the User entity.
func (uuo *UpscaleUpdateOne) SetUser(u *User) *UpscaleUpdateOne {
	return uuo.SetUserID(u.ID)
}

// Mutation returns the UpscaleMutation object of the builder.
func (uuo *UpscaleUpdateOne) Mutation() *UpscaleMutation {
	return uuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uuo *UpscaleUpdateOne) ClearUser() *UpscaleUpdateOne {
	uuo.mutation.ClearUser()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UpscaleUpdateOne) Select(field string, fields ...string) *UpscaleUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Upscale entity.
func (uuo *UpscaleUpdateOne) Save(ctx context.Context) (*Upscale, error) {
	uuo.defaults()
	return withHooks[*Upscale, UpscaleMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UpscaleUpdateOne) SaveX(ctx context.Context) *Upscale {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UpscaleUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UpscaleUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UpscaleUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := upscale.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UpscaleUpdateOne) check() error {
	if v, ok := uuo.mutation.Status(); ok {
		if err := upscale.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Upscale.status": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.UserTier(); ok {
		if err := upscale.UserTierValidator(v); err != nil {
			return &ValidationError{Name: "user_tier", err: fmt.Errorf(`ent: validator failed for field "Upscale.user_tier": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.UserID(); uuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upscale.user"`)
	}
	return nil
}

func (uuo *UpscaleUpdateOne) sqlSave(ctx context.Context) (_node *Upscale, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscale.Table,
			Columns: upscale.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscale.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Upscale.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upscale.FieldID)
		for _, f := range fields {
			if !upscale.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != upscale.FieldID {
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
	if value, ok := uuo.mutation.Width(); ok {
		_spec.SetField(upscale.FieldWidth, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedWidth(); ok {
		_spec.AddField(upscale.FieldWidth, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Height(); ok {
		_spec.SetField(upscale.FieldHeight, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedHeight(); ok {
		_spec.AddField(upscale.FieldHeight, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Scale(); ok {
		_spec.SetField(upscale.FieldScale, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedScale(); ok {
		_spec.AddField(upscale.FieldScale, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(upscale.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.ServerURL(); ok {
		_spec.SetField(upscale.FieldServerURL, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DurationMsg(); ok {
		_spec.SetField(upscale.FieldDurationMsg, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedDurationMsg(); ok {
		_spec.AddField(upscale.FieldDurationMsg, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.GetType(); ok {
		_spec.SetField(upscale.FieldType, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Prompt(); ok {
		_spec.SetField(upscale.FieldPrompt, field.TypeString, value)
	}
	if value, ok := uuo.mutation.NegativePrompt(); ok {
		_spec.SetField(upscale.FieldNegativePrompt, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Seed(); ok {
		_spec.SetField(upscale.FieldSeed, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedSeed(); ok {
		_spec.AddField(upscale.FieldSeed, field.TypeInt, value)
	}
	if uuo.mutation.SeedCleared() {
		_spec.ClearField(upscale.FieldSeed, field.TypeInt)
	}
	if value, ok := uuo.mutation.NumInferenceSteps(); ok {
		_spec.SetField(upscale.FieldNumInferenceSteps, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedNumInferenceSteps(); ok {
		_spec.AddField(upscale.FieldNumInferenceSteps, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.GuidanceScale(); ok {
		_spec.SetField(upscale.FieldGuidanceScale, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.AddedGuidanceScale(); ok {
		_spec.AddField(upscale.FieldGuidanceScale, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.CountryCode(); ok {
		_spec.SetField(upscale.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DeviceType(); ok {
		_spec.SetField(upscale.FieldDeviceType, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DeviceOs(); ok {
		_spec.SetField(upscale.FieldDeviceOs, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DeviceBrowser(); ok {
		_spec.SetField(upscale.FieldDeviceBrowser, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UserAgent(); ok {
		_spec.SetField(upscale.FieldUserAgent, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(upscale.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UserTier(); ok {
		_spec.SetField(upscale.FieldUserTier, field.TypeEnum, value)
	}
	if uuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Upscale{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upscale.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
