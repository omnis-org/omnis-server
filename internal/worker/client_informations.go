package worker

import (
	"fmt"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/internal/net"

	log "github.com/sirupsen/logrus"
)

func doClientInformations(infos *client_informations.Informations, updateMachineID int32) error {
	locationID, perimeterID, err := doOtherInformations(infos.OtherInformations)
	if err != nil {
		return fmt.Errorf("doOtherInformations failed <- %v", err)
	}

	log.Debug(fmt.Sprintf("location id : %d", locationID))
	log.Debug(fmt.Sprintf("perimeter id : %d", perimeterID))

	machineID, err := doSystemInformations(infos.SystemInformations, locationID, perimeterID, updateMachineID)
	if err != nil {
		return fmt.Errorf("doSystemInformations failed <- %v", err)
	}

	log.Debug(fmt.Sprintf("machine id : %d", machineID))

	err = doNetworkInformations(infos.NetworkInformations, machineID, perimeterID)
	if err != nil {
		return fmt.Errorf("doNetworkInformations failed <- %v", err)
	}

	return nil
}

func AnalyzeClientInformations(i interface{}) {
	infos := i.(*client_informations.Informations)

	log.Info(fmt.Sprintf("get new informations from : %s", infos.SystemInformations.Hostname))

	mac, err := getOneMacInterfaces(infos)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug(fmt.Sprintf("machine as interface with mac : %s", mac))

	itf, err := net.GetInterfaceByMac(mac)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug(fmt.Sprintf("interface id : %d", itf.Id.Int32))

	err = doClientInformations(infos, itf.MachineId.Int32)
	if err != nil {
		log.Error(err)
		return
	}

}
