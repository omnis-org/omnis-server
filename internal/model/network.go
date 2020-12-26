package model

import (
	"encoding/json"
	"fmt"
)

type Network struct {
	Id          NullInt32  `json:"id"`
	Name        NullString `json:"name"`
	Ipv4        NullString `json:"ipv4"`
	Ipv4Mask    NullInt32  `json:"ipv4_mask"`
	IsDMZ       NullBool   `json:"is_dmz"`
	HasWifi     NullBool   `json:"has_wifi"`
	PerimeterId NullInt32  `json:"perimeter_id"`
}

type Networks []Network

func (network *Network) String() string {
	return fmt.Sprintf("Network {%d, %s, %s, %d, %t, %t, %d}",
		network.Id.Int32,
		network.Name.String,
		network.Ipv4.String,
		network.Ipv4Mask.Int32,
		network.IsDMZ.Bool,
		network.HasWifi.Bool,
		network.PerimeterId.Int32)
}

func (network *Network) New() Object {
	return new(Network)
}

func (network *Network) Valid() bool {
	return network.Name.Valid && network.Ipv4.Valid && network.Ipv4Mask.Valid && network.PerimeterId.Valid
}

func (network *Network) Json() ([]byte, error) {
	return json.Marshal(network)
}

func (networks Networks) Json() ([]byte, error) {
	return json.Marshal(networks)
}
