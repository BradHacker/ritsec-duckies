// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/BradHacker/titan/ent/migrate"

	"github.com/BradHacker/titan/ent/action"
	"github.com/BradHacker/titan/ent/agent"
	"github.com/BradHacker/titan/ent/beacon"
	"github.com/BradHacker/titan/ent/instruction"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Action is the client for interacting with the Action builders.
	Action *ActionClient
	// Agent is the client for interacting with the Agent builders.
	Agent *AgentClient
	// Beacon is the client for interacting with the Beacon builders.
	Beacon *BeaconClient
	// Instruction is the client for interacting with the Instruction builders.
	Instruction *InstructionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Action = NewActionClient(c.config)
	c.Agent = NewAgentClient(c.config)
	c.Beacon = NewBeaconClient(c.config)
	c.Instruction = NewInstructionClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Action:      NewActionClient(cfg),
		Agent:       NewAgentClient(cfg),
		Beacon:      NewBeaconClient(cfg),
		Instruction: NewInstructionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:      cfg,
		Action:      NewActionClient(cfg),
		Agent:       NewAgentClient(cfg),
		Beacon:      NewBeaconClient(cfg),
		Instruction: NewInstructionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Action.
//		Query().
//		Count(ctx)
//
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
	c.Action.Use(hooks...)
	c.Agent.Use(hooks...)
	c.Beacon.Use(hooks...)
	c.Instruction.Use(hooks...)
}

// ActionClient is a client for the Action schema.
type ActionClient struct {
	config
}

