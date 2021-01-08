package model

// OperatingSystem should have a comment.
type OperatingSystem struct {
	ID              NullInt32  `json:"id"`
	Name            NullString `json:"name"`
	Platform        NullString `json:"platform"`
	PlatformFamily  NullString `json:"platformFamily"`
	PlatformVersion NullString `json:"platformVersion"`
	KernelVersion   NullString `json:"kernelVersion"`
}

// OperatingSystems should have a comment.
type OperatingSystems []OperatingSystem

// New should have a comment.
func (operatingSystem *OperatingSystem) New() Object {
	return new(OperatingSystem)
}

// Valid should have a comment.
func (operatingSystem *OperatingSystem) Valid() bool {
	if operatingSystem == nil {
		return false
	}
	return operatingSystem.Name.Valid
}
