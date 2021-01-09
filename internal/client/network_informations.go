package client

import (
	"fmt"

	"github.com/omnis-org/omnis-server/internal/db"
	"github.com/omnis-org/omnis-server/internal/model"
	"github.com/omnis-org/omnis-server/internal/utils"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	log "github.com/sirupsen/logrus"
)

func newNetwork(networkPart string, mask int, perimeterID int32) (int32, error) {
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
	err = perimeter.Scan(perimeterID)
	if err != nil {
		return 0, fmt.Errorf("perimeter.Scan failed <- %v", err)
	}
	networkO := model.Network{Name: &name,
		Ipv4:        &ipv4,
		Ipv4Mask:    &ipv4Mask,
		PerimeterID: &perimeter,
	}
	networkID, err := db.InsertNetwork(&networkO, true)
	if err != nil {
		return 0, fmt.Errorf("net.InsertNetwork failed <- %v", err)
	}
	return networkID, nil
}

func doNetwork(ip string, mask int, perimeterID int32) (int32, error) {
	var networkID int32 = 0
	networkPart := utils.GetNetworkPart(ip, mask)

	networks, err := db.GetNetworksByIP(networkPart, true)
	if err != nil {
		return 0, fmt.Errorf("net.GetNetworksByIp failed <- %v", err)
	}

	for _, network := range networks {
		if network.Ipv4.String == networkPart &&
			network.Ipv4Mask.Int32 == int32(mask) &&
			network.PerimeterID.Int32 == perimeterID {
			networkID = network.ID.Int32
		}
	}

	if networkID == 0 {
		networkID, err = newNetwork(networkPart, mask, perimeterID)
		if err != nil {
			return 0, fmt.Errorf("newNetwork failed <- %v", err)
		}
	}

	return networkID, nil
}

func newGateway(ip string, maskI int, interfaceID int32) (int32, error) {
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
	err = interfaceO.Scan(interfaceID)
	if err != nil {
		return 0, fmt.Errorf("interfaceO.Scan failed <- %v", err)
	}

	gateway := model.Gateway{Ipv4: &ipv4,
		Mask:        &mask,
		InterfaceID: &interfaceO}

	gatewayID, err := db.InsertGateway(&gateway, true)
	if err != nil {
		return 0, fmt.Errorf("net.InsertGateway failed <- %v", err)
	}
	return gatewayID, nil
}

func doGateways(interfaceID int32, gateways []string, mask int) error {
	oldGateways, err := db.GetGatewaysByInterfaceID(interfaceID, true)
	if err != nil {
		return fmt.Errorf("net.GetGatewaysByInterfaceId failed <- %v", err)
	}

	for _, gtw := range gateways {
		var gatewayID int32 = 0
		for _, oldGtw := range oldGateways {
			if oldGtw.Ipv4.String == gtw && oldGtw.Mask.Int32 == int32(mask) {
				gatewayID = oldGtw.ID.Int32
				break
			}
		}
		if gatewayID == 0 {
			_, err = newGateway(gtw, mask, interfaceID)
			if err != nil {
				return fmt.Errorf("newGateway failed <- %v", err)
			}
		}
	}

	for _, oldGtw := range oldGateways {
		var found bool = false
		for _, gtw := range gateways {
			if oldGtw.Ipv4.String == gtw && oldGtw.Mask.Int32 == int32(mask) {
				found = true
				break
			}
		}

		if !found {
			rowsAffected, err := db.DeleteGateway(oldGtw.ID.Int32)
			if err != nil {
				return fmt.Errorf("net.DeleteGateway failed <- %v", err)
			}

			if rowsAffected != 0 {
				return fmt.Errorf("net.DeleteGateway no rows affected <- %v", err)
			}

			log.Debug(fmt.Sprintf("delete gateway : %d", oldGtw.ID.Int32))
		}

	}

	return nil
}

func doInterface(itf *client_informations.InterfaceInformations, machineID int32, perimeterID int32, updateInterfaceID int32) error {
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
	err = machine.Scan(machineID)
	if err != nil {
		return fmt.Errorf("machine.Scan failed <- %v", err)
	}

	networkID, err := doNetwork(itf.Ipv4, itf.Ipv4Mask, perimeterID)
	if err != nil {
		return fmt.Errorf("doNetwork failed <- %v", err)
	}

	err = network.Scan(networkID)
	if err != nil {
		return fmt.Errorf("network.Scan failed <- %v", err)
	}

	itfO := model.InterfaceO{Name: &name,
		Ipv4:          &ipv4,
		Ipv4Mask:      &ipv4Mask,
		MAC:           &mac,
		InterfaceType: &interfaceType,
		MachineID:     &machine,
		NetworkID:     &network}

	var itfID int32 = 0

	if updateInterfaceID == 0 {
		itfID, err = db.InsertInterface(&itfO, true)
		if err != nil {
			return fmt.Errorf("net.InsertInterface failed <- %v", err)
		}
		log.Debug(fmt.Sprintf("new interface %s %s : %d", itf.Name, itf.Ipv4, itfID))
	} else {
		itfID = updateInterfaceID
		_, err := db.UpdateInterface(updateInterfaceID, &itfO, true)
		if err != nil {
			return fmt.Errorf("net.UpdateInterface failed <- %v", err)
		}

		log.Debug(fmt.Sprintf("update interface %s %s : %d", itf.Name, itf.Ipv4, itfID))
	}

	if len(itf.Gateways) != 0 {
		err = doGateways(itfID, itf.Gateways, itf.Ipv4Mask)
		if err != nil {
			return fmt.Errorf("doGateways failed <- %v", err)
		}
	}

	return nil
}

func doInterfaces(interfaces []client_informations.InterfaceInformations, machineID int32, perimeterID int32) error {
	oldInterfaces, err := db.GetInterfacesByMachineID(machineID, true)

	if err != nil {
		return fmt.Errorf("net.GetInterfacesByMachineId failed <- %v", err)
	}

	for _, itf := range interfaces {

		var updateInterface model.InterfaceO

		for _, oldInterface := range oldInterfaces {
			if oldInterface.MAC.String == itf.MAC {
				updateInterface = oldInterface
				break
			}
		}

		err = doInterface(&itf, machineID, perimeterID, updateInterface.ID.Int32)
		if err != nil {
			return fmt.Errorf("doInterface failed <- %v", err)
		}
	}

	for _, oldInterface := range oldInterfaces {
		var found bool = false
		for _, itf := range interfaces {

			if oldInterface.MAC.String == itf.MAC {
				found = true
				break
			}
		}

		if !found {
			rowsAffected, err := db.DeleteInterface(oldInterface.ID.Int32)
			if err != nil {
				return fmt.Errorf("db.DeleteInterface failed <- %v", err)
			}
			if rowsAffected == 0 {
				return fmt.Errorf("db.DeleteInterface no rows affected <- %v", err)
			}

			log.Debug(fmt.Sprintf("delete interface : %d", oldInterface.ID.Int32))
		}
	}
	return nil
}

func doNetworkInformations(networkInformation *client_informations.NetworkInformations, machineID int32, perimeterID int32) error {
	err := doInterfaces(networkInformation.Interfaces, machineID, perimeterID)
	if err != nil {
		return fmt.Errorf("doInterfaces failed <- %v", err)
	}
	return nil
}
