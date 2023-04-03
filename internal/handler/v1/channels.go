package v1

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"sort"
	"strconv"
	"strings"

	"github.com/amagimedia/nordend/internal/database/schema"
	"github.com/amagimedia/nordend/internal/exceptions"
	"github.com/amagimedia/nordend/internal/models"
	"github.com/amagimedia/nordend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ChannelParams struct {
	Page     int    `query:"page"`
	Offset   int    `query:"offset"`
	AMGID    string `query:amgid`
	Language string `query:language`
	Timezone string `query:timezone`
	Genre    string `query:genre`
}

func AddChannel(c *fiber.Ctx) error {
	channel := new(models.Channel)

	if err := c.BodyParser(channel); err != nil {
		errorResponse := exceptions.ErrorResponse{
			Message:  "internal server error",
			ErrorRes: err.Error(),
		}
		return c.Status(fiber.ErrInternalServerError.Code).JSON(errorResponse)
	}

	account, err := CheckIfAccountExists(channel.Amgid, c)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: "AMGID not found",
		})
	}

	if err := channel.Validate(); err != nil {
		errorResponse := exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		}
		return c.Status(fiber.ErrBadRequest.Code).JSON(errorResponse)
	}

	dbArtworks := channel.GetArtworks()
	dbAudios := channel.GetAudios()
	dbCaptions := channel.GetCaption()
	dbRatings := channel.GetRating()

	artworks, _ := json.Marshal(&dbArtworks)
	audios, _ := json.Marshal(&dbAudios)
	captions, _ := json.Marshal(&dbCaptions)
	ratings, _ := json.Marshal(&dbRatings)

	tags, _ := json.Marshal(&channel.Tags)
	dbChannel := &schema.Channel{
		Amgid:            channel.GetAmgid(),
		Channelid:        channel.GetAmgid() + "C" + strconv.Itoa(len(account.Channel)+1),
		Artworks:         datatypes.JSON(json.RawMessage(artworks)),
		Audios:           datatypes.JSON(json.RawMessage(audios)),
		Caption:          datatypes.JSON(json.RawMessage(captions)),
		Rating:           datatypes.JSON(json.RawMessage(ratings)),
		Genre:            channel.GetGenre(),
		LongDescription:  channel.GetLongDescription(),
		ShortDescription: channel.GetShortDescription(),
		Timezone:         string(channel.GetTimezone()),
		Tags:             datatypes.JSON(json.RawMessage(tags)),
		AccountID:        account.ID,
	}

	if _, err := dbChannel.Get(); err == nil {
		return c.Status(fiber.ErrConflict.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: "Channel ID already exists",
		})
	}

	if channel, err := dbChannel.Save(); err != nil {
		// TBD: v1.24.6 does not have ErrDuplicatedKey, this was added on Mar 6th 2023, no qualified release of Gorm is released yet. Post release the error handling needs to be added.
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorToMap(err))
	} else {
		return c.Status(fiber.StatusCreated).JSON(generateChannelResposne(channel))
	}
}

func ListChannels(c *fiber.Ctx) error {
	p := new(ChannelParams)
	if err := c.QueryParser(p); err != nil {
		errorResponse := exceptions.ErrorResponse{
			Message: "internal server error",
		}
		return c.Status(500).JSON(errorResponse)
	}

	total := 0
	dbChannels := schema.Channel{}
	channelResponse := []models.ChannelResponse{}

	if channels, err := dbChannels.List(); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else {
		for _, channel := range channels {
			channelResponse = append(channelResponse, createChannelResponse(channel))
		}

		total = len(channels)

		channelResponse = sortAccordingToChannelID("asc", channelResponse)
		if p.AMGID != "" {
			channelResponse = filterChannelByAmgid(strings.Split(p.AMGID, ","), channelResponse)
		}

		if p.Language != "" {
			channelResponse = filterChannelByPrimaryAudioLang(strings.Split(p.Language, ","), channelResponse)
		}

		if p.Timezone != "" {
			channelResponse = filterChannelByTimezone(strings.Split(p.Timezone, ","), channelResponse)
		}

		if p.Genre != "" {
			channelResponse = filterChannelByGenre(strings.Split(p.Genre, ","), channelResponse)
		}

	}

	if p.Offset == 0 {
		p.Offset = 10
	}

	start, end := utils.Paginate(len(channelResponse), p.Page, p.Offset)
	if start > len(channelResponse) || end > len(channelResponse) {
		return c.Status(fiber.StatusOK).JSON(models.ChannelPaginatedResponse{
			Total:  float32(total),
			Offset: float32(p.Offset),
			Page:   float32(p.Page),
			Result: []models.ChannelResponse{},
		})
	}
	channelResponse = channelResponse[start:end]
	if len(channelResponse) == 0 {
		channelResponse = []models.ChannelResponse{}
	}

	listResponse := models.ChannelPaginatedResponse{
		Total:  float32(total),
		Offset: float32(p.Offset),
		Page:   float32(p.Page),
		Result: channelResponse,
	}

	return c.Status(fiber.StatusOK).JSON(listResponse)
}

