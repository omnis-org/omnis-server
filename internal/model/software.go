package model

// Software should have a comment.
type Software struct {
	ID       NullInt32  `json:"id"`
	Name     NullString `json:"name"`
	Version  NullString `json:"version"`
	IsIntern NullBool   `json:"isIntern"`
}

// Softwares should have a comment.
type Softwares []Software

// New should have a comment.
func (software *Software) New() Object {
	return new(Software)
}

// Valid should have a comment.
func (software *Software) Valid() bool {
	if software == nil {
		return false
	}
	return software.Name.Valid
}
