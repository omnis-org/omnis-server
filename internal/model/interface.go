package model

import (
	"fmt"
)

// InterfaceO should have a comment.
type InterfaceO struct {
	ID            NullInt32  `json:"id"`
	Name          NullString `json:"name"`
	Ipv4          NullString `json:"ipv4"`
	Ipv4Mask      NullInt32  `json:"ipv4Mask"`
	MAC           NullString `json:"mac"`
	InterfaceType NullString `json:"interfaceType"`
	MachineID     NullInt32  `json:"machineId"`
	NetworkID     NullInt32  `json:"networkId"`
}

// InterfaceOs should have a comment.
type InterfaceOs []InterfaceO

// String should have a comment.
func (interfaceO *InterfaceO) String() string {
	return fmt.Sprintf("InterfaceO {%d, %s, %s, %d, %s, %s, %d, %d}",
		interfaceO.ID.Int32,
		interfaceO.Name.String,
		interfaceO.Ipv4.String,
		interfaceO.Ipv4Mask.Int32,
		interfaceO.MAC.String,
		interfaceO.InterfaceType.String,
		interfaceO.MachineID.Int32,
		interfaceO.NetworkID.Int32)
}

// New should have a comment.
func (interfaceO *InterfaceO) New() Object {
	return new(InterfaceO)
}

// Valid should have a comment.
func (interfaceO *InterfaceO) Valid() bool {
	if interfaceO == nil {
		return false
	}
	return interfaceO.Name.Valid && interfaceO.Ipv4.Valid && interfaceO.Ipv4Mask.Valid && interfaceO.MachineID.Valid && interfaceO.NetworkID.Valid
}
