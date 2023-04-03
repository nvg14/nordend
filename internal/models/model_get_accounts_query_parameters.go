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

// GetAccountsQueryParameters is an object.
type GetAccountsQueryParameters struct {
	// Page: This signifies the page number.
	//
	// `Starts with 1`
	Page float32 `json:"page,omitempty" mapstructure:"page,omitempty"`
	// Offset: Number of accounts in a single page.
	//
	// `Default is 20`
	Offset float32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
}

// Validate implements basic validation for this model
func (m GetAccountsQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPage returns the Page property
func (m GetAccountsQueryParameters) GetPage() float32 {
	return m.Page
}

// SetPage sets the Page property
func (m *GetAccountsQueryParameters) SetPage(val float32) {
	m.Page = val
}

// GetOffset returns the Offset property
func (m GetAccountsQueryParameters) GetOffset() float32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *GetAccountsQueryParameters) SetOffset(val float32) {
	m.Offset = val
}
