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

// PlatformSpec is an object.
type PlatformSpec map[string]interface{}

// Validate implements basic validation for this model
func (m PlatformSpec) Validate() error {
	return validation.Errors{}.Filter()
}
