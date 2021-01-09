package model

// Location should have a comment.
type Location struct {
	ID                          *NullInt32  `json:"id"`
	Name                        *NullString `json:"name"`
	Description                 *NullString `json:"description"`
	NameLastModification        *NullTime   `json:"nameLastModification"`
	DescriptionLastModification *NullTime   `json:"descriptionLastModification"`
}

// Locations should have a comment.
type Locations []Location

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
