// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/titan/ent/agent"
	"github.com/BradHacker/titan/ent/instruction"
	"github.com/BradHacker/titan/ent/predicate"
)

// AgentUpdate is the builder for updating Agent entities.
type AgentUpdate struct {
	config
	hooks    []Hook
	mutation *AgentMutation
}

// Where adds a new predicate for the AgentUpdate builder.
func (au *AgentUpdate) Where(ps ...predicate.Agent) *AgentUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetUUID sets the "uuid" field.
func (au *AgentUpdate) SetUUID(s string) *AgentUpdate {
	au.mutation.SetUUID(s)
	return au
}

// SetHostname sets the "hostname" field.
func (au *AgentUpdate) SetHostname(s string) *AgentUpdate {
	au.mutation.SetHostname(s)
	return au
}

// SetIP sets the "ip" field.
func (au *AgentUpdate) SetIP(s string) *AgentUpdate {
	au.mutation.SetIP(s)
	return au
}

// SetPort sets the "port" field.
func (au *AgentUpdate) SetPort(s string) *AgentUpdate {
	au.mutation.SetPort(s)
	return au
}

// SetPid sets the "pid" field.
func (au *AgentUpdate) SetPid(i int) *AgentUpdate {
	au.mutation.ResetPid()
	au.mutation.SetPid(i)
	return au
}

// AddPid adds i to the "pid" field.
func (au *AgentUpdate) AddPid(i int) *AgentUpdate {
	au.mutation.AddPid(i)
	return au
}

// SetInstructionID sets the "instruction" edge to the Instruction entity by ID.
func (au *AgentUpdate) SetInstructionID(id int) *AgentUpdate {
	au.mutation.SetInstructionID(id)
	return au
}

// SetNillableInstructionID sets the "instruction" edge to the Instruction entity by ID if the given value is not nil.
func (au *AgentUpdate) SetNillableInstructionID(id *int) *AgentUpdate {
	if id != nil {
		au = au.SetInstructionID(*id)
	}
	return au
}

// SetInstruction sets the "instruction" edge to the Instruction entity.
func (au *AgentUpdate) SetInstruction(i *Instruction) *AgentUpdate {
	return au.SetInstructionID(i.ID)
}

// Mutation returns the AgentMutation object of the builder.
func (au *AgentUpdate) Mutation() *AgentMutation {
	return au.mutation
}

// ClearInstruction clears the "instruction" edge to the Instruction entity.
func (au *AgentUpdate) ClearInstruction() *AgentUpdate {
	au.mutation.ClearInstruction()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AgentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AgentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AgentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AgentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AgentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agent.Table,
			Columns: agent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agent.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldUUID,
		})
	}
	if value, ok := au.mutation.Hostname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldHostname,
		})
	}
	if value, ok := au.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldIP,
		})
	}
	if value, ok := au.mutation.Port(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldPort,
		})
	}
	if value, ok := au.mutation.Pid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: agent.FieldPid,
		})
	}
	if value, ok := au.mutation.AddedPid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: agent.FieldPid,
		})
	}
	if au.mutation.InstructionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   agent.InstructionTable,
			Columns: []string{agent.InstructionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: instruction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.InstructionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   agent.InstructionTable,
			Columns: []string{agent.InstructionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: instruction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agent.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AgentUpdateOne is the builder for updating a single Agent entity.
type AgentUpdateOne struct {
	config
	hooks    []Hook
	mutation *AgentMutation
}

// SetUUID sets the "uuid" field.
func (auo *AgentUpdateOne) SetUUID(s string) *AgentUpdateOne {
	auo.mutation.SetUUID(s)
	return auo
}

// SetHostname sets the "hostname" field.
func (auo *AgentUpdateOne) SetHostname(s string) *AgentUpdateOne {
	auo.mutation.SetHostname(s)
	return auo
}

// SetIP sets the "ip" field.
func (auo *AgentUpdateOne) SetIP(s string) *AgentUpdateOne {
	auo.mutation.SetIP(s)
	return auo
}

// SetPort sets the "port" field.
func (auo *AgentUpdateOne) SetPort(s string) *AgentUpdateOne {
	auo.mutation.SetPort(s)
	return auo
}

// SetPid sets the "pid" field.
func (auo *AgentUpdateOne) SetPid(i int) *AgentUpdateOne {
	auo.mutation.ResetPid()
	auo.mutation.SetPid(i)
	return auo
}

// AddPid adds i to the "pid" field.
func (auo *AgentUpdateOne) AddPid(i int) *AgentUpdateOne {
	auo.mutation.AddPid(i)
	return auo
}

// SetInstructionID sets the "instruction" edge to the Instruction entity by ID.
func (auo *AgentUpdateOne) SetInstructionID(id int) *AgentUpdateOne {
	auo.mutation.SetInstructionID(id)
	return auo
}

// SetNillableInstructionID sets the "instruction" edge to the Instruction entity by ID if the given value is not nil.
func (auo *AgentUpdateOne) SetNillableInstructionID(id *int) *AgentUpdateOne {
	if id != nil {
		auo = auo.SetInstructionID(*id)
	}
	return auo
}

// SetInstruction sets the "instruction" edge to the Instruction entity.
func (auo *AgentUpdateOne) SetInstruction(i *Instruction) *AgentUpdateOne {
	return auo.SetInstructionID(i.ID)
}

// Mutation returns the AgentMutation object of the builder.
func (auo *AgentUpdateOne) Mutation() *AgentMutation {
	return auo.mutation
}

// ClearInstruction clears the "instruction" edge to the Instruction entity.
func (auo *AgentUpdateOne) ClearInstruction() *AgentUpdateOne {
	auo.mutation.ClearInstruction()
	return auo
}

// Save executes the query and returns the updated Agent entity.
func (auo *AgentUpdateOne) Save(ctx context.Context) (*Agent, error) {
	var (
		err  error
		node *Agent
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AgentUpdateOne) SaveX(ctx context.Context) *Agent {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AgentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AgentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AgentUpdateOne) sqlSave(ctx context.Context) (_node *Agent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agent.Table,
			Columns: agent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agent.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Agent.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldUUID,
		})
	}
	if value, ok := auo.mutation.Hostname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldHostname,
		})
	}
	if value, ok := auo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldIP,
		})
	}
	if value, ok := auo.mutation.Port(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldPort,
		})
	}
	if value, ok := auo.mutation.Pid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: agent.FieldPid,
		})
	}
	if value, ok := auo.mutation.AddedPid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: agent.FieldPid,
		})
	}
	if auo.mutation.InstructionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   agent.InstructionTable,
			Columns: []string{agent.InstructionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: instruction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.InstructionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   agent.InstructionTable,
			Columns: []string{agent.InstructionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: instruction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Agent{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agent.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
