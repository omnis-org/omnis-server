package model

import (
	"encoding/json"
	"fmt"
)

// Location should have a comment.
type Location struct {
	ID          NullInt32  `json:"id"`
	Name        NullString `json:"name"`
	Description NullString `json:"description"`
}

// Locations should have a comment.
type Locations []Location

// String should have a comment.
func (location *Location) String() string {
	return fmt.Sprintf("Location {%d, %s, %s}",
		location.ID.Int32,
		location.Name.String,
		location.Description.String)
}

// New should have a comment.
func (location *Location) New() Object {
	return new(Location)
}

// Valid should have a comment.
func (location *Location) Valid() bool {
	if location == nil {
		return false
	}
	return location.Name.Valid
}

// JSON should have a comment.
func (location *Location) JSON() ([]byte, error) {
	return json.Marshal(location)
}

// JSON should have a comment.
func (locations Locations) JSON() ([]byte, error) {
	return json.Marshal(locations)
}
