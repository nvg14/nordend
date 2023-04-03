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

// MonetizationSpecifications is an object.
type MonetizationSpecifications map[string]interface{}

// Validate implements basic validation for this model
func (m MonetizationSpecifications) Validate() error {
	return validation.Errors{}.Filter()
}
