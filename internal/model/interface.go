package model

import (
	"encoding/json"
	"fmt"
)

type InterfaceO struct {
	Id            NullInt32  `json:"id"`
	Name          NullString `json:"name"`
	Ipv4          NullString `json:"ipv4"`
	Ipv4Mask      NullInt32  `json:"ipv4_mask"`
	MAC           NullString `json:"mac"`
	InterfaceType NullString `json:"interface_type"`
	MachineId     NullInt32  `json:"machine_id"`
	NetworkId     NullInt32  `json:"network_id"`
}

type InterfaceOs []InterfaceO

func (interfaceO *InterfaceO) String() string {
	return fmt.Sprintf("InterfaceO {%d, %s, %s, %d, %s, %d, %s, %s, %d, %d}",
		interfaceO.Id.Int32,
		interfaceO.Name.String,
		interfaceO.Ipv4.String,
		interfaceO.Ipv4Mask.Int32,
		interfaceO.MAC.String,
		interfaceO.InterfaceType.String,
		interfaceO.MachineId.Int32,
		interfaceO.NetworkId.Int32)
}

func (interfaceO *InterfaceO) New() Object {
	return new(InterfaceO)
}

func (interfaceO *InterfaceO) Valid() bool {
	return interfaceO.Name.Valid && interfaceO.Ipv4.Valid && interfaceO.Ipv4Mask.Valid && interfaceO.MachineId.Valid && interfaceO.NetworkId.Valid
}

func (interfaceO *InterfaceO) Json() ([]byte, error) {
	return json.Marshal(interfaceO)
}

func (interfaceOs InterfaceOs) Json() ([]byte, error) {
	return json.Marshal(interfaceOs)
}
