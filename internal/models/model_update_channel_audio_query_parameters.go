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

// UpdateChannelAudioQueryParameters is an object.
type UpdateChannelAudioQueryParameters struct {
	// Channelid:
	Channelid float32 `json:"channelid" mapstructure:"channelid"`
}

// Validate implements basic validation for this model
func (m UpdateChannelAudioQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetChannelid returns the Channelid property
func (m UpdateChannelAudioQueryParameters) GetChannelid() float32 {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *UpdateChannelAudioQueryParameters) SetChannelid(val float32) {
	m.Channelid = val
}
