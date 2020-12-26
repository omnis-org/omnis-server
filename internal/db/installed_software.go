package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func GetInstalledSoftwares(automatic bool) (model.InstalledSoftwares, error) {
	log.Debug(fmt.Sprintf("GetInstalledSoftwares(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_installed_softwares(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var installedSoftwares model.InstalledSoftwares

	for rows.Next() {
		var installedSoftware model.InstalledSoftware

		err := rows.Scan(&installedSoftware.Id, &installedSoftware.SoftwareId, &installedSoftware.MachineId)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		installedSoftwares = append(installedSoftwares, installedSoftware)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return installedSoftwares, nil
}

func GetInstalledSoftware(id int32, automatic bool) (*model.InstalledSoftware, error) {
	log.Debug(fmt.Sprintf("GetInstalledSoftware(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var installedSoftware model.InstalledSoftware
	err = db.QueryRow("CALL get_installed_software_by_id(?,?);", id, automatic).Scan(&installedSoftware.Id, &installedSoftware.SoftwareId, &installedSoftware.MachineId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &installedSoftware, nil
}

func InsertInstalledSoftware(installedSoftware *model.InstalledSoftware, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertInstalledSoftware(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_installed_software(?,?,?);"

	err = db.QueryRow(sqlStr, installedSoftware.SoftwareId, installedSoftware.MachineId, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

func UpdateInstalledSoftware(id int32, installedSoftware *model.InstalledSoftware, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateInstalledSoftware(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_installed_software(?,?,?,?);"

	res, err := db.Exec(sqlStr, id, installedSoftware.SoftwareId, installedSoftware.MachineId, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func DeleteInstalledSoftware(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteInstalledSoftware(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_installed_software(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func GetInstalledSoftwaresO(automatic bool) (model.Objects, error) {
	return GetInstalledSoftwares(automatic)
}

func GetInstalledSoftwareO(id int32, automatic bool) (model.Object, error) {
	return GetInstalledSoftware(id, automatic)
}

func InsertInstalledSoftwareO(object *model.Object, automatic bool) (int32, error) {
	var installedSoftware *model.InstalledSoftware = (*object).(*model.InstalledSoftware)
	return InsertInstalledSoftware(installedSoftware, automatic)
}

func UpdateInstalledSoftwareO(id int32, object *model.Object, automatic bool) (int64, error) {
	var installedSoftware *model.InstalledSoftware = (*object).(*model.InstalledSoftware)
	return UpdateInstalledSoftware(id, installedSoftware, automatic)
}