// NewActionClient returns a client for the Action from the given config.
func NewActionClient(c config) *ActionClient {
	return &ActionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `action.Hooks(f(g(h())))`.
func (c *ActionClient) Use(hooks ...Hook) {
	c.hooks.Action = append(c.hooks.Action, hooks...)
}

// Create returns a create builder for Action.
func (c *ActionClient) Create() *ActionCreate {
	mutation := newActionMutation(c.config, OpCreate)
	return &ActionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Action entities.
func (c *ActionClient) CreateBulk(builders ...*ActionCreate) *ActionCreateBulk {
	return &ActionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Action.
func (c *ActionClient) Update() *ActionUpdate {
	mutation := newActionMutation(c.config, OpUpdate)
	return &ActionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActionClient) UpdateOne(a *Action) *ActionUpdateOne {
	mutation := newActionMutation(c.config, OpUpdateOne, withAction(a))
	return &ActionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActionClient) UpdateOneID(id int) *ActionUpdateOne {
	mutation := newActionMutation(c.config, OpUpdateOne, withActionID(id))
	return &ActionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Action.
func (c *ActionClient) Delete() *ActionDelete {
	mutation := newActionMutation(c.config, OpDelete)
	return &ActionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ActionClient) DeleteOne(a *Action) *ActionDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ActionClient) DeleteOneID(id int) *ActionDeleteOne {
	builder := c.Delete().Where(action.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActionDeleteOne{builder}
}

// Query returns a query builder for Action.
func (c *ActionClient) Query() *ActionQuery {
	return &ActionQuery{config: c.config}
}

// Get returns a Action entity by its id.
func (c *ActionClient) Get(ctx context.Context, id int) (*Action, error) {
	return c.Query().Where(action.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActionClient) GetX(ctx context.Context, id int) *Action {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstruction queries the instruction edge of a Action.
func (c *ActionClient) QueryInstruction(a *Action) *InstructionQuery {
	query := &InstructionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(action.Table, action.FieldID, id),
			sqlgraph.To(instruction.Table, instruction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, action.InstructionTable, action.InstructionColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActionClient) Hooks() []Hook {
	return c.hooks.Action
}

// AgentClient is a client for the Agent schema.
type AgentClient struct {
	config
}

// NewAgentClient returns a client for the Agent from the given config.
func NewAgentClient(c config) *AgentClient {
	return &AgentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `agent.Hooks(f(g(h())))`.
func (c *AgentClient) Use(hooks ...Hook) {
	c.hooks.Agent = append(c.hooks.Agent, hooks...)
}

// Create returns a create builder for Agent.
func (c *AgentClient) Create() *AgentCreate {
	mutation := newAgentMutation(c.config, OpCreate)
	return &AgentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Agent entities.
func (c *AgentClient) CreateBulk(builders ...*AgentCreate) *AgentCreateBulk {
	return &AgentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Agent.
func (c *AgentClient) Update() *AgentUpdate {
	mutation := newAgentMutation(c.config, OpUpdate)
	return &AgentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AgentClient) UpdateOne(a *Agent) *AgentUpdateOne {
	mutation := newAgentMutation(c.config, OpUpdateOne, withAgent(a))
	return &AgentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AgentClient) UpdateOneID(id int) *AgentUpdateOne {
	mutation := newAgentMutation(c.config, OpUpdateOne, withAgentID(id))
	return &AgentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Agent.
func (c *AgentClient) Delete() *AgentDelete {
	mutation := newAgentMutation(c.config, OpDelete)
	return &AgentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AgentClient) DeleteOne(a *Agent) *AgentDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AgentClient) DeleteOneID(id int) *AgentDeleteOne {
	builder := c.Delete().Where(agent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AgentDeleteOne{builder}
}

// Query returns a query builder for Agent.
func (c *AgentClient) Query() *AgentQuery {
	return &AgentQuery{config: c.config}
}

// Get returns a Agent entity by its id.
func (c *AgentClient) Get(ctx context.Context, id int) (*Agent, error) {
	return c.Query().Where(agent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AgentClient) GetX(ctx context.Context, id int) *Agent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstruction queries the instruction edge of a Agent.
func (c *AgentClient) QueryInstruction(a *Agent) *InstructionQuery {
	query := &InstructionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(agent.Table, agent.FieldID, id),
			sqlgraph.To(instruction.Table, instruction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, agent.InstructionTable, agent.InstructionColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AgentClient) Hooks() []Hook {
	return c.hooks.Agent
}

// BeaconClient is a client for the Beacon schema.
type BeaconClient struct {
	config
}

// NewBeaconClient returns a client for the Beacon from the given config.
func NewBeaconClient(c config) *BeaconClient {
	return &BeaconClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `beacon.Hooks(f(g(h())))`.
func (c *BeaconClient) Use(hooks ...Hook) {
	c.hooks.Beacon = append(c.hooks.Beacon, hooks...)
}

// Create returns a create builder for Beacon.
func (c *BeaconClient) Create() *BeaconCreate {
	mutation := newBeaconMutation(c.config, OpCreate)
	return &BeaconCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Beacon entities.
func (c *BeaconClient) CreateBulk(builders ...*BeaconCreate) *BeaconCreateBulk {
	return &BeaconCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Beacon.
func (c *BeaconClient) Update() *BeaconUpdate {
	mutation := newBeaconMutation(c.config, OpUpdate)
	return &BeaconUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BeaconClient) UpdateOne(b *Beacon) *BeaconUpdateOne {
	mutation := newBeaconMutation(c.config, OpUpdateOne, withBeacon(b))
	return &BeaconUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BeaconClient) UpdateOneID(id int) *BeaconUpdateOne {
	mutation := newBeaconMutation(c.config, OpUpdateOne, withBeaconID(id))
	return &BeaconUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Beacon.
func (c *BeaconClient) Delete() *BeaconDelete {
	mutation := newBeaconMutation(c.config, OpDelete)
	return &BeaconDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BeaconClient) DeleteOne(b *Beacon) *BeaconDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BeaconClient) DeleteOneID(id int) *BeaconDeleteOne {
	builder := c.Delete().Where(beacon.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BeaconDeleteOne{builder}
}

// Query returns a query builder for Beacon.
func (c *BeaconClient) Query() *BeaconQuery {
	return &BeaconQuery{config: c.config}
}

// Get returns a Beacon entity by its id.
func (c *BeaconClient) Get(ctx context.Context, id int) (*Beacon, error) {
	return c.Query().Where(beacon.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BeaconClient) GetX(ctx context.Context, id int) *Beacon {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstruction queries the instruction edge of a Beacon.
func (c *BeaconClient) QueryInstruction(b *Beacon) *InstructionQuery {
	query := &InstructionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(beacon.Table, beacon.FieldID, id),
			sqlgraph.To(instruction.Table, instruction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, beacon.InstructionTable, beacon.InstructionColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BeaconClient) Hooks() []Hook {
	return c.hooks.Beacon
}

// InstructionClient is a client for the Instruction schema.
type InstructionClient struct {
	config
}

// NewInstructionClient returns a client for the Instruction from the given config.
func NewInstructionClient(c config) *InstructionClient {
	return &InstructionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `instruction.Hooks(f(g(h())))`.
func (c *InstructionClient) Use(hooks ...Hook) {
	c.hooks.Instruction = append(c.hooks.Instruction, hooks...)
}

// Create returns a create builder for Instruction.
func (c *InstructionClient) Create() *InstructionCreate {
	mutation := newInstructionMutation(c.config, OpCreate)
	return &InstructionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Instruction entities.
func (c *InstructionClient) CreateBulk(builders ...*InstructionCreate) *InstructionCreateBulk {
	return &InstructionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Instruction.
func (c *InstructionClient) Update() *InstructionUpdate {
	mutation := newInstructionMutation(c.config, OpUpdate)
	return &InstructionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstructionClient) UpdateOne(i *Instruction) *InstructionUpdateOne {
	mutation := newInstructionMutation(c.config, OpUpdateOne, withInstruction(i))
	return &InstructionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstructionClient) UpdateOneID(id int) *InstructionUpdateOne {
	mutation := newInstructionMutation(c.config, OpUpdateOne, withInstructionID(id))
	return &InstructionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Instruction.
func (c *InstructionClient) Delete() *InstructionDelete {
	mutation := newInstructionMutation(c.config, OpDelete)
	return &InstructionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *InstructionClient) DeleteOne(i *Instruction) *InstructionDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *InstructionClient) DeleteOneID(id int) *InstructionDeleteOne {
	builder := c.Delete().Where(instruction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstructionDeleteOne{builder}
}

// Query returns a query builder for Instruction.
func (c *InstructionClient) Query() *InstructionQuery {
	return &InstructionQuery{config: c.config}
}

// Get returns a Instruction entity by its id.
func (c *InstructionClient) Get(ctx context.Context, id int) (*Instruction, error) {
	return c.Query().Where(instruction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstructionClient) GetX(ctx context.Context, id int) *Instruction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAgent queries the agent edge of a Instruction.
func (c *InstructionClient) QueryAgent(i *Instruction) *AgentQuery {
	query := &AgentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instruction.Table, instruction.FieldID, id),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, instruction.AgentTable, instruction.AgentColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAction queries the action edge of a Instruction.
func (c *InstructionClient) QueryAction(i *Instruction) *ActionQuery {
	query := &ActionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instruction.Table, instruction.FieldID, id),
			sqlgraph.To(action.Table, action.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, instruction.ActionTable, instruction.ActionColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBeacon queries the beacon edge of a Instruction.
func (c *InstructionClient) QueryBeacon(i *Instruction) *BeaconQuery {
	query := &BeaconQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instruction.Table, instruction.FieldID, id),
			sqlgraph.To(beacon.Table, beacon.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, instruction.BeaconTable, instruction.BeaconColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InstructionClient) Hooks() []Hook {
	return c.hooks.Instruction
}
