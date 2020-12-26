package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func GetTags(automatic bool) (model.Tags, error) {
	log.Debug(fmt.Sprintf("GetTags(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_tags(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var tags model.Tags

	for rows.Next() {
		var tag model.Tag

		err := rows.Scan(&tag.Id, &tag.Name, &tag.Color)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		tags = append(tags, tag)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return tags, nil
}

func GetTag(id int32, automatic bool) (*model.Tag, error) {
	log.Debug(fmt.Sprintf("GetTag(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var tag model.Tag
	err = db.QueryRow("CALL get_tag_by_id(?,?);", id, automatic).Scan(&tag.Id, &tag.Name, &tag.Color)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &tag, nil
}

func InsertTag(tag *model.Tag, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertTag(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_tag(?,?,?);"

	err = db.QueryRow(sqlStr, tag.Name, tag.Color, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

func UpdateTag(id int32, tag *model.Tag, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateTag(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_tag(?,?,?,?);"

	res, err := db.Exec(sqlStr, id, tag.Name, tag.Color, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func DeleteTag(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteTag(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_tag(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func GetTagsO(automatic bool) (model.Objects, error) {
	return GetTags(automatic)
}

func GetTagO(id int32, automatic bool) (model.Object, error) {
	return GetTag(id, automatic)
}

func InsertTagO(object *model.Object, automatic bool) (int32, error) {
	var tag *model.Tag = (*object).(*model.Tag)
	return InsertTag(tag, automatic)
}

func UpdateTagO(id int32, object *model.Object, automatic bool) (int64, error) {
	var tag *model.Tag = (*object).(*model.Tag)
	return UpdateTag(id, tag, automatic)
}
