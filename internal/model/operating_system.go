package model

import (
	"encoding/json"
	"fmt"
)

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

// String should have a comment.
func (operatingSystem *OperatingSystem) String() string {
	return fmt.Sprintf("OperatingSystem {%d, %s, %s, %s, %s, %s}",
		operatingSystem.ID.Int32,
		operatingSystem.Name.String,
		operatingSystem.Platform.String,
		operatingSystem.PlatformFamily.String,
		operatingSystem.PlatformVersion.String,
		operatingSystem.KernelVersion.String)
}

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

// JSON should have a comment.
func (operatingSystem *OperatingSystem) JSON() ([]byte, error) {
	return json.Marshal(operatingSystem)
}

// JSON should have a comment.
func (operatingSystems OperatingSystems) JSON() ([]byte, error) {
	return json.Marshal(operatingSystems)
}
