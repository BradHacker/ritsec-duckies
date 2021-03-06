// Code generated by entc, DO NOT EDIT.

package action

import (
	"fmt"
)

const (
	// Label holds the string label denoting the action type in the database.
	Label = "action"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActionType holds the string denoting the actiontype field in the database.
	FieldActionType = "action_type"
	// FieldCmd holds the string denoting the cmd field in the database.
	FieldCmd = "cmd"
	// FieldArgs holds the string denoting the args field in the database.
	FieldArgs = "args"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// EdgeInstruction holds the string denoting the instruction edge name in mutations.
	EdgeInstruction = "instruction"
	// Table holds the table name of the action in the database.
	Table = "actions"
	// InstructionTable is the table the holds the instruction relation/edge.
	InstructionTable = "actions"
	// InstructionInverseTable is the table name for the Instruction entity.
	// It exists in this package in order to avoid circular dependency with the "instruction" package.
	InstructionInverseTable = "instructions"
	// InstructionColumn is the table column denoting the instruction relation/edge.
	InstructionColumn = "instruction_action"
)

// Columns holds all SQL columns for action fields.
var Columns = []string{
	FieldID,
	FieldActionType,
	FieldCmd,
	FieldArgs,
	FieldOutput,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "actions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"instruction_action",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// ActionType defines the type for the "actionType" enum field.
type ActionType string

// ActionType values.
const (
	ActionTypeEXEC ActionType = "EXEC"
)

func (at ActionType) String() string {
	return string(at)
}

// ActionTypeValidator is a validator for the "actionType" field enum values. It is called by the builders before save.
func ActionTypeValidator(at ActionType) error {
	switch at {
	case ActionTypeEXEC:
		return nil
	default:
		return fmt.Errorf("action: invalid enum value for actionType field: %q", at)
	}
}
