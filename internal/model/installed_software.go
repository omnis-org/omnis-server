package model

import (
	"encoding/json"
	"fmt"
)

// InstalledSoftware should have a comment.
type InstalledSoftware struct {
	ID         NullInt32 `json:"id"`
	SoftwareID NullInt32 `json:"softwareId"`
	MachineID  NullInt32 `json:"machineId"`
}

// InstalledSoftwares should have a comment.
type InstalledSoftwares []InstalledSoftware

// String should have a comment.
func (installedSoftware *InstalledSoftware) String() string {
	return fmt.Sprintf("InstalledSoftware {%d, %d, %d}",
		installedSoftware.ID.Int32,
		installedSoftware.SoftwareID.Int32,
		installedSoftware.MachineID.Int32)
}

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

// JSON should have a comment.
func (installedSoftware *InstalledSoftware) JSON() ([]byte, error) {
	return json.Marshal(installedSoftware)
}

// JSON should have a comment.
func (installedSoftwares InstalledSoftwares) JSON() ([]byte, error) {
	return json.Marshal(installedSoftwares)
}
