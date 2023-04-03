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

// AccountPaginatedResponse is an object.
type AccountPaginatedResponse struct {
	// Offset:
	Offset float32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Page:
	Page float32 `json:"page,omitempty" mapstructure:"page,omitempty"`
	// Result:
	Result []AccountResponse `json:"result,omitempty" mapstructure:"result,omitempty"`
}

// Validate implements basic validation for this model
func (m AccountPaginatedResponse) Validate() error {
	return validation.Errors{
		"result": validation.Validate(
			m.Result,
		),
	}.Filter()
}

// GetOffset returns the Offset property
func (m AccountPaginatedResponse) GetOffset() float32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *AccountPaginatedResponse) SetOffset(val float32) {
	m.Offset = val
}

// GetPage returns the Page property
func (m AccountPaginatedResponse) GetPage() float32 {
	return m.Page
}

// SetPage sets the Page property
func (m *AccountPaginatedResponse) SetPage(val float32) {
	m.Page = val
}

// GetResult returns the Result property
func (m AccountPaginatedResponse) GetResult() []AccountResponse {
	return m.Result
}

// SetResult sets the Result property
func (m *AccountPaginatedResponse) SetResult(val []AccountResponse) {
	m.Result = val
}
