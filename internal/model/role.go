package model

// Role should have a comment.
type Role struct {
	ID                         *NullInt32  `json:"id"`
	Name                       *NullString `json:"name"`
	OmnisPermissions           *NullInt32  `json:"omnisPermissions"`
	RolesPermissions           *NullInt32  `json:"rolesPermissions"`
	UsersPermissions           *NullInt32  `json:"usersPermissions"`
	PendingMachinesPermissions *NullInt32  `json:"pendingMachinesPermissions"`
}

// Roles should have a comment.
type Roles []Role

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
