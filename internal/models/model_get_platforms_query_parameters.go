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

// GetPlatformsQueryParameters is an object.
type GetPlatformsQueryParameters struct {
	// Page: This signifies the page number.
	//
	// `Starts with 1`
	Page float32 `json:"page,omitempty" mapstructure:"page,omitempty"`
	// Offset: Number of accounts in a single page.
	//
	// `Default is 20`
	Offset float32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Region: Filter list for a given region
	Region string `json:"region,omitempty" mapstructure:"region,omitempty"`
}

// Validate implements basic validation for this model
func (m GetPlatformsQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPage returns the Page property
func (m GetPlatformsQueryParameters) GetPage() float32 {
	return m.Page
}

// SetPage sets the Page property
func (m *GetPlatformsQueryParameters) SetPage(val float32) {
	m.Page = val
}

// GetOffset returns the Offset property
func (m GetPlatformsQueryParameters) GetOffset() float32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *GetPlatformsQueryParameters) SetOffset(val float32) {
	m.Offset = val
}

// GetRegion returns the Region property
func (m GetPlatformsQueryParameters) GetRegion() string {
	return m.Region
}

// SetRegion sets the Region property
func (m *GetPlatformsQueryParameters) SetRegion(val string) {
	m.Region = val
}
