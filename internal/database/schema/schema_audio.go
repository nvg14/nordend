package schema

import (
	"github.com/amagimedia/nordend/internal/database"
	"gorm.io/gorm"
)

type Audio struct {
	CustomGormModel
	IsPrimary bool    `json:"is_primary,omitempty" mapstructure:"is_primary,omitempty"`
	Language  string  `json:"language,omitempty" mapstructure:"language,omitempty"`
	Track     float32 `json:"track,omitempty" mapstructure:"track,omitempty"`
	ChannelID uint
}

func (a *Audio) Save() (*gorm.DB, error) {
	resp := database.Database.Create(a)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}

func (a *Audio) Get() (Audio, error) {
	var result Audio
	resp := database.Database.First(&result, a.CustomGormModel.ID)
	if resp.Error != nil {
		return result, resp.Error
	}
	return result, nil
}

func (c *Audio) Update() (Audio, error) {
	var result Audio
	resp := database.Database.First(&result, c.CustomGormModel.ID).Updates(c)
	if resp.Error != nil {
		return result, resp.Error
	}

	result, err := c.Get()
	if err != nil {
		return result, err
	}

	return result, nil
}

func (a *Audio) Delete() (*gorm.DB, error) {
	var result Audio
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
