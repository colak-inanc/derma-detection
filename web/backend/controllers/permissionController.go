package controllers

import (
	"VeriMadenciligi/database"
	"VeriMadenciligi/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
