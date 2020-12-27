package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

// GetUsers should have a comment.
func GetUsers() (model.Users, error) {
	log.Debug("GetUsers()")

	db, err := GetAdminConnection()
	if err != nil {
		return nil, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	rows, err := db.Query("SELECT * FROM User;")
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var users model.Users

	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Admin)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return users, nil
}

// GetUser should have a comment.
func GetUser(id int32) (*model.User, error) {
	log.Debug(fmt.Sprintf("GetUser(%d)", id))

	db, err := GetAdminConnection()
	if err != nil {
		return nil, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	var user model.User
	err = db.QueryRow("SELECT * FROM User WHERE id=?;", id).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Admin)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &user, nil
}

// InsertUser should have a comment.
func InsertUser(user *model.User) (int32, error) {
	log.Debug("InsertUser()")

	db, err := GetAdminConnection()
	if err != nil {
		return 0, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	sqlStr := "INSERT INTO User(username,password,first_name,last_name,admin) VALUES(?,?,?,?,?);"

	res, err := db.Exec(sqlStr, user.Username, user.Password, user.FirstName, user.LastName, user.Admin)

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

// UpdateUser should have a comment.
func UpdateUser(id int32, user *model.User) (int64, error) {
	log.Debug("UpdateUser()")

	db, err := GetAdminConnection()
	if err != nil {
		return 0, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	sqlStr := "UPDATE User SET username = COALESCE(?, username), password = COALESCE(?, password),"
	sqlStr += "first_name = COALESCE(?, first_name), last_name = COALESCE(?, last_name), admin = COALESCE(?, admin) WHERE id = ?;"

	res, err := db.Exec(sqlStr, user.Username, user.Password, user.FirstName, user.LastName, user.Admin, id)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// DeleteUser should have a comment.
func DeleteUser(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteUser(%d)", id))

	db, err := GetAdminConnection()
	if err != nil {
		return 0, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	res, err := db.Exec("DELETE FROM User WHERE id=?;", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// GetUserByUsername should have a comment.
func GetUserByUsername(username string) (*model.User, error) {
	log.Debug(fmt.Sprintf("GetUserByUsername(%s)", username))

	db, err := GetAdminConnection()
	if err != nil {
		return nil, fmt.Errorf("GetAdminConnection failed <- %v", err)
	}

	var user model.User
	err = db.QueryRow("SELECT * FROM User WHERE username=?;", username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Admin)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &user, nil
}

// GetUsersO should have a comment.
func GetUsersO(automatic bool) (model.Objects, error) {
	return GetUsers()
}

// GetUserO should have a comment.
func GetUserO(id int32, automatic bool) (model.Object, error) {
	return GetUser(id)
}

// InsertUserO should have a comment.
func InsertUserO(object *model.Object, automatic bool) (int32, error) {
	var user *model.User = (*object).(*model.User)
	return InsertUser(user)
}

// UpdateUserO should have a comment.
func UpdateUserO(id int32, object *model.Object, automatic bool) (int64, error) {
	var user *model.User = (*object).(*model.User)
	return UpdateUser(id, user)
}

// GetUserByUsernameO should have a comment.
func GetUserByUsernameO(username string, automatic bool) (model.Object, error) {
	return GetUserByUsername(username)
}
