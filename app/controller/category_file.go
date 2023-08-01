package controller

import (
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategoryFile(c *fiber.Ctx) error {
	data := []model.CategoryFile{}
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
		result = result.Where("category_files.name ILIKE ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)

	var total int64

	if search != "" {
		services.DB.Db.Where("category_files.name ILIKE ?", "%"+search+"%").Count(&total)
	} else {
		services.DB.Db.Model(&model.CategoryFile{}).Count(&total)
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

	// Pengecekan agar tidak menampilkan nomor halaman minus
	if startPage < 1 {
		endPage = endPage + (1 - startPage)
		startPage = 1

		// Pengecekan lagi agar halaman akhir tidak melebihi total halaman
		if endPage > totalPages {
			endPage = totalPages
		}
	}

	if endPage > totalPages {
		startPage = startPage - (endPage - totalPages)
		endPage = totalPages

		// Pengecekan lagi agar halaman awal tidak kurang dari 1
		if startPage < 1 {
			startPage = 1
		}
	}

	var pages []int
	for i := startPage; i <= endPage; i++ {
		pages = append(pages, i)
	}

	return c.Render("categoryfile/index_categoryfile", fiber.Map{
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

func GetCategoryFile(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.CategoryFile{}
	result := services.DB.Db.Where("category_files.id = ?", id).First(&data)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("categoryfile/update_categoryfile", fiber.Map{
		"Data":    data,
	})
}

func CreateCategoryFileView(c *fiber.Ctx) error {
	data := []model.Skill{}
	services.DB.Db.Find(&data)
	return c.Render("categoryfile/create_categoryfile", fiber.Map{
		"Data": data,
	})
}

func CreateCategoryFile(c *fiber.Ctx) error {
	data := new(model.CategoryFile)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Name == "" {
		return c.Render("", fiber.Map{"Error": "Name is required"})
	}
	services.DB.Db.Create(&data)
	return c.Redirect("/categoryfile")
}

func UpdateCategoryFile(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.CategoryFile)

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

func DeleteCategoryFile(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.CategoryFile{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}