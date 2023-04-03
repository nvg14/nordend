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

// Platform is an object.
type Platform struct {
	// LinearSpecification:
	LinearSpecification *LinearSpecifications `json:"linear_specification,omitempty" mapstructure:"linear_specification,omitempty"`
	// Logo:
	Logo string `json:"logo,omitempty" mapstructure:"logo,omitempty"`
	// LongDescription:
	LongDescription string `json:"long_description,omitempty" mapstructure:"long_description,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// OndemandSpecification:
	OndemandSpecification *OnDemandSpecifications `json:"ondemand_specification,omitempty" mapstructure:"ondemand_specification,omitempty"`
	// Region:
	Region []string `json:"region,omitempty" mapstructure:"region,omitempty"`
	// ShortDescription:
	ShortDescription string `json:"short_description,omitempty" mapstructure:"short_description,omitempty"`
}

// Validate implements basic validation for this model
func (m Platform) Validate() error {
	return validation.Errors{
		"linearSpecification": validation.Validate(
			m.LinearSpecification,
		),
		"ondemandSpecification": validation.Validate(
			m.OndemandSpecification,
		),
		"region": validation.Validate(
			m.Region,
		),
	}.Filter()
}

// GetLinearSpecification returns the LinearSpecification property
func (m Platform) GetLinearSpecification() *LinearSpecifications {
	return m.LinearSpecification
}

// SetLinearSpecification sets the LinearSpecification property
func (m *Platform) SetLinearSpecification(val *LinearSpecifications) {
	m.LinearSpecification = val
}

// GetLogo returns the Logo property
func (m Platform) GetLogo() string {
	return m.Logo
}

// SetLogo sets the Logo property
func (m *Platform) SetLogo(val string) {
	m.Logo = val
}

// GetLongDescription returns the LongDescription property
func (m Platform) GetLongDescription() string {
	return m.LongDescription
}

// SetLongDescription sets the LongDescription property
func (m *Platform) SetLongDescription(val string) {
	m.LongDescription = val
}

// GetName returns the Name property
func (m Platform) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Platform) SetName(val string) {
	m.Name = val
}

// GetOndemandSpecification returns the OndemandSpecification property
func (m Platform) GetOndemandSpecification() *OnDemandSpecifications {
	return m.OndemandSpecification
}

// SetOndemandSpecification sets the OndemandSpecification property
func (m *Platform) SetOndemandSpecification(val *OnDemandSpecifications) {
	m.OndemandSpecification = val
}

// GetRegion returns the Region property
func (m Platform) GetRegion() []string {
	return m.Region
}

// SetRegion sets the Region property
func (m *Platform) SetRegion(val []string) {
	m.Region = val
}

// GetShortDescription returns the ShortDescription property
func (m Platform) GetShortDescription() string {
	return m.ShortDescription
}

// SetShortDescription sets the ShortDescription property
func (m *Platform) SetShortDescription(val string) {
	m.ShortDescription = val
}
