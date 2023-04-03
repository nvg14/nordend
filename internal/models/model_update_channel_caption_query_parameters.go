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

// UpdateChannelCaptionQueryParameters is an object.
type UpdateChannelCaptionQueryParameters struct {
	// Channelid:
	Channelid float32 `json:"channelid" mapstructure:"channelid"`
}

// Validate implements basic validation for this model
func (m UpdateChannelCaptionQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetChannelid returns the Channelid property
func (m UpdateChannelCaptionQueryParameters) GetChannelid() float32 {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *UpdateChannelCaptionQueryParameters) SetChannelid(val float32) {
	m.Channelid = val
}
