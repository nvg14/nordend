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

// Conflict is an object.
type Conflict struct {
	// Errors:
	Errors map[string]interface{} `json:"errors,omitempty" mapstructure:"errors,omitempty"`
	// Message:
	Message string `json:"message,omitempty" mapstructure:"message,omitempty"`
}

// Validate implements basic validation for this model
func (m Conflict) Validate() error {
	return validation.Errors{
		"errors": validation.Validate(
			m.Errors,
		),
	}.Filter()
}

// GetErrors returns the Errors property
func (m Conflict) GetErrors() map[string]interface{} {
	return m.Errors
}

// SetErrors sets the Errors property
func (m *Conflict) SetErrors(val map[string]interface{}) {
	m.Errors = val
}

// GetMessage returns the Message property
func (m Conflict) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m *Conflict) SetMessage(val string) {
	m.Message = val
}
