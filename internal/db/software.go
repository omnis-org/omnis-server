package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func GetSoftwares(automatic bool) (model.Softwares, error) {
	log.Debug(fmt.Sprintf("GetSoftwares(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_softwares(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var softwares model.Softwares

	for rows.Next() {
		var software model.Software

		err := rows.Scan(&software.Id, &software.Name, &software.Version, &software.IsIntern)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		softwares = append(softwares, software)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return softwares, nil
}

func GetSoftware(id int32, automatic bool) (*model.Software, error) {
	log.Debug(fmt.Sprintf("GetSoftware(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var software model.Software
	err = db.QueryRow("CALL get_software_by_id(?,?);", id, automatic).Scan(&software.Id, &software.Name, &software.Version, &software.IsIntern)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &software, nil
}

func InsertSoftware(software *model.Software, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertSoftware(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_software(?,?,?,?);"

	err = db.QueryRow(sqlStr, software.Name, software.Version, software.IsIntern, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

func UpdateSoftware(id int32, software *model.Software, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateSoftware(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_software(?,?,?,?,?);"

	res, err := db.Exec(sqlStr, id, software.Name, software.Version, software.IsIntern, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func DeleteSoftware(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteSoftware(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_software(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func GetSoftwaresO(automatic bool) (model.Objects, error) {
	return GetSoftwares(automatic)
}

func GetSoftwareO(id int32, automatic bool) (model.Object, error) {
	return GetSoftware(id, automatic)
}

func InsertSoftwareO(object *model.Object, automatic bool) (int32, error) {
	var software *model.Software = (*object).(*model.Software)
	return InsertSoftware(software, automatic)
}

func UpdateSoftwareO(id int32, object *model.Object, automatic bool) (int64, error) {
	var software *model.Software = (*object).(*model.Software)
	return UpdateSoftware(id, software, automatic)
}
