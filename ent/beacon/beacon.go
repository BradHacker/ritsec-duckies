// Code generated by entc, DO NOT EDIT.

package beacon

const (
	// Label holds the string label denoting the beacon type in the database.
	Label = "beacon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSentAt holds the string denoting the sentat field in the database.
	FieldSentAt = "sent_at"
	// FieldReceivedAt holds the string denoting the receivedat field in the database.
	FieldReceivedAt = "received_at"
	// EdgeInstruction holds the string denoting the instruction edge name in mutations.
	EdgeInstruction = "instruction"
	// Table holds the table name of the beacon in the database.
	Table = "beacons"
	// InstructionTable is the table the holds the instruction relation/edge.
	InstructionTable = "instructions"
	// InstructionInverseTable is the table name for the Instruction entity.
	// It exists in this package in order to avoid circular dependency with the "instruction" package.
	InstructionInverseTable = "instructions"
	// InstructionColumn is the table column denoting the instruction relation/edge.
	InstructionColumn = "beacon_instruction"
)

// Columns holds all SQL columns for beacon fields.
var Columns = []string{
	FieldID,
	FieldSentAt,
	FieldReceivedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
