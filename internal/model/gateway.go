package model

import (
	"encoding/json"
	"fmt"
)

type Gateway struct {
	Id          NullInt32  `json:"id"`
	Ipv4        NullString `json:"ipv4"`
	Mask        NullInt32  `json:"mask"`
	InterfaceId NullInt32  `json:"interface_id"`
}

type Gateways []Gateway

func (gateway *Gateway) String() string {
	return fmt.Sprintf("Gateway {%d, %s, %d, %d}",
		gateway.Id.Int32,
		gateway.Ipv4.String,
		gateway.Mask.Int32,
		gateway.InterfaceId.Int32)
}

func (gateway *Gateway) New() Object {
	return new(Gateway)
}

func (gateway *Gateway) Valid() bool {
	return gateway.Ipv4.Valid && gateway.Mask.Valid && gateway.InterfaceId.Valid
}

func (gateway *Gateway) Json() ([]byte, error) {
	return json.Marshal(gateway)
}

func (gateways Gateways) Json() ([]byte, error) {
	return json.Marshal(gateways)
}
