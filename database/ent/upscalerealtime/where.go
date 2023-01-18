// Code generated by ent, DO NOT EDIT.

package upscalerealtime

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLTE(FieldID, id))
}

// CountryCode applies equality check predicate on the "country_code" field. It's identical to CountryCodeEQ.
func CountryCode(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldCountryCode, v))
}

// UsesDefaultServer applies equality check predicate on the "uses_default_server" field. It's identical to UsesDefaultServerEQ.
func UsesDefaultServer(v bool) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldUsesDefaultServer, v))
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldWidth, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldHeight, v))
}

// Scale applies equality check predicate on the "scale" field. It's identical to ScaleEQ.
func Scale(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldScale, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldStatus, vs...))
}

// CountryCodeEQ applies the EQ predicate on the "country_code" field.
func CountryCodeEQ(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldCountryCode, v))
}

// CountryCodeNEQ applies the NEQ predicate on the "country_code" field.
func CountryCodeNEQ(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldCountryCode, v))
}

// CountryCodeIn applies the In predicate on the "country_code" field.
func CountryCodeIn(vs ...string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldCountryCode, vs...))
}

// CountryCodeNotIn applies the NotIn predicate on the "country_code" field.
func CountryCodeNotIn(vs ...string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldCountryCode, vs...))
}

// CountryCodeGT applies the GT predicate on the "country_code" field.
func CountryCodeGT(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGT(FieldCountryCode, v))
}

// CountryCodeGTE applies the GTE predicate on the "country_code" field.
func CountryCodeGTE(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGTE(FieldCountryCode, v))
}

// CountryCodeLT applies the LT predicate on the "country_code" field.
func CountryCodeLT(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLT(FieldCountryCode, v))
}

// CountryCodeLTE applies the LTE predicate on the "country_code" field.
func CountryCodeLTE(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLTE(FieldCountryCode, v))
}

// CountryCodeContains applies the Contains predicate on the "country_code" field.
func CountryCodeContains(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldContains(FieldCountryCode, v))
}

// CountryCodeHasPrefix applies the HasPrefix predicate on the "country_code" field.
func CountryCodeHasPrefix(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldHasPrefix(FieldCountryCode, v))
}

// CountryCodeHasSuffix applies the HasSuffix predicate on the "country_code" field.
func CountryCodeHasSuffix(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldHasSuffix(FieldCountryCode, v))
}

// CountryCodeEqualFold applies the EqualFold predicate on the "country_code" field.
func CountryCodeEqualFold(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEqualFold(FieldCountryCode, v))
}

// CountryCodeContainsFold applies the ContainsFold predicate on the "country_code" field.
func CountryCodeContainsFold(v string) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldContainsFold(FieldCountryCode, v))
}

// UsesDefaultServerEQ applies the EQ predicate on the "uses_default_server" field.
func UsesDefaultServerEQ(v bool) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldUsesDefaultServer, v))
}

// UsesDefaultServerNEQ applies the NEQ predicate on the "uses_default_server" field.
func UsesDefaultServerNEQ(v bool) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldUsesDefaultServer, v))
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldWidth, v))
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldWidth, v))
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldWidth, vs...))
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldWidth, vs...))
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGT(FieldWidth, v))
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGTE(FieldWidth, v))
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLT(FieldWidth, v))
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLTE(FieldWidth, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLTE(FieldHeight, v))
}

// ScaleEQ applies the EQ predicate on the "scale" field.
func ScaleEQ(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldScale, v))
}

// ScaleNEQ applies the NEQ predicate on the "scale" field.
func ScaleNEQ(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldScale, v))
}

// ScaleIn applies the In predicate on the "scale" field.
func ScaleIn(vs ...int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldScale, vs...))
}

// ScaleNotIn applies the NotIn predicate on the "scale" field.
func ScaleNotIn(vs ...int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldScale, vs...))
}

// ScaleGT applies the GT predicate on the "scale" field.
func ScaleGT(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGT(FieldScale, v))
}

// ScaleGTE applies the GTE predicate on the "scale" field.
func ScaleGTE(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGTE(FieldScale, v))
}

// ScaleLT applies the LT predicate on the "scale" field.
func ScaleLT(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLT(FieldScale, v))
}

// ScaleLTE applies the LTE predicate on the "scale" field.
func ScaleLTE(v int) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLTE(FieldScale, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserTierEQ applies the EQ predicate on the "user_tier" field.
func UserTierEQ(v UserTier) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldEQ(FieldUserTier, v))
}

// UserTierNEQ applies the NEQ predicate on the "user_tier" field.
func UserTierNEQ(v UserTier) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNEQ(FieldUserTier, v))
}

// UserTierIn applies the In predicate on the "user_tier" field.
func UserTierIn(vs ...UserTier) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldIn(FieldUserTier, vs...))
}

// UserTierNotIn applies the NotIn predicate on the "user_tier" field.
func UserTierNotIn(vs ...UserTier) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(sql.FieldNotIn(FieldUserTier, vs...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UpscaleRealtime) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UpscaleRealtime) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UpscaleRealtime) predicate.UpscaleRealtime {
	return predicate.UpscaleRealtime(func(s *sql.Selector) {
		p(s.Not())
	})
}