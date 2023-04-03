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

// Artwork is an object.
type Artwork struct {
	// File:
	File string `json:"file,omitempty" mapstructure:"file,omitempty"`
	// Height:
	Height float32 `json:"height,omitempty" mapstructure:"height,omitempty"`
	// Labels:
	Labels map[string]interface{} `json:"labels,omitempty" mapstructure:"labels,omitempty"`
	// Width:
	Width float32 `json:"width,omitempty" mapstructure:"width,omitempty"`
}

// Validate implements basic validation for this model
func (m Artwork) Validate() error {
	return validation.Errors{
		"labels": validation.Validate(
			m.Labels,
		),
	}.Filter()
}

// GetFile returns the File property
func (m Artwork) GetFile() string {
	return m.File
}

// SetFile sets the File property
func (m *Artwork) SetFile(val string) {
	m.File = val
}

// GetHeight returns the Height property
func (m Artwork) GetHeight() float32 {
	return m.Height
}

// SetHeight sets the Height property
func (m *Artwork) SetHeight(val float32) {
	m.Height = val
}

// GetLabels returns the Labels property
func (m Artwork) GetLabels() map[string]interface{} {
	return m.Labels
}

// SetLabels sets the Labels property
func (m *Artwork) SetLabels(val map[string]interface{}) {
	m.Labels = val
}

// GetWidth returns the Width property
func (m Artwork) GetWidth() float32 {
	return m.Width
}

// SetWidth sets the Width property
func (m *Artwork) SetWidth(val float32) {
	m.Width = val
}
