package model

import (
	"encoding/json"
	"fmt"
)

type Perimeter struct {
	Id          NullInt32  `json:"id"`
	Name        NullString `json:"name"`
	Description NullString `json:"description"`
}

type Perimeters []Perimeter

func (perimeter *Perimeter) String() string {
	return fmt.Sprintf("Perimeter {%d, %s, %s}",
		perimeter.Id.Int32,
		perimeter.Name.String,
		perimeter.Description.String)
}

func (perimeter *Perimeter) New() Object {
	return new(Perimeter)
}

func (perimeter *Perimeter) Valid() bool {
	return perimeter.Name.Valid
}

func (perimeter *Perimeter) Json() ([]byte, error) {
	return json.Marshal(perimeter)
}

func (perimeters Perimeters) Json() ([]byte, error) {
	return json.Marshal(perimeters)
}
