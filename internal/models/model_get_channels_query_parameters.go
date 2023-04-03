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

// GetChannelsQueryParameters is an object.
type GetChannelsQueryParameters struct {
	// Page: This signifies the page number.
	//
	// `Starts with 1`
	Page float32 `json:"page,omitempty" mapstructure:"page,omitempty"`
	// Offset: Number of accounts in a single page.
	//
	// `Default is 10`
	Offset float32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Amgid: Filter list for a given amgid
	Amgid string `json:"amgid,omitempty" mapstructure:"amgid,omitempty"`
	// Language: Filter list for a primary audio language
	Language string `json:"language,omitempty" mapstructure:"language,omitempty"`
	// Timezone: Filter list for a timezone
	Timezone string `json:"timezone,omitempty" mapstructure:"timezone,omitempty"`
	// Rating: Filter list for a rating
	Rating string `json:"rating,omitempty" mapstructure:"rating,omitempty"`
	// Genre: Filter list for a genre
	Genre string `json:"genre,omitempty" mapstructure:"genre,omitempty"`
	// Tags: Filter list for given tags
	Tags string `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m GetChannelsQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPage returns the Page property
func (m GetChannelsQueryParameters) GetPage() float32 {
	return m.Page
}

// SetPage sets the Page property
func (m *GetChannelsQueryParameters) SetPage(val float32) {
	m.Page = val
}

// GetOffset returns the Offset property
func (m GetChannelsQueryParameters) GetOffset() float32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *GetChannelsQueryParameters) SetOffset(val float32) {
	m.Offset = val
}

// GetAmgid returns the Amgid property
func (m GetChannelsQueryParameters) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *GetChannelsQueryParameters) SetAmgid(val string) {
	m.Amgid = val
}

// GetLanguage returns the Language property
func (m GetChannelsQueryParameters) GetLanguage() string {
	return m.Language
}

// SetLanguage sets the Language property
func (m *GetChannelsQueryParameters) SetLanguage(val string) {
	m.Language = val
}

// GetTimezone returns the Timezone property
func (m GetChannelsQueryParameters) GetTimezone() string {
	return m.Timezone
}

// SetTimezone sets the Timezone property
func (m *GetChannelsQueryParameters) SetTimezone(val string) {
	m.Timezone = val
}

// GetRating returns the Rating property
func (m GetChannelsQueryParameters) GetRating() string {
	return m.Rating
}

// SetRating sets the Rating property
func (m *GetChannelsQueryParameters) SetRating(val string) {
	m.Rating = val
}

// GetGenre returns the Genre property
func (m GetChannelsQueryParameters) GetGenre() string {
	return m.Genre
}

// SetGenre sets the Genre property
func (m *GetChannelsQueryParameters) SetGenre(val string) {
	m.Genre = val
}

// GetTags returns the Tags property
func (m GetChannelsQueryParameters) GetTags() string {
	return m.Tags
}

// SetTags sets the Tags property
func (m *GetChannelsQueryParameters) SetTags(val string) {
	m.Tags = val
}
