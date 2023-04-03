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

// GetSpecificPlatformQueryParameters is an object.
type GetSpecificPlatformQueryParameters struct {
	// Platformid:
	Platformid float32 `json:"platformid" mapstructure:"platformid"`
}

// Validate implements basic validation for this model
func (m GetSpecificPlatformQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPlatformid returns the Platformid property
func (m GetSpecificPlatformQueryParameters) GetPlatformid() float32 {
	return m.Platformid
}

// SetPlatformid sets the Platformid property
func (m *GetSpecificPlatformQueryParameters) SetPlatformid(val float32) {
	m.Platformid = val
}
