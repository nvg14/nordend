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

// ChannelPatch is an object.
type ChannelPatch struct {
	// Op:
	Op string `json:"op,omitempty" mapstructure:"op,omitempty"`
	// Path:
	Path string `json:"path,omitempty" mapstructure:"path,omitempty"`
	// Value:
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}

// Validate implements basic validation for this model
func (m ChannelPatch) Validate() error {
	return validation.Errors{}.Filter()
}

// GetOp returns the Op property
func (m ChannelPatch) GetOp() string {
	return m.Op
}

// SetOp sets the Op property
func (m *ChannelPatch) SetOp(val string) {
	m.Op = val
}

// GetPath returns the Path property
func (m ChannelPatch) GetPath() string {
	return m.Path
}

// SetPath sets the Path property
func (m *ChannelPatch) SetPath(val string) {
	m.Path = val
}

// GetValue returns the Value property
func (m ChannelPatch) GetValue() string {
	return m.Value
}

// SetValue sets the Value property
func (m *ChannelPatch) SetValue(val string) {
	m.Value = val
}
