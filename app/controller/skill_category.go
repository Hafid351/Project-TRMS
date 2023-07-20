package controller

import (
	_ "log"
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
	_ "golang.org/x/crypto/bcrypt"
)

func GetAllSkillCategory(c *fiber.Ctx) error {
	data := []model.SkillCategory{}
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
		services.DB.Db.Model(&model.SkillCategory{}).Count(&total)
	}
	return c.Render("skill_category/index_skillcategory", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func GetSkillCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.SkillCategory{}
	result := services.DB.Db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("skill_category/update_skillcategory", fiber.Map{
		"Data": data,
	})
}

func CreateSkillCategoryView(c *fiber.Ctx) error {
	return c.Render("skill_category/create_skillcategory", fiber.Map{})
}

func CreateSkillCategory(c *fiber.Ctx) error {
	data := new(model.SkillCategory)
	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if data.Name == "" {
		return c.JSON(fiber.Map{
			"Error": "Skill Category Name must be filled",
		})
	}

	services.DB.Db.Create(&data)
	return c.JSON(fiber.Map{
		"Message": "Skill Category created successfully",
	})
}

func UpdateSkillCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.SkillCategory)
	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Name == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Username is required"})
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

func DeleteSkillCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.SkillCategory{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
