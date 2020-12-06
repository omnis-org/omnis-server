package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"sync"
)

type ServerConfig struct {
	Ip   string `json:"ip"`
	Port int64  `json:"port"`
}

type WorkerConfig struct {
	WaitWorkTime int64 `json:"wait_work_time"`
}

type AdministrationConfig struct {
	ExpirationTokenTime int64  `json:"expiration_token_time"`
	AuthKeyFile         string `json:"auth_key_file"`
	AuthPubFile         string `json:"auth_pub_file"`
	AuthSimpleKey       []byte `json:"AuthSimpleKey"`
}

type RestApiConfig struct {
	Ip                 string `json:"ip"`
	Port               int64  `json:"port"`
	AdminPath          string `json:"admin_path"`
	OmnisPath          string `json:"omnis_path"`
	TLS                bool   `json:"tls"`
	InsecureSkipVerify bool   `json:"insecure_skip_verify"`
}

type TlsConfig struct {
	Activated     bool   `json:"activated"`
	ServerKeyFile string `json:"server_key_file"`
	ServerCrtFile string `json:"server_crt_file"`
}

type Config struct {
	Server  *ServerConfig         `json:"server"`
	Worker  *WorkerConfig         `json:"worker"`
	Admin   *AdministrationConfig `json:"admin"`
	RestApi *RestApiConfig        `json:"rest_api"`
	TLS     *TlsConfig            `json:"tls"`
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
	ac := AdministrationConfig{5, "", "", []byte("SECRET_KEY")}
	rc := RestApiConfig{"127.0.0.1", 4320, "/api/auto", "/admin", false, false}
	tc := TlsConfig{Activated: false}
	return &Config{&sc, &wc, &ac, &rc, &tc}
}

func GetConfig() *Config {
	lockConfig.Lock()
	defer lockConfig.Unlock()
	if loadedConfig == nil {
		return defaultConfig()
	}
	return loadedConfig
}

func GetRestApiScheme() string {
	protocol := "http://"
	if GetConfig().RestApi.TLS {
		protocol = "https://"
	}
	return protocol
}

func GetRestApiStringUrl() string {
	return fmt.Sprintf("%s:%d", GetConfig().RestApi.Ip, GetConfig().RestApi.Port)
}

func GetRestApiUrl() (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s%s", GetRestApiScheme(), GetRestApiStringUrl()))
}
