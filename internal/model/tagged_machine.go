package model

import (
	"fmt"
)

// TaggedMachine should have a comment.
type TaggedMachine struct {
	ID        NullInt32 `json:"id"`
	TagID     NullInt32 `json:"tagId"`
	MachineID NullInt32 `json:"machineId"`
}

// TaggedMachines should have a comment.
type TaggedMachines []TaggedMachine

// String should have a comment.
func (taggedMachine *TaggedMachine) String() string {
	return fmt.Sprintf("TaggedMachine {%d, %d, %d}",
		taggedMachine.ID.Int32,
		taggedMachine.TagID.Int32,
		taggedMachine.MachineID.Int32)
}

// New should have a comment.
func (taggedMachine *TaggedMachine) New() Object {
	return new(TaggedMachine)
}

// Valid should have a comment.
func (taggedMachine *TaggedMachine) Valid() bool {
	if taggedMachine == nil {
		return false
	}
	return taggedMachine.TagID.Valid && taggedMachine.MachineID.Valid
}
