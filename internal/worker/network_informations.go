package worker

import (
	"fmt"

	"github.com/omnis-org/omnis-server/internal/utils"

	"github.com/omnis-org/omnis-rest-api/pkg/model"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/internal/net"
	log "github.com/sirupsen/logrus"
)

func getOneMacInterfaces(infos *client_informations.Informations) (string, error) {
	for _, itf := range infos.NetworkInformations.Interfaces {
		if itf.MAC != "" {
			return itf.MAC, nil
		}
	}

	return "", fmt.Errorf("Not found interfaces for  %s %s machine", infos.SystemInformations.Hostname, infos.SystemInformations.SerialNumber)
}

func newNetwork(networkPart string, mask int, perimeterId int32) (int32, error) {
	var name model.NullString
	var ipv4 model.NullString
	var ipv4Mask model.NullInt32
	var perimeter model.NullInt32

	err := name.Scan(networkPart)
	if err != nil {
		return 0, fmt.Errorf("name.Scan failed <- %v", err)
	}
	err = ipv4.Scan(networkPart)
	if err != nil {
		return 0, fmt.Errorf("ipv4.Scan failed <- %v", err)
	}
	err = ipv4Mask.Scan(mask)
	if err != nil {
		return 0, fmt.Errorf("ipv4Mask.Scan failed <- %v", err)
	}
	err = perimeter.Scan(perimeterId)
	if err != nil {
		return 0, fmt.Errorf("perimeter.Scan failed <- %v", err)
	}
	networkO := model.Network{Name: name,
		Ipv4:        ipv4,
		Ipv4Mask:    ipv4Mask,
		PerimeterId: perimeter,
	}
	networkId, err := net.InsertNetwork(&networkO)
	if err != nil {
		return 0, fmt.Errorf("net.InsertNetwork failed <- %v", err)
	}
	return networkId, nil
}

func doNetwork(ip string, mask int, perimeterId int32) (int32, error) {
	var networkId int32 = 0
	networkPart := utils.GetNetworkPart(ip, mask)

	networks, err := net.GetNetworksByIp(networkPart)
	if err != nil {
		return 0, fmt.Errorf("net.GetNetworksByIp failed <- %v", err)
	}

	for _, network := range networks {
		if network.Ipv4.String == networkPart &&
			network.Ipv4Mask.Int32 == int32(mask) &&
			network.PerimeterId.Int32 == perimeterId {
			networkId = network.Id.Int32
		}
	}

	if networkId == 0 {
		networkId, err = newNetwork(networkPart, mask, perimeterId)
		if err != nil {
			return 0, fmt.Errorf("newNetwork failed <- %v", err)
		}
	}

	return networkId, nil
}

func newGateway(ip string, maskI int, interfaceId int32) (int32, error) {
	var ipv4 model.NullString
	var mask model.NullInt32
	var interfaceO model.NullInt32

	err := ipv4.Scan(ip)
	if err != nil {
		return 0, fmt.Errorf("ipv4.Scan failed <- %v", err)
	}
	err = mask.Scan(maskI)
	if err != nil {
		return 0, fmt.Errorf("mask.Scan failed <- %v", err)
	}
	err = interfaceO.Scan(interfaceId)
	if err != nil {
		return 0, fmt.Errorf("interfaceO.Scan failed <- %v", err)
	}

	gateway := model.Gateway{Ipv4: ipv4,
		Mask:        mask,
		InterfaceId: interfaceO}

	gatewayId, err := net.InsertGateway(&gateway)
	if err != nil {
		return 0, fmt.Errorf("net.InsertGateway failed <- %v", err)
	}
	return gatewayId, nil
}

func doGateways(interfaceId int32, gateways []string, mask int) error {

	oldGateways, err := net.GetGatewaysByInterfaceId(interfaceId)
	if err != nil {
		return fmt.Errorf("net.GetGatewaysByInterfaceId failed <- %v", err)
	}

	for _, gtw := range gateways {
		var gatewayId int32 = 0
		for _, oldGtw := range oldGateways {
			if oldGtw.Ipv4.String == gtw &&
				oldGtw.Mask.Int32 == int32(mask) {
				gatewayId = oldGtw.Id.Int32
			}
		}

		if gatewayId == 0 {
			gatewayId, err = newGateway(gtw, mask, interfaceId)
			if err != nil {
				return fmt.Errorf("newGateway failed <- %v", err)
			}
		}

	}

	return nil

}

func newInterfaces(networkInformation []client_informations.InterfaceInformations, machineId int32, perimeterId int32) error {
	for _, itf := range networkInformation {
		var name model.NullString
		var ipv4 model.NullString
		var ipv4Mask model.NullInt32
		var mac model.NullString
		var interfaceType model.NullString
		var machine model.NullInt32
		var network model.NullInt32

		err := name.Scan(itf.Name)
		if err != nil {
			return fmt.Errorf("name.Scan failed <- %v", err)
		}
		err = ipv4.Scan(itf.Ipv4)
		if err != nil {
			return fmt.Errorf("ipv4.Scan failed <- %v", err)
		}
		err = ipv4Mask.Scan(itf.Ipv4Mask)
		if err != nil {
			return fmt.Errorf("ipv4Mask.Scan failed <- %v", err)
		}
		err = mac.Scan(itf.MAC)
		if err != nil {
			return fmt.Errorf("mac.Scan failed <- %v", err)
		}
		// interface type
		err = interfaceType.Scan("eth") // TODO
		if err != nil {
			return fmt.Errorf("interfaceType.Scan failed <- %v", err)
		}
		err = machine.Scan(machineId)
		if err != nil {
			return fmt.Errorf("machine.Scan failed <- %v", err)
		}

		networkId, err := doNetwork(itf.Ipv4, itf.Ipv4Mask, perimeterId)
		if err != nil {
			return fmt.Errorf("doNetwork failed <- %v", err)
		}

		err = network.Scan(networkId)
		if err != nil {
			return fmt.Errorf("network.Scan failed <- %v", err)
		}

		itfO := model.InterfaceO{Name: name,
			Ipv4:          ipv4,
			Ipv4Mask:      ipv4Mask,
			MAC:           mac,
			InterfaceType: interfaceType,
			MachineId:     machine,
			NetworkId:     network}

		interfaceId, err := net.InsertInterface(&itfO)
		if err != nil {
			return fmt.Errorf("InsertInterfnetworkace failed <- %v", err)
		}

		log.Debug(fmt.Sprintf("Create new interface %s %s : %d", itf.Name, itf.Ipv4, interfaceId))

		if len(itf.Gateways) != 0 {
			err = doGateways(interfaceId, itf.Gateways, itf.Ipv4Mask)
			if err != nil {
				return fmt.Errorf("doGateways failed <- %v", err)
			}
		}

	}
	return nil
}

func doNetworkInformations(networkInformation *client_informations.NetworkInformations, machineId int32, perimeterId int32) error {
	itfsDb, err := net.GetInterfacesByMachineId(machineId)

	if err != nil {
		return fmt.Errorf("net.GetInterfacesByMachineId failed <- %v", err)
	}

	if len(itfsDb) == 0 {
		err = newInterfaces(networkInformation.Interfaces, machineId, perimeterId)
		if err != nil {
			return fmt.Errorf("newInterfaces failed <- %v", err)
		}
	} else {
		return fmt.Errorf("Update not implemented <- %v", err) // TODO
	}

	return nil

}
