package controller

import (
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllDepartement(c *fiber.Ctx) error {
	data := []model.Departement{}
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
		services.DB.Db.Where("name ILIKE ?", "%"+search+"%").Count(&total)
	} else {
		services.DB.Db.Model(&model.Departement{}).Count(&total)
	}
	return c.Render("departement/index_departement", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func GetDepartement(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Departement{}
	result := services.DB.Db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("departement/update_departement", fiber.Map{
		"Data": data,
	})
}

func CreateDepartementView(c *fiber.Ctx) error {
	return c.Render("departement/create_departement", fiber.Map{})
}

func CreateDepartement(c *fiber.Ctx) error {
	data := new(model.Departement)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Name == "" {
		return c.Render("", fiber.Map{"Error": "Name is required"})
	}

	services.DB.Db.Create(&data)
	return c.Redirect("/departement")
}

func UpdateDepartement(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.Departement)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Name == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Name is required"})
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

func DeleteDepartement(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Departement{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
