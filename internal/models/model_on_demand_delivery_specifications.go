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

// OnDemandDeliverySpecifications is an object.
type OnDemandDeliverySpecifications map[string]interface{}

// Validate implements basic validation for this model
func (m OnDemandDeliverySpecifications) Validate() error {
	return validation.Errors{}.Filter()
}
