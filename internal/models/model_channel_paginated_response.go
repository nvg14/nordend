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

// ChannelPaginatedResponse is an object.
type ChannelPaginatedResponse struct {
	// Offset:
	Offset float32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Page:
	Page float32 `json:"page,omitempty" mapstructure:"page,omitempty"`
	// Result:
	Result []ChannelResponse `json:"result,omitempty" mapstructure:"result,omitempty"`
	// Total:
	Total float32 `json:"total,omitempty" mapstructure:"total,omitempty"`
}

// Validate implements basic validation for this model
func (m ChannelPaginatedResponse) Validate() error {
	return validation.Errors{
		"result": validation.Validate(
			m.Result,
		),
	}.Filter()
}

// GetOffset returns the Offset property
func (m ChannelPaginatedResponse) GetOffset() float32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ChannelPaginatedResponse) SetOffset(val float32) {
	m.Offset = val
}

// GetPage returns the Page property
func (m ChannelPaginatedResponse) GetPage() float32 {
	return m.Page
}

// SetPage sets the Page property
func (m *ChannelPaginatedResponse) SetPage(val float32) {
	m.Page = val
}

// GetResult returns the Result property
func (m ChannelPaginatedResponse) GetResult() []ChannelResponse {
	return m.Result
}

// SetResult sets the Result property
func (m *ChannelPaginatedResponse) SetResult(val []ChannelResponse) {
	m.Result = val
}

// GetTotal returns the Total property
func (m ChannelPaginatedResponse) GetTotal() float32 {
	return m.Total
}

// SetTotal sets the Total property
func (m *ChannelPaginatedResponse) SetTotal(val float32) {
	m.Total = val
}
