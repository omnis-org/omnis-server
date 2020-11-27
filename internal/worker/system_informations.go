package worker

import (
	"fmt"

	"github.com/omnis-org/omnis-rest-api/pkg/model"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/internal/net"
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

	os := model.OperatingSystem{Name: name,
		Platform:        platform,
		PlatformFamily:  platformFamily,
		PlatformVersion: platformVersion,
		KernelVersion:   kernelVersion}

	id, err := net.InsertOperatingSystem(&os)
	if err != nil {
		return 0, fmt.Errorf("net.InsertOperatingSystem failed <- %v", err)
	}
	return id, nil
}

func doOperatingSystem(osInfos *client_informations.OperatingSystemInformations) (int32, error) {
	var operatingSystemID int32 = 0
	var err error

	if osInfos.OS != "" {
		operatingSystems, err := net.GetOperatingSystemsByName(osInfos.OS)
		if err != nil {
			return 0, fmt.Errorf("net.GetOperatingSystemsByName <- %v", err)
		}

		for _, os := range operatingSystems {
			if os.Platform.String == osInfos.Platform &&
				os.PlatformFamily.String == osInfos.PlatformFamily &&
				os.PlatformVersion.String == osInfos.PlatformVersion &&
				os.KernelVersion.String == osInfos.KernelVersion {
				operatingSystemID = os.Id.Int32
			}
		}
	}

	if operatingSystemID == 0 {
		// new
		operatingSystemID, err = newOperatingSystem(osInfos)
		if err != nil {
			return 0, fmt.Errorf("newOperatingSystem failed <- %v", err)
		}
	} else {
		// update
		return 0, fmt.Errorf("Update not implemented <- %v", err) // TODO
	}

	return operatingSystemID, nil
}

func newMachine(systemInformations *client_informations.SystemInformations, locationID int32, perimeterID int32, osID int32) (int32, error) {
	var hostname model.NullString
	var label model.NullString
	var isVirtualized model.NullBool
	var serialNumber model.NullString
	var perimeter model.NullInt32
	var location model.NullInt32
	var operatingSystem model.NullInt32
	var machineType model.NullString
	var omnisVersion model.NullString

	err := hostname.Scan(systemInformations.Hostname)
	if err != nil {
		return 0, fmt.Errorf("hostname.Scan failed <- %v", err)
	}
	err = label.Scan(systemInformations.Hostname)
	if err != nil {
		return 0, fmt.Errorf("label.Scan failed <- %v", err)
	}
	err = isVirtualized.Scan(systemInformations.VirtualizationInformations.IsVirtualized)
	if err != nil {
		return 0, fmt.Errorf("isVirtualized.Scan failed <- %v", err)
	}
	// TODO : Add virtualization system
	err = serialNumber.Scan(systemInformations.SerialNumber)
	if err != nil {
		return 0, fmt.Errorf("serialNumber.Scan failed <- %v", err)
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
	// TODO : MACHINE TYPE
	err = machineType.Scan("client")
	if err != nil {
		return 0, fmt.Errorf("machineType.Scan failed <- %v", err)
	}

	err = omnisVersion.Scan(systemInformations.OmnisVersion)
	if err != nil {
		return 0, fmt.Errorf("omnisVersion.Scan failed <- %v", err)
	}

	machine := model.Machine{Hostname: hostname,
		Label:             label,
		IsVirtualized:     isVirtualized,
		SerialNumber:      serialNumber,
		PerimeterId:       perimeter,
		LocationId:        location,
		OperatingSystemId: operatingSystem,
		MachineType:       machineType,
		OmnisVersion:      omnisVersion}

	machineID, err := net.InsertMachine(&machine)
	if err != nil {
		return 0, fmt.Errorf("net.InsertMachine failed <- %v", err)
	}

	return machineID, nil
}

func doSystemInformations(systemInformations *client_informations.SystemInformations, machineID int32, locationID int32, perimeterID int32) (int32, error) {
	osID, err := doOperatingSystem(systemInformations.OperatingSystem)

	if err != nil {
		return 0, err
	}

	if machineID == 0 {
		machineID, err = newMachine(systemInformations, locationID, perimeterID, osID)
		if err != nil {
			return 0, fmt.Errorf("newMachine failed <- %v", err)
		}
	} else {
		return 0, fmt.Errorf("Update not implemented <- %v", err) // TODO - Update
	}

	return machineID, nil
}
