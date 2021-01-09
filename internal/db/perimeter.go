package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

// GetPerimeters should have a comment.
func GetPerimeters(automatic bool) (model.Perimeters, error) {
	log.Debug(fmt.Sprintf("GetPerimeters(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_perimeters(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var perimeters model.Perimeters

	for rows.Next() {
		var perimeter model.Perimeter

		err := rows.Scan(&perimeter.ID, &perimeter.Name, &perimeter.Description)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		perimeters = append(perimeters, perimeter)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return perimeters, nil
}

// GetPerimeter should have a comment.
func GetPerimeter(id int32, automatic bool) (*model.Perimeter, error) {
	log.Debug(fmt.Sprintf("GetPerimeter(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var perimeter model.Perimeter
	err = db.QueryRow("CALL get_perimeter_by_id(?,?);", id, automatic).Scan(&perimeter.ID, &perimeter.Name, &perimeter.Description)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &perimeter, nil
}

// InsertPerimeter should have a comment.
func InsertPerimeter(perimeter *model.Perimeter, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertPerimeter(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_perimeter(?,?,?);"

	err = db.QueryRow(sqlStr, perimeter.Name, perimeter.Description, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

// UpdatePerimeter should have a comment.
func UpdatePerimeter(id int32, perimeter *model.Perimeter, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdatePerimeter(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_perimeter(?,?,?,?);"

	res, err := db.Exec(sqlStr, id, perimeter.Name, perimeter.Description, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// DeletePerimeter should have a comment.
func DeletePerimeter(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeletePerimeter(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_perimeter(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// GetPerimeterByName should have a comment.
func GetPerimeterByName(name string, automatic bool) (*model.Perimeter, error) {
	log.Debug(fmt.Sprintf("GetPerimeterByName(%s,%t)", name, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var perimeter model.Perimeter
	err = db.QueryRow("CALL get_perimeter_by_name(?,?);", name, automatic).Scan(&perimeter.ID, &perimeter.Name, &perimeter.Description)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &perimeter, nil
}

// GetOutdatedPerimeters only return authorized machines
func GetOutdatedPerimeters(outdatedDay int) (model.Perimeters, error) {
	log.Debug(fmt.Sprintf("GetOutdatedPerimeters(%d)", outdatedDay))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_outdated_perimeters(?);", outdatedDay)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var perimeters model.Perimeters

	for rows.Next() {
		var perimeter model.Perimeter

		err := rows.Scan(&perimeter.ID,
			&perimeter.Name,
			&perimeter.Description,
			&perimeter.NameLastModification,
			&perimeter.DescriptionLastModification)

		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		perimeters = append(perimeters, perimeter)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return perimeters, nil
}

// GetPerimetersO should have a comment.
func GetPerimetersO(automatic bool) (model.Objects, error) {
	return GetPerimeters(automatic)
}

// GetPerimeterO should have a comment.
func GetPerimeterO(id int32, automatic bool) (model.Object, error) {
	return GetPerimeter(id, automatic)
}

// InsertPerimeterO should have a comment.
func InsertPerimeterO(object *model.Object, automatic bool) (int32, error) {
	var perimeter *model.Perimeter = (*object).(*model.Perimeter)
	return InsertPerimeter(perimeter, automatic)
}

// UpdatePerimeterO should have a comment.
func UpdatePerimeterO(id int32, object *model.Object, automatic bool) (int64, error) {
	var perimeter *model.Perimeter = (*object).(*model.Perimeter)
	return UpdatePerimeter(id, perimeter, automatic)
}

// GetPerimeterByNameO should have a comment.
func GetPerimeterByNameO(name string, automatic bool) (model.Object, error) {
	return GetPerimeterByName(name, automatic)
}

// GetOutdatedPerimetersO should have a comment.
func GetOutdatedPerimetersO(outdatedDay int) (model.Objects, error) {
	return GetOutdatedPerimeters(outdatedDay)
}
