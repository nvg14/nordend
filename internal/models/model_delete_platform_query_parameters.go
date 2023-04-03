// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Nordend
//	Version: 1.0.0
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// DeletePlatformQueryParameters is an object.
type DeletePlatformQueryParameters struct {
	// Platformid:
	Platformid string `json:"platformid" mapstructure:"platformid"`
}

// Validate implements basic validation for this model
func (m DeletePlatformQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPlatformid returns the Platformid property
func (m DeletePlatformQueryParameters) GetPlatformid() string {
	return m.Platformid
}

// SetPlatformid sets the Platformid property
func (m *DeletePlatformQueryParameters) SetPlatformid(val string) {
	m.Platformid = val
}
