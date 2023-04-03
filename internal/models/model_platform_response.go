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

// platformResponseIdPattern is the validation pattern for PlatformResponse.Id
var platformResponseIdPattern = regexp.MustCompile(`PLT[0-9]{5}`)

// PlatformResponse is an object.
type PlatformResponse struct {
	// Id:
	Id string `json:"id,omitempty" mapstructure:"id,omitempty"`
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
func (m PlatformResponse) Validate() error {
	return validation.Errors{
		"id": validation.Validate(
			m.Id, validation.Match(platformResponseIdPattern),
		),
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

// GetId returns the Id property
func (m PlatformResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m *PlatformResponse) SetId(val string) {
	m.Id = val
}

// GetLinearSpecification returns the LinearSpecification property
func (m PlatformResponse) GetLinearSpecification() *LinearSpecifications {
	return m.LinearSpecification
}

// SetLinearSpecification sets the LinearSpecification property
func (m *PlatformResponse) SetLinearSpecification(val *LinearSpecifications) {
	m.LinearSpecification = val
}

// GetLogo returns the Logo property
func (m PlatformResponse) GetLogo() string {
	return m.Logo
}

// SetLogo sets the Logo property
func (m *PlatformResponse) SetLogo(val string) {
	m.Logo = val
}

// GetLongDescription returns the LongDescription property
func (m PlatformResponse) GetLongDescription() string {
	return m.LongDescription
}

// SetLongDescription sets the LongDescription property
func (m *PlatformResponse) SetLongDescription(val string) {
	m.LongDescription = val
}

// GetName returns the Name property
func (m PlatformResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PlatformResponse) SetName(val string) {
	m.Name = val
}

// GetOndemandSpecification returns the OndemandSpecification property
func (m PlatformResponse) GetOndemandSpecification() *OnDemandSpecifications {
	return m.OndemandSpecification
}

// SetOndemandSpecification sets the OndemandSpecification property
func (m *PlatformResponse) SetOndemandSpecification(val *OnDemandSpecifications) {
	m.OndemandSpecification = val
}

// GetRegion returns the Region property
func (m PlatformResponse) GetRegion() []string {
	return m.Region
}

// SetRegion sets the Region property
func (m *PlatformResponse) SetRegion(val []string) {
	m.Region = val
}

// GetShortDescription returns the ShortDescription property
func (m PlatformResponse) GetShortDescription() string {
	return m.ShortDescription
}

// SetShortDescription sets the ShortDescription property
func (m *PlatformResponse) SetShortDescription(val string) {
	m.ShortDescription = val
}
