// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/titan/ent/action"
	"github.com/BradHacker/titan/ent/agent"
	"github.com/BradHacker/titan/ent/beacon"
	"github.com/BradHacker/titan/ent/instruction"
	"github.com/BradHacker/titan/ent/predicate"
)

// InstructionUpdate is the builder for updating Instruction entities.
type InstructionUpdate struct {
	config
	hooks    []Hook
	mutation *InstructionMutation
}

// Where adds a new predicate for the InstructionUpdate builder.
func (iu *InstructionUpdate) Where(ps ...predicate.Instruction) *InstructionUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetSentAt sets the "sentAt" field.
func (iu *InstructionUpdate) SetSentAt(t time.Time) *InstructionUpdate {
	iu.mutation.SetSentAt(t)
	return iu
}

// SetAgentID sets the "agent" edge to the Agent entity by ID.
func (iu *InstructionUpdate) SetAgentID(id int) *InstructionUpdate {
	iu.mutation.SetAgentID(id)
	return iu
}

// SetAgent sets the "agent" edge to the Agent entity.
func (iu *InstructionUpdate) SetAgent(a *Agent) *InstructionUpdate {
	return iu.SetAgentID(a.ID)
}

// SetActionID sets the "action" edge to the Action entity by ID.
func (iu *InstructionUpdate) SetActionID(id int) *InstructionUpdate {
	iu.mutation.SetActionID(id)
	return iu
}

// SetAction sets the "action" edge to the Action entity.
func (iu *InstructionUpdate) SetAction(a *Action) *InstructionUpdate {
	return iu.SetActionID(a.ID)
}

// SetBeaconID sets the "beacon" edge to the Beacon entity by ID.
func (iu *InstructionUpdate) SetBeaconID(id int) *InstructionUpdate {
	iu.mutation.SetBeaconID(id)
	return iu
}

// SetNillableBeaconID sets the "beacon" edge to the Beacon entity by ID if the given value is not nil.
func (iu *InstructionUpdate) SetNillableBeaconID(id *int) *InstructionUpdate {
	if id != nil {
		iu = iu.SetBeaconID(*id)
	}
	return iu
}

// SetBeacon sets the "beacon" edge to the Beacon entity.
func (iu *InstructionUpdate) SetBeacon(b *Beacon) *InstructionUpdate {
	return iu.SetBeaconID(b.ID)
}

// Mutation returns the InstructionMutation object of the builder.
func (iu *InstructionUpdate) Mutation() *InstructionMutation {
	return iu.mutation
}

// ClearAgent clears the "agent" edge to the Agent entity.
func (iu *InstructionUpdate) ClearAgent() *InstructionUpdate {
	iu.mutation.ClearAgent()
	return iu
}

// ClearAction clears the "action" edge to the Action entity.
func (iu *InstructionUpdate) ClearAction() *InstructionUpdate {
	iu.mutation.ClearAction()
	return iu
}

// ClearBeacon clears the "beacon" edge to the Beacon entity.
func (iu *InstructionUpdate) ClearBeacon() *InstructionUpdate {
	iu.mutation.ClearBeacon()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InstructionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstructionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InstructionUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InstructionUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InstructionUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *InstructionUpdate) check() error {
	if _, ok := iu.mutation.AgentID(); iu.mutation.AgentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"agent\"")
	}
	if _, ok := iu.mutation.ActionID(); iu.mutation.ActionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"action\"")
	}
	return nil
}

