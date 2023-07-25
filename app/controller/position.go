package controller

import (
	_ "log"
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllPosition(c *fiber.Ctx) error {
	data := []model.Position{}
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
	result := services.DB.Db.Select("positions.id, positions.name, position_categories.name as category").Joins("JOIN position_categories ON positions.category_id = position_categories.id")
	if search != "" {
		result = result.Where("positions.name ILIKE ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)
	var total int64
	if search != "" {
		services.DB.Db.Where("positions.name ILIKE ?", "%"+search+"%").Select("positions.id, positions.name, position_categories.name as category").Joins("JOIN position_categories ON positions.category_id = position_categories.id").Count(&total)
	} else {
		services.DB.Db.Model(&model.Position{}).Select("positions.id, positions.name, position_categories.name as category").Joins("JOIN position_categories ON positions.category_id = position_categories.id").Count(&total)
		// SELECT positions.id, positions.name, position_categories.name as category FROM positions JOIN position_categories ON positions.category_id = position_categories.id
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


	return c.Render("position/index_position", fiber.Map{
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

func GetPosition(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Position{}
	payload := []model.PositionCategory{}
	result := services.DB.Db.Select("positions.id, positions.name, position_categories.name as category").Joins("JOIN position_categories ON positions.category_id = position_categories.id").Where("positions.id = ?", id).First(&data)
	services.DB.Db.Find(&payload)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("position/update_position", fiber.Map{
		"Data":    data,
		"Payload": payload,
	})
}

func CreatePositionView(c *fiber.Ctx) error {
	data := []model.PositionCategory{}
	services.DB.Db.Find(&data)
	return c.Render("position/create_position", fiber.Map{
		"Data": data,
	})
}

func CreatePosition(c *fiber.Ctx) error {
	data := new(model.Position)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Name == "" {
		return c.Render("", fiber.Map{"Error": "Nama Posisi is required"})
	}
	result := services.DB.Db.Create(&data)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Redirect("/position")
}

func UpdatePosition(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.Position)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Name == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Nama Posisi is required"})
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

func DeletePosition(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Position{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
