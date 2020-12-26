package model

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	Id          NullInt32  `json:"id"`
	Name        NullString `json:"name"`
	Description NullString `json:"description"`
}

type Locations []Location

func (location *Location) String() string {
	return fmt.Sprintf("Location {%d, %s, %s}",
		location.Id.Int32,
		location.Name.String,
		location.Description.String)
}

func (location *Location) New() Object {
	return new(Location)
}

func (location *Location) Valid() bool {
	return location.Name.Valid
}

func (location *Location) Json() ([]byte, error) {
	return json.Marshal(location)
}

func (locations Locations) Json() ([]byte, error) {
	return json.Marshal(locations)
}
