package client

import (
	"fmt"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/internal/db"

	log "github.com/sirupsen/logrus"
)

// TODO Refactor name
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

func AnalyzeClientInformations(i interface{}) error {
	infos := i.(*client_informations.Informations)

	log.Info(fmt.Sprintf("get new informations from : %s", infos.SystemInformations.Hostname))
	log.Info(infos)
	var machineID int32 = 0

	machine, err := db.GetMachineByUUID(infos.SystemInformations.UUID, true)
	if err != nil {
		return fmt.Errorf("db.GetMachineByUUID failed <- %v", err)
	}

	if machine != nil {
		machineID = machine.ID.Int32
	}

	log.Debug(fmt.Sprintf("machine id : %d", machineID))

	err = doClientInformations(infos, machineID)
	if err != nil {
		return fmt.Errorf("doClientInformations failed <- %v", err)
	}

	return nil
}
