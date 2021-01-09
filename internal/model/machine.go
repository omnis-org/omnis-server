package model

// Machine should have a comment.
type Machine struct {
	ID                                   *NullInt32  `json:"id"`
	UUID                                 *NullString `json:"uuid"`
	Authorized                           *NullBool   `json:"authorized"`
	Hostname                             *NullString `json:"hostname"`
	Label                                *NullString `json:"label"`
	Description                          *NullString `json:"description"`
	VirtualizationSystem                 *NullString `json:"virtualizationSystem"`
	SerialNumber                         *NullString `json:"serialNumber"`
	MachineType                          *NullString `json:"machineType"`
	PerimeterID                          *NullInt32  `json:"perimeterId"`
	LocationID                           *NullInt32  `json:"locationId"`
	OperatingSystemID                    *NullInt32  `json:"operatingSystemId"`
	OmnisVersion                         *NullString `json:"omnisVersion"`
	UUIDLastModification                 *NullTime   `json:"uuidLastModification"`
	AuthorizedLastModification           *NullTime   `json:"authorizedLastModification"`
	HostnameLastModification             *NullTime   `json:"hostnameLastModification"`
	LabelLastModification                *NullTime   `json:"labelLastModification"`
	DescriptionLastModification          *NullTime   `json:"descriptionLastModification"`
	VirtualizationSystemLastModification *NullTime   `json:"virtualizationSystemLastModification"`
	SerialNumberLastModification         *NullTime   `json:"serialNumberLastModification"`
	MachineTypeLastModification          *NullTime   `json:"machineTypeLastModification"`
	PerimeterIDLastModification          *NullTime   `json:"perimeterIdLastModification"`
	LocationIDLastModification           *NullTime   `json:"locationIdLastModification"`
	OperatingSystemIDLastModification    *NullTime   `json:"operatingSystemIdLastModification"`
	OmnisVersionLastModification         *NullTime   `json:"omnisVersionLastModification"`
}

// Machines should have a comment.
type Machines []Machine

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
