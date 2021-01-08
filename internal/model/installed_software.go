package model

// InstalledSoftware should have a comment.
type InstalledSoftware struct {
	ID         NullInt32 `json:"id"`
	SoftwareID NullInt32 `json:"softwareId"`
	MachineID  NullInt32 `json:"machineId"`
}

// InstalledSoftwares should have a comment.
type InstalledSoftwares []InstalledSoftware

// New should have a comment.
func (installedSoftware *InstalledSoftware) New() Object {
	return new(InstalledSoftware)
}

// Valid should have a comment.
func (installedSoftware *InstalledSoftware) Valid() bool {
	if installedSoftware == nil {
		return false
	}
	return installedSoftware.SoftwareID.Valid && installedSoftware.MachineID.Valid
}
