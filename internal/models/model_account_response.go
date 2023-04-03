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

// AccountResponse is an object.
type AccountResponse struct {
	// Amgid:
	Amgid string `json:"amgid,omitempty" mapstructure:"amgid,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Domain:
	Domain string `json:"domain,omitempty" mapstructure:"domain,omitempty"`
	// Logo:
	Logo string `json:"logo,omitempty" mapstructure:"logo,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
}

// Validate implements basic validation for this model
func (m AccountResponse) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAmgid returns the Amgid property
func (m AccountResponse) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *AccountResponse) SetAmgid(val string) {
	m.Amgid = val
}

// GetDescription returns the Description property
func (m AccountResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *AccountResponse) SetDescription(val string) {
	m.Description = val
}

// GetDomain returns the Domain property
func (m AccountResponse) GetDomain() string {
	return m.Domain
}

// SetDomain sets the Domain property
func (m *AccountResponse) SetDomain(val string) {
	m.Domain = val
}

// GetLogo returns the Logo property
func (m AccountResponse) GetLogo() string {
	return m.Logo
}

// SetLogo sets the Logo property
func (m *AccountResponse) SetLogo(val string) {
	m.Logo = val
}

// GetName returns the Name property
func (m AccountResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *AccountResponse) SetName(val string) {
	m.Name = val
}
