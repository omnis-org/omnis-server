package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

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

	for rows.Next() {
		var machine model.Machine

		err := rows.Scan(&machine.Id, &machine.Hostname, &machine.Label, &machine.Description, &machine.VirtualizationSystem, &machine.SerialNumber, &machine.PerimeterId, &machine.LocationId, &machine.OperatingSystemId, &machine.MachineType, &machine.OmnisVersion)
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

func GetMachine(id int32, automatic bool) (*model.Machine, error) {
	log.Debug(fmt.Sprintf("GetMachine(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var machine model.Machine
	err = db.QueryRow("CALL get_machine_by_id(?,?);", id, automatic).Scan(&machine.Id, &machine.Hostname, &machine.Label, &machine.Description, &machine.VirtualizationSystem, &machine.SerialNumber, &machine.PerimeterId, &machine.LocationId, &machine.OperatingSystemId, &machine.MachineType, &machine.OmnisVersion)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &machine, nil
}

func InsertMachine(machine *model.Machine, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertMachine(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_machine(?,?,?,?,?,?,?,?,?,?,?);"

	err = db.QueryRow(sqlStr, machine.Hostname, machine.Label, machine.Description, machine.VirtualizationSystem, machine.SerialNumber, machine.PerimeterId, machine.LocationId, machine.OperatingSystemId, machine.MachineType, machine.OmnisVersion, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

func UpdateMachine(id int32, machine *model.Machine, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateMachine(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_machine(?,?,?,?,?,?,?,?,?,?,?,?);"

	res, err := db.Exec(sqlStr, id, machine.Hostname, machine.Label, machine.Description, machine.VirtualizationSystem, machine.SerialNumber, machine.PerimeterId, machine.LocationId, machine.OperatingSystemId, machine.MachineType, machine.OmnisVersion, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

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

func GetMachinesO(automatic bool) (model.Objects, error) {
	return GetMachines(automatic)
}

func GetMachineO(id int32, automatic bool) (model.Object, error) {
	return GetMachine(id, automatic)
}

func InsertMachineO(object *model.Object, automatic bool) (int32, error) {
	var machine *model.Machine = (*object).(*model.Machine)
	return InsertMachine(machine, automatic)
}

func UpdateMachineO(id int32, object *model.Object, automatic bool) (int64, error) {
	var machine *model.Machine = (*object).(*model.Machine)
	return UpdateMachine(id, machine, automatic)
}
