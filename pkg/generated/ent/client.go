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
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/chat"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/setting"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Chat is the client for interacting with the Chat builders.
	Chat *ChatClient
	// Modelfile is the client for interacting with the Modelfile builders.
	Modelfile *ModelfileClient
	// Setting is the client for interacting with the Setting builders.
	Setting *SettingClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Chat = NewChatClient(c.config)
	c.Modelfile = NewModelfileClient(c.config)
	c.Setting = NewSettingClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Chat:      NewChatClient(cfg),
		Modelfile: NewModelfileClient(cfg),
		Setting:   NewSettingClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Chat:      NewChatClient(cfg),
		Modelfile: NewModelfileClient(cfg),
		Setting:   NewSettingClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Chat.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Chat.Use(hooks...)
	c.Modelfile.Use(hooks...)
	c.Setting.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Chat.Intercept(interceptors...)
	c.Modelfile.Intercept(interceptors...)
	c.Setting.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ChatMutation:
		return c.Chat.mutate(ctx, m)
	case *ModelfileMutation:
		return c.Modelfile.mutate(ctx, m)
	case *SettingMutation:
		return c.Setting.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ChatClient is a client for the Chat schema.
type ChatClient struct {
	config
}

// NewChatClient returns a client for the Chat from the given config.
func NewChatClient(c config) *ChatClient {
	return &ChatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chat.Hooks(f(g(h())))`.
func (c *ChatClient) Use(hooks ...Hook) {
	c.hooks.Chat = append(c.hooks.Chat, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chat.Intercept(f(g(h())))`.
func (c *ChatClient) Intercept(interceptors ...Interceptor) {
	c.inters.Chat = append(c.inters.Chat, interceptors...)
}

// Create returns a builder for creating a Chat entity.
func (c *ChatClient) Create() *ChatCreate {
	mutation := newChatMutation(c.config, OpCreate)
	return &ChatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Chat entities.
func (c *ChatClient) CreateBulk(builders ...*ChatCreate) *ChatCreateBulk {
	return &ChatCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChatClient) MapCreateBulk(slice any, setFunc func(*ChatCreate, int)) *ChatCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChatCreateBulk{err: fmt.Errorf("calling to ChatClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChatCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Chat.
func (c *ChatClient) Update() *ChatUpdate {
	mutation := newChatMutation(c.config, OpUpdate)
	return &ChatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatClient) UpdateOne(ch *Chat) *ChatUpdateOne {
	mutation := newChatMutation(c.config, OpUpdateOne, withChat(ch))
	return &ChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatClient) UpdateOneID(id uuid.UUID) *ChatUpdateOne {
	mutation := newChatMutation(c.config, OpUpdateOne, withChatID(id))
	return &ChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Chat.
func (c *ChatClient) Delete() *ChatDelete {
	mutation := newChatMutation(c.config, OpDelete)
	return &ChatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChatClient) DeleteOne(ch *Chat) *ChatDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChatClient) DeleteOneID(id uuid.UUID) *ChatDeleteOne {
	builder := c.Delete().Where(chat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChatDeleteOne{builder}
}

// Query returns a query builder for Chat.
func (c *ChatClient) Query() *ChatQuery {
	return &ChatQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChat},
		inters: c.Interceptors(),
	}
}

// Get returns a Chat entity by its id.
func (c *ChatClient) Get(ctx context.Context, id uuid.UUID) (*Chat, error) {
	return c.Query().Where(chat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatClient) GetX(ctx context.Context, id uuid.UUID) *Chat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Chat.
func (c *ChatClient) QueryOwner(ch *Chat) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chat.Table, chat.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chat.OwnerTable, chat.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChatClient) Hooks() []Hook {
	return c.hooks.Chat
}

// Interceptors returns the client interceptors.
func (c *ChatClient) Interceptors() []Interceptor {
	return c.inters.Chat
}

func (c *ChatClient) mutate(ctx context.Context, m *ChatMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChatCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChatUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChatDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Chat mutation op: %q", m.Op())
	}
}

// ModelfileClient is a client for the Modelfile schema.
type ModelfileClient struct {
	config
}

// NewModelfileClient returns a client for the Modelfile from the given config.
func NewModelfileClient(c config) *ModelfileClient {
	return &ModelfileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `modelfile.Hooks(f(g(h())))`.
func (c *ModelfileClient) Use(hooks ...Hook) {
	c.hooks.Modelfile = append(c.hooks.Modelfile, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `modelfile.Intercept(f(g(h())))`.
func (c *ModelfileClient) Intercept(interceptors ...Interceptor) {
	c.inters.Modelfile = append(c.inters.Modelfile, interceptors...)
}

// Create returns a builder for creating a Modelfile entity.
func (c *ModelfileClient) Create() *ModelfileCreate {
	mutation := newModelfileMutation(c.config, OpCreate)
	return &ModelfileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Modelfile entities.
func (c *ModelfileClient) CreateBulk(builders ...*ModelfileCreate) *ModelfileCreateBulk {
	return &ModelfileCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ModelfileClient) MapCreateBulk(slice any, setFunc func(*ModelfileCreate, int)) *ModelfileCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ModelfileCreateBulk{err: fmt.Errorf("calling to ModelfileClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ModelfileCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ModelfileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Modelfile.
func (c *ModelfileClient) Update() *ModelfileUpdate {
	mutation := newModelfileMutation(c.config, OpUpdate)
	return &ModelfileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ModelfileClient) UpdateOne(m *Modelfile) *ModelfileUpdateOne {
	mutation := newModelfileMutation(c.config, OpUpdateOne, withModelfile(m))
	return &ModelfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ModelfileClient) UpdateOneID(id uuid.UUID) *ModelfileUpdateOne {
	mutation := newModelfileMutation(c.config, OpUpdateOne, withModelfileID(id))
	return &ModelfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Modelfile.
func (c *ModelfileClient) Delete() *ModelfileDelete {
	mutation := newModelfileMutation(c.config, OpDelete)
	return &ModelfileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ModelfileClient) DeleteOne(m *Modelfile) *ModelfileDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ModelfileClient) DeleteOneID(id uuid.UUID) *ModelfileDeleteOne {
	builder := c.Delete().Where(modelfile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ModelfileDeleteOne{builder}
}

// Query returns a query builder for Modelfile.
func (c *ModelfileClient) Query() *ModelfileQuery {
	return &ModelfileQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeModelfile},
		inters: c.Interceptors(),
	}
}

// Get returns a Modelfile entity by its id.
func (c *ModelfileClient) Get(ctx context.Context, id uuid.UUID) (*Modelfile, error) {
	return c.Query().Where(modelfile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ModelfileClient) GetX(ctx context.Context, id uuid.UUID) *Modelfile {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Modelfile.
func (c *ModelfileClient) QueryOwner(m *Modelfile) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(modelfile.Table, modelfile.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, modelfile.OwnerTable, modelfile.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ModelfileClient) Hooks() []Hook {
	return c.hooks.Modelfile
}

// Interceptors returns the client interceptors.
func (c *ModelfileClient) Interceptors() []Interceptor {
	return c.inters.Modelfile
}

func (c *ModelfileClient) mutate(ctx context.Context, m *ModelfileMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ModelfileCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ModelfileUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ModelfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ModelfileDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Modelfile mutation op: %q", m.Op())
	}
}

// SettingClient is a client for the Setting schema.
type SettingClient struct {
	config
}

// NewSettingClient returns a client for the Setting from the given config.
func NewSettingClient(c config) *SettingClient {
	return &SettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `setting.Hooks(f(g(h())))`.
func (c *SettingClient) Use(hooks ...Hook) {
	c.hooks.Setting = append(c.hooks.Setting, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `setting.Intercept(f(g(h())))`.
func (c *SettingClient) Intercept(interceptors ...Interceptor) {
	c.inters.Setting = append(c.inters.Setting, interceptors...)
}

// Create returns a builder for creating a Setting entity.
func (c *SettingClient) Create() *SettingCreate {
	mutation := newSettingMutation(c.config, OpCreate)
	return &SettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Setting entities.
func (c *SettingClient) CreateBulk(builders ...*SettingCreate) *SettingCreateBulk {
	return &SettingCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SettingClient) MapCreateBulk(slice any, setFunc func(*SettingCreate, int)) *SettingCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SettingCreateBulk{err: fmt.Errorf("calling to SettingClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SettingCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Setting.
func (c *SettingClient) Update() *SettingUpdate {
	mutation := newSettingMutation(c.config, OpUpdate)
	return &SettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SettingClient) UpdateOne(s *Setting) *SettingUpdateOne {
	mutation := newSettingMutation(c.config, OpUpdateOne, withSetting(s))
	return &SettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SettingClient) UpdateOneID(id int) *SettingUpdateOne {
	mutation := newSettingMutation(c.config, OpUpdateOne, withSettingID(id))
	return &SettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Setting.
func (c *SettingClient) Delete() *SettingDelete {
	mutation := newSettingMutation(c.config, OpDelete)
	return &SettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SettingClient) DeleteOne(s *Setting) *SettingDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SettingClient) DeleteOneID(id int) *SettingDeleteOne {
	builder := c.Delete().Where(setting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SettingDeleteOne{builder}
}

// Query returns a query builder for Setting.
func (c *SettingClient) Query() *SettingQuery {
	return &SettingQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSetting},
		inters: c.Interceptors(),
	}
}

// Get returns a Setting entity by its id.
func (c *SettingClient) Get(ctx context.Context, id int) (*Setting, error) {
	return c.Query().Where(setting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SettingClient) GetX(ctx context.Context, id int) *Setting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SettingClient) Hooks() []Hook {
	return c.hooks.Setting
}

// Interceptors returns the client interceptors.
func (c *SettingClient) Interceptors() []Interceptor {
	return c.inters.Setting
}

func (c *SettingClient) mutate(ctx context.Context, m *SettingMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SettingCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SettingUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SettingDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Setting mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChats queries the chats edge of a User.
func (c *UserClient) QueryChats(u *User) *ChatQuery {
	query := (&ChatClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(chat.Table, chat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ChatsTable, user.ChatsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryModelfiles queries the modelfiles edge of a User.
func (c *UserClient) QueryModelfiles(u *User) *ModelfileQuery {
	query := (&ModelfileClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(modelfile.Table, modelfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ModelfilesTable, user.ModelfilesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Chat, Modelfile, Setting, User []ent.Hook
	}
	inters struct {
		Chat, Modelfile, Setting, User []ent.Interceptor
	}
)
