package controller

import (
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllIndustry(c *fiber.Ctx) error {
	data := []model.Industry{}
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
		result = result.Where("name like ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)

	var total int64

	if search != "" {
		services.DB.Db.Where("name like ?", "%"+search+"%").Count(&total)
	} else {
		services.DB.Db.Model(&model.Industry{}).Count(&total)
	}
	return c.Render("industry/index_industry", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func GetIndustry(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Industry{}
	result := services.DB.Db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("industry/update_industry", fiber.Map{
		"Data": data,
	})
}

func CreateIndustryView(c *fiber.Ctx) error {
	return c.Render("industry/create_industry", fiber.Map{})
}

func CreateIndustry(c *fiber.Ctx) error {
	data := new(model.Industry)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Name == "" {
		return c.Render("", fiber.Map{"Error": "name is required"})
	}

	services.DB.Db.Create(&data)
	return c.Redirect("/industry")
}

func UpdateIndustry(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.Industry)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Name == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "name is required"})
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

func DeleteIndustry(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Industry{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
