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

// CaptionObject is an object.
type CaptionObject struct {
	// Format:
	Format string `json:"format,omitempty" mapstructure:"format,omitempty"`
	// Language:
	Language []AudioLanguages `json:"language,omitempty" mapstructure:"language,omitempty"`
}

// Validate implements basic validation for this model
func (m CaptionObject) Validate() error {
	return validation.Errors{
		"language": validation.Validate(
			m.Language,
		),
	}.Filter()
}

// GetFormat returns the Format property
func (m CaptionObject) GetFormat() string {
	return m.Format
}

// SetFormat sets the Format property
func (m *CaptionObject) SetFormat(val string) {
	m.Format = val
}

// GetLanguage returns the Language property
func (m CaptionObject) GetLanguage() []AudioLanguages {
	return m.Language
}

// SetLanguage sets the Language property
func (m *CaptionObject) SetLanguage(val []AudioLanguages) {
	m.Language = val
}
