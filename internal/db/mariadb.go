package db

import (
	"database/sql"
	"fmt"
	"sync"

	// implement mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/omnis-org/omnis-server/config"
)

var lockOmnisConnection = &sync.Mutex{}
var omnisConnection *sql.DB = nil
var lockAdminConnection = &sync.Mutex{}
var adminConnection *sql.DB = nil

// CreateConnection comment
func createConnection(dbString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		return nil, fmt.Errorf("sql.Open failed <- %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db.Ping failed <- %v", err)
	}

	return db, nil
}

// CreateOmnisConnection should have a comment.
func CreateOmnisConnection() (*sql.DB, error) {
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.GetConfig().OmnisDB.Username,
		config.GetConfig().OmnisDB.Password,
		config.GetConfig().OmnisDB.Host,
		config.GetConfig().OmnisDB.Port,
		config.GetConfig().OmnisDB.Name)
	return createConnection(dbString)
}

// CreateAdminConnection should have a comment.
func CreateAdminConnection() (*sql.DB, error) {
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.GetConfig().AdminDB.Username,
		config.GetConfig().AdminDB.Password,
		config.GetConfig().AdminDB.Host,
		config.GetConfig().AdminDB.Port,
		config.GetConfig().AdminDB.Name)
	return createConnection(dbString)
}

// GetOmnisConnection should have a comment.
func GetOmnisConnection() (*sql.DB, error) {
	var err error = nil
	lockOmnisConnection.Lock()
	defer lockOmnisConnection.Unlock()
	if omnisConnection == nil {
		omnisConnection, err = CreateOmnisConnection()
	}
	return omnisConnection, err
}

// GetAdminConnection should have a comment.
func GetAdminConnection() (*sql.DB, error) {
	var err error = nil
	lockAdminConnection.Lock()
	defer lockAdminConnection.Unlock()
	if adminConnection == nil {
		adminConnection, err = CreateAdminConnection()
	}
	return adminConnection, err
}
