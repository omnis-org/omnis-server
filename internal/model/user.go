package model

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
