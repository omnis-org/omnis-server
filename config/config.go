package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

// ServerConfig should have a comment.
type ServerConfig struct {
	IP       string `json:"ip"`
	Port     int64  `json:"port"`
	OmnisAPI string `json:"omnisApi"`
	AdminAPI string `json:"adminApi"`
	Admin    string `json:"admin"`
	Client   string `json:"client"`
}

// DBConfig should have a comment.
type DBConfig struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
}

// AdminConfig should have a comment.
type AdminConfig struct {
	ExpirationTokenTime int64  `json:"expirationTokenTime"`
	AuthKeyFile         string `json:"authKeyFile"`
	AuthPubFile         string `json:"authPubFile"`
	AuthSimpleKey       []byte `json:"AuthSimpleKey"`
}

// TLSConfig should have a comment.
type TLSConfig struct {
	Activated     bool   `json:"activated"`
	ServerKeyFile string `json:"serverKeyFile"`
	ServerCrtFile string `json:"serverCrtFile"`
}

// Config should have a comment.
type Config struct {
	Server  *ServerConfig `json:"server"`
	Admin   *AdminConfig  `json:"admin"`
	OmnisDB *DBConfig     `json:"omnisDb"`
	AdminDB *DBConfig     `json:"adminDb"`
	TLS     *TLSConfig    `json:"tls"`
}

var lockConfig = &sync.Mutex{}
var loadedConfig *Config = nil

// LoadConfig should have a comment.
func LoadConfig(configFile *string) error {
	lockConfig.Lock()
	defer lockConfig.Unlock()
	var loadedConfigTmp Config
	jsonS, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile failed <- %v", err)
	}

	err = json.Unmarshal(jsonS, &loadedConfigTmp)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	loadedConfig = &loadedConfigTmp
	return nil
}

func defaultConfig() *Config {
	sc := ServerConfig{"127.0.0.1", 4320, "/api/omnis", "/api/admin", "/admin", "/client"}
	ac := AdminConfig{5, "", "", []byte("SECRET_KEY")}
	omnisDbc := DBConfig{"OMNIS", "omnis", "PASSWORD", "localhost", 3306}
	adminDbc := DBConfig{"OMNIS_ADMIN", "omnis", "PASSWORD", "localhost", 3306}
	tc := TLSConfig{Activated: false}
	return &Config{&sc, &ac, &omnisDbc, &adminDbc, &tc}
}

// GetConfig should have a comment.
func GetConfig() *Config {
	lockConfig.Lock()
	defer lockConfig.Unlock()
	if loadedConfig == nil {
		return defaultConfig()
	}
	return loadedConfig
}
