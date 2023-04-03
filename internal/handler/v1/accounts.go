package v1

import (
	"sort"

	"github.com/amagimedia/nordend/internal/database"
	"github.com/amagimedia/nordend/internal/database/schema"
	"github.com/amagimedia/nordend/internal/exceptions"
	"github.com/amagimedia/nordend/internal/models"
	"github.com/amagimedia/nordend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AccountParams struct {
	Page   int `query:"page"`
	Offset int `query:"offset"`
}

func AddAccount(c *fiber.Ctx) error {
	account := new(models.Account)

	if err := c.BodyParser(account); err != nil {
		errorResponse := exceptions.ErrorResponse{
			Message:  "internal server error",
			ErrorRes: err.Error(),
		}
		return c.Status(fiber.ErrInternalServerError.Code).JSON(errorResponse)
	}

	if err := account.Validate(); err != nil {
		errorResponse := exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		}
		return c.Status(fiber.ErrBadRequest.Code).JSON(errorResponse)
	}

	dbAccount := schema.Account{
		AMGID:       account.GetAmgid(),
		Description: account.GetDescription(),
		Logo:        account.GetLogo(),
		Name:        account.GetName(),
		Domain:      account.GetDomain(),
	}

	_, err := dbAccount.Get()
	if err == gorm.ErrRecordNotFound {
		if _, err := dbAccount.Save(); err != nil {
			// TBD: v1.24.6 does not have ErrDuplicatedKey, this was added on Mar 6th 2023, no qualified release of Gorm is released yet. Post release the error handling needs to be added.
			return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorToMap(err))
		} else {
			return c.Status(fiber.StatusCreated).JSON(models.AccountResponse{
				Amgid:       dbAccount.AMGID,
				Description: dbAccount.Description,
				Domain:      dbAccount.Domain,
				Logo:        dbAccount.Logo,
				Name:        dbAccount.Name,
			})
		}
	} else if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else {
		return c.Status(fiber.ErrConflict.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: "Record already exist",
		})
	}
}

func ListAccounts(c *fiber.Ctx) error {
	p := new(AccountParams)
	if err := c.QueryParser(p); err != nil {
		errorResponse := exceptions.ErrorResponse{
			Message: "internal server error",
		}
		return c.Status(500).JSON(errorResponse)
	}

	var accounts = []schema.Account{}

	resp := database.Database.Find(&accounts)
	if resp.Error != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: resp.Error.Error(),
		})
	}

	response := []models.AccountResponse{}

	for _, account := range accounts {
		response = append(response, models.AccountResponse{
			Amgid:       account.AMGID,
			Description: account.Description,
			Domain:      account.Domain,
			Logo:        account.Logo,
			Name:        account.Name,
		})
	}

	response = sortAccordingToAMDID("asc", response)

	if p.Offset == 0 {
		p.Offset = 20
	}

	start, end := utils.Paginate(len(response), p.Page, p.Offset)
	response = response[start:end]
	if len(response) == 0 {
		response = []models.AccountResponse{}
	}

	listResponse := models.AccountPaginatedResponse{
		Offset: float32(p.Offset),
		Page:   float32(p.Page),
		Result: response,
	}

	return c.Status(fiber.StatusOK).JSON(listResponse)
}

func GetAccount(c *fiber.Ctx) error {
	amgid := c.Params("amgid")
	dbAccount := schema.Account{
		AMGID: amgid,
	}

	dbAccount, err := dbAccount.Get()
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.ErrNotFound.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.AccountResponse{
		Amgid:       dbAccount.AMGID,
		Description: dbAccount.Description,
		Domain:      dbAccount.Domain,
		Logo:        dbAccount.Logo,
		Name:        dbAccount.Name,
	})
}

func DeleteAccount(c *fiber.Ctx) error {
	amgid := c.Params("amgid")
	dbAccount := schema.Account{
		AMGID: amgid,
	}

	dbAccount, err := dbAccount.Get()
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.ErrNotFound.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	} else if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	}

	if len(dbAccount.Channel) != 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: "Cannot delete the account as there are associated `channel` resources attached to this resource.",
		})
	}

	if _, err = dbAccount.Delete(); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(exceptions.ErrorResponse{
			ErrorRes: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func sortAccordingToAMDID(sortType string, data []models.AccountResponse) []models.AccountResponse {
	sort.Slice(data, func(i, j int) bool {
		if sortType == "asc" {
			return data[i].Amgid < data[j].Amgid
		} else {
			return data[i].Amgid > data[j].Amgid
		}
	})

	return data
}
