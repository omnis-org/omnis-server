package net

import (
	"fmt"

	"github.com/omnis-org/omnis-rest-api/pkg/model"
	"github.com/omnis-org/omnis-server/config"
)

func InsertPerimeter(perimeter *model.Perimeter) (int32, error) {
	return insertObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), perimeter, "perimeter")
}

func InsertLocation(location *model.Location) (int32, error) {
	return insertObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), location, "location")
}

func InsertMachine(machine *model.Machine) (int32, error) {
	return insertObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), machine, "machine")
}

func InsertInterface(itf *model.InterfaceO) (int32, error) {
	return insertObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), itf, "interface")
}

func InsertOperatingSystem(os *model.OperatingSystem) (int32, error) {
	return insertObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), os, "operatingSystem")
}

func InsertNetwork(network *model.Network) (int32, error) {
	return insertObject(config.GetConfig().RestApi.OmnisPath, network, "network")
}

func InsertGateway(gateway *model.Gateway) (int32, error) {
	return insertObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), gateway, "gateway")
}

func UpdateMachine(id int32, machine *model.Machine) error {
	return updateObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), machine, "machine", id)
}

func UpdateInterface(id int32, itf *model.InterfaceO) error {
	return updateObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), itf, "interface", id)
}

func DeleteGateway(id int32) error {
	return deleteObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "gateway", id)
}

func DeleteInterface(id int32) error {
	return deleteObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "interface", id)
}

func GetInterfacesByMachineId(machineId int32) (model.InterfaceOs, error) {
	itfs := model.InterfaceOs{}
	err := getObjects(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "interfaces/machineId", machineId, &itfs)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return itfs, nil
}

func GetOperatingSystemsByName(name string) (model.OperatingSystems, error) {
	operatingSystems := model.OperatingSystems{}
	err := getObjects(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "operatingSystems/name", name, &operatingSystems)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return operatingSystems, nil
}

func GetNetworksByIp(ip string) (model.Networks, error) {
	networks := model.Networks{}
	err := getObjects(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "networks/ip", ip, &networks)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return networks, nil
}

func GetGatewaysByInterfaceId(interfaceId int32) (model.Gateways, error) {
	gateways := model.Gateways{}
	err := getObjects(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "gateways/interfaceId", interfaceId, &gateways)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return gateways, nil
}

func GetPerimeterByName(name string) (*model.Perimeter, error) {
	perimeter := model.Perimeter{}

	err := getObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "perimeter/name", name, &perimeter)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &perimeter, nil
}

func GetLocationByName(name string) (*model.Location, error) {
	location := model.Location{}

	err := getObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "location/name", name, &location)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &location, nil
}

func GetInterfaceByMac(mac string) (*model.InterfaceO, error) {
	itf := model.InterfaceO{}

	err := getObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "interface/mac", mac, &itf)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &itf, nil
}

func GetMachineById(id int32) (*model.Machine, error) {
	machine := model.Machine{}

	err := getObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "machine", id, &machine)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &machine, nil
}

func GetGatewayById(id int32) (*model.Gateway, error) {
	gateway := model.Gateway{}

	err := getObject(fmt.Sprintf("%s/auto", config.GetConfig().RestApi.OmnisPath), "gateway", id, &gateway)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &gateway, nil
}
