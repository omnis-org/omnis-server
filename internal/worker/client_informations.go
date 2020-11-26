package worker

import (
	"fmt"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/internal/net"

	log "github.com/sirupsen/logrus"
)

// new informations
func new(infos *client_informations.Informations) error {
	locationId, perimeterId, err := doOtherInformations(infos.OtherInformations)

	if err != nil {
		return fmt.Errorf("doOtherInformations failed <- %v", err)
	}

	machineId, err := doSystemInformations(infos.SystemInformations, 0, locationId, perimeterId)

	if err != nil {
		return fmt.Errorf("doSystemInformations failed <- %v", err)
	}

	err = doNetworkInformations(infos.NetworkInformations, machineId, perimeterId)

	if err != nil {
		return fmt.Errorf("doNetworkInformations failed <- %v", err)
	}

	return nil
}

// update informations
func update(machineId int32, infos *client_informations.Informations) error {

	locationId, perimeterId, err := doOtherInformations(infos.OtherInformations)

	if err != nil {
		return fmt.Errorf("doNetworkInformations failed <- %v", err)
	}

	machineId, err = doSystemInformations(infos.SystemInformations, machineId, locationId, perimeterId)
	if err != nil {
		return fmt.Errorf("doSystemInformations failed <- %v", err)
	}

	err = doNetworkInformations(infos.NetworkInformations, machineId, perimeterId)

	if err != nil {
		return fmt.Errorf("doNetworkInformations failed <- %v", err)
	}

	return nil
}

func AnalyzeClientInformations(i interface{}) {
	infos := i.(*client_informations.Informations)

	mac, err := getOneMacInterfaces(infos)
	if err != nil {
		log.Error(err)
		return
	}

	itf, err := net.GetInterfaceByMac(mac)
	if err != nil {
		log.Error(err)
		return
	}

	if !itf.Id.Valid {
		err = new(infos)
		if err != nil {
			log.Error(err)
		}
		return
	}

	err = update(itf.MachineId.Int32, infos)
	if err != nil {
		log.Error(err)
		return
	}

}
