package model

// TaggedMachine should have a comment.
type TaggedMachine struct {
	ID                        *NullInt32 `json:"id"`
	TagID                     *NullInt32 `json:"tagId"`
	MachineID                 *NullInt32 `json:"machineId"`
	TagIDLastModification     *NullTime  `json:"tagIdLastModification"`
	MachineIDLastModification *NullTime  `json:"machineIdLastModification"`
}

// TaggedMachines should have a comment.
type TaggedMachines []TaggedMachine

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
