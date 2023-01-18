// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/generationrealtime"
)

// GenerationRealtime is the model entity for the GenerationRealtime schema.
type GenerationRealtime struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CountryCode holds the value of the "country_code" field.
	CountryCode *string `json:"country_code,omitempty"`
	// DurationMs holds the value of the "duration_ms" field.
	DurationMs *int `json:"duration_ms,omitempty"`
	// Status holds the value of the "status" field.
	Status *generationrealtime.Status `json:"status,omitempty"`
	// UsesDefaultServer holds the value of the "uses_default_server" field.
	UsesDefaultServer bool `json:"uses_default_server,omitempty"`
	// Width holds the value of the "width" field.
	Width *int `json:"width,omitempty"`
	// Height holds the value of the "height" field.
	Height *int `json:"height,omitempty"`
	// NumInterferenceSteps holds the value of the "num_interference_steps" field.
	NumInterferenceSteps *int `json:"num_interference_steps,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserTier holds the value of the "user_tier" field.
	UserTier generationrealtime.UserTier `json:"user_tier,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GenerationRealtime) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case generationrealtime.FieldUsesDefaultServer:
			values[i] = new(sql.NullBool)
		case generationrealtime.FieldDurationMs, generationrealtime.FieldWidth, generationrealtime.FieldHeight, generationrealtime.FieldNumInterferenceSteps:
			values[i] = new(sql.NullInt64)
		case generationrealtime.FieldCountryCode, generationrealtime.FieldStatus, generationrealtime.FieldUserTier:
			values[i] = new(sql.NullString)
		case generationrealtime.FieldCreatedAt, generationrealtime.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case generationrealtime.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GenerationRealtime", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GenerationRealtime fields.
func (gr *GenerationRealtime) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case generationrealtime.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gr.ID = *value
			}
		case generationrealtime.FieldCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_code", values[i])
			} else if value.Valid {
				gr.CountryCode = new(string)
				*gr.CountryCode = value.String
			}
		case generationrealtime.FieldDurationMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_ms", values[i])
			} else if value.Valid {
				gr.DurationMs = new(int)
				*gr.DurationMs = int(value.Int64)
			}
		case generationrealtime.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				gr.Status = new(generationrealtime.Status)
				*gr.Status = generationrealtime.Status(value.String)
			}
		case generationrealtime.FieldUsesDefaultServer:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field uses_default_server", values[i])
			} else if value.Valid {
				gr.UsesDefaultServer = value.Bool
			}
		case generationrealtime.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				gr.Width = new(int)
				*gr.Width = int(value.Int64)
			}
		case generationrealtime.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				gr.Height = new(int)
				*gr.Height = int(value.Int64)
			}
		case generationrealtime.FieldNumInterferenceSteps:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num_interference_steps", values[i])
			} else if value.Valid {
				gr.NumInterferenceSteps = new(int)
				*gr.NumInterferenceSteps = int(value.Int64)
			}
		case generationrealtime.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case generationrealtime.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gr.UpdatedAt = value.Time
			}
		case generationrealtime.FieldUserTier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_tier", values[i])
			} else if value.Valid {
				gr.UserTier = generationrealtime.UserTier(value.String)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GenerationRealtime.
// Note that you need to call GenerationRealtime.Unwrap() before calling this method if this GenerationRealtime
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *GenerationRealtime) Update() *GenerationRealtimeUpdateOne {
	return (&GenerationRealtimeClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the GenerationRealtime entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *GenerationRealtime) Unwrap() *GenerationRealtime {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: GenerationRealtime is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *GenerationRealtime) String() string {
	var builder strings.Builder
	builder.WriteString("GenerationRealtime(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	if v := gr.CountryCode; v != nil {
		builder.WriteString("country_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := gr.DurationMs; v != nil {
		builder.WriteString("duration_ms=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := gr.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("uses_default_server=")
	builder.WriteString(fmt.Sprintf("%v", gr.UsesDefaultServer))
	builder.WriteString(", ")
	if v := gr.Width; v != nil {
		builder.WriteString("width=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := gr.Height; v != nil {
		builder.WriteString("height=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := gr.NumInterferenceSteps; v != nil {
		builder.WriteString("num_interference_steps=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_tier=")
	builder.WriteString(fmt.Sprintf("%v", gr.UserTier))
	builder.WriteByte(')')
	return builder.String()
}

// GenerationRealtimes is a parsable slice of GenerationRealtime.
type GenerationRealtimes []*GenerationRealtime

func (gr GenerationRealtimes) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}