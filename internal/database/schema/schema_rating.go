package schema

import (
	"github.com/amagimedia/nordend/internal/database"
	"gorm.io/gorm"
)

type Rating struct {
	CustomGormModel
	Body      string `json:"body,omitempty" mapstructure:"body,omitempty"`
	Value     string `json:"value,omitempty" mapstructure:"value,omitempty"`
	ChannelID uint
}

func (r *Rating) Save() (*gorm.DB, error) {
	resp := database.Database.Create(r)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}

func (r *Rating) Get() (*gorm.DB, error) {
	var result Account
	resp := database.Database.First(&result, r.CustomGormModel.ID)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}

func (a *Rating) Delete() (*gorm.DB, error) {
	var result Rating
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
