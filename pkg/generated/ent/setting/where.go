/*
Copyright YEAR 1block.ai.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package setting

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldName, v))
}

// Default applies equality check predicate on the "default" field. It's identical to DefaultEQ.
func Default(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldDefault, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldValue, v))
}

// IsActive applies equality check predicate on the "isActive" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldIsActive, v))
}

// ReadOnly applies equality check predicate on the "readOnly" field. It's identical to ReadOnlyEQ.
func ReadOnly(v bool) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldReadOnly, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContainsFold(FieldName, v))
}

// DefaultEQ applies the EQ predicate on the "default" field.
func DefaultEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldDefault, v))
}

// DefaultNEQ applies the NEQ predicate on the "default" field.
func DefaultNEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldDefault, v))
}

// DefaultIn applies the In predicate on the "default" field.
func DefaultIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldDefault, vs...))
}

// DefaultNotIn applies the NotIn predicate on the "default" field.
func DefaultNotIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldDefault, vs...))
}

// DefaultGT applies the GT predicate on the "default" field.
func DefaultGT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldDefault, v))
}

// DefaultGTE applies the GTE predicate on the "default" field.
func DefaultGTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldDefault, v))
}

// DefaultLT applies the LT predicate on the "default" field.
func DefaultLT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldDefault, v))
}

// DefaultLTE applies the LTE predicate on the "default" field.
func DefaultLTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldDefault, v))
}

// DefaultContains applies the Contains predicate on the "default" field.
func DefaultContains(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContains(FieldDefault, v))
}

// DefaultHasPrefix applies the HasPrefix predicate on the "default" field.
func DefaultHasPrefix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasPrefix(FieldDefault, v))
}

// DefaultHasSuffix applies the HasSuffix predicate on the "default" field.
func DefaultHasSuffix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasSuffix(FieldDefault, v))
}

// DefaultEqualFold applies the EqualFold predicate on the "default" field.
func DefaultEqualFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEqualFold(FieldDefault, v))
}

// DefaultContainsFold applies the ContainsFold predicate on the "default" field.
func DefaultContainsFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContainsFold(FieldDefault, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasSuffix(FieldValue, v))
}

// ValueIsNil applies the IsNil predicate on the "value" field.
func ValueIsNil() predicate.Setting {
	return predicate.Setting(sql.FieldIsNull(FieldValue))
}

// ValueNotNil applies the NotNil predicate on the "value" field.
func ValueNotNil() predicate.Setting {
	return predicate.Setting(sql.FieldNotNull(FieldValue))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContainsFold(FieldValue, v))
}

// IsActiveEQ applies the EQ predicate on the "isActive" field.
func IsActiveEQ(v bool) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "isActive" field.
func IsActiveNEQ(v bool) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldIsActive, v))
}

// ReadOnlyEQ applies the EQ predicate on the "readOnly" field.
func ReadOnlyEQ(v bool) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldReadOnly, v))
}

// ReadOnlyNEQ applies the NEQ predicate on the "readOnly" field.
func ReadOnlyNEQ(v bool) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldReadOnly, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Setting) predicate.Setting {
	return predicate.Setting(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Setting) predicate.Setting {
	return predicate.Setting(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Setting) predicate.Setting {
	return predicate.Setting(sql.NotPredicates(p))
}
