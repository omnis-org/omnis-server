package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

// GetNetworks should have a comment.
func GetNetworks(automatic bool) (model.Networks, error) {
	log.Debug(fmt.Sprintf("GetNetworks(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_networks(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var networks model.Networks

	for rows.Next() {
		var network model.Network

		err := rows.Scan(&network.ID, &network.Name, &network.Ipv4, &network.Ipv4Mask, &network.IsDMZ, &network.HasWifi, &network.PerimeterID)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		networks = append(networks, network)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return networks, nil
}

// GetNetwork should have a comment.
func GetNetwork(id int32, automatic bool) (*model.Network, error) {
	log.Debug(fmt.Sprintf("GetNetwork(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var network model.Network
	err = db.QueryRow("CALL get_network_by_id(?,?);", id, automatic).Scan(&network.ID, &network.Name, &network.Ipv4, &network.Ipv4Mask, &network.IsDMZ, &network.HasWifi, &network.PerimeterID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &network, nil
}

// InsertNetwork should have a comment.
func InsertNetwork(network *model.Network, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertNetwork(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_network(?,?,?,?,?,?,?);"

	err = db.QueryRow(sqlStr, network.Name, network.Ipv4, network.Ipv4Mask, network.IsDMZ, network.HasWifi, network.PerimeterID, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

// UpdateNetwork should have a comment.
func UpdateNetwork(id int32, network *model.Network, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateNetwork(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_network(?,?,?,?,?,?,?,?);"

	res, err := db.Exec(sqlStr, id, network.Name, network.Ipv4, network.Ipv4Mask, network.IsDMZ, network.HasWifi, network.PerimeterID, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// DeleteNetwork should have a comment.
func DeleteNetwork(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteNetwork(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_network(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// GetNetworksByIP should have a comment.
func GetNetworksByIP(ip string, automatic bool) (model.Networks, error) {
	log.Debug(fmt.Sprintf("GetNetworksByIp(%s,%t)", ip, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_networks_by_ip(?,?);", ip, automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var networks model.Networks

	for rows.Next() {
		var network model.Network

		err := rows.Scan(&network.ID, &network.Name, &network.Ipv4, &network.Ipv4Mask, &network.IsDMZ, &network.HasWifi, &network.PerimeterID)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		networks = append(networks, network)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return networks, nil
}

// GetNetworksO should have a comment.
func GetNetworksO(automatic bool) (model.Objects, error) {
	return GetNetworks(automatic)
}

// GetNetworkO should have a comment.
func GetNetworkO(id int32, automatic bool) (model.Object, error) {
	return GetNetwork(id, automatic)
}

// InsertNetworkO should have a comment.
func InsertNetworkO(object *model.Object, automatic bool) (int32, error) {
	var network *model.Network = (*object).(*model.Network)
	return InsertNetwork(network, automatic)
}

// UpdateNetworkO should have a comment.
func UpdateNetworkO(id int32, object *model.Object, automatic bool) (int64, error) {
	var network *model.Network = (*object).(*model.Network)
	return UpdateNetwork(id, network, automatic)
}

// GetNetworksByIPO should have a comment.
func GetNetworksByIPO(ip string, automatic bool) (model.Objects, error) {
	return GetNetworksByIP(ip, automatic)
}
