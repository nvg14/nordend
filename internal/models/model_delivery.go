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

// Delivery is an object.
type Delivery struct {
	// Amgid:
	Amgid string `json:"amgid,omitempty" mapstructure:"amgid,omitempty"`
	// Channelid:
	Channelid string `json:"channelid,omitempty" mapstructure:"channelid,omitempty"`
	// Platform:
	Platform *struct {
		Deliveries PlatformSpec `json:"deliveries,omitempty" mapstructure:"deliveries,omitempty"`
		Id         string       `json:"id,omitempty" mapstructure:"id,omitempty"`
	} `json:"platform,omitempty" mapstructure:"platform,omitempty"`
}

// Validate implements basic validation for this model
func (m Delivery) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAmgid returns the Amgid property
func (m Delivery) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *Delivery) SetAmgid(val string) {
	m.Amgid = val
}

// GetChannelid returns the Channelid property
func (m Delivery) GetChannelid() string {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *Delivery) SetChannelid(val string) {
	m.Channelid = val
}

// GetPlatform returns the Platform property
func (m Delivery) GetPlatform() *struct {
	Deliveries PlatformSpec `json:"deliveries,omitempty" mapstructure:"deliveries,omitempty"`
	Id         string       `json:"id,omitempty" mapstructure:"id,omitempty"`
} {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *Delivery) SetPlatform(val *struct {
	Deliveries PlatformSpec `json:"deliveries,omitempty" mapstructure:"deliveries,omitempty"`
	Id         string       `json:"id,omitempty" mapstructure:"id,omitempty"`
}) {
	m.Platform = val
}
