package model

// Tag should have a comment.
type Tag struct {
	ID                    *NullInt32  `json:"id"`
	Name                  *NullString `json:"name"`
	Color                 *NullString `json:"color"`
	NameLastModification  *NullTime   `json:"nameLastModification"`
	ColorLastModification *NullTime   `json:"colorLastModification"`
}

// Tags should have a comment.
type Tags []Tag

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
