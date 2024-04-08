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

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/chat"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/predicate"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/setting"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
	v1 "github.com/llmos-ai/llmos-dashboard/pkg/types/v1"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeChat      = "Chat"
	TypeModelfile = "Modelfile"
	TypeSetting   = "Setting"
	TypeUser      = "User"
)

// ChatMutation represents an operation that mutates the Chat nodes in the graph.
type ChatMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	title          *string
	models         *[]string
	appendmodels   []string
	tags           *[]string
	appendtags     []string
	history        *v1.Histroy
	messages       *[]v1.Message
	appendmessages []v1.Message
	createdAt      *time.Time
	clearedFields  map[string]struct{}
	owner          *uuid.UUID
	clearedowner   bool
	done           bool
	oldValue       func(context.Context) (*Chat, error)
	predicates     []predicate.Chat
}

var _ ent.Mutation = (*ChatMutation)(nil)

// chatOption allows management of the mutation configuration using functional options.
type chatOption func(*ChatMutation)

// newChatMutation creates new mutation for the Chat entity.
func newChatMutation(c config, op Op, opts ...chatOption) *ChatMutation {
	m := &ChatMutation{
		config:        c,
		op:            op,
		typ:           TypeChat,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withChatID sets the ID field of the mutation.
func withChatID(id uuid.UUID) chatOption {
	return func(m *ChatMutation) {
		var (
			err   error
			once  sync.Once
			value *Chat
		)
		m.oldValue = func(ctx context.Context) (*Chat, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Chat.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withChat sets the old Chat of the mutation.
func withChat(node *Chat) chatOption {
	return func(m *ChatMutation) {
		m.oldValue = func(context.Context) (*Chat, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ChatMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ChatMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Chat entities.
func (m *ChatMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ChatMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ChatMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Chat.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *ChatMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *ChatMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *ChatMutation) ResetTitle() {
	m.title = nil
}

// SetUserId sets the "userId" field.
func (m *ChatMutation) SetUserId(u uuid.UUID) {
	m.owner = &u
}

// UserId returns the value of the "userId" field in the mutation.
func (m *ChatMutation) UserId() (r uuid.UUID, exists bool) {
	v := m.owner
	if v == nil {
		return
	}
	return *v, true
}

// OldUserId returns the old "userId" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldUserId(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserId: %w", err)
	}
	return oldValue.UserId, nil
}

// ResetUserId resets all changes to the "userId" field.
func (m *ChatMutation) ResetUserId() {
	m.owner = nil
}

// SetModels sets the "models" field.
func (m *ChatMutation) SetModels(s []string) {
	m.models = &s
	m.appendmodels = nil
}

// Models returns the value of the "models" field in the mutation.
func (m *ChatMutation) Models() (r []string, exists bool) {
	v := m.models
	if v == nil {
		return
	}
	return *v, true
}

// OldModels returns the old "models" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldModels(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldModels is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldModels requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldModels: %w", err)
	}
	return oldValue.Models, nil
}

// AppendModels adds s to the "models" field.
func (m *ChatMutation) AppendModels(s []string) {
	m.appendmodels = append(m.appendmodels, s...)
}

// AppendedModels returns the list of values that were appended to the "models" field in this mutation.
func (m *ChatMutation) AppendedModels() ([]string, bool) {
	if len(m.appendmodels) == 0 {
		return nil, false
	}
	return m.appendmodels, true
}

// ResetModels resets all changes to the "models" field.
func (m *ChatMutation) ResetModels() {
	m.models = nil
	m.appendmodels = nil
}

// SetTags sets the "tags" field.
func (m *ChatMutation) SetTags(s []string) {
	m.tags = &s
	m.appendtags = nil
}

// Tags returns the value of the "tags" field in the mutation.
func (m *ChatMutation) Tags() (r []string, exists bool) {
	v := m.tags
	if v == nil {
		return
	}
	return *v, true
}

// OldTags returns the old "tags" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldTags(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTags is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTags requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTags: %w", err)
	}
	return oldValue.Tags, nil
}

// AppendTags adds s to the "tags" field.
func (m *ChatMutation) AppendTags(s []string) {
	m.appendtags = append(m.appendtags, s...)
}

// AppendedTags returns the list of values that were appended to the "tags" field in this mutation.
func (m *ChatMutation) AppendedTags() ([]string, bool) {
	if len(m.appendtags) == 0 {
		return nil, false
	}
	return m.appendtags, true
}

// ResetTags resets all changes to the "tags" field.
func (m *ChatMutation) ResetTags() {
	m.tags = nil
	m.appendtags = nil
}

// SetHistory sets the "history" field.
func (m *ChatMutation) SetHistory(v v1.Histroy) {
	m.history = &v
}

// History returns the value of the "history" field in the mutation.
func (m *ChatMutation) History() (r v1.Histroy, exists bool) {
	v := m.history
	if v == nil {
		return
	}
	return *v, true
}

// OldHistory returns the old "history" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldHistory(ctx context.Context) (v v1.Histroy, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHistory is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHistory requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHistory: %w", err)
	}
	return oldValue.History, nil
}

