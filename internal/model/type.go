package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Object should have a comment.
type Object interface {
	Valid() bool
	New() Object
}

// Objects should have a comment.
type Objects interface{}

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

// NullTime should have a comment.
type NullTime struct {
	sql.NullTime
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

// WARNING : MarshalJSON doesn't return val if t is not valid
func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON should have a comment.
func (t *NullTime) UnmarshalJSON(data []byte) error {
	var x *time.Time
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if t != nil {
		t.Valid = true
		t.Time = *x
	} else {
		t.Valid = false
	}
	return nil
}
