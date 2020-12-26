package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func GetGateways(automatic bool) (model.Gateways, error) {
	log.Debug(fmt.Sprintf("GetGateways(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_gateways(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var gateways model.Gateways

	for rows.Next() {
		var gateway model.Gateway

		err := rows.Scan(&gateway.Id, &gateway.Ipv4, &gateway.Mask, &gateway.InterfaceId)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		gateways = append(gateways, gateway)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return gateways, nil
}

func GetGateway(id int32, automatic bool) (*model.Gateway, error) {
	log.Debug(fmt.Sprintf("GetGateway(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var gateway model.Gateway
	err = db.QueryRow("CALL get_gateway_by_id(?,?);", id, automatic).Scan(&gateway.Id, &gateway.Ipv4, &gateway.Mask, &gateway.InterfaceId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &gateway, nil
}

func InsertGateway(gateway *model.Gateway, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertGateway(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_gateway(?,?,?,?);"

	err = db.QueryRow(sqlStr, gateway.Ipv4, gateway.Mask, gateway.InterfaceId, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

func UpdateGateway(id int32, gateway *model.Gateway, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateGateway(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_gateway(?,?,?,?,?);"

	res, err := db.Exec(sqlStr, id, gateway.Ipv4, gateway.Mask, gateway.InterfaceId, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func DeleteGateway(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteGateway(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_gateway(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

func GetGatewaysByInterfaceId(interfaceId int32, automatic bool) (model.Gateways, error) {
	log.Debug(fmt.Sprintf("GetGatewaysByInterfaceId(%d,%t)", interfaceId, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_gateways_by_interface_id(?,?);", interfaceId, automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var gateways model.Gateways

	for rows.Next() {
		var gateway model.Gateway

		err := rows.Scan(&gateway.Id, &gateway.Ipv4, &gateway.Mask, &gateway.InterfaceId)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		gateways = append(gateways, gateway)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return gateways, nil
}

func GetGatewaysO(automatic bool) (model.Objects, error) {
	return GetGateways(automatic)
}

func GetGatewayO(id int32, automatic bool) (model.Object, error) {
	return GetGateway(id, automatic)
}

func InsertGatewayO(object *model.Object, automatic bool) (int32, error) {
	var gateway *model.Gateway = (*object).(*model.Gateway)
	return InsertGateway(gateway, automatic)
}

func UpdateGatewayO(id int32, object *model.Object, automatic bool) (int64, error) {
	var gateway *model.Gateway = (*object).(*model.Gateway)
	return UpdateGateway(id, gateway, automatic)
}

func GetGatewaysByInterfaceIdO(interfaceId int32, automatic bool) (model.Objects, error) {
	return GetGatewaysByInterfaceId(interfaceId, automatic)
}