// ResetHistory resets all changes to the "history" field.
func (m *ChatMutation) ResetHistory() {
	m.history = nil
}

// SetMessages sets the "messages" field.
func (m *ChatMutation) SetMessages(v []v1.Message) {
	m.messages = &v
	m.appendmessages = nil
}

// Messages returns the value of the "messages" field in the mutation.
func (m *ChatMutation) Messages() (r []v1.Message, exists bool) {
	v := m.messages
	if v == nil {
		return
	}
	return *v, true
}

// OldMessages returns the old "messages" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldMessages(ctx context.Context) (v []v1.Message, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessages is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessages requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessages: %w", err)
	}
	return oldValue.Messages, nil
}

// AppendMessages adds v to the "messages" field.
func (m *ChatMutation) AppendMessages(v []v1.Message) {
	m.appendmessages = append(m.appendmessages, v...)
}

// AppendedMessages returns the list of values that were appended to the "messages" field in this mutation.
func (m *ChatMutation) AppendedMessages() ([]v1.Message, bool) {
	if len(m.appendmessages) == 0 {
		return nil, false
	}
	return m.appendmessages, true
}

// ResetMessages resets all changes to the "messages" field.
func (m *ChatMutation) ResetMessages() {
	m.messages = nil
	m.appendmessages = nil
}

// SetCreatedAt sets the "createdAt" field.
func (m *ChatMutation) SetCreatedAt(t time.Time) {
	m.createdAt = &t
}

