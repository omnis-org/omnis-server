package model

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id        NullInt32  `json:"id"`
	Username  NullString `json:"username"`
	Password  NullString `json:"password"`
	FirstName NullString `json:"firstName"`
	LastName  NullString `json:"lastName"`
	Admin     NullBool   `json:"admin"`
}

type Users []User

func (user *User) String() string {
	return fmt.Sprintf("User {%d, %s, %s, %s, %t}",
		user.Id.Int32,
		user.Username.String,
		user.FirstName.String,
		user.LastName.String,
		user.Admin.Bool)
}

func (user *User) New() Object {
	return new(User)
}

func (user *User) Valid() bool {
	return user.Username.Valid && user.Password.Valid
}

func (user *User) Json() ([]byte, error) {
	return json.Marshal(user)
}

func (users Users) Json() ([]byte, error) {
	return json.Marshal(users)
}
