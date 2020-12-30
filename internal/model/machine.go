package model

import (
	"encoding/json"
	"fmt"
)

// Machine should have a comment.
type Machine struct {
	ID                   NullInt32  `json:"id"`
	UUID                 NullString `json:"uuid"`
	Authorized           NullBool   `json:"authorized"`
	Hostname             NullString `json:"hostname"`
	Label                NullString `json:"label"`
	Description          NullString `json:"description"`
	VirtualizationSystem NullString `json:"virtualizationSystem"`
	SerialNumber         NullString `json:"serialNumber"`
	PerimeterID          NullInt32  `json:"perimeterId"`
	LocationID           NullInt32  `json:"locationId"`
	OperatingSystemID    NullInt32  `json:"operatingSystemId"`
	MachineType          NullString `json:"machineType"`
	OmnisVersion         NullString `json:"omnisVersion"`
}

// Machines should have a comment.
type Machines []Machine

// String should have a comment.
func (machine *Machine) String() string {
	return fmt.Sprintf("Machine {%d, %s, %t, %s, %s, %s, %s, %d, %d, %d, %s, %s}",
		machine.ID.Int32,
		machine.UUID.String,
		machine.Authorized.Bool,
		machine.Hostname.String,
		machine.Label.String,
		machine.Description.String,
		machine.VirtualizationSystem.String,
		machine.SerialNumber.String,
		machine.PerimeterID.Int32,
		machine.LocationID.Int32,
		machine.OperatingSystemID.Int32,
		machine.MachineType.String,
		machine.OmnisVersion.String)
}

// New should have a comment.
func (machine *Machine) New() Object {
	return new(Machine)
}

// Valid should have a comment.
func (machine *Machine) Valid() bool {
	if machine == nil {
		return false
	}
	return machine.Hostname.Valid && machine.Label.Valid && machine.PerimeterID.Valid && machine.LocationID.Valid
}

// JSON should have a comment.
func (machine *Machine) JSON() ([]byte, error) {
	return json.Marshal(machine)
}

// JSON should have a comment.
func (machines Machines) JSON() ([]byte, error) {
	return json.Marshal(machines)
}
