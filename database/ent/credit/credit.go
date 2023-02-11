// Code generated by ent, DO NOT EDIT.

package credit

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the credit type in the database.
	Label = "credit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRemainingAmount holds the string denoting the remaining_amount field in the database.
	FieldRemainingAmount = "remaining_amount"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCreditTypeID holds the string denoting the credit_type_id field in the database.
	FieldCreditTypeID = "credit_type_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeCreditTypes holds the string denoting the credit_types edge name in mutations.
	EdgeCreditTypes = "credit_types"
	// Table holds the table name of the credit in the database.
	Table = "credits"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "credits"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
	// CreditTypesTable is the table that holds the credit_types relation/edge.
	CreditTypesTable = "credits"
	// CreditTypesInverseTable is the table name for the CreditType entity.
	// It exists in this package in order to avoid circular dependency with the "credittype" package.
	CreditTypesInverseTable = "credit_types"
	// CreditTypesColumn is the table column denoting the credit_types relation/edge.
	CreditTypesColumn = "credit_type_id"
)

// Columns holds all SQL columns for credit fields.
var Columns = []string{
	FieldID,
	FieldRemainingAmount,
	FieldExpiresAt,
	FieldUserID,
	FieldCreditTypeID,
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