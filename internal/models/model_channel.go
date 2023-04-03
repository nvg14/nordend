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

// channelAmgidPattern is the validation pattern for Channel.Amgid
var channelAmgidPattern = regexp.MustCompile(`^AMG[0-9]{5}$`)

// channelStartTimePattern is the validation pattern for Channel.StartTime
var channelStartTimePattern = regexp.MustCompile(`([01]?[0-9]|2[0-3]):[0-5][0-9]?`)

// Channel is an object.
type Channel struct {
	// Amgid:
	Amgid string `json:"amgid,omitempty" mapstructure:"amgid,omitempty"`
	// Artworks:
	Artworks []Artwork `json:"artworks,omitempty" mapstructure:"artworks,omitempty"`
	// Audios:
	Audios []Audio `json:"audios,omitempty" mapstructure:"audios,omitempty"`
	// Caption:
	Caption []CaptionObject `json:"caption,omitempty" mapstructure:"caption,omitempty"`
	// Genre:
	Genre string `json:"genre,omitempty" mapstructure:"genre,omitempty"`
	// LongDescription:
	LongDescription string `json:"long_description,omitempty" mapstructure:"long_description,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Rating:
	Rating []Rating `json:"rating,omitempty" mapstructure:"rating,omitempty"`
	// ShortDescription:
	ShortDescription string `json:"short_description,omitempty" mapstructure:"short_description,omitempty"`
	// StartTime:
	StartTime string `json:"start_time,omitempty" mapstructure:"start_time,omitempty"`
	// Tags:
	Tags map[string]interface{} `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Timezone:
	Timezone Timezone `json:"timezone,omitempty" mapstructure:"timezone,omitempty"`
}

// Validate implements basic validation for this model
func (m Channel) Validate() error {
	return validation.Errors{
		"amgid": validation.Validate(
			m.Amgid, validation.Match(channelAmgidPattern),
		),
		"artworks": validation.Validate(
			m.Artworks,
		),
		"audios": validation.Validate(
			m.Audios,
		),
		"caption": validation.Validate(
			m.Caption,
		),
		"rating": validation.Validate(
			m.Rating,
		),
		"startTime": validation.Validate(
			m.StartTime, validation.Match(channelStartTimePattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"timezone": validation.Validate(
			m.Timezone,
		),
	}.Filter()
}

// GetAmgid returns the Amgid property
func (m Channel) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *Channel) SetAmgid(val string) {
	m.Amgid = val
}

// GetArtworks returns the Artworks property
func (m Channel) GetArtworks() []Artwork {
	return m.Artworks
}

// SetArtworks sets the Artworks property
func (m *Channel) SetArtworks(val []Artwork) {
	m.Artworks = val
}

// GetAudios returns the Audios property
func (m Channel) GetAudios() []Audio {
	return m.Audios
}

// SetAudios sets the Audios property
func (m *Channel) SetAudios(val []Audio) {
	m.Audios = val
}

// GetCaption returns the Caption property
func (m Channel) GetCaption() []CaptionObject {
	return m.Caption
}

// SetCaption sets the Caption property
func (m *Channel) SetCaption(val []CaptionObject) {
	m.Caption = val
}

// GetGenre returns the Genre property
func (m Channel) GetGenre() string {
	return m.Genre
}

// SetGenre sets the Genre property
func (m *Channel) SetGenre(val string) {
	m.Genre = val
}

// GetLongDescription returns the LongDescription property
func (m Channel) GetLongDescription() string {
	return m.LongDescription
}

// SetLongDescription sets the LongDescription property
func (m *Channel) SetLongDescription(val string) {
	m.LongDescription = val
}

// GetName returns the Name property
func (m Channel) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Channel) SetName(val string) {
	m.Name = val
}

// GetRating returns the Rating property
func (m Channel) GetRating() []Rating {
	return m.Rating
}

// SetRating sets the Rating property
func (m *Channel) SetRating(val []Rating) {
	m.Rating = val
}

// GetShortDescription returns the ShortDescription property
func (m Channel) GetShortDescription() string {
	return m.ShortDescription
}

// SetShortDescription sets the ShortDescription property
func (m *Channel) SetShortDescription(val string) {
	m.ShortDescription = val
}

// GetStartTime returns the StartTime property
func (m Channel) GetStartTime() string {
	return m.StartTime
}

// SetStartTime sets the StartTime property
func (m *Channel) SetStartTime(val string) {
	m.StartTime = val
}

// GetTags returns the Tags property
func (m Channel) GetTags() map[string]interface{} {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Channel) SetTags(val map[string]interface{}) {
	m.Tags = val
}

// GetTimezone returns the Timezone property
func (m Channel) GetTimezone() Timezone {
	return m.Timezone
}

// SetTimezone sets the Timezone property
func (m *Channel) SetTimezone(val Timezone) {
	m.Timezone = val
}