func FetchChannel(c *fiber.Ctx) error {
	dbChannels := schema.Channel{
		Channelid: c.Params("channelid"),
	}

	if channel, err := dbChannels.Get(); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(generateChannelResposne(channel))
	}
}

// func PatchChannel(c *fiber.Ctx) error {
// 	// channelPatch := new(models.ChannelPatch)

// 	// dbChannels := schema.Channel{
// 	// 	Channelid: c.Params("channelid"),
// 	// }

// 	// if err := c.BodyParser(channelPatch); err != nil {
// 	// 	errorResponse := exceptions.ErrorResponse{
// 	// 		Message:  "internal server error",
// 	// 		ErrorRes: err.Error(),
// 	// 	}
// 	// 	return c.Status(fiber.ErrInternalServerError.Code).JSON(errorResponse)
// 	// }

// 	// if channel, err := dbChannels.Get(); err != nil {
// 	// 	return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
// 	// 		ErrorRes: err.Error(),
// 	// 	})
// 	// } else {
// 	// 	switch channelPatch.Op {
// 	// 	case "add":

// 	// 	case "remove":
// 	// 	case "replace":
// 	// 	case "move":
// 	// 	case "copy":
// 	// 	default:
// 	// 		return c.Status(fiber.ErrBadRequest.Code).JSON(exceptions.ErrorResponse{
// 	// 			ErrorRes: "Invalid op in the request body",
// 	// 		})
// 	// 	}
// 	// 	return nil
// 	// }
// 	return nil
// }

func DeleteChannel(c *fiber.Ctx) error {
	dbChannels := schema.Channel{
		Channelid: c.Params("channelid"),
	}

	if channel, err := dbChannels.Get(); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else {
		if _, err = channel.Delete(); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
				ErrorRes: err.Error(),
			})
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

func UpdateCaptions(c *fiber.Ctx) error {
	caption := []models.CaptionObject{}

	if err := c.BodyParser(&caption); err != nil {
		errorResponse := exceptions.ErrorResponse{
			Message:  "internal server error",
			ErrorRes: err.Error(),
		}
		return c.Status(fiber.ErrInternalServerError.Code).JSON(errorResponse)
	}

	dbChannels := schema.Channel{
		Channelid: c.Params("channelid"),
	}

	_, err := dbChannels.Get()
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.ErrNotFound.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	}

	dbCaptions, _ := json.Marshal(&caption)
	dbChannels.Caption = datatypes.JSON(json.RawMessage(dbCaptions))

	if channel, err := dbChannels.Update(); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(generateChannelResposne(channel))
	}
}

func UpdateRatings(c *fiber.Ctx) error {
	ratings := []models.Rating{}

	if err := c.BodyParser(&ratings); err != nil {
		errorResponse := exceptions.ErrorResponse{
			Message:  "internal server error",
			ErrorRes: err.Error(),
		}
		return c.Status(fiber.ErrInternalServerError.Code).JSON(errorResponse)
	}

	dbChannels := schema.Channel{
		Channelid: c.Params("channelid"),
	}

	_, err := dbChannels.Get()
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.ErrNotFound.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	}

	dbRatings, _ := json.Marshal(&ratings)

	dbChannels.Rating = datatypes.JSON(json.RawMessage(dbRatings))
	if channel, err := dbChannels.Update(); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(generateChannelResposne(channel))
	}
}

func UpdateAudios(c *fiber.Ctx) error {
	audios := []models.Audio{}

	if err := c.BodyParser(&audios); err != nil {
		errorResponse := exceptions.ErrorResponse{
			Message:  "internal server error",
			ErrorRes: err.Error(),
		}
		return c.Status(fiber.ErrInternalServerError.Code).JSON(errorResponse)
	}

	dbChannels := schema.Channel{
		Channelid: c.Params("channelid"),
	}

	_, err := dbChannels.Get()
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.ErrNotFound.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	}

	dbAudios, _ := json.Marshal(&audios)
	dbChannels.Audios = datatypes.JSON(json.RawMessage(dbAudios))

	if channel, err := dbChannels.Update(); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(generateChannelResposne(channel))
	}
}

func UpdateArtworks(c *fiber.Ctx) error {
	artworks := []models.Artwork{}

	if err := c.BodyParser(&artworks); err != nil {
		errorResponse := exceptions.ErrorResponse{
			Message:  "internal server error",
			ErrorRes: err.Error(),
		}
		return c.Status(fiber.ErrInternalServerError.Code).JSON(errorResponse)
	}

	dbChannels := schema.Channel{
		Channelid: c.Params("channelid"),
	}

	_, err := dbChannels.Get()
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.ErrNotFound.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	}

	dbArtworks, _ := json.Marshal(&artworks)
	dbChannels.Artworks = datatypes.JSON(json.RawMessage(dbArtworks))

	if channel, err := dbChannels.Update(); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(generateChannelResposne(channel))
	}
}

