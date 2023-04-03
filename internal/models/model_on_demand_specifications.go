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

// OnDemandSpecifications is an object.
type OnDemandSpecifications struct {
	// DeliverySpecification:
	DeliverySpecification *struct {
		OndemandDeliverySpecification OnDemandDeliverySpecifications `json:"ondemand_delivery_specification,omitempty" mapstructure:"ondemand_delivery_specification,omitempty"`
	} `json:"delivery_specification,omitempty" mapstructure:"delivery_specification,omitempty"`
	// MonetizationSpecification:
	MonetizationSpecification MonetizationSpecifications `json:"monetization_specification,omitempty" mapstructure:"monetization_specification,omitempty"`
	// TranscodingSpecification:
	TranscodingSpecification TranscodingSpecifications `json:"transcoding_specification,omitempty" mapstructure:"transcoding_specification,omitempty"`
}

// Validate implements basic validation for this model
func (m OnDemandSpecifications) Validate() error {
	return validation.Errors{
		"monetizationSpecification": validation.Validate(
			m.MonetizationSpecification,
		),
		"transcodingSpecification": validation.Validate(
			m.TranscodingSpecification,
		),
	}.Filter()
}

// GetDeliverySpecification returns the DeliverySpecification property
func (m OnDemandSpecifications) GetDeliverySpecification() *struct {
	OndemandDeliverySpecification OnDemandDeliverySpecifications `json:"ondemand_delivery_specification,omitempty" mapstructure:"ondemand_delivery_specification,omitempty"`
} {
	return m.DeliverySpecification
}

// SetDeliverySpecification sets the DeliverySpecification property
func (m *OnDemandSpecifications) SetDeliverySpecification(val *struct {
	OndemandDeliverySpecification OnDemandDeliverySpecifications `json:"ondemand_delivery_specification,omitempty" mapstructure:"ondemand_delivery_specification,omitempty"`
}) {
	m.DeliverySpecification = val
}

// GetMonetizationSpecification returns the MonetizationSpecification property
func (m OnDemandSpecifications) GetMonetizationSpecification() MonetizationSpecifications {
	return m.MonetizationSpecification
}

// SetMonetizationSpecification sets the MonetizationSpecification property
func (m *OnDemandSpecifications) SetMonetizationSpecification(val MonetizationSpecifications) {
	m.MonetizationSpecification = val
}

// GetTranscodingSpecification returns the TranscodingSpecification property
func (m OnDemandSpecifications) GetTranscodingSpecification() TranscodingSpecifications {
	return m.TranscodingSpecification
}

// SetTranscodingSpecification sets the TranscodingSpecification property
func (m *OnDemandSpecifications) SetTranscodingSpecification(val TranscodingSpecifications) {
	m.TranscodingSpecification = val
}
