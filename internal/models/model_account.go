// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Nordend
//	Version: 1.0.0
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

// accountAmgidPattern is the validation pattern for Account.Amgid
var accountAmgidPattern = regexp.MustCompile(`^AMG[0-9]{5}$`)

// Account is an object.
type Account struct {
	// Amgid:
	Amgid string `json:"amgid" mapstructure:"amgid"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Domain:
	Domain string `json:"domain,omitempty" mapstructure:"domain,omitempty"`
	// Logo:
	Logo string `json:"logo,omitempty" mapstructure:"logo,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
}

// Validate implements basic validation for this model
func (m Account) Validate() error {
	return validation.Errors{
		"amgid": validation.Validate(
			m.Amgid, validation.Match(accountAmgidPattern),
		),
	}.Filter()
}

// GetAmgid returns the Amgid property
func (m Account) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *Account) SetAmgid(val string) {
	m.Amgid = val
}

// GetDescription returns the Description property
func (m Account) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Account) SetDescription(val string) {
	m.Description = val
}

// GetDomain returns the Domain property
func (m Account) GetDomain() string {
	return m.Domain
}

// SetDomain sets the Domain property
func (m *Account) SetDomain(val string) {
	m.Domain = val
}

// GetLogo returns the Logo property
func (m Account) GetLogo() string {
	return m.Logo
}

// SetLogo sets the Logo property
func (m *Account) SetLogo(val string) {
	m.Logo = val
}

// GetName returns the Name property
func (m Account) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Account) SetName(val string) {
	m.Name = val
}