func (iu *InstructionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instruction.Table,
			Columns: instruction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instruction.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.SentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instruction.FieldSentAt,
		})
	}
	if iu.mutation.AgentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instruction.AgentTable,
			Columns: []string{instruction.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.AgentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instruction.AgentTable,
			Columns: []string{instruction.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.ActionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instruction.ActionTable,
			Columns: []string{instruction.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: action.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ActionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instruction.ActionTable,
			Columns: []string{instruction.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: action.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.BeaconCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instruction.BeaconTable,
			Columns: []string{instruction.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.BeaconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instruction.BeaconTable,
			Columns: []string{instruction.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instruction.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// InstructionUpdateOne is the builder for updating a single Instruction entity.
type InstructionUpdateOne struct {
	config
	hooks    []Hook
	mutation *InstructionMutation
}

// SetSentAt sets the "sentAt" field.
func (iuo *InstructionUpdateOne) SetSentAt(t time.Time) *InstructionUpdateOne {
	iuo.mutation.SetSentAt(t)
	return iuo
}

// SetAgentID sets the "agent" edge to the Agent entity by ID.
func (iuo *InstructionUpdateOne) SetAgentID(id int) *InstructionUpdateOne {
	iuo.mutation.SetAgentID(id)
	return iuo
}

// SetAgent sets the "agent" edge to the Agent entity.
func (iuo *InstructionUpdateOne) SetAgent(a *Agent) *InstructionUpdateOne {
	return iuo.SetAgentID(a.ID)
}

// SetActionID sets the "action" edge to the Action entity by ID.
func (iuo *InstructionUpdateOne) SetActionID(id int) *InstructionUpdateOne {
	iuo.mutation.SetActionID(id)
	return iuo
}

// SetAction sets the "action" edge to the Action entity.
func (iuo *InstructionUpdateOne) SetAction(a *Action) *InstructionUpdateOne {
	return iuo.SetActionID(a.ID)
}

// SetBeaconID sets the "beacon" edge to the Beacon entity by ID.
func (iuo *InstructionUpdateOne) SetBeaconID(id int) *InstructionUpdateOne {
	iuo.mutation.SetBeaconID(id)
	return iuo
}

// SetNillableBeaconID sets the "beacon" edge to the Beacon entity by ID if the given value is not nil.
func (iuo *InstructionUpdateOne) SetNillableBeaconID(id *int) *InstructionUpdateOne {
	if id != nil {
		iuo = iuo.SetBeaconID(*id)
	}
	return iuo
}

// SetBeacon sets the "beacon" edge to the Beacon entity.
func (iuo *InstructionUpdateOne) SetBeacon(b *Beacon) *InstructionUpdateOne {
	return iuo.SetBeaconID(b.ID)
}

// Mutation returns the InstructionMutation object of the builder.
func (iuo *InstructionUpdateOne) Mutation() *InstructionMutation {
	return iuo.mutation
}

// ClearAgent clears the "agent" edge to the Agent entity.
func (iuo *InstructionUpdateOne) ClearAgent() *InstructionUpdateOne {
	iuo.mutation.ClearAgent()
	return iuo
}

// ClearAction clears the "action" edge to the Action entity.
func (iuo *InstructionUpdateOne) ClearAction() *InstructionUpdateOne {
	iuo.mutation.ClearAction()
	return iuo
}

// ClearBeacon clears the "beacon" edge to the Beacon entity.
func (iuo *InstructionUpdateOne) ClearBeacon() *InstructionUpdateOne {
	iuo.mutation.ClearBeacon()
	return iuo
}

// Save executes the query and returns the updated Instruction entity.
func (iuo *InstructionUpdateOne) Save(ctx context.Context) (*Instruction, error) {
	var (
		err  error
		node *Instruction
	)
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstructionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InstructionUpdateOne) SaveX(ctx context.Context) *Instruction {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InstructionUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InstructionUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *InstructionUpdateOne) check() error {
	if _, ok := iuo.mutation.AgentID(); iuo.mutation.AgentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"agent\"")
	}
	if _, ok := iuo.mutation.ActionID(); iuo.mutation.ActionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"action\"")
	}
	return nil
}

func (iuo *InstructionUpdateOne) sqlSave(ctx context.Context) (_node *Instruction, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instruction.Table,
			Columns: instruction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instruction.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Instruction.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.SentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instruction.FieldSentAt,
		})
	}
	if iuo.mutation.AgentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instruction.AgentTable,
			Columns: []string{instruction.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.AgentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instruction.AgentTable,
			Columns: []string{instruction.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.ActionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instruction.ActionTable,
			Columns: []string{instruction.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: action.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ActionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instruction.ActionTable,
			Columns: []string{instruction.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: action.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.BeaconCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instruction.BeaconTable,
			Columns: []string{instruction.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.BeaconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instruction.BeaconTable,
			Columns: []string{instruction.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Instruction{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instruction.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}