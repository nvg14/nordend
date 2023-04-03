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

// DeliverySpecifications is an object.
type DeliverySpecifications map[string]interface{}

// Validate implements basic validation for this model
func (m DeliverySpecifications) Validate() error {
	return validation.Errors{}.Filter()
}
