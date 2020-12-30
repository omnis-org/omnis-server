package model

import (
	"encoding/json"
	"fmt"
)

// Gateway should have a comment.
type Gateway struct {
	ID          NullInt32  `json:"id"`
	Ipv4        NullString `json:"ipv4"`
	Mask        NullInt32  `json:"mask"`
	InterfaceID NullInt32  `json:"interfaceId"`
}

// Gateways should have a comment.
type Gateways []Gateway

// String should have a comment.
func (gateway *Gateway) String() string {
	return fmt.Sprintf("Gateway {%d, %s, %d, %d}",
		gateway.ID.Int32,
		gateway.Ipv4.String,
		gateway.Mask.Int32,
		gateway.InterfaceID.Int32)
}

// New should have a comment.
func (gateway *Gateway) New() Object {
	return new(Gateway)
}

// Valid should have a comment.
func (gateway *Gateway) Valid() bool {
	if gateway == nil {
		return false
	}
	return gateway.Ipv4.Valid && gateway.Mask.Valid && gateway.InterfaceID.Valid
}

// JSON should have a comment.
func (gateway *Gateway) JSON() ([]byte, error) {
	return json.Marshal(gateway)
}

// JSON should have a comment.
func (gateways Gateways) JSON() ([]byte, error) {
	return json.Marshal(gateways)
}
