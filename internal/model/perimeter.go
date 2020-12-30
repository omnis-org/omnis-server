package model

import (
	"encoding/json"
	"fmt"
)

// Perimeter should have a comment.
type Perimeter struct {
	ID          NullInt32  `json:"id"`
	Name        NullString `json:"name"`
	Description NullString `json:"description"`
}

// Perimeters should have a comment.
type Perimeters []Perimeter

// String should have a comment.
func (perimeter *Perimeter) String() string {
	return fmt.Sprintf("Perimeter {%d, %s, %s}",
		perimeter.ID.Int32,
		perimeter.Name.String,
		perimeter.Description.String)
}

// New should have a comment.
func (perimeter *Perimeter) New() Object {
	return new(Perimeter)
}

// Valid should have a comment.
func (perimeter *Perimeter) Valid() bool {
	if perimeter == nil {
		return false
	}
	return perimeter.Name.Valid
}

// JSON should have a comment.
func (perimeter *Perimeter) JSON() ([]byte, error) {
	return json.Marshal(perimeter)
}

// JSON should have a comment.
func (perimeters Perimeters) JSON() ([]byte, error) {
	return json.Marshal(perimeters)
}
