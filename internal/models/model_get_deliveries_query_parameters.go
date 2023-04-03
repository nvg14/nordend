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

// GetDeliveriesQueryParameters is an object.
type GetDeliveriesQueryParameters struct {
	// Page: This signifies the page number.
	//
	// `Starts with 1`
	Page float32 `json:"page,omitempty" mapstructure:"page,omitempty"`
	// Offset: Number of accounts in a single page.
	//
	// `Default is 20`
	Offset float32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Amgid: Filter list for a given amgid
	Amgid string `json:"amgid" mapstructure:"amgid"`
	// Platformid: Filter list for a given platformid
	Platformid string `json:"platformid" mapstructure:"platformid"`
	// Channelid: Filter list for a given channelid
	Channelid string `json:"channelid,omitempty" mapstructure:"channelid,omitempty"`
}

// Validate implements basic validation for this model
func (m GetDeliveriesQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPage returns the Page property
func (m GetDeliveriesQueryParameters) GetPage() float32 {
	return m.Page
}

// SetPage sets the Page property
func (m *GetDeliveriesQueryParameters) SetPage(val float32) {
	m.Page = val
}

// GetOffset returns the Offset property
func (m GetDeliveriesQueryParameters) GetOffset() float32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *GetDeliveriesQueryParameters) SetOffset(val float32) {
	m.Offset = val
}

// GetAmgid returns the Amgid property
func (m GetDeliveriesQueryParameters) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *GetDeliveriesQueryParameters) SetAmgid(val string) {
	m.Amgid = val
}

// GetPlatformid returns the Platformid property
func (m GetDeliveriesQueryParameters) GetPlatformid() string {
	return m.Platformid
}

// SetPlatformid sets the Platformid property
func (m *GetDeliveriesQueryParameters) SetPlatformid(val string) {
	m.Platformid = val
}

// GetChannelid returns the Channelid property
func (m GetDeliveriesQueryParameters) GetChannelid() string {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *GetDeliveriesQueryParameters) SetChannelid(val string) {
	m.Channelid = val
}
