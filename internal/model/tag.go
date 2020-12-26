package model

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	Id    NullInt32  `json:"id"`
	Name  NullString `json:"name"`
	Color NullString `json:"color"`
}

type Tags []Tag

func (tag *Tag) String() string {
	return fmt.Sprintf("Tag {%d, %s, %s}",
		tag.Id.Int32,
		tag.Name.String,
		tag.Color.String)
}

func (tag *Tag) New() Object {
	return new(Tag)
}

func (tag *Tag) Valid() bool {
	return tag.Name.Valid
}

func (tag *Tag) Json() ([]byte, error) {
	return json.Marshal(tag)
}

func (tags Tags) Json() ([]byte, error) {
	return json.Marshal(tags)
}