// CreatedAt returns the value of the "createdAt" field in the mutation.
func (m *ChatMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.createdAt
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "createdAt" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "createdAt" field.
func (m *ChatMutation) ResetCreatedAt() {
	m.createdAt = nil
}

// SetOwnerID sets the "owner" edge to the User entity by id.
func (m *ChatMutation) SetOwnerID(id uuid.UUID) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *ChatMutation) ClearOwner() {
	m.clearedowner = true
	m.clearedFields[chat.FieldUserId] = struct{}{}
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *ChatMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *ChatMutation) OwnerID() (id uuid.UUID, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *ChatMutation) OwnerIDs() (ids []uuid.UUID) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *ChatMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the ChatMutation builder.
func (m *ChatMutation) Where(ps ...predicate.Chat) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ChatMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ChatMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Chat, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ChatMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ChatMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Chat).
func (m *ChatMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ChatMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.title != nil {
		fields = append(fields, chat.FieldTitle)
	}
	if m.owner != nil {
		fields = append(fields, chat.FieldUserId)
	}
	if m.models != nil {
		fields = append(fields, chat.FieldModels)
	}
	if m.tags != nil {
		fields = append(fields, chat.FieldTags)
	}
	if m.history != nil {
		fields = append(fields, chat.FieldHistory)
	}
	if m.messages != nil {
		fields = append(fields, chat.FieldMessages)
	}
	if m.createdAt != nil {
		fields = append(fields, chat.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ChatMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case chat.FieldTitle:
		return m.Title()
	case chat.FieldUserId:
		return m.UserId()
	case chat.FieldModels:
		return m.Models()
	case chat.FieldTags:
		return m.Tags()
	case chat.FieldHistory:
		return m.History()
	case chat.FieldMessages:
		return m.Messages()
	case chat.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ChatMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case chat.FieldTitle:
		return m.OldTitle(ctx)
	case chat.FieldUserId:
		return m.OldUserId(ctx)
	case chat.FieldModels:
		return m.OldModels(ctx)
	case chat.FieldTags:
		return m.OldTags(ctx)
	case chat.FieldHistory:
		return m.OldHistory(ctx)
	case chat.FieldMessages:
		return m.OldMessages(ctx)
	case chat.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Chat field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMutation) SetField(name string, value ent.Value) error {
	switch name {
	case chat.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case chat.FieldUserId:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserId(v)
		return nil
	case chat.FieldModels:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetModels(v)
		return nil
	case chat.FieldTags:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTags(v)
		return nil
	case chat.FieldHistory:
		v, ok := value.(v1.Histroy)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHistory(v)
		return nil
	case chat.FieldMessages:
		v, ok := value.([]v1.Message)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessages(v)
		return nil
	case chat.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Chat field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ChatMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ChatMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Chat numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ChatMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ChatMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ChatMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Chat nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ChatMutation) ResetField(name string) error {
	switch name {
	case chat.FieldTitle:
		m.ResetTitle()
		return nil
	case chat.FieldUserId:
		m.ResetUserId()
		return nil
	case chat.FieldModels:
		m.ResetModels()
		return nil
	case chat.FieldTags:
		m.ResetTags()
		return nil
	case chat.FieldHistory:
		m.ResetHistory()
		return nil
	case chat.FieldMessages:
		m.ResetMessages()
		return nil
	case chat.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Chat field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ChatMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, chat.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ChatMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case chat.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ChatMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ChatMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ChatMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, chat.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ChatMutation) EdgeCleared(name string) bool {
	switch name {
	case chat.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ChatMutation) ClearEdge(name string) error {
	switch name {
	case chat.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Chat unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ChatMutation) ResetEdge(name string) error {
	switch name {
	case chat.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Chat edge %s", name)
}

// ModelfileMutation represents an operation that mutates the Modelfile nodes in the graph.
type ModelfileMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	tagName       *string
	modelfile     *string
	createdAt     *time.Time
	clearedFields map[string]struct{}
	owner         *uuid.UUID
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Modelfile, error)
	predicates    []predicate.Modelfile
}

var _ ent.Mutation = (*ModelfileMutation)(nil)

// modelfileOption allows management of the mutation configuration using functional options.
type modelfileOption func(*ModelfileMutation)

// newModelfileMutation creates new mutation for the Modelfile entity.
func newModelfileMutation(c config, op Op, opts ...modelfileOption) *ModelfileMutation {
	m := &ModelfileMutation{
		config:        c,
		op:            op,
		typ:           TypeModelfile,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withModelfileID sets the ID field of the mutation.
func withModelfileID(id uuid.UUID) modelfileOption {
	return func(m *ModelfileMutation) {
		var (
			err   error
			once  sync.Once
			value *Modelfile
		)
		m.oldValue = func(ctx context.Context) (*Modelfile, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Modelfile.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withModelfile sets the old Modelfile of the mutation.
func withModelfile(node *Modelfile) modelfileOption {
	return func(m *ModelfileMutation) {
		m.oldValue = func(context.Context) (*Modelfile, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ModelfileMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ModelfileMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Modelfile entities.
func (m *ModelfileMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ModelfileMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ModelfileMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Modelfile.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTagName sets the "tagName" field.
func (m *ModelfileMutation) SetTagName(s string) {
	m.tagName = &s
}

// TagName returns the value of the "tagName" field in the mutation.
func (m *ModelfileMutation) TagName() (r string, exists bool) {
	v := m.tagName
	if v == nil {
		return
	}
	return *v, true
}

// OldTagName returns the old "tagName" field's value of the Modelfile entity.
// If the Modelfile object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ModelfileMutation) OldTagName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTagName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTagName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTagName: %w", err)
	}
	return oldValue.TagName, nil
}

// ResetTagName resets all changes to the "tagName" field.
func (m *ModelfileMutation) ResetTagName() {
	m.tagName = nil
}

// SetModelfile sets the "modelfile" field.
func (m *ModelfileMutation) SetModelfile(s string) {
	m.modelfile = &s
}

// Modelfile returns the value of the "modelfile" field in the mutation.
func (m *ModelfileMutation) Modelfile() (r string, exists bool) {
	v := m.modelfile
	if v == nil {
		return
	}
	return *v, true
}

// OldModelfile returns the old "modelfile" field's value of the Modelfile entity.
// If the Modelfile object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ModelfileMutation) OldModelfile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldModelfile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldModelfile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldModelfile: %w", err)
	}
	return oldValue.Modelfile, nil
}

// ResetModelfile resets all changes to the "modelfile" field.
func (m *ModelfileMutation) ResetModelfile() {
	m.modelfile = nil
}

// SetUserId sets the "userId" field.
func (m *ModelfileMutation) SetUserId(u uuid.UUID) {
	m.owner = &u
}

// UserId returns the value of the "userId" field in the mutation.
func (m *ModelfileMutation) UserId() (r uuid.UUID, exists bool) {
	v := m.owner
	if v == nil {
		return
	}
	return *v, true
}

// OldUserId returns the old "userId" field's value of the Modelfile entity.
// If the Modelfile object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ModelfileMutation) OldUserId(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserId: %w", err)
	}
	return oldValue.UserId, nil
}

// ResetUserId resets all changes to the "userId" field.
func (m *ModelfileMutation) ResetUserId() {
	m.owner = nil
}

// SetCreatedAt sets the "createdAt" field.
func (m *ModelfileMutation) SetCreatedAt(t time.Time) {
	m.createdAt = &t
}

// CreatedAt returns the value of the "createdAt" field in the mutation.
func (m *ModelfileMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.createdAt
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "createdAt" field's value of the Modelfile entity.
// If the Modelfile object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ModelfileMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "createdAt" field.
func (m *ModelfileMutation) ResetCreatedAt() {
	m.createdAt = nil
}

// SetOwnerID sets the "owner" edge to the User entity by id.
func (m *ModelfileMutation) SetOwnerID(id uuid.UUID) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *ModelfileMutation) ClearOwner() {
	m.clearedowner = true
	m.clearedFields[modelfile.FieldUserId] = struct{}{}
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *ModelfileMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *ModelfileMutation) OwnerID() (id uuid.UUID, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *ModelfileMutation) OwnerIDs() (ids []uuid.UUID) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *ModelfileMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the ModelfileMutation builder.
func (m *ModelfileMutation) Where(ps ...predicate.Modelfile) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ModelfileMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ModelfileMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Modelfile, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ModelfileMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ModelfileMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Modelfile).
func (m *ModelfileMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ModelfileMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.tagName != nil {
		fields = append(fields, modelfile.FieldTagName)
	}
	if m.modelfile != nil {
		fields = append(fields, modelfile.FieldModelfile)
	}
	if m.owner != nil {
		fields = append(fields, modelfile.FieldUserId)
	}
	if m.createdAt != nil {
		fields = append(fields, modelfile.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ModelfileMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case modelfile.FieldTagName:
		return m.TagName()
	case modelfile.FieldModelfile:
		return m.Modelfile()
	case modelfile.FieldUserId:
		return m.UserId()
	case modelfile.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ModelfileMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case modelfile.FieldTagName:
		return m.OldTagName(ctx)
	case modelfile.FieldModelfile:
		return m.OldModelfile(ctx)
	case modelfile.FieldUserId:
		return m.OldUserId(ctx)
	case modelfile.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Modelfile field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ModelfileMutation) SetField(name string, value ent.Value) error {
	switch name {
	case modelfile.FieldTagName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTagName(v)
		return nil
	case modelfile.FieldModelfile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetModelfile(v)
		return nil
	case modelfile.FieldUserId:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserId(v)
		return nil
	case modelfile.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Modelfile field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ModelfileMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ModelfileMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ModelfileMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Modelfile numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ModelfileMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ModelfileMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ModelfileMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Modelfile nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ModelfileMutation) ResetField(name string) error {
	switch name {
	case modelfile.FieldTagName:
		m.ResetTagName()
		return nil
	case modelfile.FieldModelfile:
		m.ResetModelfile()
		return nil
	case modelfile.FieldUserId:
		m.ResetUserId()
		return nil
	case modelfile.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Modelfile field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ModelfileMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, modelfile.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ModelfileMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case modelfile.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ModelfileMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ModelfileMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ModelfileMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, modelfile.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ModelfileMutation) EdgeCleared(name string) bool {
	switch name {
	case modelfile.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ModelfileMutation) ClearEdge(name string) error {
	switch name {
	case modelfile.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Modelfile unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ModelfileMutation) ResetEdge(name string) error {
	switch name {
	case modelfile.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Modelfile edge %s", name)
}

// SettingMutation represents an operation that mutates the Setting nodes in the graph.
type SettingMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	_default      *string
	value         *string
	isActive      *bool
	readOnly      *bool
	createdAt     *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Setting, error)
	predicates    []predicate.Setting
}

var _ ent.Mutation = (*SettingMutation)(nil)

// settingOption allows management of the mutation configuration using functional options.
type settingOption func(*SettingMutation)

// newSettingMutation creates new mutation for the Setting entity.
func newSettingMutation(c config, op Op, opts ...settingOption) *SettingMutation {
	m := &SettingMutation{
		config:        c,
		op:            op,
		typ:           TypeSetting,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSettingID sets the ID field of the mutation.
func withSettingID(id int) settingOption {
	return func(m *SettingMutation) {
		var (
			err   error
			once  sync.Once
			value *Setting
		)
		m.oldValue = func(ctx context.Context) (*Setting, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Setting.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSetting sets the old Setting of the mutation.
func withSetting(node *Setting) settingOption {
	return func(m *SettingMutation) {
		m.oldValue = func(context.Context) (*Setting, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SettingMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SettingMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SettingMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SettingMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Setting.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *SettingMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *SettingMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Setting entity.
// If the Setting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SettingMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *SettingMutation) ResetName() {
	m.name = nil
}

// SetDefault sets the "default" field.
func (m *SettingMutation) SetDefault(s string) {
	m._default = &s
}

// Default returns the value of the "default" field in the mutation.
func (m *SettingMutation) Default() (r string, exists bool) {
	v := m._default
	if v == nil {
		return
	}
	return *v, true
}

// OldDefault returns the old "default" field's value of the Setting entity.
// If the Setting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SettingMutation) OldDefault(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDefault is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDefault requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDefault: %w", err)
	}
	return oldValue.Default, nil
}

// ClearDefault clears the value of the "default" field.
func (m *SettingMutation) ClearDefault() {
	m._default = nil
	m.clearedFields[setting.FieldDefault] = struct{}{}
}

// DefaultCleared returns if the "default" field was cleared in this mutation.
func (m *SettingMutation) DefaultCleared() bool {
	_, ok := m.clearedFields[setting.FieldDefault]
	return ok
}

// ResetDefault resets all changes to the "default" field.
func (m *SettingMutation) ResetDefault() {
	m._default = nil
	delete(m.clearedFields, setting.FieldDefault)
}

// SetValue sets the "value" field.
func (m *SettingMutation) SetValue(s string) {
	m.value = &s
}

// Value returns the value of the "value" field in the mutation.
func (m *SettingMutation) Value() (r string, exists bool) {
	v := m.value
	if v == nil {
		return
	}
	return *v, true
}

// OldValue returns the old "value" field's value of the Setting entity.
// If the Setting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SettingMutation) OldValue(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldValue is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldValue requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldValue: %w", err)
	}
	return oldValue.Value, nil
}

// ClearValue clears the value of the "value" field.
func (m *SettingMutation) ClearValue() {
	m.value = nil
	m.clearedFields[setting.FieldValue] = struct{}{}
}

// ValueCleared returns if the "value" field was cleared in this mutation.
func (m *SettingMutation) ValueCleared() bool {
	_, ok := m.clearedFields[setting.FieldValue]
	return ok
}

// ResetValue resets all changes to the "value" field.
func (m *SettingMutation) ResetValue() {
	m.value = nil
	delete(m.clearedFields, setting.FieldValue)
}

// SetIsActive sets the "isActive" field.
func (m *SettingMutation) SetIsActive(b bool) {
	m.isActive = &b
}

// IsActive returns the value of the "isActive" field in the mutation.
func (m *SettingMutation) IsActive() (r bool, exists bool) {
	v := m.isActive
	if v == nil {
		return
	}
	return *v, true
}

// OldIsActive returns the old "isActive" field's value of the Setting entity.
// If the Setting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SettingMutation) OldIsActive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsActive is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsActive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsActive: %w", err)
	}
	return oldValue.IsActive, nil
}

// ResetIsActive resets all changes to the "isActive" field.
func (m *SettingMutation) ResetIsActive() {
	m.isActive = nil
}

// SetReadOnly sets the "readOnly" field.
func (m *SettingMutation) SetReadOnly(b bool) {
	m.readOnly = &b
}

// ReadOnly returns the value of the "readOnly" field in the mutation.
func (m *SettingMutation) ReadOnly() (r bool, exists bool) {
	v := m.readOnly
	if v == nil {
		return
	}
	return *v, true
}

// OldReadOnly returns the old "readOnly" field's value of the Setting entity.
// If the Setting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SettingMutation) OldReadOnly(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldReadOnly is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldReadOnly requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReadOnly: %w", err)
	}
	return oldValue.ReadOnly, nil
}

// ResetReadOnly resets all changes to the "readOnly" field.
func (m *SettingMutation) ResetReadOnly() {
	m.readOnly = nil
}

// SetCreatedAt sets the "createdAt" field.
func (m *SettingMutation) SetCreatedAt(t time.Time) {
	m.createdAt = &t
}

// CreatedAt returns the value of the "createdAt" field in the mutation.
func (m *SettingMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.createdAt
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "createdAt" field's value of the Setting entity.
// If the Setting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SettingMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "createdAt" field.
func (m *SettingMutation) ResetCreatedAt() {
	m.createdAt = nil
}

// Where appends a list predicates to the SettingMutation builder.
func (m *SettingMutation) Where(ps ...predicate.Setting) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SettingMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SettingMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Setting, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SettingMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SettingMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Setting).
func (m *SettingMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SettingMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.name != nil {
		fields = append(fields, setting.FieldName)
	}
	if m._default != nil {
		fields = append(fields, setting.FieldDefault)
	}
	if m.value != nil {
		fields = append(fields, setting.FieldValue)
	}
	if m.isActive != nil {
		fields = append(fields, setting.FieldIsActive)
	}
	if m.readOnly != nil {
		fields = append(fields, setting.FieldReadOnly)
	}
	if m.createdAt != nil {
		fields = append(fields, setting.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SettingMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case setting.FieldName:
		return m.Name()
	case setting.FieldDefault:
		return m.Default()
	case setting.FieldValue:
		return m.Value()
	case setting.FieldIsActive:
		return m.IsActive()
	case setting.FieldReadOnly:
		return m.ReadOnly()
	case setting.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SettingMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case setting.FieldName:
		return m.OldName(ctx)
	case setting.FieldDefault:
		return m.OldDefault(ctx)
	case setting.FieldValue:
		return m.OldValue(ctx)
	case setting.FieldIsActive:
		return m.OldIsActive(ctx)
	case setting.FieldReadOnly:
		return m.OldReadOnly(ctx)
	case setting.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Setting field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SettingMutation) SetField(name string, value ent.Value) error {
	switch name {
	case setting.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case setting.FieldDefault:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDefault(v)
		return nil
	case setting.FieldValue:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetValue(v)
		return nil
	case setting.FieldIsActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsActive(v)
		return nil
	case setting.FieldReadOnly:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReadOnly(v)
		return nil
	case setting.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Setting field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SettingMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SettingMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SettingMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Setting numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SettingMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(setting.FieldDefault) {
		fields = append(fields, setting.FieldDefault)
	}
	if m.FieldCleared(setting.FieldValue) {
		fields = append(fields, setting.FieldValue)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SettingMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SettingMutation) ClearField(name string) error {
	switch name {
	case setting.FieldDefault:
		m.ClearDefault()
		return nil
	case setting.FieldValue:
		m.ClearValue()
		return nil
	}
	return fmt.Errorf("unknown Setting nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SettingMutation) ResetField(name string) error {
	switch name {
	case setting.FieldName:
		m.ResetName()
		return nil
	case setting.FieldDefault:
		m.ResetDefault()
		return nil
	case setting.FieldValue:
		m.ResetValue()
		return nil
	case setting.FieldIsActive:
		m.ResetIsActive()
		return nil
	case setting.FieldReadOnly:
		m.ResetReadOnly()
		return nil
	case setting.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Setting field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SettingMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SettingMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SettingMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SettingMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SettingMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SettingMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SettingMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Setting unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SettingMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Setting edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                Op
	typ               string
	id                *uuid.UUID
	name              *string
	email             *string
	password          *string
	role              *user.Role
	profileImageUrl   *string
	createdAt         *time.Time
	clearedFields     map[string]struct{}
	chats             map[uuid.UUID]struct{}
	removedchats      map[uuid.UUID]struct{}
	clearedchats      bool
	modelfiles        map[uuid.UUID]struct{}
	removedmodelfiles map[uuid.UUID]struct{}
	clearedmodelfiles bool
	done              bool
	oldValue          func(context.Context) (*User, error)
	predicates        []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetRole sets the "role" field.
func (m *UserMutation) SetRole(u user.Role) {
	m.role = &u
}

// Role returns the value of the "role" field in the mutation.
func (m *UserMutation) Role() (r user.Role, exists bool) {
	v := m.role
	if v == nil {
		return
	}
	return *v, true
}

// OldRole returns the old "role" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRole(ctx context.Context) (v user.Role, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRole is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRole requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRole: %w", err)
	}
	return oldValue.Role, nil
}

// ResetRole resets all changes to the "role" field.
func (m *UserMutation) ResetRole() {
	m.role = nil
}

// SetProfileImageUrl sets the "profileImageUrl" field.
func (m *UserMutation) SetProfileImageUrl(s string) {
	m.profileImageUrl = &s
}

// ProfileImageUrl returns the value of the "profileImageUrl" field in the mutation.
func (m *UserMutation) ProfileImageUrl() (r string, exists bool) {
	v := m.profileImageUrl
	if v == nil {
		return
	}
	return *v, true
}

// OldProfileImageUrl returns the old "profileImageUrl" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldProfileImageUrl(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProfileImageUrl is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProfileImageUrl requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProfileImageUrl: %w", err)
	}
	return oldValue.ProfileImageUrl, nil
}

// ResetProfileImageUrl resets all changes to the "profileImageUrl" field.
func (m *UserMutation) ResetProfileImageUrl() {
	m.profileImageUrl = nil
}

// SetCreatedAt sets the "createdAt" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.createdAt = &t
}

// CreatedAt returns the value of the "createdAt" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.createdAt
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "createdAt" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "createdAt" field.
func (m *UserMutation) ResetCreatedAt() {
	m.createdAt = nil
}

// AddChatIDs adds the "chats" edge to the Chat entity by ids.
func (m *UserMutation) AddChatIDs(ids ...uuid.UUID) {
	if m.chats == nil {
		m.chats = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.chats[ids[i]] = struct{}{}
	}
}

// ClearChats clears the "chats" edge to the Chat entity.
func (m *UserMutation) ClearChats() {
	m.clearedchats = true
}

// ChatsCleared reports if the "chats" edge to the Chat entity was cleared.
func (m *UserMutation) ChatsCleared() bool {
	return m.clearedchats
}

// RemoveChatIDs removes the "chats" edge to the Chat entity by IDs.
func (m *UserMutation) RemoveChatIDs(ids ...uuid.UUID) {
	if m.removedchats == nil {
		m.removedchats = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.chats, ids[i])
		m.removedchats[ids[i]] = struct{}{}
	}
}

// RemovedChats returns the removed IDs of the "chats" edge to the Chat entity.
func (m *UserMutation) RemovedChatsIDs() (ids []uuid.UUID) {
	for id := range m.removedchats {
		ids = append(ids, id)
	}
	return
}

// ChatsIDs returns the "chats" edge IDs in the mutation.
func (m *UserMutation) ChatsIDs() (ids []uuid.UUID) {
	for id := range m.chats {
		ids = append(ids, id)
	}
	return
}

// ResetChats resets all changes to the "chats" edge.
func (m *UserMutation) ResetChats() {
	m.chats = nil
	m.clearedchats = false
	m.removedchats = nil
}

// AddModelfileIDs adds the "modelfiles" edge to the Modelfile entity by ids.
func (m *UserMutation) AddModelfileIDs(ids ...uuid.UUID) {
	if m.modelfiles == nil {
		m.modelfiles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.modelfiles[ids[i]] = struct{}{}
	}
}

// ClearModelfiles clears the "modelfiles" edge to the Modelfile entity.
func (m *UserMutation) ClearModelfiles() {
	m.clearedmodelfiles = true
}

// ModelfilesCleared reports if the "modelfiles" edge to the Modelfile entity was cleared.
func (m *UserMutation) ModelfilesCleared() bool {
	return m.clearedmodelfiles
}

// RemoveModelfileIDs removes the "modelfiles" edge to the Modelfile entity by IDs.
func (m *UserMutation) RemoveModelfileIDs(ids ...uuid.UUID) {
	if m.removedmodelfiles == nil {
		m.removedmodelfiles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.modelfiles, ids[i])
		m.removedmodelfiles[ids[i]] = struct{}{}
	}
}

// RemovedModelfiles returns the removed IDs of the "modelfiles" edge to the Modelfile entity.
func (m *UserMutation) RemovedModelfilesIDs() (ids []uuid.UUID) {
	for id := range m.removedmodelfiles {
		ids = append(ids, id)
	}
	return
}

// ModelfilesIDs returns the "modelfiles" edge IDs in the mutation.
func (m *UserMutation) ModelfilesIDs() (ids []uuid.UUID) {
	for id := range m.modelfiles {
		ids = append(ids, id)
	}
	return
}

// ResetModelfiles resets all changes to the "modelfiles" edge.
func (m *UserMutation) ResetModelfiles() {
	m.modelfiles = nil
	m.clearedmodelfiles = false
	m.removedmodelfiles = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.role != nil {
		fields = append(fields, user.FieldRole)
	}
	if m.profileImageUrl != nil {
		fields = append(fields, user.FieldProfileImageUrl)
	}
	if m.createdAt != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	case user.FieldRole:
		return m.Role()
	case user.FieldProfileImageUrl:
		return m.ProfileImageUrl()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldRole:
		return m.OldRole(ctx)
	case user.FieldProfileImageUrl:
		return m.OldProfileImageUrl(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldRole:
		v, ok := value.(user.Role)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRole(v)
		return nil
	case user.FieldProfileImageUrl:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProfileImageUrl(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldRole:
		m.ResetRole()
		return nil
	case user.FieldProfileImageUrl:
		m.ResetProfileImageUrl()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.chats != nil {
		edges = append(edges, user.EdgeChats)
	}
	if m.modelfiles != nil {
		edges = append(edges, user.EdgeModelfiles)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeChats:
		ids := make([]ent.Value, 0, len(m.chats))
		for id := range m.chats {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeModelfiles:
		ids := make([]ent.Value, 0, len(m.modelfiles))
		for id := range m.modelfiles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedchats != nil {
		edges = append(edges, user.EdgeChats)
	}
	if m.removedmodelfiles != nil {
		edges = append(edges, user.EdgeModelfiles)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeChats:
		ids := make([]ent.Value, 0, len(m.removedchats))
		for id := range m.removedchats {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeModelfiles:
		ids := make([]ent.Value, 0, len(m.removedmodelfiles))
		for id := range m.removedmodelfiles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedchats {
		edges = append(edges, user.EdgeChats)
	}
	if m.clearedmodelfiles {
		edges = append(edges, user.EdgeModelfiles)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeChats:
		return m.clearedchats
	case user.EdgeModelfiles:
		return m.clearedmodelfiles
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeChats:
		m.ResetChats()
		return nil
	case user.EdgeModelfiles:
		m.ResetModelfiles()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
