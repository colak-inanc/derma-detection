package controllers

import (
	"VeriMadenciligi/database"
	"VeriMadenciligi/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {

	var users []models.User

	database.DB.Preload("Role").Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
