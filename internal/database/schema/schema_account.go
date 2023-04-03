package schema

import (
	"time"

	"github.com/amagimedia/nordend/internal/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomGormModel struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Account struct {
	CustomGormModel
	AMGID       string `gorm:"unique" json:"amgid,omitempty" mapstructure:"amgid,omitempty"`
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	Domain      string `json:"domain,omitempty" mapstructure:"domain,omitempty"`
	Logo        string `json:"logo,omitempty" mapstructure:"logo,omitempty"`
	Name        string `json:"name,omitempty" mapstructure:"name,omitempty"`
	Channel     []Channel
}

func (a *Account) Save() (*gorm.DB, error) {
	resp := database.Database.Create(a)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}

func (a *Account) Get() (Account, error) {
	var result Account
	resp := database.Database.Preload("Channel").Preload(clause.Associations).Where("amg_id = ?", a.AMGID).First(&result)
	if resp.Error != nil {
		return Account{}, resp.Error
	}
	return result, nil
}

func (a *Account) Delete() (*gorm.DB, error) {
	var result Account
	resp := database.Database.Where("amg_id = ?", a.AMGID).First(&result)
	if resp.Error != nil {
		return nil, resp.Error
	}

	resp = database.Database.Delete(a)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}
