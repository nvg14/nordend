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

// GetAccountByAmgIdQueryParameters is an object.
type GetAccountByAmgIdQueryParameters struct {
	// Amgid:
	Amgid string `json:"amgid" mapstructure:"amgid"`
}

// Validate implements basic validation for this model
func (m GetAccountByAmgIdQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAmgid returns the Amgid property
func (m GetAccountByAmgIdQueryParameters) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *GetAccountByAmgIdQueryParameters) SetAmgid(val string) {
	m.Amgid = val
}
