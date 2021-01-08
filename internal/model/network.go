package model

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
