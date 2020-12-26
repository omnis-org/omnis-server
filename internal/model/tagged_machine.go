package model

import (
	"encoding/json"
	"fmt"
)

type TaggedMachine struct {
	Id        NullInt32 `json:"id"`
	TagId     NullInt32 `json:"tag_id"`
	MachineId NullInt32 `json:"machine_id"`
}

type TaggedMachines []TaggedMachine

func (taggedMachine *TaggedMachine) String() string {
	return fmt.Sprintf("TaggedMachine {%d, %d, %d}",
		taggedMachine.Id.Int32,
		taggedMachine.TagId.Int32,
		taggedMachine.MachineId.Int32)
}

func (taggedMachine *TaggedMachine) New() Object {
	return new(TaggedMachine)
}

func (taggedMachine *TaggedMachine) Valid() bool {
	return taggedMachine.TagId.Valid && taggedMachine.MachineId.Valid
}

func (taggedMachine *TaggedMachine) Json() ([]byte, error) {
	return json.Marshal(taggedMachine)
}

func (taggedMachines TaggedMachines) Json() ([]byte, error) {
	return json.Marshal(taggedMachines)
}
