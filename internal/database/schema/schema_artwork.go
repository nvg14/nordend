package schema

import (
	"github.com/amagimedia/nordend/internal/database"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Artwork struct {
	CustomGormModel
	File      string         `json:"file,omitempty" mapstructure:"file,omitempty"`
	Height    float32        `json:"height,omitempty" mapstructure:"height,omitempty"`
	Labels    datatypes.JSON `json:"labels,omitempty" mapstructure:"labels,omitempty"`
	Width     float32        `json:"width,omitempty" mapstructure:"width,omitempty"`
	ChannelID uint
}

func (a *Artwork) Save() (*gorm.DB, error) {
	resp := database.Database.Create(a)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}

func (a *Artwork) Get() (*gorm.DB, error) {
	var result Account
	resp := database.Database.First(&result, a.CustomGormModel.ID)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}

func (a *Artwork) Delete() (*gorm.DB, error) {
	var result Artwork
	resp := database.Database.First(&result, a.CustomGormModel.ID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	resp = database.Database.Delete(a)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}
