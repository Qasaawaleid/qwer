// Code generated by ent, DO NOT EDIT.

package negativeprompt

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the negativeprompt type in the database.
	Label = "negative_prompt"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeGeneration holds the string denoting the generation edge name in mutations.
	EdgeGeneration = "generation"
	// EdgeGenerationG holds the string denoting the generation_g edge name in mutations.
	EdgeGenerationG = "generation_g"
	// Table holds the table name of the negativeprompt in the database.
	Table = "negative_prompt"
	// GenerationTable is the table that holds the generation relation/edge.
	GenerationTable = "generation"
	// GenerationInverseTable is the table name for the Generation entity.
	// It exists in this package in order to avoid circular dependency with the "generation" package.
	GenerationInverseTable = "generation"
	// GenerationColumn is the table column denoting the generation relation/edge.
	GenerationColumn = "negative_prompt_id"
	// GenerationGTable is the table that holds the generation_g relation/edge.
	GenerationGTable = "generation_g"
	// GenerationGInverseTable is the table name for the GenerationG entity.
	// It exists in this package in order to avoid circular dependency with the "generationg" package.
	GenerationGInverseTable = "generation_g"
	// GenerationGColumn is the table column denoting the generation_g relation/edge.
	GenerationGColumn = "negative_prompt_id"
)

// Columns holds all SQL columns for negativeprompt fields.
var Columns = []string{
	FieldID,
	FieldText,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
