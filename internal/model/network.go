package model

import (
	"encoding/json"
	"fmt"
)

// Network should have a comment.
type Network struct {
	ID          NullInt32  `json:"id"`
	Name        NullString `json:"name"`
	Ipv4        NullString `json:"ipv4"`
	Ipv4Mask    NullInt32  `json:"ipv4Mask"`
	IsDMZ       NullBool   `json:"isDmz"`
	HasWifi     NullBool   `json:"hasWifi"`
	PerimeterID NullInt32  `json:"perimeterId"`
}

// Networks should have a comment.
type Networks []Network

// String should have a comment.
func (network *Network) String() string {
	return fmt.Sprintf("Network {%d, %s, %s, %d, %t, %t, %d}",
		network.ID.Int32,
		network.Name.String,
		network.Ipv4.String,
		network.Ipv4Mask.Int32,
		network.IsDMZ.Bool,
		network.HasWifi.Bool,
		network.PerimeterID.Int32)
}

// New should have a comment.
func (network *Network) New() Object {
	return new(Network)
}

// Valid should have a comment.
func (network *Network) Valid() bool {
	if network == nil {
		return false
	}
	return network.Name.Valid && network.Ipv4.Valid && network.Ipv4Mask.Valid && network.PerimeterID.Valid
}

// JSON should have a comment.
func (network *Network) JSON() ([]byte, error) {
	return json.Marshal(network)
}

// JSON should have a comment.
func (networks Networks) JSON() ([]byte, error) {
	return json.Marshal(networks)
}
