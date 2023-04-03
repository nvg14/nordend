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

// UpdatePlatformConfigQueryParameters is an object.
type UpdatePlatformConfigQueryParameters struct {
	// Platformid:
	Platformid float32 `json:"platformid" mapstructure:"platformid"`
}

// Validate implements basic validation for this model
func (m UpdatePlatformConfigQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPlatformid returns the Platformid property
func (m UpdatePlatformConfigQueryParameters) GetPlatformid() float32 {
	return m.Platformid
}

// SetPlatformid sets the Platformid property
func (m *UpdatePlatformConfigQueryParameters) SetPlatformid(val float32) {
	m.Platformid = val
}
