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

// LinearSpecifications is an object.
type LinearSpecifications struct {
	// DeliverySpecification:
	DeliverySpecification *struct {
		EpgDeliverySpecification       EPGDeliverySpecifications       `json:"epg_delivery_specification,omitempty" mapstructure:"epg_delivery_specification,omitempty"`
		StreamingDeliverySpecification StreamingDeliverySpecifications `json:"streaming_delivery_specification,omitempty" mapstructure:"streaming_delivery_specification,omitempty"`
	} `json:"delivery_specification,omitempty" mapstructure:"delivery_specification,omitempty"`
	// EpgSpecification:
	EpgSpecification EPGSpecifications `json:"epg_specification,omitempty" mapstructure:"epg_specification,omitempty"`
	// MonetizationSpecification:
	MonetizationSpecification MonetizationSpecifications `json:"monetization_specification,omitempty" mapstructure:"monetization_specification,omitempty"`
	// TranscodingSpecification:
	TranscodingSpecification TranscodingSpecifications `json:"transcoding_specification,omitempty" mapstructure:"transcoding_specification,omitempty"`
}

// Validate implements basic validation for this model
func (m LinearSpecifications) Validate() error {
	return validation.Errors{
		"epgSpecification": validation.Validate(
			m.EpgSpecification,
		),
		"monetizationSpecification": validation.Validate(
			m.MonetizationSpecification,
		),
		"transcodingSpecification": validation.Validate(
			m.TranscodingSpecification,
		),
	}.Filter()
}

// GetDeliverySpecification returns the DeliverySpecification property
func (m LinearSpecifications) GetDeliverySpecification() *struct {
	EpgDeliverySpecification       EPGDeliverySpecifications       `json:"epg_delivery_specification,omitempty" mapstructure:"epg_delivery_specification,omitempty"`
	StreamingDeliverySpecification StreamingDeliverySpecifications `json:"streaming_delivery_specification,omitempty" mapstructure:"streaming_delivery_specification,omitempty"`
} {
	return m.DeliverySpecification
}

// SetDeliverySpecification sets the DeliverySpecification property
func (m *LinearSpecifications) SetDeliverySpecification(val *struct {
	EpgDeliverySpecification       EPGDeliverySpecifications       `json:"epg_delivery_specification,omitempty" mapstructure:"epg_delivery_specification,omitempty"`
	StreamingDeliverySpecification StreamingDeliverySpecifications `json:"streaming_delivery_specification,omitempty" mapstructure:"streaming_delivery_specification,omitempty"`
}) {
	m.DeliverySpecification = val
}

// GetEpgSpecification returns the EpgSpecification property
func (m LinearSpecifications) GetEpgSpecification() EPGSpecifications {
	return m.EpgSpecification
}

// SetEpgSpecification sets the EpgSpecification property
func (m *LinearSpecifications) SetEpgSpecification(val EPGSpecifications) {
	m.EpgSpecification = val
}

// GetMonetizationSpecification returns the MonetizationSpecification property
func (m LinearSpecifications) GetMonetizationSpecification() MonetizationSpecifications {
	return m.MonetizationSpecification
}

// SetMonetizationSpecification sets the MonetizationSpecification property
func (m *LinearSpecifications) SetMonetizationSpecification(val MonetizationSpecifications) {
	m.MonetizationSpecification = val
}

// GetTranscodingSpecification returns the TranscodingSpecification property
func (m LinearSpecifications) GetTranscodingSpecification() TranscodingSpecifications {
	return m.TranscodingSpecification
}

// SetTranscodingSpecification sets the TranscodingSpecification property
func (m *LinearSpecifications) SetTranscodingSpecification(val TranscodingSpecifications) {
	m.TranscodingSpecification = val
}
