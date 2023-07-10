package controller

import (
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllCertificateCategory(c *fiber.Ctx) error {
	data := []model.CertificateCategory{}
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
	result := services.DB.Db.Select("certificate_categories.id, certificate_categories.name, skills.name as skill, certificate_categories.files").Joins("JOIN skills ON certificate_categories.skills_id = skills.id")
	//SELECT certificate_categories.id, certificate_categories.name, skills.name as skill, certificate_categories.files FROM "certificate_categories" JOIN skills ON certificate_categories.skills_id = skills.id
	if search != "" {
		result = result.Where("certificate_categories.name like ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)

	var total int64

	if search != "" {
		services.DB.Db.Where("certificate_categories.name like ?", "%"+search+"%").Select("certificate_categories.id, certificate_categories.name, skills.name as skill, certificate_categories.files").Joins("JOIN skills ON certificate_categories.skills_id = skills.id").Count(&total)
	} else {
		services.DB.Db.Model(&model.CertificateCategory{}).Select("certificate_categories.id, certificate_categories.name, skills.name as skill, certificate_categories.files").Joins("JOIN skills ON certificate_categories.skills_id = skills.id").Count(&total)
	}
	return c.Render("certificate_category/index_certificatecategory", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func GetCertificateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.CertificateCategory{}
	payload := []model.Skill{}
	result := services.DB.Db.Select("certificate_categories.id, certificate_categories.name, skills.name as skill, certificate_categories.files").Joins("JOIN skills ON certificate_categories.skills_id = skills.id").Where("certificate_categories.id = ?", id).First(&data)
	services.DB.Db.Find(&payload)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("certificate_category/update_certificatecategory", fiber.Map{
		"Data":    data,
		"Payload": payload,
	})
}

func CreateCertificateCategoryView(c *fiber.Ctx) error {
	data := []model.Skill{}
	services.DB.Db.Find(&data)
	return c.Render("certificate_category/create_certificatecategory", fiber.Map{
		"Data": data,
	})
}

func CreateCertificateCategory(c *fiber.Ctx) error {
	data := new(model.CertificateCategory)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Name == "" {
		return c.Render("", fiber.Map{"Error": "Name is required"})
	}
	if data.Files == "" {
		return c.Render("", fiber.Map{"Error": "Files is required"})
	}
	services.DB.Db.Create(&data)
	return c.Redirect("/certificatecategory")
}

func UpdateCertificateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.CertificateCategory)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Name == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Name is required"})
	}

	if data.Files == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Files is required"})
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

func DeleteCertificateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.CertificateCategory{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
