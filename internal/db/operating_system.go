package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

// GetOperatingSystems should have a comment.
func GetOperatingSystems(automatic bool) (model.OperatingSystems, error) {
	log.Debug(fmt.Sprintf("GetOperatingSystems(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_operating_systems(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var operatingSystems model.OperatingSystems

	for rows.Next() {
		var operatingSystem model.OperatingSystem

		err := rows.Scan(&operatingSystem.ID, &operatingSystem.Name, &operatingSystem.Platform, &operatingSystem.PlatformFamily, &operatingSystem.PlatformVersion, &operatingSystem.KernelVersion)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		operatingSystems = append(operatingSystems, operatingSystem)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return operatingSystems, nil
}

// GetOperatingSystem should have a comment.
func GetOperatingSystem(id int32, automatic bool) (*model.OperatingSystem, error) {
	log.Debug(fmt.Sprintf("GetOperatingSystem(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var operatingSystem model.OperatingSystem
	err = db.QueryRow("CALL get_operating_system_by_id(?,?);", id, automatic).Scan(&operatingSystem.ID, &operatingSystem.Name, &operatingSystem.Platform, &operatingSystem.PlatformFamily, &operatingSystem.PlatformVersion, &operatingSystem.KernelVersion)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &operatingSystem, nil
}

// InsertOperatingSystem should have a comment.
func InsertOperatingSystem(operatingSystem *model.OperatingSystem, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertOperatingSystem(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_operating_system(?,?,?,?,?,?);"

	err = db.QueryRow(sqlStr, operatingSystem.Name, operatingSystem.Platform, operatingSystem.PlatformFamily, operatingSystem.PlatformVersion, operatingSystem.KernelVersion, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

// UpdateOperatingSystem should have a comment.
func UpdateOperatingSystem(id int32, operatingSystem *model.OperatingSystem, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateOperatingSystem(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_operating_system(?,?,?,?,?,?,?);"

	res, err := db.Exec(sqlStr, id, operatingSystem.Name, operatingSystem.Platform, operatingSystem.PlatformFamily, operatingSystem.PlatformVersion, operatingSystem.KernelVersion, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// DeleteOperatingSystem should have a comment.
func DeleteOperatingSystem(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteOperatingSystem(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_operating_system(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// GetOperatingSystemsByName should have a comment.
func GetOperatingSystemsByName(name string, automatic bool) (model.OperatingSystems, error) {
	log.Debug(fmt.Sprintf("GetOperatingSystemsByName(%s,%t)", name, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_operating_systems_by_name(?,?);", name, automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var operatingSystems model.OperatingSystems

	for rows.Next() {
		var operatingSystem model.OperatingSystem

		err := rows.Scan(&operatingSystem.ID, &operatingSystem.Name, &operatingSystem.Platform, &operatingSystem.PlatformFamily, &operatingSystem.PlatformVersion, &operatingSystem.KernelVersion)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		operatingSystems = append(operatingSystems, operatingSystem)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return operatingSystems, nil
}

// GetOperatingSystemsO should have a comment.
func GetOperatingSystemsO(automatic bool) (model.Objects, error) {
	return GetOperatingSystems(automatic)
}

// GetOperatingSystemO should have a comment.
func GetOperatingSystemO(id int32, automatic bool) (model.Object, error) {
	return GetOperatingSystem(id, automatic)
}

// InsertOperatingSystemO should have a comment.
func InsertOperatingSystemO(object *model.Object, automatic bool) (int32, error) {
	var operatingSystem *model.OperatingSystem = (*object).(*model.OperatingSystem)
	return InsertOperatingSystem(operatingSystem, automatic)
}

// UpdateOperatingSystemO should have a comment.
func UpdateOperatingSystemO(id int32, object *model.Object, automatic bool) (int64, error) {
	var operatingSystem *model.OperatingSystem = (*object).(*model.OperatingSystem)
	return UpdateOperatingSystem(id, operatingSystem, automatic)
}

// GetOperatingSystemsByNameO should have a comment.
func GetOperatingSystemsByNameO(name string, automatic bool) (model.Objects, error) {
	return GetOperatingSystemsByName(name, automatic)
}
