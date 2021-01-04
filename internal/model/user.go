package model

import (
	"encoding/json"
	"fmt"
)

// User should have a comment.
type User struct {
	ID        NullInt32  `json:"id"`
	Username  NullString `json:"username"`
	Password  NullString `json:"password"`
	FirstName NullString `json:"firstName"`
	LastName  NullString `json:"lastName"`
	RoleID    NullInt32  `json:"roleId"`
}

// Users should have a comment.
type Users []User

// String should have a comment.
func (user *User) String() string {
	return fmt.Sprintf("User {%d, %s, %s, %s, %d}",
		user.ID.Int32,
		user.Username.String,
		user.FirstName.String,
		user.LastName.String,
		user.RoleID.Int32)
}

// New should have a comment.
func (user *User) New() Object {
	return new(User)
}

// Valid should have a comment.
func (user *User) Valid() bool {
	if user == nil {
		return false
	}
	return user.Username.Valid && user.Password.Valid && user.RoleID.Valid
}

// JSON should have a comment.
func (user *User) JSON() ([]byte, error) {
	return json.Marshal(user)
}

// JSON should have a comment.
func (users Users) JSON() ([]byte, error) {
	return json.Marshal(users)
}
