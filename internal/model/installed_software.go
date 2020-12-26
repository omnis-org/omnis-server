package model

import (
	"encoding/json"
	"fmt"
)

type InstalledSoftware struct {
	Id         NullInt32 `json:"id"`
	SoftwareId NullInt32 `json:"software_id"`
	MachineId  NullInt32 `json:"machine_id"`
}

type InstalledSoftwares []InstalledSoftware

func (installedSoftware *InstalledSoftware) String() string {
	return fmt.Sprintf("InstalledSoftware {%d, %d, %d}",
		installedSoftware.Id.Int32,
		installedSoftware.SoftwareId.Int32,
		installedSoftware.MachineId.Int32)
}

func (installedSoftware *InstalledSoftware) New() Object {
	return new(InstalledSoftware)
}

func (installedSoftware *InstalledSoftware) Valid() bool {
	return installedSoftware.SoftwareId.Valid && installedSoftware.MachineId.Valid
}

func (installedSoftware *InstalledSoftware) Json() ([]byte, error) {
	return json.Marshal(installedSoftware)
}

func (installedSoftwares InstalledSoftwares) Json() ([]byte, error) {
	return json.Marshal(installedSoftwares)
}
