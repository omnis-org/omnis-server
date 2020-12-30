package model

import (
	"database/sql"
	"encoding/json"
)

// Object should have a comment.
type Object interface {
	JSON() ([]byte, error)
	Valid() bool
	New() Object
}

// Objects should have a comment.
type Objects interface {
	JSON() ([]byte, error)
}

// IDJSON should have a comment.
type IDJSON struct {
	ID int32 `json:"id"`
}

// NullInt32 should have a comment.
type NullInt32 struct {
	sql.NullInt32
}

// NullInt64 should have a comment.
type NullInt64 struct {
	sql.NullInt64
}

// NullString should have a comment.
type NullString struct {
	sql.NullString
}

// NullBool should have a comment.
type NullBool struct {
	sql.NullBool
}

// MarshalJSON should have a comment.
func (v NullInt32) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int32)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON should have a comment.
func (v *NullInt32) UnmarshalJSON(data []byte) error {
	var x *int32
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int32 = *x
	} else {
		v.Valid = false
	}
	return nil
}

// MarshalJSON should have a comment.
func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON should have a comment.
func (v *NullInt64) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

// MarshalJSON should have a comment.
func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON should have a comment.
func (s *NullString) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		s.Valid = true
		s.String = *x
	} else {
		s.Valid = false
	}
	return nil
}

// MarshalJSON should have a comment.
func (b NullBool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Bool)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON should have a comment.
func (b *NullBool) UnmarshalJSON(data []byte) error {
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		b.Valid = true
		b.Bool = *x
	} else {
		b.Valid = false
	}
	return nil
}

// Compare should have a comment.
func (v *NullInt32) Compare(c int32) bool {
	if v == nil || !v.Valid {
		return false
	}

	if v.Int32 == c {
		return true
	}

	return false
}

// Compare should have a comment.
func (v *NullInt64) Compare(c int64) bool {
	if v == nil || !v.Valid {
		return false
	}

	if v.Int64 == c {
		return true
	}

	return false
}

// Compare should have a comment.
func (s *NullString) Compare(c string) bool {
	if s == nil || !s.Valid {
		return false
	}

	if s.String == c {
		return true
	}

	return false
}

// Compare should have a comment.
func (b *NullBool) Compare(c bool) bool {
	if b == nil || !b.Valid {
		return false
	}

	if b.Bool == c {
		return true
	}

	return false
}
