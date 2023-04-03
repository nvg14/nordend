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

// Rating is an object.
type Rating struct {
	// Body:
	Body string `json:"body,omitempty" mapstructure:"body,omitempty"`
	// Value:
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}

// Validate implements basic validation for this model
func (m Rating) Validate() error {
	return validation.Errors{}.Filter()
}

// GetBody returns the Body property
func (m Rating) GetBody() string {
	return m.Body
}

// SetBody sets the Body property
func (m *Rating) SetBody(val string) {
	m.Body = val
}

// GetValue returns the Value property
func (m Rating) GetValue() string {
	return m.Value
}

// SetValue sets the Value property
func (m *Rating) SetValue(val string) {
	m.Value = val
}
