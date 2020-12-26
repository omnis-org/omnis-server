package model

import (
	"database/sql"
	"encoding/json"
)

type Object interface {
	Json() ([]byte, error)
	Valid() bool
	New() Object
}
type Objects interface {
	Json() ([]byte, error)
}

type IdJSON struct {
	Id int32 `json:"id"`
}

type NullInt32 struct {
	sql.NullInt32
}

type NullInt64 struct {
	sql.NullInt64
}

type NullString struct {
	sql.NullString
}

type NullBool struct {
	sql.NullBool
}

func (v NullInt32) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int32)
	}

	return json.Marshal(nil)
}

func (i *NullInt32) UnmarshalJSON(data []byte) error {
	var x *int32
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		i.Valid = true
		i.Int32 = *x
	} else {
		i.Valid = false
	}
	return nil
}

func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}

	return json.Marshal(nil)
}

func (i *NullInt64) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		i.Valid = true
		i.Int64 = *x
	} else {
		i.Valid = false
	}
	return nil
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}

	return json.Marshal(nil)
}

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

func (b NullBool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Bool)
	}

	return json.Marshal(nil)
}

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

func (i *NullInt32) Compare(c int32) bool {
	if i == nil || !i.Valid {
		return false
	}

	if i.Int32 == c {
		return true
	}

	return false
}

func (i *NullInt64) Compare(c int64) bool {
	if i == nil || !i.Valid {
		return false
	}

	if i.Int64 == c {
		return true
	}

	return false
}

func (s *NullString) Compare(c string) bool {
	if s == nil || !s.Valid {
		return false
	}

	if s.String == c {
		return true
	}

	return false
}

func (b *NullBool) Compare(c bool) bool {
	if b == nil || !b.Valid {
		return false
	}

	if b.Bool == c {
		return true
	}

	return false
}
