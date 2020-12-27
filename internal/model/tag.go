package model

import (
	"encoding/json"
	"fmt"
)

// Tag should have a comment.
type Tag struct {
	ID    NullInt32  `json:"id"`
	Name  NullString `json:"name"`
	Color NullString `json:"color"`
}

// Tags should have a comment.
type Tags []Tag

// String should have a comment.
func (tag *Tag) String() string {
	return fmt.Sprintf("Tag {%d, %s, %s}",
		tag.ID.Int32,
		tag.Name.String,
		tag.Color.String)
}

// New should have a comment.
func (tag *Tag) New() Object {
	return new(Tag)
}

// Valid should have a comment.
func (tag *Tag) Valid() bool {
	if tag == nil {
		return false
	}
	return tag.Name.Valid
}

// JSON should have a comment.
func (tag *Tag) JSON() ([]byte, error) {
	return json.Marshal(tag)
}

// JSON should have a comment.
func (tags Tags) JSON() ([]byte, error) {
	return json.Marshal(tags)
}
