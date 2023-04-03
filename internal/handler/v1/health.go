package v1

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var Lggr *zap.SugaredLogger

type HealthMessage struct {
	Status string `json:"status"`
}

func Health(c *fiber.Ctx) error {
	return c.Status(200).JSON(HealthMessage{
		Status: "healthy",
	})
}
