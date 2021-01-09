package model

// OperatingSystem should have a comment.
type OperatingSystem struct {
	ID                              *NullInt32  `json:"id"`
	Name                            *NullString `json:"name"`
	Platform                        *NullString `json:"platform"`
	PlatformFamily                  *NullString `json:"platformFamily"`
	PlatformVersion                 *NullString `json:"platformVersion"`
	KernelVersion                   *NullString `json:"kernelVersion"`
	NameLastModification            *NullTime   `json:"nameLastModification"`
	PlatformLastModification        *NullTime   `json:"platformLastModification"`
	PlatformFamilyLastModification  *NullTime   `json:"platformFamilyLastModification"`
	PlatformVersionLastModification *NullTime   `json:"platformVersionLastModification"`
	KernelVersionLastModification   *NullTime   `json:"kernelVersionLastModification"`
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
