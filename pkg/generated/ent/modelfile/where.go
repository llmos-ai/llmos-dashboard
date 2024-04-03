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

package modelfile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldUserID, v))
}

// TagName applies equality check predicate on the "tag_name" field. It's identical to TagNameEQ.
func TagName(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldTagName, v))
}

// Modelfile applies equality check predicate on the "modelfile" field. It's identical to ModelfileEQ.
func Modelfile(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldModelfile, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLTE(FieldUserID, v))
}

// TagNameEQ applies the EQ predicate on the "tag_name" field.
func TagNameEQ(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldTagName, v))
}

// TagNameNEQ applies the NEQ predicate on the "tag_name" field.
func TagNameNEQ(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldTagName, v))
}

// TagNameIn applies the In predicate on the "tag_name" field.
func TagNameIn(vs ...string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldTagName, vs...))
}

// TagNameNotIn applies the NotIn predicate on the "tag_name" field.
func TagNameNotIn(vs ...string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldTagName, vs...))
}

// TagNameGT applies the GT predicate on the "tag_name" field.
func TagNameGT(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGT(FieldTagName, v))
}

// TagNameGTE applies the GTE predicate on the "tag_name" field.
func TagNameGTE(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGTE(FieldTagName, v))
}

// TagNameLT applies the LT predicate on the "tag_name" field.
func TagNameLT(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLT(FieldTagName, v))
}

// TagNameLTE applies the LTE predicate on the "tag_name" field.
func TagNameLTE(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLTE(FieldTagName, v))
}

// TagNameContains applies the Contains predicate on the "tag_name" field.
func TagNameContains(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldContains(FieldTagName, v))
}

// TagNameHasPrefix applies the HasPrefix predicate on the "tag_name" field.
func TagNameHasPrefix(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldHasPrefix(FieldTagName, v))
}

// TagNameHasSuffix applies the HasSuffix predicate on the "tag_name" field.
func TagNameHasSuffix(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldHasSuffix(FieldTagName, v))
}

// TagNameEqualFold applies the EqualFold predicate on the "tag_name" field.
func TagNameEqualFold(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEqualFold(FieldTagName, v))
}

// TagNameContainsFold applies the ContainsFold predicate on the "tag_name" field.
func TagNameContainsFold(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldContainsFold(FieldTagName, v))
}

// ModelfileEQ applies the EQ predicate on the "modelfile" field.
func ModelfileEQ(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldModelfile, v))
}

// ModelfileNEQ applies the NEQ predicate on the "modelfile" field.
func ModelfileNEQ(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldModelfile, v))
}

// ModelfileIn applies the In predicate on the "modelfile" field.
func ModelfileIn(vs ...string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldModelfile, vs...))
}

// ModelfileNotIn applies the NotIn predicate on the "modelfile" field.
func ModelfileNotIn(vs ...string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldModelfile, vs...))
}

// ModelfileGT applies the GT predicate on the "modelfile" field.
func ModelfileGT(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGT(FieldModelfile, v))
}

// ModelfileGTE applies the GTE predicate on the "modelfile" field.
func ModelfileGTE(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGTE(FieldModelfile, v))
}

// ModelfileLT applies the LT predicate on the "modelfile" field.
func ModelfileLT(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLT(FieldModelfile, v))
}

// ModelfileLTE applies the LTE predicate on the "modelfile" field.
func ModelfileLTE(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLTE(FieldModelfile, v))
}

// ModelfileContains applies the Contains predicate on the "modelfile" field.
func ModelfileContains(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldContains(FieldModelfile, v))
}

// ModelfileHasPrefix applies the HasPrefix predicate on the "modelfile" field.
func ModelfileHasPrefix(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldHasPrefix(FieldModelfile, v))
}

// ModelfileHasSuffix applies the HasSuffix predicate on the "modelfile" field.
func ModelfileHasSuffix(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldHasSuffix(FieldModelfile, v))
}

// ModelfileEqualFold applies the EqualFold predicate on the "modelfile" field.
func ModelfileEqualFold(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEqualFold(FieldModelfile, v))
}

// ModelfileContainsFold applies the ContainsFold predicate on the "modelfile" field.
func ModelfileContainsFold(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldContainsFold(FieldModelfile, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Modelfile) predicate.Modelfile {
	return predicate.Modelfile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Modelfile) predicate.Modelfile {
	return predicate.Modelfile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Modelfile) predicate.Modelfile {
	return predicate.Modelfile(sql.NotPredicates(p))
}
