// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/titan/ent/action"
	"github.com/BradHacker/titan/ent/agent"
	"github.com/BradHacker/titan/ent/beacon"
	"github.com/BradHacker/titan/ent/instruction"
	"github.com/BradHacker/titan/ent/predicate"
)

// InstructionQuery is the builder for querying Instruction entities.
type InstructionQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Instruction
	// eager-loading edges.
	withAgent  *AgentQuery
	withAction *ActionQuery
	withBeacon *BeaconQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InstructionQuery builder.
func (iq *InstructionQuery) Where(ps ...predicate.Instruction) *InstructionQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *InstructionQuery) Limit(limit int) *InstructionQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *InstructionQuery) Offset(offset int) *InstructionQuery {
	iq.offset = &offset
	return iq
}

// Order adds an order step to the query.
func (iq *InstructionQuery) Order(o ...OrderFunc) *InstructionQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryAgent chains the current query on the "agent" edge.
func (iq *InstructionQuery) QueryAgent() *AgentQuery {
	query := &AgentQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instruction.Table, instruction.FieldID, selector),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, instruction.AgentTable, instruction.AgentColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAction chains the current query on the "action" edge.
func (iq *InstructionQuery) QueryAction() *ActionQuery {
	query := &ActionQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instruction.Table, instruction.FieldID, selector),
			sqlgraph.To(action.Table, action.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, instruction.ActionTable, instruction.ActionColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBeacon chains the current query on the "beacon" edge.
func (iq *InstructionQuery) QueryBeacon() *BeaconQuery {
	query := &BeaconQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instruction.Table, instruction.FieldID, selector),
			sqlgraph.To(beacon.Table, beacon.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, instruction.BeaconTable, instruction.BeaconColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Instruction entity from the query.
// Returns a *NotFoundError when no Instruction was found.
func (iq *InstructionQuery) First(ctx context.Context) (*Instruction, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{instruction.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InstructionQuery) FirstX(ctx context.Context) *Instruction {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Instruction ID from the query.
// Returns a *NotFoundError when no Instruction ID was found.
func (iq *InstructionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{instruction.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InstructionQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Instruction entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Instruction entity is not found.
// Returns a *NotFoundError when no Instruction entities are found.
func (iq *InstructionQuery) Only(ctx context.Context) (*Instruction, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{instruction.Label}
	default:
		return nil, &NotSingularError{instruction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InstructionQuery) OnlyX(ctx context.Context) *Instruction {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Instruction ID in the query.
// Returns a *NotSingularError when exactly one Instruction ID is not found.
// Returns a *NotFoundError when no entities are found.
func (iq *InstructionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = &NotSingularError{instruction.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InstructionQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Instructions.
func (iq *InstructionQuery) All(ctx context.Context) ([]*Instruction, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *InstructionQuery) AllX(ctx context.Context) []*Instruction {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Instruction IDs.
func (iq *InstructionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := iq.Select(instruction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InstructionQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InstructionQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InstructionQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InstructionQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InstructionQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InstructionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InstructionQuery) Clone() *InstructionQuery {
	if iq == nil {
		return nil
	}
	return &InstructionQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]OrderFunc{}, iq.order...),
		predicates: append([]predicate.Instruction{}, iq.predicates...),
		withAgent:  iq.withAgent.Clone(),
		withAction: iq.withAction.Clone(),
		withBeacon: iq.withBeacon.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithAgent tells the query-builder to eager-load the nodes that are connected to
// the "agent" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstructionQuery) WithAgent(opts ...func(*AgentQuery)) *InstructionQuery {
	query := &AgentQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withAgent = query
	return iq
}

// WithAction tells the query-builder to eager-load the nodes that are connected to
// the "action" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstructionQuery) WithAction(opts ...func(*ActionQuery)) *InstructionQuery {
	query := &ActionQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withAction = query
	return iq
}

// WithBeacon tells the query-builder to eager-load the nodes that are connected to
// the "beacon" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstructionQuery) WithBeacon(opts ...func(*BeaconQuery)) *InstructionQuery {
	query := &BeaconQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withBeacon = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SentAt time.Time `json:"sentAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Instruction.Query().
//		GroupBy(instruction.FieldSentAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *InstructionQuery) GroupBy(field string, fields ...string) *InstructionGroupBy {
	group := &InstructionGroupBy{config: iq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SentAt time.Time `json:"sentAt,omitempty"`
//	}
//
//	client.Instruction.Query().
//		Select(instruction.FieldSentAt).
//		Scan(ctx, &v)
//
func (iq *InstructionQuery) Select(field string, fields ...string) *InstructionSelect {
	iq.fields = append([]string{field}, fields...)
	return &InstructionSelect{InstructionQuery: iq}
}

func (iq *InstructionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !instruction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InstructionQuery) sqlAll(ctx context.Context) ([]*Instruction, error) {
	var (
		nodes       = []*Instruction{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [3]bool{
			iq.withAgent != nil,
			iq.withAction != nil,
			iq.withBeacon != nil,
		}
	)
	if iq.withAgent != nil || iq.withBeacon != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, instruction.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Instruction{config: iq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withAgent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Instruction)
		for i := range nodes {
			fk := nodes[i].instruction_agent
			if fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(agent.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "instruction_agent" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Agent = n
			}
		}
	}

	if query := iq.withAction; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Instruction)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Action(func(s *sql.Selector) {
			s.Where(sql.InValues(instruction.ActionColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.instruction_action
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "instruction_action" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "instruction_action" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Action = n
		}
	}

	if query := iq.withBeacon; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Instruction)
		for i := range nodes {
			fk := nodes[i].beacon_instruction
			if fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(beacon.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "beacon_instruction" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Beacon = n
			}
		}
	}

	return nodes, nil
}

func (iq *InstructionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InstructionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (iq *InstructionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instruction.Table,
			Columns: instruction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instruction.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instruction.FieldID)
		for i := range fields {
			if fields[i] != instruction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, instruction.ValidColumn)
			}
		}
	}
	return _spec
}

func (iq *InstructionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(instruction.Table)
	selector := builder.Select(t1.Columns(instruction.Columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(instruction.Columns...)...)
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector, instruction.ValidColumn)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstructionGroupBy is the group-by builder for Instruction entities.
type InstructionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InstructionGroupBy) Aggregate(fns ...AggregateFunc) *InstructionGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *InstructionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (igb *InstructionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := igb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *InstructionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstructionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (igb *InstructionGroupBy) StringsX(ctx context.Context) []string {
	v, err := igb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *InstructionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = igb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = fmt.Errorf("ent: InstructionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (igb *InstructionGroupBy) StringX(ctx context.Context) string {
	v, err := igb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *InstructionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstructionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (igb *InstructionGroupBy) IntsX(ctx context.Context) []int {
	v, err := igb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *InstructionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = igb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = fmt.Errorf("ent: InstructionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (igb *InstructionGroupBy) IntX(ctx context.Context) int {
	v, err := igb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *InstructionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstructionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (igb *InstructionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := igb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *InstructionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = igb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = fmt.Errorf("ent: InstructionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (igb *InstructionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := igb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *InstructionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstructionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (igb *InstructionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := igb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *InstructionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = igb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = fmt.Errorf("ent: InstructionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (igb *InstructionGroupBy) BoolX(ctx context.Context) bool {
	v, err := igb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (igb *InstructionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range igb.fields {
		if !instruction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *InstructionGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql
	columns := make([]string, 0, len(igb.fields)+len(igb.fns))
	columns = append(columns, igb.fields...)
	for _, fn := range igb.fns {
		columns = append(columns, fn(selector, instruction.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(igb.fields...)
}

// InstructionSelect is the builder for selecting fields of Instruction entities.
type InstructionSelect struct {
	*InstructionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *InstructionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.InstructionQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (is *InstructionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := is.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (is *InstructionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstructionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (is *InstructionSelect) StringsX(ctx context.Context) []string {
	v, err := is.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (is *InstructionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = is.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = fmt.Errorf("ent: InstructionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (is *InstructionSelect) StringX(ctx context.Context) string {
	v, err := is.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (is *InstructionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstructionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (is *InstructionSelect) IntsX(ctx context.Context) []int {
	v, err := is.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (is *InstructionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = is.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = fmt.Errorf("ent: InstructionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (is *InstructionSelect) IntX(ctx context.Context) int {
	v, err := is.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (is *InstructionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstructionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (is *InstructionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := is.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (is *InstructionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = is.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = fmt.Errorf("ent: InstructionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (is *InstructionSelect) Float64X(ctx context.Context) float64 {
	v, err := is.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (is *InstructionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstructionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (is *InstructionSelect) BoolsX(ctx context.Context) []bool {
	v, err := is.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (is *InstructionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = is.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instruction.Label}
	default:
		err = fmt.Errorf("ent: InstructionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (is *InstructionSelect) BoolX(ctx context.Context) bool {
	v, err := is.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (is *InstructionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sqlQuery().Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (is *InstructionSelect) sqlQuery() sql.Querier {
	selector := is.sql
	selector.Select(selector.Columns(is.fields...)...)
	return selector
}
