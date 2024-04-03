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

package chat

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldTitle, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldUserID, v))
}

// Chat applies equality check predicate on the "chat" field. It's identical to ChatEQ.
func Chat(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldChat, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldCreatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContainsFold(FieldTitle, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldUserID, vs...))
}

// ChatEQ applies the EQ predicate on the "chat" field.
func ChatEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldChat, v))
}

// ChatNEQ applies the NEQ predicate on the "chat" field.
func ChatNEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldChat, v))
}

// ChatIn applies the In predicate on the "chat" field.
func ChatIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldChat, vs...))
}

// ChatNotIn applies the NotIn predicate on the "chat" field.
func ChatNotIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldChat, vs...))
}

// ChatGT applies the GT predicate on the "chat" field.
func ChatGT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldChat, v))
}

// ChatGTE applies the GTE predicate on the "chat" field.
func ChatGTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldChat, v))
}

// ChatLT applies the LT predicate on the "chat" field.
func ChatLT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldChat, v))
}

// ChatLTE applies the LTE predicate on the "chat" field.
func ChatLTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldChat, v))
}

// ChatContains applies the Contains predicate on the "chat" field.
func ChatContains(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContains(FieldChat, v))
}

// ChatHasPrefix applies the HasPrefix predicate on the "chat" field.
func ChatHasPrefix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasPrefix(FieldChat, v))
}

// ChatHasSuffix applies the HasSuffix predicate on the "chat" field.
func ChatHasSuffix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasSuffix(FieldChat, v))
}

// ChatEqualFold applies the EqualFold predicate on the "chat" field.
func ChatEqualFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEqualFold(FieldChat, v))
}

// ChatContainsFold applies the ContainsFold predicate on the "chat" field.
func ChatContainsFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContainsFold(FieldChat, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Chat) predicate.Chat {
	return predicate.Chat(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Chat) predicate.Chat {
	return predicate.Chat(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Chat) predicate.Chat {
	return predicate.Chat(sql.NotPredicates(p))
}
