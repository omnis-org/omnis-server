package model

// Machine should have a comment.
type Machine struct {
	ID                                   NullInt32  `json:"id"`
	UUID                                 NullString `json:"uuid"`
	Authorized                           NullBool   `json:"authorized"`
	Hostname                             NullString `json:"hostname"`
	Label                                NullString `json:"label"`
	Description                          NullString `json:"description"`
	VirtualizationSystem                 NullString `json:"virtualizationSystem"`
	SerialNumber                         NullString `json:"serialNumber"`
	MachineType                          NullString `json:"machineType"`
	PerimeterID                          NullInt32  `json:"perimeterId"`
	LocationID                           NullInt32  `json:"locationId"`
	OperatingSystemID                    NullInt32  `json:"operatingSystemId"`
	OmnisVersion                         NullString `json:"omnisVersion"`
	UUIDLastModification                 NullTime   `json:"uuidLastModification,ommitempty"`
	AuthorizedLastModification           NullTime   `json:"authorizedLastModification,ommitempty"`
	HostnameLastModification             NullTime   `json:"hostnameLastModification,ommitempty"`
	LabelLastModification                NullTime   `json:"labelLastModification,ommitempty"`
	DescriptionLastModification          NullTime   `json:"descriptionLastModification,ommitempty"`
	VirtualizationSystemLastModification NullTime   `json:"virtualizationSystemLastModification,ommitempty"`
	SerialNumberLastModification         NullTime   `json:"serialNumberLastModification,ommitempty"`
	MachineTypeLastModification          NullTime   `json:"machineTypeLastModification,ommitempty"`
	PerimeterIDLastModification          NullTime   `json:"perimeterIdLastModification,ommitempty"`
	LocationIDLastModification           NullTime   `json:"locationIdLastModification,ommitempty"`
	OperatingSystemIDLastModification    NullTime   `json:"operatingSystemIdLastModification,ommitempty"`
	OmnisVersionLastModification         NullTime   `json:"omnisVersionLastModification,ommitempty"`
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
