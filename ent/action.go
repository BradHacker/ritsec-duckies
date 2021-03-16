// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/BradHacker/titan/ent/action"
	"github.com/BradHacker/titan/ent/instruction"
)

// Action is the model entity for the Action schema.
type Action struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ActionType holds the value of the "actionType" field.
	ActionType action.ActionType `json:"actionType,omitempty"`
	// Cmd holds the value of the "cmd" field.
	Cmd string `json:"cmd,omitempty"`
	// Args holds the value of the "args" field.
	Args []string `json:"args,omitempty"`
	// Output holds the value of the "output" field.
	Output *string `json:"output,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActionQuery when eager-loading is set.
	Edges              ActionEdges `json:"edges"`
	instruction_action *int
}

// ActionEdges holds the relations/edges for other nodes in the graph.
type ActionEdges struct {
	// Instruction holds the value of the instruction edge.
	Instruction *Instruction `json:"instruction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InstructionOrErr returns the Instruction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActionEdges) InstructionOrErr() (*Instruction, error) {
	if e.loadedTypes[0] {
		if e.Instruction == nil {
			// The edge instruction was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instruction.Label}
		}
		return e.Instruction, nil
	}
	return nil, &NotLoadedError{edge: "instruction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Action) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case action.FieldArgs:
			values[i] = &[]byte{}
		case action.FieldID:
			values[i] = &sql.NullInt64{}
		case action.FieldActionType, action.FieldCmd, action.FieldOutput:
			values[i] = &sql.NullString{}
		case action.ForeignKeys[0]: // instruction_action
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Action", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Action fields.
func (a *Action) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case action.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case action.FieldActionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actionType", values[i])
			} else if value.Valid {
				a.ActionType = action.ActionType(value.String)
			}
		case action.FieldCmd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cmd", values[i])
			} else if value.Valid {
				a.Cmd = value.String
			}
		case action.FieldArgs:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field args", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Args); err != nil {
					return fmt.Errorf("unmarshal field args: %w", err)
				}
			}
		case action.FieldOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[i])
			} else if value.Valid {
				a.Output = new(string)
				*a.Output = value.String
			}
		case action.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field instruction_action", value)
			} else if value.Valid {
				a.instruction_action = new(int)
				*a.instruction_action = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryInstruction queries the "instruction" edge of the Action entity.
func (a *Action) QueryInstruction() *InstructionQuery {
	return (&ActionClient{config: a.config}).QueryInstruction(a)
}

// Update returns a builder for updating this Action.
// Note that you need to call Action.Unwrap() before calling this method if this Action
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Action) Update() *ActionUpdateOne {
	return (&ActionClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Action entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Action) Unwrap() *Action {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Action is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Action) String() string {
	var builder strings.Builder
	builder.WriteString("Action(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", actionType=")
	builder.WriteString(fmt.Sprintf("%v", a.ActionType))
	builder.WriteString(", cmd=")
	builder.WriteString(a.Cmd)
	builder.WriteString(", args=")
	builder.WriteString(fmt.Sprintf("%v", a.Args))
	if v := a.Output; v != nil {
		builder.WriteString(", output=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Actions is a parsable slice of Action.
type Actions []*Action

func (a Actions) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
