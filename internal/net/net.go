package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/omnis-org/omnis-rest-api/pkg/model"
	"github.com/omnis-org/omnis-server/config"
)

func getProtocol() string {
	protocol := "http"
	if config.GetConfig().RestApi.TLS {
		protocol = "https"
	}
	return protocol
}

func get(path string, i interface{}) ([]byte, error) {
	var url string
	switch v := i.(type) {
	case int32:
		url = fmt.Sprintf("%s://%s:%d/%s/%d", getProtocol(), config.GetConfig().RestApi.Ip, config.GetConfig().RestApi.Port, path, v)
	case string:
		url = fmt.Sprintf("%s://%s:%d/%s/%s", getProtocol(), config.GetConfig().RestApi.Ip, config.GetConfig().RestApi.Port, path, v)
	default:
		url = fmt.Sprintf("%s://%s:%d/%s", getProtocol(), config.GetConfig().RestApi.Ip, config.GetConfig().RestApi.Port, path)
	}

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Get failed <- %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error rest api : %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll failed <- %v", err)
	}
	return body, nil
}

func postB(path string, jsonB []byte) ([]byte, error) {
	url := fmt.Sprintf("%s://%s:%d/%s", getProtocol(), config.GetConfig().RestApi.Ip, config.GetConfig().RestApi.Port, path)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonB))
	if err != nil {
		return nil, fmt.Errorf("http.Post failed <- %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error rest api : %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll failed <- %v", err)
	}

	return body, nil
}

func insertObject(o model.Object, apiPath string) (int32, error) {
	jsonB, err := json.Marshal(o)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal failed <- %v", err)
	}

	body, err := postB(apiPath, jsonB)
	if err != nil {
		return 0, fmt.Errorf("postB failed <- %v", err)
	}

	var jsonID model.IdJSON
	err = json.Unmarshal(body, &jsonID)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	id32 := int32(jsonID.Id)

	return id32, nil
}

func getObjects(apiPath string, i interface{}, objects model.Objects) error {
	data, err := get(apiPath, i)
	if err != nil {
		return fmt.Errorf("get failed <- %v", err)
	}

	err = json.Unmarshal(data, &objects)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	return nil
}

func getObject(apiPath string, i interface{}, object model.Object) error {
	data, err := get(apiPath, i)
	if err != nil {
		return fmt.Errorf("get failed <- %v", err)
	}

	err = json.Unmarshal(data, &object)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	return nil
}

func InsertPerimeter(perimeter *model.Perimeter) (int32, error) {
	return insertObject(perimeter, "api/perimeter")
}

func InsertLocation(location *model.Location) (int32, error) {
	return insertObject(location, "api/location")
}

func InsertMachine(machine *model.Machine) (int32, error) {
	return insertObject(machine, "api/machine")
}

func InsertInterface(itf *model.InterfaceO) (int32, error) {
	return insertObject(itf, "api/interface")
}

func InsertOperatingSystem(os *model.OperatingSystem) (int32, error) {
	return insertObject(os, "api/operatingSystem")
}

func InsertNetwork(network *model.Network) (int32, error) {
	return insertObject(network, "api/network")
}

func InsertGateway(gateway *model.Gateway) (int32, error) {
	return insertObject(gateway, "api/gateway")
}

func InsertInterfaceGateway(interfaceGateway *model.InterfaceGateway) (int32, error) {
	return insertObject(interfaceGateway, "api/interfaceGateway")
}

func GetInterfacesByMachineId(machineId int32) (model.InterfaceOs, error) {
	itfs := model.InterfaceOs{}
	err := getObjects("/api/interfaces/machineId", machineId, &itfs)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return itfs, nil
}

func GetOperatingSystemsByName(name string) (model.OperatingSystems, error) {
	operatingSystems := model.OperatingSystems{}
	err := getObjects("/api/operatingSystems/name", name, &operatingSystems)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return operatingSystems, nil
}

func GetNetworksByIp(ip string) (model.Networks, error) {
	networks := model.Networks{}
	err := getObjects("/api/networks/ip", ip, &networks)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return networks, nil
}

func GetInterfaceGatewaysByInterfaceId(interfaceId int32) (model.InterfaceGateways, error) {
	interfaceGateways := model.InterfaceGateways{}

	err := getObjects("/api/interfacesGateway/interfaceId/", interfaceId, &interfaceGateways)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return interfaceGateways, nil
}

func GetPerimeterByName(name string) (*model.Perimeter, error) {
	perimeter := model.Perimeter{}

	err := getObject("/api/perimeter/name", name, &perimeter)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &perimeter, nil
}

func GetLocationByName(name string) (*model.Location, error) {
	location := model.Location{}

	err := getObject("/api/location/name", name, &location)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &location, nil
}

func GetInterfaceByMac(mac string) (*model.InterfaceO, error) {
	itf := model.InterfaceO{}

	err := getObject("/api/interface/mac", mac, &itf)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &itf, nil
}

func GetMachineById(id int32) (*model.Machine, error) {
	machine := model.Machine{}

	err := getObject("/api/machine", id, &machine)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &machine, nil
}

func GetGatewayById(id int32) (*model.Gateway, error) {
	gateway := model.Gateway{}

	err := getObject("/api/gateway", id, &gateway)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &gateway, nil
}
