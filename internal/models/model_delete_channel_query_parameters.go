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

// DeleteChannelQueryParameters is an object.
type DeleteChannelQueryParameters struct {
	// Channelid:
	Channelid string `json:"channelid" mapstructure:"channelid"`
}

// Validate implements basic validation for this model
func (m DeleteChannelQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetChannelid returns the Channelid property
func (m DeleteChannelQueryParameters) GetChannelid() string {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *DeleteChannelQueryParameters) SetChannelid(val string) {
	m.Channelid = val
}
