package model

import (
	"encoding/json"
	"fmt"
)

type Software struct {
	Id       NullInt32  `json:"id"`
	Name     NullString `json:"name"`
	Version  NullString `json:"version"`
	IsIntern NullBool   `json:"is_intern"`
}

type Softwares []Software

func (software *Software) String() string {
	return fmt.Sprintf("Software {%d, %s, %s, %t}",
		software.Id.Int32,
		software.Name.String,
		software.Version.String,
		software.IsIntern.Bool)
}

func (software *Software) New() Object {
	return new(Software)
}

func (software *Software) Valid() bool {
	return software.Name.Valid
}

func (software *Software) Json() ([]byte, error) {
	return json.Marshal(software)
}

func (softwares Softwares) Json() ([]byte, error) {
	return json.Marshal(softwares)
}
