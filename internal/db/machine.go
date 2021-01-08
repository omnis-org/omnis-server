package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

// GetMachines only return authorized machines
func GetMachines(automatic bool) (model.Machines, error) {
	log.Debug(fmt.Sprintf("GetMachines(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_machines(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var machines model.Machines

	var ignoreAuthorized model.NullBool
	for rows.Next() {
		var machine model.Machine

		err := rows.Scan(&machine.ID, &machine.UUID, &ignoreAuthorized, &machine.Hostname, &machine.Label, &machine.Description, &machine.VirtualizationSystem, &machine.SerialNumber, &machine.PerimeterID, &machine.LocationID, &machine.OperatingSystemID, &machine.MachineType, &machine.OmnisVersion)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		machines = append(machines, machine)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return machines, nil
}

// GetMachine should have a comment.
func GetMachine(id int32, automatic bool) (*model.Machine, error) {
	log.Debug(fmt.Sprintf("GetMachine(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var machine model.Machine
	var ignoreAuthorized model.NullBool
	err = db.QueryRow("CALL get_machine_by_id(?,?);", id, automatic).Scan(&machine.ID, &machine.UUID, &ignoreAuthorized, &machine.Hostname, &machine.Label, &machine.Description, &machine.VirtualizationSystem, &machine.SerialNumber, &machine.PerimeterID, &machine.LocationID, &machine.OperatingSystemID, &machine.MachineType, &machine.OmnisVersion)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &machine, nil
}

// GetPendingMachines should have a comment.
func GetPendingMachines() (model.Machines, error) {
	log.Debug("GetPendingMachines()")

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_pending_machines();")
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var machines model.Machines

	for rows.Next() {
		var machine model.Machine

		err := rows.Scan(&machine.ID, &machine.UUID, &machine.Authorized, &machine.Hostname, &machine.Label, &machine.Description, &machine.VirtualizationSystem, &machine.SerialNumber, &machine.PerimeterID, &machine.LocationID, &machine.OperatingSystemID, &machine.MachineType, &machine.OmnisVersion)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		machines = append(machines, machine)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return machines, nil
}

// InsertMachine should have a comment.
func InsertMachine(machine *model.Machine, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertMachine(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_machine(?,?,NULL,?,?,?,?,?,?,?,?,?,?);"

	err = db.QueryRow(sqlStr, machine.UUID, machine.Hostname, machine.Label, machine.Description, machine.VirtualizationSystem, machine.SerialNumber, machine.PerimeterID, machine.LocationID, machine.OperatingSystemID, machine.MachineType, machine.OmnisVersion, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

// UpdateMachine should have a comment.
func UpdateMachine(id int32, machine *model.Machine, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateMachine(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_machine(?,?,NULL,?,?,?,?,?,?,?,?,?,?,?);"

	res, err := db.Exec(sqlStr, id, machine.UUID, machine.Hostname, machine.Label, machine.Description, machine.VirtualizationSystem, machine.SerialNumber, machine.PerimeterID, machine.LocationID, machine.OperatingSystemID, machine.MachineType, machine.OmnisVersion, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// DeleteMachine should have a comment.
func DeleteMachine(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteMachine(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_machine(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// AuthorizeMachine should have a comment.
func AuthorizeMachine(id int32, authorize bool) (int64, error) {
	log.Debug(fmt.Sprintf("AuthorizeMachine(%d, %t)", id, authorize))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_machine(?,NULL,?,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,0);"

	res, err := db.Exec(sqlStr, id, authorize)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// GetOutdatedMachines only return authorized machines
func GetOutdatedMachines(outdatedDay int) (model.Machines, error) {
	log.Debug(fmt.Sprintf("GetMachines(%d)", outdatedDay))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_outdated_machines(?);", outdatedDay)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var machines model.Machines

	var ignoreAuthorized model.NullBool
	for rows.Next() {
		var machine model.Machine

		err := rows.Scan(&machine.ID, &machine.UUID, &ignoreAuthorized, &machine.Hostname, &machine.Label, &machine.Description, &machine.VirtualizationSystem, &machine.SerialNumber, &machine.PerimeterID, &machine.LocationID, &machine.OperatingSystemID, &machine.MachineType, &machine.OmnisVersion)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		machines = append(machines, machine)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return machines, nil
}

// GetMachinesO should have a comment.
func GetMachinesO(automatic bool) (model.Objects, error) {
	return GetMachines(automatic)
}

// GetMachineO should have a comment.
func GetMachineO(id int32, automatic bool) (model.Object, error) {
	return GetMachine(id, automatic)
}

// InsertMachineO should have a comment.
func InsertMachineO(object *model.Object, automatic bool) (int32, error) {
	var machine *model.Machine = (*object).(*model.Machine)
	return InsertMachine(machine, automatic)
}

// UpdateMachineO should have a comment.
func UpdateMachineO(id int32, object *model.Object, automatic bool) (int64, error) {
	var machine *model.Machine = (*object).(*model.Machine)
	return UpdateMachine(id, machine, automatic)
}

// GetOutdatedMachinesO should have a comment.
func GetOutdatedMachinesO(outdatedDay int) (model.Objects, error) {
	return GetOutdatedMachines(outdatedDay)
}
