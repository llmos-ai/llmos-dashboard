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
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLTE(FieldID, id))
}

// TagName applies equality check predicate on the "tagName" field. It's identical to TagNameEQ.
func TagName(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldTagName, v))
}

// Modelfile applies equality check predicate on the "modelfile" field. It's identical to ModelfileEQ.
func Modelfile(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldModelfile, v))
}

// UserId applies equality check predicate on the "userId" field. It's identical to UserIdEQ.
func UserId(v uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldUserId, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldCreatedAt, v))
}

// TagNameEQ applies the EQ predicate on the "tagName" field.
func TagNameEQ(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldTagName, v))
}

// TagNameNEQ applies the NEQ predicate on the "tagName" field.
func TagNameNEQ(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldTagName, v))
}

// TagNameIn applies the In predicate on the "tagName" field.
func TagNameIn(vs ...string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldTagName, vs...))
}

// TagNameNotIn applies the NotIn predicate on the "tagName" field.
func TagNameNotIn(vs ...string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldTagName, vs...))
}

// TagNameGT applies the GT predicate on the "tagName" field.
func TagNameGT(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGT(FieldTagName, v))
}

// TagNameGTE applies the GTE predicate on the "tagName" field.
func TagNameGTE(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGTE(FieldTagName, v))
}

// TagNameLT applies the LT predicate on the "tagName" field.
func TagNameLT(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLT(FieldTagName, v))
}

// TagNameLTE applies the LTE predicate on the "tagName" field.
func TagNameLTE(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLTE(FieldTagName, v))
}

// TagNameContains applies the Contains predicate on the "tagName" field.
func TagNameContains(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldContains(FieldTagName, v))
}

// TagNameHasPrefix applies the HasPrefix predicate on the "tagName" field.
func TagNameHasPrefix(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldHasPrefix(FieldTagName, v))
}

// TagNameHasSuffix applies the HasSuffix predicate on the "tagName" field.
func TagNameHasSuffix(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldHasSuffix(FieldTagName, v))
}

// TagNameEqualFold applies the EqualFold predicate on the "tagName" field.
func TagNameEqualFold(v string) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEqualFold(FieldTagName, v))
}

// TagNameContainsFold applies the ContainsFold predicate on the "tagName" field.
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

// UserIdEQ applies the EQ predicate on the "userId" field.
func UserIdEQ(v uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldUserId, v))
}

// UserIdNEQ applies the NEQ predicate on the "userId" field.
func UserIdNEQ(v uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldUserId, v))
}

// UserIdIn applies the In predicate on the "userId" field.
func UserIdIn(vs ...uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldUserId, vs...))
}

// UserIdNotIn applies the NotIn predicate on the "userId" field.
func UserIdNotIn(vs ...uuid.UUID) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldUserId, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Modelfile {
	return predicate.Modelfile(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Modelfile {
	return predicate.Modelfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Modelfile {
	return predicate.Modelfile(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
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
