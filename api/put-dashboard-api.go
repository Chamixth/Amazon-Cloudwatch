package api

import (
	"amazon-cloudwatch/dao"
	"amazon-cloudwatch/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PutDashboard(c *fiber.Ctx) error {
	var application dto.Dashboard

	if err := c.BodyParser(&application); err != nil {
		return err
	}

	result, err := dao.PutDashboard(application)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