func generateChannelResposne(channel schema.Channel) models.ChannelResponse {

	tagBody, _ := json.Marshal(&channel.Tags)
	tags := make(map[string]interface{})
	json.Unmarshal(tagBody, &tags)

	artworksBody, _ := json.Marshal(&channel.Artworks)
	artworks := []models.Artwork{}
	json.Unmarshal(artworksBody, &artworks)

	audiosBody, _ := json.Marshal(&channel.Audios)
	audios := []models.Audio{}
	json.Unmarshal(audiosBody, &audios)

	captionsBody, _ := json.Marshal(&channel.Caption)
	captions := []models.CaptionObject{}
	json.Unmarshal(captionsBody, &captions)

	ratingsBody, _ := json.Marshal(&channel.Rating)
	ratings := []models.Rating{}
	json.Unmarshal(ratingsBody, &ratings)

	return models.ChannelResponse{
		Amgid:            channel.Amgid,
		Channelid:        channel.Channelid,
		Genre:            channel.Genre,
		LongDescription:  channel.LongDescription,
		ShortDescription: channel.ShortDescription,
		Timezone:         models.Timezone(channel.Timezone),
		Tags:             tags,
		Artworks:         artworks,
		Audios:           audios,
		Caption:          captions,
		Rating:           ratings,
	}
}

func CheckIfAccountExists(amgid string, c *fiber.Ctx) (schema.Account, error) {
	dbAccount := schema.Account{
		AMGID: amgid,
	}

	dbAccount, err := dbAccount.Get()
	if err == gorm.ErrRecordNotFound {
		return schema.Account{}, err
	}

	return dbAccount, nil
}

func GetBytes(key map[string]interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(key)
	return buf.Bytes()
}

func sortAccordingToChannelID(sortType string, data []models.ChannelResponse) []models.ChannelResponse {
	sort.Slice(data, func(i, j int) bool {
		if sortType == "asc" {
			return data[i].Channelid < data[j].Channelid
		} else {
			return data[i].Channelid > data[j].Channelid
		}
	})

	return data
}

func filterChannelByAmgid(amgids []string, channels []models.ChannelResponse) []models.ChannelResponse {
	filteredChannels := make(map[string]models.ChannelResponse)
	for _, amgid := range amgids {
		for _, channel := range channels {
			if channel.Amgid == amgid {
				filteredChannels[channel.Channelid] = channel
			}
		}
	}

	return generateFilteredChannels(filteredChannels)
}

func filterChannelByPrimaryAudioLang(languages []string, channels []models.ChannelResponse) []models.ChannelResponse {
	filteredChannels := make(map[string]models.ChannelResponse)
	for _, language := range languages {
		for _, channel := range channels {
			for _, audio := range channel.Audios {
				if audio.IsPrimary {
					if audio.Language == models.AudioLanguages(language) {
						filteredChannels[channel.Channelid] = channel
					}
				}
			}
		}
	}
	return generateFilteredChannels(filteredChannels)
}

func filterChannelByTimezone(timezones []string, channels []models.ChannelResponse) []models.ChannelResponse {
	filteredChannels := make(map[string]models.ChannelResponse)
	for _, timezone := range timezones {
		for _, channel := range channels {
			if channel.Timezone == models.Timezone(timezone) {
				filteredChannels[channel.Channelid] = channel
			}
		}
	}
	return generateFilteredChannels(filteredChannels)
}

func filterChannelByGenre(genres []string, channels []models.ChannelResponse) []models.ChannelResponse {
	filteredChannels := make(map[string]models.ChannelResponse)
	for _, genre := range genres {
		for _, channel := range channels {
			for _, channelGenre := range strings.Split(channel.Genre, ",") {
				if channelGenre == genre {
					filteredChannels[channel.Channelid] = channel
				}
			}
		}
	}
	return generateFilteredChannels(filteredChannels)
}

func generateFilteredChannels(filteredChannels map[string]models.ChannelResponse) []models.ChannelResponse {
	channels := []models.ChannelResponse{}
	for _, channel := range filteredChannels {
		channels = append(channels, channel)
	}
	return sortAccordingToChannelID("asc", channels)
}

func createChannelResponse(channel schema.Channel) models.ChannelResponse {

	tagBody, _ := json.Marshal(&channel.Tags)
	tags := make(map[string]interface{})
	json.Unmarshal(tagBody, &tags)

	artworksBody, _ := json.Marshal(&channel.Artworks)
	artworks := []models.Artwork{}
	json.Unmarshal(artworksBody, &artworks)

	audiosBody, _ := json.Marshal(&channel.Audios)
	audios := []models.Audio{}
	json.Unmarshal(audiosBody, &audios)

	captionsBody, _ := json.Marshal(&channel.Caption)
	captions := []models.CaptionObject{}
	json.Unmarshal(captionsBody, &captions)

	ratingsBody, _ := json.Marshal(&channel.Rating)
	ratings := []models.Rating{}
	json.Unmarshal(ratingsBody, &ratings)

	return models.ChannelResponse{
		Amgid:            channel.Amgid,
		Channelid:        channel.Channelid,
		Genre:            channel.Genre,
		LongDescription:  channel.LongDescription,
		ShortDescription: channel.ShortDescription,
		Timezone:         models.Timezone(channel.Timezone),
		Tags:             tags,
		Artworks:         artworks,
		Audios:           audios,
		Caption:          captions,
		Rating:           ratings,
	}
}
