package controller

import (
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func IndexLogin(c *fiber.Ctx) error {
	return c.Render("login/login", fiber.Map{
		"Title": "Login TRMS",
	})
}

func IndexTRMS(c *fiber.Ctx) error {
	return c.Render("login/login", fiber.Map{
		"Title": "TOG Recruitment",
	})
}

func Login(c *fiber.Ctx) error {
	data := new(model.User)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	var user model.User
	result := services.DB.Db.Where("username = ?", data.Username).First(&user)
	if result.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.Status(404).SendString("Username or Password is Wrong!")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.PasswordHash)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Status(404).SendString("Username or Password is Wrong!")
	}
	return c.Render("dashboard/dist/index", fiber.Map{
		"Username": data.Username,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.Render("login/login", fiber.Map{})
}
