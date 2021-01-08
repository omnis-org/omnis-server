package model

// Perimeter should have a comment.
type Perimeter struct {
	ID          NullInt32  `json:"id"`
	Name        NullString `json:"name"`
	Description NullString `json:"description"`
}

// Perimeters should have a comment.
type Perimeters []Perimeter

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
