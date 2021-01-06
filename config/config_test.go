package config

import (
	"testing"
)

func initTest() {
	loadedConfig = nil
}

func TestLoadConfigEmpty(t *testing.T) {
	initTest()
	err := LoadConfig("")
	if err == nil {
		t.Error("Expected error")
	}
}

func TestLoadConfigError(t *testing.T) {
	initTest()
	err := LoadConfig("testdata/bad_example.json")
	if err == nil {
		t.Error("Expected error")
	}
}

func TestLoadConfig(t *testing.T) {
	initTest()
	err := LoadConfig("testdata/example.json")
	if err != nil {
		t.Errorf("Error : %v", err)
	}

	if loadedConfig == nil {
		t.Errorf("loadedConfig should not be nil")
	}

	// Server

	if loadedConfig.Server == nil {
		t.Errorf("loadedConfig.Server should not be nil")
	}

	if loadedConfig.Server.IP != "0.0.0.0" {
		t.Errorf("loadedConfig.Server.IP should be '0.0.0.0'")
	}

	if loadedConfig.Server.Port != 4320 {
		t.Errorf("loadedConfig.Server.IP should be 4200")
	}

	if loadedConfig.Server.OmnisAPI != "/api/omnis" {
		t.Errorf("loadedConfig.Server.IP should be '/api/omnis'")
	}

	if loadedConfig.Server.AdminAPI != "/api/admin" {
		t.Errorf("loadedConfig.Server.IP should be '/api/admin'")
	}

	if loadedConfig.Server.Admin != "/admin" {
		t.Errorf("loadedConfig.Server.IP should be '/admin'")
	}

	if loadedConfig.Server.Client != "/client" {
		t.Errorf("loadedConfig.Server.IP should be '/client'")
	}

	// Admin

	if loadedConfig.Admin == nil {
		t.Errorf("loadedConfig.Admin should not be nil")
	}

	if loadedConfig.Admin.ExpirationTokenTime != 10 {
		t.Errorf("loadedConfig.Admin should be 10")
	}

	if loadedConfig.Admin.AuthKeyFile != "../keys/auth.key" {
		t.Errorf("loadedConfig.Admin should be '../keys/auth.key'")
	}

	if loadedConfig.Admin.AuthPubFile != "../keys/auth.pub" {
		t.Errorf("loadedConfig.Admin should be '../keys/auth.key'")
	}

	// OmnisDB

	if loadedConfig.OmnisDB == nil {
		t.Errorf("loadedConfig.OmnisDB should not be nil")
	}

	if loadedConfig.OmnisDB.Name != "OMNIS" {
		t.Errorf("loadedConfig.OmnisDB.Name  should be 'OMNIS'")
	}

	if loadedConfig.OmnisDB.Username != "omnis" {
		t.Errorf("loadedConfig.OmnisDB.Username should be 'omnis'")
	}

	if loadedConfig.OmnisDB.Password != "MyBeautifulPassword8273!" {
		t.Errorf("loadedConfig.OmnisDB.Password should be 'MyBeautifulPassword8273!'")
	}

	if loadedConfig.OmnisDB.Host != "localhost" {
		t.Errorf("loadedConfig.OmnisDB.Host should be 'localhost'")
	}

	if loadedConfig.OmnisDB.Port != 3306 {
		t.Errorf("loadedConfig.OmnisDB.Port should be 3306")
	}

	// AdminDB

	if loadedConfig.AdminDB == nil {
		t.Errorf("loadedConfig.AdminDB should not be nil")
	}

	if loadedConfig.AdminDB.Name != "OMNIS_ADMIN" {
		t.Errorf("loadedConfig.AdminDB.Name  should be 'OMNIS_ADMIN'")
	}

	if loadedConfig.AdminDB.Username != "omnis" {
		t.Errorf("loadedConfig.AdminDB.Username should be 'omnis'")
	}

	if loadedConfig.AdminDB.Password != "MyBeautifulPassword8273!" {
		t.Errorf("loadedConfig.AdminDB.Password should be 'MyBeautifulPassword8273!'")
	}

	if loadedConfig.AdminDB.Host != "localhost" {
		t.Errorf("loadedConfig.AdminDB.Host should be 'localhost'")
	}

	if loadedConfig.AdminDB.Port != 3306 {
		t.Errorf("loadedConfig.AdminDB.Port should be 3306")
	}

	// TLS

	if loadedConfig.TLS == nil {
		t.Errorf("loadedConfig.TLS should not be nil")
	}

	if loadedConfig.TLS.Activated != true {
		t.Errorf("loadedConfig.AdminDB.Port should be true")
	}

	if loadedConfig.TLS.ServerKeyFile != "../keys/server.key" {
		t.Errorf("loadedConfig.AdminDB.Port should be '../keys/server.key'")
	}

	if loadedConfig.TLS.ServerCrtFile != "../keys/server.crt" {
		t.Errorf("loadedConfig.AdminDB.Port should be '../keys/server.crt'")
	}
}

func TestGetConfigNotLoaded(t *testing.T) {
	initTest()
	conf := GetConfig()

	if conf == nil {
		t.Error("Expected default conf")
	}

	if conf.Server.IP != "127.0.0.1" {
		t.Errorf("conf.Server.IP should be '127.0.0.1'")
	}
}

func TestGetConfigLoaded(t *testing.T) {
	initTest()
	err := LoadConfig("testdata/example.json")
	if err != nil {
		t.Errorf("Error : %v", err)
	}

	conf := GetConfig()

	if conf == nil || conf != loadedConfig {
		t.Error("Expected loaded conf")
	}

	if conf.Server.IP != "0.0.0.0" {
		t.Errorf("conf.Server.IP should be '0.0.0.0'")
	}
}
