package controller

import (
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllPositionCategory(c *fiber.Ctx) error {
	data := []model.PositionCategory{}
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
		result = result.Where("nama_kategori like ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)

	var total int64

	if search != "" {
		services.DB.Db.Where("nama_kategori like ?", "%"+search+"%").Count(&total)
	} else {
		services.DB.Db.Model(&model.PositionCategory{}).Count(&total)
	}
	return c.Render("position_category/index_positioncategory", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func GetPositionCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.PositionCategory{}
	result := services.DB.Db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("position_category/update_positioncategory", fiber.Map{
		"Data": data,
	})
}

func CreatePositionCategoryView(c *fiber.Ctx) error {
	return c.Render("position_category/create_positioncategory", fiber.Map{})
}

// func CreatePositionCategory(c *fiber.Ctx) error {
// 	data := new(model.PositionCategory)
// 	if err := c.BodyParser(data); err != nil {
// 		return c.Render("", fiber.Map{"Error": err.Error()})
// 	}
// 	if data.Name == "" {
// 		return c.Render("", fiber.Map{"Error": "Nama Kategori is required"})
// 	}

// 	services.DB.Db.Create(&data)
// 	return c.Redirect("/positioncategory")
// }

func CreatePositionCategory(c *fiber.Ctx) error {
	data := new(model.PositionCategory)
	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if data.Name == "" {
		return c.JSON(fiber.Map{
			"Error": "Position Category Name must be filled",
		})
	}

	services.DB.Db.Create(&data)
	return c.JSON(fiber.Map{
		"Message": "Position Category created successfully",
	})
}

func UpdatePositionCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.PositionCategory)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Name == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Nama Kategori is required"})
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

func DeletePositionCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.PositionCategory{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
