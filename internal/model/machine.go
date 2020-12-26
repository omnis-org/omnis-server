package model

import (
	"encoding/json"
	"fmt"
)

type Machine struct {
	Id                   NullInt32  `json:"id"`
	Hostname             NullString `json:"hostname"`
	Label                NullString `json:"label"`
	Description          NullString `json:"description"`
	VirtualizationSystem NullString `json:"virtualization_system"`
	SerialNumber         NullString `json:"serial_number"`
	PerimeterId          NullInt32  `json:"perimeter_id"`
	LocationId           NullInt32  `json:"location_id"`
	OperatingSystemId    NullInt32  `json:"operating_system_id"`
	MachineType          NullString `json:"machine_type"`
	OmnisVersion         NullString `json:"omnis_version"`
}

type Machines []Machine

func (machine *Machine) String() string {
	return fmt.Sprintf("Machine {%d, %s, %s, %s, %s, %s, %d, %d, %d, %s, %s}",
		machine.Id.Int32,
		machine.Hostname.String,
		machine.Label.String,
		machine.Description.String,
		machine.VirtualizationSystem.String,
		machine.SerialNumber.String,
		machine.PerimeterId.Int32,
		machine.LocationId.Int32,
		machine.OperatingSystemId.Int32,
		machine.MachineType.String,
		machine.OmnisVersion.String)
}

func (machine *Machine) New() Object {
	return new(Machine)
}

func (machine *Machine) Valid() bool {
	return machine.Hostname.Valid && machine.Label.Valid && machine.PerimeterId.Valid && machine.LocationId.Valid
}

func (machine *Machine) Json() ([]byte, error) {
	return json.Marshal(machine)
}

func (machines Machines) Json() ([]byte, error) {
	return json.Marshal(machines)
}
