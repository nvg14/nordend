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

// GetSpecificChannelQueryParameters is an object.
type GetSpecificChannelQueryParameters struct {
	// Channelid:
	Channelid string `json:"channelid" mapstructure:"channelid"`
}

// Validate implements basic validation for this model
func (m GetSpecificChannelQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetChannelid returns the Channelid property
func (m GetSpecificChannelQueryParameters) GetChannelid() string {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *GetSpecificChannelQueryParameters) SetChannelid(val string) {
	m.Channelid = val
}
