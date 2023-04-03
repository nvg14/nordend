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

// UpdateChannelRatingsQueryParameters is an object.
type UpdateChannelRatingsQueryParameters struct {
	// Channelid:
	Channelid float32 `json:"channelid" mapstructure:"channelid"`
}

// Validate implements basic validation for this model
func (m UpdateChannelRatingsQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetChannelid returns the Channelid property
func (m UpdateChannelRatingsQueryParameters) GetChannelid() float32 {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *UpdateChannelRatingsQueryParameters) SetChannelid(val float32) {
	m.Channelid = val
}
