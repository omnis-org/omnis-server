package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type ServerConfig struct {
	Ip       string `json:"ip"`
	Port     int64  `json:"port"`
	OmnisApi string `json:"omnis_api"`
	AdminApi string `json:"admin_api"`
}

type DbConfig struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
}

type AdminConfig struct {
	ExpirationTokenTime int64  `json:"expiration_token_time"`
	AuthKeyFile         string `json:"auth_key_file"`
	AuthPubFile         string `json:"auth_pub_file"`
	AuthSimpleKey       []byte `json:"AuthSimpleKey"`
}

type TlsConfig struct {
	Activated     bool   `json:"activated"`
	ServerKeyFile string `json:"server_key_file"`
	ServerCrtFile string `json:"server_crt_file"`
}

type Config struct {
	Server  *ServerConfig `json:"server"`
	Admin   *AdminConfig  `json:"admin"`
	OmnisDB *DbConfig     `json:"omnis_db"`
	AdminDB *DbConfig     `json:"admin_db"`
	TLS     *TlsConfig    `json:"tls"`
}

var lockConfig = &sync.Mutex{}
var loadedConfig *Config = nil

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
	sc := ServerConfig{"127.0.0.1", 4320, "/api/omnis", "/api/admin"}
	ac := AdminConfig{5, "", "", []byte("SECRET_KEY")}
	omnisDbc := DbConfig{"OMNIS", "omnis", "PASSWORD", "localhost", 3306}
	adminDbc := DbConfig{"OMNIS_ADMIN", "omnis", "PASSWORD", "localhost", 3306}
	tc := TlsConfig{Activated: false}
	return &Config{&sc, &ac, &omnisDbc, &adminDbc, &tc}
}

func GetConfig() *Config {
	lockConfig.Lock()
	defer lockConfig.Unlock()
	if loadedConfig == nil {
		return defaultConfig()
	}
	return loadedConfig
}
