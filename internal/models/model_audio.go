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

// Audio is an object.
type Audio struct {
	// IsPrimary:
	IsPrimary bool `json:"is_primary,omitempty" mapstructure:"is_primary,omitempty"`
	// Language:
	Language AudioLanguages `json:"language,omitempty" mapstructure:"language,omitempty"`
	// Track:
	Track float32 `json:"track,omitempty" mapstructure:"track,omitempty"`
}

// Validate implements basic validation for this model
func (m Audio) Validate() error {
	return validation.Errors{
		"language": validation.Validate(
			m.Language,
		),
	}.Filter()
}

// GetIsPrimary returns the IsPrimary property
func (m Audio) GetIsPrimary() bool {
	return m.IsPrimary
}

// SetIsPrimary sets the IsPrimary property
func (m *Audio) SetIsPrimary(val bool) {
	m.IsPrimary = val
}

// GetLanguage returns the Language property
func (m Audio) GetLanguage() AudioLanguages {
	return m.Language
}

// SetLanguage sets the Language property
func (m *Audio) SetLanguage(val AudioLanguages) {
	m.Language = val
}

// GetTrack returns the Track property
func (m Audio) GetTrack() float32 {
	return m.Track
}

// SetTrack sets the Track property
func (m *Audio) SetTrack(val float32) {
	m.Track = val
}
