package model

import (
	"encoding/json"
	"fmt"
)

type OperatingSystem struct {
	Id              NullInt32  `json:"id"`
	Name            NullString `json:"name"`
	Platform        NullString `json:"platform"`
	PlatformFamily  NullString `json:"platform_family"`
	PlatformVersion NullString `json:"platform_version"`
	KernelVersion   NullString `json:"kernel_version"`
}

type OperatingSystems []OperatingSystem

func (operatingSystem *OperatingSystem) String() string {
	return fmt.Sprintf("OperatingSystem {%d, %s, %s, %s, %s, %s}",
		operatingSystem.Id.Int32,
		operatingSystem.Name.String,
		operatingSystem.Platform.String,
		operatingSystem.PlatformFamily.String,
		operatingSystem.PlatformVersion.String,
		operatingSystem.KernelVersion.String)
}

func (operatingSystem *OperatingSystem) New() Object {
	return new(OperatingSystem)
}

func (operatingSystem *OperatingSystem) Valid() bool {
	return operatingSystem.Name.Valid
}

func (operatingSystem *OperatingSystem) Json() ([]byte, error) {
	return json.Marshal(operatingSystem)
}

func (operatingSystems OperatingSystems) Json() ([]byte, error) {
	return json.Marshal(operatingSystems)
}
