// Code generated by ent, DO NOT EDIT.

package scheduler

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the scheduler type in the database.
	Label = "scheduler"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsFree holds the string denoting the is_free field in the database.
	FieldIsFree = "is_free"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeGenerations holds the string denoting the generations edge name in mutations.
	EdgeGenerations = "generations"
	// Table holds the table name of the scheduler in the database.
	Table = "schedulers"
	// GenerationsTable is the table that holds the generations relation/edge.
	GenerationsTable = "generations"
	// GenerationsInverseTable is the table name for the Generation entity.
	// It exists in this package in order to avoid circular dependency with the "generation" package.
	GenerationsInverseTable = "generations"
	// GenerationsColumn is the table column denoting the generations relation/edge.
	GenerationsColumn = "scheduler_id"
)

// Columns holds all SQL columns for scheduler fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldIsFree,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultIsFree holds the default value on creation for the "is_free" field.
	DefaultIsFree bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
