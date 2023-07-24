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

func GetAllSkill(c *fiber.Ctx) error {
	data := []model.Skill{}
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
	result := services.DB.Db.Select("skills.id, skills.name, skill_categories.name as category").Joins("JOIN skill_categories ON skills.category_id = skill_categories.id")
	if search != "" {
		result = result.Where("skills.name ILIKE ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)
	var total int64
	if search != "" {
		services.DB.Db.Where("skills.name ILIKE ?", "%"+search+"%").Select("skills.id, skills.name, skill_categories.name as category").Joins("JOIN skill_categories ON skills.category_id = skill_categories.id").Count(&total)
	} else {
		services.DB.Db.Model(&model.Skill{}).Select("skills.id, skills.name, skill_categories.name as category").Joins("JOIN skill_categories ON skills.category_id = skill_categories.id").Count(&total)
		//SELECT skills.id, skills.name, skill_categories.name as category FROM skills JOIN skill_categories ON skills.kategori_id = skill_categories.id
	}
	// Perbaikan paginasi halaman
	currentPage := int(page)
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	prevPage := currentPage - 1
	nextPage := currentPage + 1

	if prevPage < 1 {
		prevPage = 1
	}

	if nextPage > totalPages {
		nextPage = totalPages
	}

	const maxPagesToShow = 5
	startPage := currentPage - maxPagesToShow/2
	endPage := currentPage + maxPagesToShow/2

	if startPage < 1 {
		endPage = endPage + (1 - startPage)
		startPage = 1
	}

	if endPage > totalPages {
		startPage = startPage - (endPage - totalPages)
		endPage = totalPages
	}

	var pages []int
	for i := startPage; i <= endPage; i++ {
		pages = append(pages, i)
	}

	return c.Render("skill/index_skill", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
		"PrevPage":   prevPage,
		"NextPage":   nextPage,
		"Pages":      pages,
		"Search":     search,
	})
}

func GetSkill(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Skill{}
	payload := []model.SkillCategory{}
	result := services.DB.Db.Select("skills.id, skills.name, skill_categories.name as category").Joins("JOIN skill_categories ON skills.category_id = skill_categories.id").Where("skills.id = ?", id).First(&data)
	services.DB.Db.Find(&payload)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("skill/update_skill", fiber.Map{
		"Data":    data,
		"Payload": payload,
	})
}

func CreateSkillView(c *fiber.Ctx) error {
	data := []model.SkillCategory{}
	services.DB.Db.Find(&data)
	return c.Render("skill/create_skill", fiber.Map{
		"Data": data,
	})
}

func CreateSkill(c *fiber.Ctx) error {
	data := new(model.Skill)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Name == "" {
		return c.Render("", fiber.Map{"Error": "Name is required"})
	}
	result := services.DB.Db.Create(&data)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Redirect("/skill")
}

func UpdateSkill(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.Skill)
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

func DeleteSkill(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Skill{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
