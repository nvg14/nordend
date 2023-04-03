package schema

import (
	"github.com/amagimedia/nordend/internal/database"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Usera struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number  string
	UseraID uint
}

type Channel struct {
	CustomGormModel
	Channelid        string         `gorm:"uniqueIndex,size:10" json:"channelid,omitempty" mapstructure:"channelid,omitempty"`
	Amgid            string         `json:"amgid,omitempty" mapstructure:"amgid,omitempty"`
	Artworks         datatypes.JSON `json:"artworks,omitempty" mapstructure:"artworks,omitempty"`
	Audios           datatypes.JSON `json:"audios,omitempty" mapstructure:"audios,omitempty"`
	Caption          datatypes.JSON `json:"caption,omitempty" mapstructure:"caption,omitempty"`
	Genre            string         `json:"genre,omitempty" mapstructure:"genre,omitempty"`
	LongDescription  string         `json:"long_description,omitempty" mapstructure:"long_description,omitempty"`
	Name             string         `json:"name,omitempty" mapstructure:"name,omitempty"`
	Rating           datatypes.JSON `json:"rating,omitempty" mapstructure:"rating,omitempty"`
	ShortDescription string         `json:"short_description,omitempty" mapstructure:"short_description,omitempty"`
	StartTime        string         `json:"start_time,omitempty" mapstructure:"start_time,omitempty"`
	Tags             datatypes.JSON `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	Timezone         string         `json:"timezone,omitempty" mapstructure:"timezone,omitempty"`
	AccountID        uint
}

func (c *Channel) Save() (Channel, error) {
	resp := database.Database.Create(c)
	if resp.Error != nil {
		return Channel{}, resp.Error
	}

	channel, err := c.Get()
	if err != nil {
		return Channel{}, err
	}

	return channel, nil
}

func (c *Channel) Get() (Channel, error) {
	var result Channel
	resp := database.Database.Where("channelid = ?", c.Channelid).First(&result)
	if resp.Error != nil {
		return Channel{}, resp.Error
	}
	return result, nil
}

func (c *Channel) List() ([]Channel, error) {
	var channels = []Channel{}
	resp := database.Database.Find(&channels)
	if resp.Error != nil {
		return []Channel{}, resp.Error
	}

	return channels, nil
}

func (c *Channel) Update() (Channel, error) {
	resp := database.Database.Where("channelid = ?", c.Channelid).Updates(c)
	if resp.Error != nil {
		return Channel{}, resp.Error
	}

	channel, err := c.Get()
	if err != nil {
		return Channel{}, err
	}

	return channel, nil
}

func (a *Channel) Delete() (*gorm.DB, error) {
	var c Channel
	resp := database.Database.Where("channelid = ?", c.Channelid).Updates(c)
	if resp.Error != nil {
		return nil, resp.Error
	}

	resp = database.Database.Delete(a)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp, nil
}
