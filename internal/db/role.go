package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

// GetRoles should have a comment.
func GetRoles() (model.Roles, error) {
	log.Debug("GetRoles()")

	db, err := GetAdminConnection()
	if err != nil {
		return nil, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	rows, err := db.Query("SELECT * FROM Role;")
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var roles model.Roles

	for rows.Next() {
		var role model.Role

		err := rows.Scan(&role.ID, &role.Name, &role.OmnisPermissions, &role.RolesPermissions, &role.UsersPermissions, &role.PendingMachinesPermissions)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return roles, nil
}

// GetRole should have a comment.
func GetRole(id int32) (*model.Role, error) {
	log.Debug(fmt.Sprintf("GetRole(%d)", id))

	db, err := GetAdminConnection()
	if err != nil {
		return nil, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	var role model.Role
	err = db.QueryRow("SELECT * FROM Role WHERE id=?;", id).Scan(&role.ID, &role.Name, &role.OmnisPermissions, &role.RolesPermissions, &role.UsersPermissions, &role.PendingMachinesPermissions)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &role, nil
}

// InsertRole should have a comment.
func InsertRole(role *model.Role) (int32, error) {
	log.Debug("InsertRole()")

	db, err := GetAdminConnection()
	if err != nil {
		return 0, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	sqlStr := "INSERT INTO Role(name,omnis_permissions,roles_permissions,users_permissions,pending_machines_permissions) VALUES(?,?,?,?,?);"

	res, err := db.Exec(sqlStr, role.Name, role.OmnisPermissions, role.RolesPermissions, role.UsersPermissions, role.PendingMachinesPermissions)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	var id int64 = 0
	id, err = res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("res.LastInsertId failed <- %v", err)
	}

	return int32(id), nil
}

// UpdateRole should have a comment.
func UpdateRole(id int32, role *model.Role) (int64, error) {
	log.Debug("UpdateRole()")

	db, err := GetAdminConnection()
	if err != nil {
		return 0, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	sqlStr := "UPDATE Role SET name = COALESCE(?, name), omnis_permissions = COALESCE(?, omnis_permissions),"
	sqlStr += "roles_permissions = COALESCE(?, roles_permissions), users_permissions = COALESCE(?, users_permissions),"
	sqlStr += "pending_machines_permissions = COALESCE(?, pending_machines_permissions) WHERE id = ?;"

	res, err := db.Exec(sqlStr, role.Name, role.OmnisPermissions, role.RolesPermissions, role.UsersPermissions, role.PendingMachinesPermissions, id)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// DeleteRole should have a comment.
func DeleteRole(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteRole(%d)", id))

	db, err := GetAdminConnection()
	if err != nil {
		return 0, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	res, err := db.Exec("DELETE FROM Role WHERE id=?;", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// GetRolesO should have a comment.
func GetRolesO(automatic bool) (model.Objects, error) {
	return GetRoles()
}

// GetRoleO should have a comment.
func GetRoleO(id int32, automatic bool) (model.Object, error) {
	return GetRole(id)
}

// InsertRoleO should have a comment.
func InsertRoleO(object *model.Object, automatic bool) (int32, error) {
	var role *model.Role = (*object).(*model.Role)
	return InsertRole(role)
}

// UpdateRoleO should have a comment.
func UpdateRoleO(id int32, object *model.Object, automatic bool) (int64, error) {
	var role *model.Role = (*object).(*model.Role)
	return UpdateRole(id, role)
}
