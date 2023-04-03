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

// DeleteDeliveryQueryParameters is an object.
type DeleteDeliveryQueryParameters struct {
	// Amgid:
	Amgid string `json:"amgid" mapstructure:"amgid"`
	// Platformid:
	Platformid string `json:"platformid" mapstructure:"platformid"`
	// Channelid:
	Channelid string `json:"channelid,omitempty" mapstructure:"channelid,omitempty"`
}

// Validate implements basic validation for this model
func (m DeleteDeliveryQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAmgid returns the Amgid property
func (m DeleteDeliveryQueryParameters) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *DeleteDeliveryQueryParameters) SetAmgid(val string) {
	m.Amgid = val
}

// GetPlatformid returns the Platformid property
func (m DeleteDeliveryQueryParameters) GetPlatformid() string {
	return m.Platformid
}

// SetPlatformid sets the Platformid property
func (m *DeleteDeliveryQueryParameters) SetPlatformid(val string) {
	m.Platformid = val
}

// GetChannelid returns the Channelid property
func (m DeleteDeliveryQueryParameters) GetChannelid() string {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *DeleteDeliveryQueryParameters) SetChannelid(val string) {
	m.Channelid = val
}
