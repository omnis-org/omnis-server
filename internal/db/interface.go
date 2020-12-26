package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func GetInterfaces(automatic bool) (model.InterfaceOs, error) {
	log.Debug(fmt.Sprintf("GetInterfaceOs(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_interfaces(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var interfaceOs model.InterfaceOs

	for rows.Next() {
		var interfaceO model.InterfaceO

		err := rows.Scan(&interfaceO.Id, &interfaceO.Name, &interfaceO.Ipv4, &interfaceO.Ipv4Mask, &interfaceO.MAC, &interfaceO.InterfaceType, &interfaceO.MachineId, &interfaceO.NetworkId)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		interfaceOs = append(interfaceOs, interfaceO)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return interfaceOs, nil
}

func GetInterface(id int32, automatic bool) (*model.InterfaceO, error) {
	log.Debug(fmt.Sprintf("GetInterfaceO(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var interfaceO model.InterfaceO
	err = db.QueryRow("CALL get_interface_by_id(?,?);", id, automatic).Scan(&interfaceO.Id, &interfaceO.Name, &interfaceO.Ipv4, &interfaceO.Ipv4Mask, &interfaceO.MAC, &interfaceO.InterfaceType, &interfaceO.MachineId, &interfaceO.NetworkId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &interfaceO, nil
}

func InsertInterface(interfaceO *model.InterfaceO, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertInterfaceO(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_interface(?,?,?,?,?,?,?,?);"

	err = db.QueryRow(sqlStr, interfaceO.Name, interfaceO.Ipv4, interfaceO.Ipv4Mask, interfaceO.MAC, interfaceO.InterfaceType, interfaceO.MachineId, interfaceO.NetworkId, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

func UpdateInterface(id int32, interfaceO *model.InterfaceO, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateInterfaceO(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_interface(?,?,?,?,?,?,?,?,?);"

	res, err := db.Exec(sqlStr, id, interfaceO.Name, interfaceO.Ipv4, interfaceO.Ipv4Mask, interfaceO.MAC, interfaceO.InterfaceType, interfaceO.MachineId, interfaceO.NetworkId, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func DeleteInterface(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteInterfaceO(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_interface(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func GetInterfaceByMac(mac string, automatic bool) (*model.InterfaceO, error) {
	log.Debug(fmt.Sprintf("GetInterfaceOByMac(%s,%t)", mac, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var interfaceO model.InterfaceO
	err = db.QueryRow("CALL get_interface_by_mac(?,?);", mac, automatic).Scan(&interfaceO.Id, &interfaceO.Name, &interfaceO.Ipv4, &interfaceO.Ipv4Mask, &interfaceO.MAC, &interfaceO.InterfaceType, &interfaceO.MachineId, &interfaceO.NetworkId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &interfaceO, nil
}

func GetInterfacesByMachineId(machineId int32, automatic bool) (model.InterfaceOs, error) {
	log.Debug(fmt.Sprintf("GetInterfaceOsByMachineId(%d,%t)", machineId, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_interfaces_by_machine_id(?,?);", machineId, automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var interfaceOs model.InterfaceOs

	for rows.Next() {
		var interfaceO model.InterfaceO

		err := rows.Scan(&interfaceO.Id, &interfaceO.Name, &interfaceO.Ipv4, &interfaceO.Ipv4Mask, &interfaceO.MAC, &interfaceO.InterfaceType, &interfaceO.MachineId, &interfaceO.NetworkId)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		interfaceOs = append(interfaceOs, interfaceO)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return interfaceOs, nil
}

func GetInterfacesO(automatic bool) (model.Objects, error) {
	return GetInterfaces(automatic)
}

func GetInterfaceO(id int32, automatic bool) (model.Object, error) {
	return GetInterface(id, automatic)
}

func InsertInterfaceO(object *model.Object, automatic bool) (int32, error) {
	var interfaceO *model.InterfaceO = (*object).(*model.InterfaceO)
	return InsertInterface(interfaceO, automatic)
}

func UpdateInterfaceO(id int32, object *model.Object, automatic bool) (int64, error) {
	var interfaceO *model.InterfaceO = (*object).(*model.InterfaceO)
	return UpdateInterface(id, interfaceO, automatic)
}

func GetInterfaceByMacO(mac string, automatic bool) (model.Object, error) {
	return GetInterfaceByMac(mac, automatic)
}

func GetInterfacesByMachineIdO(machineId int32, automatic bool) (model.Objects, error) {
	return GetInterfacesByMachineId(machineId, automatic)
}
