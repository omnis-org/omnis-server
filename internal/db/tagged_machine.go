package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func GetTaggedMachines(automatic bool) (model.TaggedMachines, error) {
	log.Debug(fmt.Sprintf("GetTaggedMachines(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_tagged_machines(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var taggedMachines model.TaggedMachines

	for rows.Next() {
		var taggedMachine model.TaggedMachine

		err := rows.Scan(&taggedMachine.Id, &taggedMachine.TagId, &taggedMachine.MachineId)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		taggedMachines = append(taggedMachines, taggedMachine)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return taggedMachines, nil
}

func GetTaggedMachine(id int32, automatic bool) (*model.TaggedMachine, error) {
	log.Debug(fmt.Sprintf("GetTaggedMachine(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var taggedMachine model.TaggedMachine
	err = db.QueryRow("CALL get_tagged_machine_by_id(?,?);", id, automatic).Scan(&taggedMachine.Id, &taggedMachine.TagId, &taggedMachine.MachineId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &taggedMachine, nil
}

func InsertTaggedMachine(taggedMachine *model.TaggedMachine, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertTaggedMachine(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_tagged_machine(?,?,?);"

	err = db.QueryRow(sqlStr, taggedMachine.TagId, taggedMachine.MachineId, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

func UpdateTaggedMachine(id int32, taggedMachine *model.TaggedMachine, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateTaggedMachine(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_tagged_machine(?,?,?,?);"

	res, err := db.Exec(sqlStr, id, taggedMachine.TagId, taggedMachine.MachineId, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func DeleteTaggedMachine(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteTaggedMachine(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_tagged_machine(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func GetTaggedMachinesO(automatic bool) (model.Objects, error) {
	return GetTaggedMachines(automatic)
}

func GetTaggedMachineO(id int32, automatic bool) (model.Object, error) {
	return GetTaggedMachine(id, automatic)
}

func InsertTaggedMachineO(object *model.Object, automatic bool) (int32, error) {
	var taggedMachine *model.TaggedMachine = (*object).(*model.TaggedMachine)
	return InsertTaggedMachine(taggedMachine, automatic)
}

func UpdateTaggedMachineO(id int32, object *model.Object, automatic bool) (int64, error) {
	var taggedMachine *model.TaggedMachine = (*object).(*model.TaggedMachine)
	return UpdateTaggedMachine(id, taggedMachine, automatic)
}
