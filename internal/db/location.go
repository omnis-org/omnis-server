package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func GetLocations(automatic bool) (model.Locations, error) {
	log.Debug(fmt.Sprintf("GetLocations(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_locations(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var locations model.Locations

	for rows.Next() {
		var location model.Location

		err := rows.Scan(&location.Id, &location.Name, &location.Description)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		locations = append(locations, location)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return locations, nil
}

func GetLocation(id int32, automatic bool) (*model.Location, error) {
	log.Debug(fmt.Sprintf("GetLocation(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var location model.Location
	err = db.QueryRow("CALL get_location_by_id(?,?);", id, automatic).Scan(&location.Id, &location.Name, &location.Description)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &location, nil
}

func InsertLocation(location *model.Location, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertLocation(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_location(?,?,?);"

	err = db.QueryRow(sqlStr, location.Name, location.Description, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

func UpdateLocation(id int32, location *model.Location, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateLocation(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_location(?,?,?,?);"

	res, err := db.Exec(sqlStr, id, location.Name, location.Description, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func DeleteLocation(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteLocation(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_location(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func GetLocationByName(name string, automatic bool) (*model.Location, error) {
	log.Debug(fmt.Sprintf("GetLocationByName(%s,%t)", name, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var location model.Location
	err = db.QueryRow("CALL get_location_by_name(?,?);", name, automatic).Scan(&location.Id, &location.Name, &location.Description)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &location, nil
}

func GetLocationsO(automatic bool) (model.Objects, error) {
	return GetLocations(automatic)
}

func GetLocationO(id int32, automatic bool) (model.Object, error) {
	return GetLocation(id, automatic)
}

func InsertLocationO(object *model.Object, automatic bool) (int32, error) {
	var location *model.Location = (*object).(*model.Location)
	return InsertLocation(location, automatic)
}

func UpdateLocationO(id int32, object *model.Object, automatic bool) (int64, error) {
	var location *model.Location = (*object).(*model.Location)
	return UpdateLocation(id, location, automatic)
}

func GetLocationByNameO(name string, automatic bool) (model.Object, error) {
	return GetLocationByName(name, automatic)
}
