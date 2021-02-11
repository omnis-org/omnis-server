package client

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/internal/db"
	"github.com/omnis-org/omnis-server/internal/model"
)

func newOperatingSystem(osInfos *client_informations.OperatingSystemInformations) (int32, error) {
	var name model.NullString
	var platform model.NullString
	var platformFamily model.NullString
	var platformVersion model.NullString
	var kernelVersion model.NullString

	err := name.Scan(osInfos.OS)
	if err != nil {
		return 0, fmt.Errorf("name.Scan failed <- %v", err)
	}
	err = platform.Scan(osInfos.Platform)
	if err != nil {
		return 0, fmt.Errorf("platform.Scan failed <- %v", err)
	}
	err = platformFamily.Scan(osInfos.PlatformFamily)
	if err != nil {
		return 0, fmt.Errorf("platformFamily.Scan failed <- %v", err)
	}
	err = platformVersion.Scan(osInfos.PlatformVersion)
	if err != nil {
		return 0, fmt.Errorf("platformVersion.Scan failed <- %v", err)
	}
	err = kernelVersion.Scan(osInfos.KernelVersion)
	if err != nil {
		return 0, fmt.Errorf("kernelVersion.Scan failed <- %v", err)
	}

	os := model.OperatingSystem{Name: &name,
		Platform:        &platform,
		PlatformFamily:  &platformFamily,
		PlatformVersion: &platformVersion,
		KernelVersion:   &kernelVersion}

	id, err := db.InsertOperatingSystem(&os, true)
	if err != nil {
		return 0, fmt.Errorf("net.InsertOperatingSystem failed <- %v", err)
	}
	return id, nil
}

func doOperatingSystem(osInfos *client_informations.OperatingSystemInformations) (int32, error) {
	var operatingSystemID int32 = 0
	var err error

	operatingSystems, err := db.GetOperatingSystemsByName(osInfos.OS, true)
	if err != nil {
		return 0, fmt.Errorf("net.GetOperatingSystemsByName <- %v", err)
	}

	for _, os := range operatingSystems {
		if os.Platform.String == osInfos.Platform &&
			os.PlatformFamily.String == osInfos.PlatformFamily &&
			os.PlatformVersion.String == osInfos.PlatformVersion &&
			os.KernelVersion.String == osInfos.KernelVersion {
			operatingSystemID = os.ID.Int32
		}
	}

	if operatingSystemID == 0 {
		operatingSystemID, err = newOperatingSystem(osInfos)
		if err != nil {
			return 0, fmt.Errorf("newOperatingSystem failed <- %v", err)
		}

		log.Debug(fmt.Sprintf("new os : %s %s %s %s", osInfos.Platform, osInfos.PlatformFamily, osInfos.PlatformVersion, osInfos.KernelVersion))
	}
	return operatingSystemID, nil
}

func doMachine(systemInformations *client_informations.SystemInformations, locationID int32, perimeterID int32, osID int32, updateMachineID int32) (int32, error) {
	var hostname model.NullString
	var label model.NullString
	var virtualizationSystem model.NullString
	var serialNumber model.NullString
	var machineType model.NullString
	var perimeter model.NullInt32
	var location model.NullInt32
	var operatingSystem model.NullInt32
	var omnisVersion model.NullString
	var uuid model.NullString

	err := hostname.Scan(systemInformations.Hostname)
	if err != nil {
		return 0, fmt.Errorf("hostname.Scan failed <- %v", err)
	}

	err = uuid.Scan(systemInformations.UUID)
	if err != nil {
		return 0, fmt.Errorf("uuid.Scan failed <- %v", err)
	}

	err = label.Scan(systemInformations.Hostname)
	if err != nil {
		return 0, fmt.Errorf("label.Scan failed <- %v", err)
	}

	if !systemInformations.VirtualizationInformations.IsVirtualized {
		err = virtualizationSystem.Scan(systemInformations.VirtualizationInformations.VirtualizationSystem)
		if err != nil {
			return 0, fmt.Errorf("virtualizationSystem.Scan failed <- %v", err)
		}
	}

	err = serialNumber.Scan(systemInformations.SerialNumber)
	if err != nil {
		return 0, fmt.Errorf("serialNumber.Scan failed <- %v", err)
	}

	if strings.Contains(strings.ToLower(systemInformations.OperatingSystem.PlatformVersion), "server") ||
		strings.Contains(strings.ToLower(systemInformations.OperatingSystem.PlatformFamily), "server") { // TO DO : check work on all server type
		err = machineType.Scan("server")
	} else {
		err = machineType.Scan("client")
	}

	if err != nil {
		return 0, fmt.Errorf("machineType.Scan failed <- %v", err)
	}

	err = perimeter.Scan(perimeterID)
	if err != nil {
		return 0, fmt.Errorf("perimeter.Scan failed <- %v", err)
	}
	err = location.Scan(locationID)
	if err != nil {
		return 0, fmt.Errorf("location.Scan failed <- %v", err)
	}
	err = operatingSystem.Scan(osID)
	if err != nil {
		return 0, fmt.Errorf("operatingSystem.Scan failed <- %v", err)
	}

	err = omnisVersion.Scan(systemInformations.OmnisVersion)
	if err != nil {
		return 0, fmt.Errorf("omnisVersion.Scan failed <- %v", err)
	}

	machine := model.Machine{Hostname: &hostname,
		UUID:                 &uuid,
		Label:                &label,
		VirtualizationSystem: &virtualizationSystem,
		SerialNumber:         &serialNumber,
		MachineType:          &machineType,
		PerimeterID:          &perimeter,
		LocationID:           &location,
		OperatingSystemID:    &operatingSystem,
		OmnisVersion:         &omnisVersion}

	var machineID int32 = 0
	if updateMachineID == 0 { // new
		machineID, err = db.InsertMachine(&machine, true)
		if err != nil {
			return 0, fmt.Errorf("net.InsertMachine failed <- %v", err)
		}

		log.Debug(fmt.Sprintf("new machine : %s", systemInformations.Hostname))

	} else { // update
		machineID = updateMachineID
		_, err := db.UpdateMachine(machineID, &machine, true)
		if err != nil {
			return 0, fmt.Errorf("db.UpdateMachine failed <- %v", err)
		}

		log.Debug(fmt.Sprintf("update machine : %s", systemInformations.Hostname))
	}

	return machineID, nil
}

func doSystemInformations(systemInformations *client_informations.SystemInformations, locationID int32, perimeterID int32, updateMachineID int32) (int32, error) {
	var machineID int32 = 0

	osID, err := doOperatingSystem(systemInformations.OperatingSystem)
	if err != nil {
		return 0, fmt.Errorf("doOperatingSystem failed <- %v", err)
	}

	log.Debug(fmt.Sprintf("os id : %d", osID))

	machineID, err = doMachine(systemInformations, locationID, perimeterID, osID, updateMachineID)
	if err != nil {
		return 0, fmt.Errorf("doMachine failed <- %v", err)
	}

	return machineID, nil
}
