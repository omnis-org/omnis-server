package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type ServerConfig struct {
	Ip   string `json:"ip"`
	Port int64  `json:"port"`
}

type WorkerConfig struct {
	WaitWorkTime int64 `json:"wait_work_time"`
}

type RestApiConfig struct {
	Ip                 string `json:"ip"`
	Port               int64  `json:"port"`
	RootPath           string `json:"root_path"`
	TLS                bool   `json:"tls"`
	InsecureSkipVerify bool   `json:"insecure_skip_verify"`
}

type TlsConfig struct {
	Activated bool   `json:"activated"`
	KeyFile   string `json:"key_file"`
	CrtFile   string `json:"crt_file"`
}

type Config struct {
	Server  *ServerConfig  `json:"server"`
	Worker  *WorkerConfig  `json:"worker"`
	RestApi *RestApiConfig `json:"rest_api"`
	TLS     *TlsConfig     `json:"tls"`
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
	sc := ServerConfig{"127.0.0.1", 4320}
	wc := WorkerConfig{60}
	rc := RestApiConfig{"127.0.0.1", 4320, "/api/auto", false, false}
	tc := TlsConfig{Activated: false}
	return &Config{&sc, &wc, &rc, &tc}
}

func GetConfig() *Config {
	lockConfig.Lock()
	defer lockConfig.Unlock()
	if loadedConfig == nil {
		return defaultConfig()
	}
	return loadedConfig
}
