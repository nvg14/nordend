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

// UpdateDeliveryConfigQueryParameters is an object.
type UpdateDeliveryConfigQueryParameters struct {
	// Amgid: Filter list for a given amgid
	Amgid string `json:"amgid" mapstructure:"amgid"`
	// Platformid: Filter list for a given platformid
	Platformid string `json:"platformid" mapstructure:"platformid"`
	// Channelid: Filter list for a given channelid
	Channelid string `json:"channelid,omitempty" mapstructure:"channelid,omitempty"`
}

// Validate implements basic validation for this model
func (m UpdateDeliveryConfigQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAmgid returns the Amgid property
func (m UpdateDeliveryConfigQueryParameters) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *UpdateDeliveryConfigQueryParameters) SetAmgid(val string) {
	m.Amgid = val
}

// GetPlatformid returns the Platformid property
func (m UpdateDeliveryConfigQueryParameters) GetPlatformid() string {
	return m.Platformid
}

// SetPlatformid sets the Platformid property
func (m *UpdateDeliveryConfigQueryParameters) SetPlatformid(val string) {
	m.Platformid = val
}

// GetChannelid returns the Channelid property
func (m UpdateDeliveryConfigQueryParameters) GetChannelid() string {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *UpdateDeliveryConfigQueryParameters) SetChannelid(val string) {
	m.Channelid = val
}
