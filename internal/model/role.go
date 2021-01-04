package model

import (
	"encoding/json"
	"fmt"
)

// Role should have a comment.
type Role struct {
	ID                         NullInt32  `json:"id"`
	Name                       NullString `json:"name"`
	OmnisPermissions           NullInt32  `json:"omnisPermissions"`
	RolesPermissions           NullInt32  `json:"rolesPermissions"`
	UsersPermissions           NullInt32  `json:"usersPermissions"`
	PendingMachinesPermissions NullInt32  `json:"pendingMachinesPermissions"`
}

// Roles should have a comment.
type Roles []Role

// String should have a comment.
func (role *Role) String() string {
	return fmt.Sprintf("Role {%d, %s, %d, %d, %d, %d}",
		role.ID.Int32,
		role.Name.String,
		role.OmnisPermissions.Int32,
		role.RolesPermissions.Int32,
		role.UsersPermissions.Int32,
		role.PendingMachinesPermissions.Int32)
}

// New should have a comment.
func (role *Role) New() Object {
	return new(Role)
}

// Valid should have a comment.
func (role *Role) Valid() bool {
	if role == nil {
		return false
	}
	return role.Name.Valid && role.OmnisPermissions.Valid && role.RolesPermissions.Valid && role.UsersPermissions.Valid && role.PendingMachinesPermissions.Valid
}

// JSON should have a comment.
func (role *Role) JSON() ([]byte, error) {
	return json.Marshal(role)
}

// JSON should have a comment.
func (roles Roles) JSON() ([]byte, error) {
	return json.Marshal(roles)
}
