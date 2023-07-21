package controller

import (
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllCountry(c *fiber.Ctx) error {
	data := []model.Country{}
	search := c.Query("search")
	perPage, err := strconv.Atoi(c.Query("perPage", "20"))
	if err != nil {
		return c.SendString("perPage harus angka")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendString("page harus angka")
	}
	offset := (page - 1) * perPage
	result := services.DB.Db
	if search != "" {
		result = result.Where("name ILIKE ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)

	var total int64

	if search != "" {
		services.DB.Db.Where("country_name ILIKE ?", "%"+search+"%").Count(&total)
	} else {
		services.DB.Db.Model(&model.Country{}).Count(&total)
	}
	return c.Render("country/index_country", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func GetCountry(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Country{}
	result := services.DB.Db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("country/update_country", fiber.Map{
		"Data": data,
	})
}

func CreateCountryView(c *fiber.Ctx) error {
	return c.Render("country/create_country", fiber.Map{})
}

func CreateCountry(c *fiber.Ctx) error {
	data := new(model.Country)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Code == "" {
		return c.Render("", fiber.Map{"Error": "Country Code is required"})
	}
	if data.Name == "" {
		return c.Render("", fiber.Map{"Error": "Country Name is required"})
	}

	services.DB.Db.Create(&data)
	return c.Redirect("/country")
}

func UpdateCountry(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.Country)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Code == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Country Code is required"})
	}
	if data.Name == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Country Name is required"})
	}

	result := services.DB.Db.Model(&data).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": result.Error,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"Message": "Success",
	})
}

func DeleteCountry(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Country{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
