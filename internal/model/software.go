package model

import (
	"encoding/json"
	"fmt"
)

// Software should have a comment.
type Software struct {
	ID       NullInt32  `json:"id"`
	Name     NullString `json:"name"`
	Version  NullString `json:"version"`
	IsIntern NullBool   `json:"isIntern"`
}

// Softwares should have a comment.
type Softwares []Software

// String should have a comment.
func (software *Software) String() string {
	return fmt.Sprintf("Software {%d, %s, %s, %t}",
		software.ID.Int32,
		software.Name.String,
		software.Version.String,
		software.IsIntern.Bool)
}

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

// JSON should have a comment.
func (software *Software) JSON() ([]byte, error) {
	return json.Marshal(software)
}

// JSON should have a comment.
func (softwares Softwares) JSON() ([]byte, error) {
	return json.Marshal(softwares)
}
