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

// UpdateChannelArtworkQueryParameters is an object.
type UpdateChannelArtworkQueryParameters struct {
	// Channelid:
	Channelid float32 `json:"channelid" mapstructure:"channelid"`
}

// Validate implements basic validation for this model
func (m UpdateChannelArtworkQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetChannelid returns the Channelid property
func (m UpdateChannelArtworkQueryParameters) GetChannelid() float32 {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *UpdateChannelArtworkQueryParameters) SetChannelid(val float32) {
	m.Channelid = val
}
