// Code generated by ent, DO NOT EDIT.

package generationrealtime

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the generationrealtime type in the database.
	Label = "generation_realtime"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldDurationMs holds the string denoting the duration_ms field in the database.
	FieldDurationMs = "duration_ms"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUsesDefaultServer holds the string denoting the uses_default_server field in the database.
	FieldUsesDefaultServer = "uses_default_server"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldNumInterferenceSteps holds the string denoting the num_interference_steps field in the database.
	FieldNumInterferenceSteps = "num_interference_steps"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserTier holds the string denoting the user_tier field in the database.
	FieldUserTier = "user_tier"
	// Table holds the table name of the generationrealtime in the database.
	Table = "generation_realtime"
)

// Columns holds all SQL columns for generationrealtime fields.
var Columns = []string{
	FieldID,
	FieldCountryCode,
	FieldDurationMs,
	FieldStatus,
	FieldUsesDefaultServer,
	FieldWidth,
	FieldHeight,
	FieldNumInterferenceSteps,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserTier,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusStarted   Status = "started"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
	StatusRejected  Status = "rejected"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusStarted, StatusSucceeded, StatusFailed, StatusRejected:
		return nil
	default:
		return fmt.Errorf("generationrealtime: invalid enum value for status field: %q", s)
	}
}

// UserTier defines the type for the "user_tier" enum field.
type UserTier string

// UserTierFREE is the default value of the UserTier enum.
const DefaultUserTier = UserTierFREE

// UserTier values.
const (
	UserTierFREE UserTier = "FREE"
	UserTierPRO  UserTier = "PRO"
)

func (ut UserTier) String() string {
	return string(ut)
}

// UserTierValidator is a validator for the "user_tier" field enum values. It is called by the builders before save.
func UserTierValidator(ut UserTier) error {
	switch ut {
	case UserTierFREE, UserTierPRO:
		return nil
	default:
		return fmt.Errorf("generationrealtime: invalid enum value for user_tier field: %q", ut)
	}
}