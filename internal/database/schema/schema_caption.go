package schema

import (
	"github.com/amagimedia/nordend/internal/database"
	"gorm.io/gorm"
)

type Caption struct {
	CustomGormModel
	Format    string `json:"format,omitempty" mapstructure:"format,omitempty"`
	Languages string `json:"language,omitempty" mapstructure:"language,omitempty"`
	ChannelID uint
}

func (c *Caption) Save() (*gorm.DB, error) {
	resp := database.Database.Create(c)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}

func (c *Caption) Get() (*gorm.DB, error) {
	var result Account
	resp := database.Database.First(&result, c.CustomGormModel.ID)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}

func (a *Caption) Delete() (*gorm.DB, error) {
	var result Caption
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
