package controller

import (
	"log"
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
	data.Password = hashAndSalt([]byte(data.Password))

	result := services.DB.Db.Where("username = ?", data.Username).Select("password_hash").First(&user)
	if result.Error != nil {
		// User not found in the database, return error response
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": "Username is Wrong!",
		})
	}
	log.Print(user.PasswordHash)
	log.Print(data.Password)
	// Compare the user's entered password with the one stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.Password)); err != nil {
		// Passwords do not match, return error response
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Password is Wrong!",
		})
	}

	// If both username and password match, redirect to the dashboard
	return c.Render("dashboard/dist/index", fiber.Map{
		"Username": data.Username,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.Render("login/login", fiber.Map{})
}
