package model

// Gateway should have a comment.
type Gateway struct {
	ID          NullInt32  `json:"id"`
	Ipv4        NullString `json:"ipv4"`
	Mask        NullInt32  `json:"mask"`
	InterfaceID NullInt32  `json:"interfaceId"`
}

// Gateways should have a comment.
type Gateways []Gateway

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
