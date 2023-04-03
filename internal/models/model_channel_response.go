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

// ChannelResponse is an object.
type ChannelResponse struct {
	// Amgid:
	Amgid string `json:"amgid,omitempty" mapstructure:"amgid,omitempty"`
	// Artworks:
	Artworks []Artwork `json:"artworks,omitempty" mapstructure:"artworks,omitempty"`
	// Audios:
	Audios []Audio `json:"audios,omitempty" mapstructure:"audios,omitempty"`
	// Caption:
	Caption []CaptionObject `json:"caption,omitempty" mapstructure:"caption,omitempty"`
	// Channelid:
	Channelid string `json:"channelid,omitempty" mapstructure:"channelid,omitempty"`
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
	// Status:
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags map[string]interface{} `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Timezone:
	Timezone Timezone `json:"timezone,omitempty" mapstructure:"timezone,omitempty"`
}

// Validate implements basic validation for this model
func (m ChannelResponse) Validate() error {
	return validation.Errors{
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
		"tags": validation.Validate(
			m.Tags,
		),
		"timezone": validation.Validate(
			m.Timezone,
		),
	}.Filter()
}

// GetAmgid returns the Amgid property
func (m ChannelResponse) GetAmgid() string {
	return m.Amgid
}

// SetAmgid sets the Amgid property
func (m *ChannelResponse) SetAmgid(val string) {
	m.Amgid = val
}

// GetArtworks returns the Artworks property
func (m ChannelResponse) GetArtworks() []Artwork {
	return m.Artworks
}

// SetArtworks sets the Artworks property
func (m *ChannelResponse) SetArtworks(val []Artwork) {
	m.Artworks = val
}

// GetAudios returns the Audios property
func (m ChannelResponse) GetAudios() []Audio {
	return m.Audios
}

// SetAudios sets the Audios property
func (m *ChannelResponse) SetAudios(val []Audio) {
	m.Audios = val
}

// GetCaption returns the Caption property
func (m ChannelResponse) GetCaption() []CaptionObject {
	return m.Caption
}

// SetCaption sets the Caption property
func (m *ChannelResponse) SetCaption(val []CaptionObject) {
	m.Caption = val
}

// GetChannelid returns the Channelid property
func (m ChannelResponse) GetChannelid() string {
	return m.Channelid
}

// SetChannelid sets the Channelid property
func (m *ChannelResponse) SetChannelid(val string) {
	m.Channelid = val
}

// GetGenre returns the Genre property
func (m ChannelResponse) GetGenre() string {
	return m.Genre
}

// SetGenre sets the Genre property
func (m *ChannelResponse) SetGenre(val string) {
	m.Genre = val
}

// GetLongDescription returns the LongDescription property
func (m ChannelResponse) GetLongDescription() string {
	return m.LongDescription
}

// SetLongDescription sets the LongDescription property
func (m *ChannelResponse) SetLongDescription(val string) {
	m.LongDescription = val
}

// GetName returns the Name property
func (m ChannelResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ChannelResponse) SetName(val string) {
	m.Name = val
}

// GetRating returns the Rating property
func (m ChannelResponse) GetRating() []Rating {
	return m.Rating
}

// SetRating sets the Rating property
func (m *ChannelResponse) SetRating(val []Rating) {
	m.Rating = val
}

// GetShortDescription returns the ShortDescription property
func (m ChannelResponse) GetShortDescription() string {
	return m.ShortDescription
}

// SetShortDescription sets the ShortDescription property
func (m *ChannelResponse) SetShortDescription(val string) {
	m.ShortDescription = val
}

// GetStartTime returns the StartTime property
func (m ChannelResponse) GetStartTime() string {
	return m.StartTime
}

// SetStartTime sets the StartTime property
func (m *ChannelResponse) SetStartTime(val string) {
	m.StartTime = val
}

// GetStatus returns the Status property
func (m ChannelResponse) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *ChannelResponse) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m ChannelResponse) GetTags() map[string]interface{} {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ChannelResponse) SetTags(val map[string]interface{}) {
	m.Tags = val
}

// GetTimezone returns the Timezone property
func (m ChannelResponse) GetTimezone() Timezone {
	return m.Timezone
}

// SetTimezone sets the Timezone property
func (m *ChannelResponse) SetTimezone(val Timezone) {
	m.Timezone = val
}
