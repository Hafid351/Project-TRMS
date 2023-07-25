package controller

import (
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllCompany(c *fiber.Ctx) error {
	data := []model.Company{}
	search := c.Query("search")
	perPage, err := strconv.Atoi(c.Query("perPage", "20"))
	if err != nil {
		return c.SendString("perPage harus angka")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendString("page harus angka")
	}
	if page < 1 {
		page = 1
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
		services.DB.Db.Model(&model.Company{}).Count(&total)
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
	
	return c.Render("company/index_company", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       currentPage,
		"TotalPages": totalPages,
		"PerPage":    perPage,
		"PrevPage":   currentPage - 1,
		"NextPage":   currentPage + 1,
		"Pages":      pages,
		"Search":     search,
	})
}

func GetCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Company{}
	industry := []model.Industry{}
	country := []model.Country{}
	province := []model.Province{}
	result := services.DB.Db.Select("companies.id, companies.name, companies.company_description, industries.name as industry, companies.website, companies.phone_number, companies.email, companies.working_hour, companies.benefit, companies.language_use, countries.name as country, provinces.name as province, cities.name as city, companies.address").Joins("FULL JOIN industries ON companies.industry_id = industries.id").Joins("FULL JOIN countries ON companies.country_id = countries.id").Joins("FULL JOIN provinces ON companies.province_id = provinces.id").Joins("FULL JOIN cities ON companies.city_id = cities.id").Where("companies.id = ?", id).First(&data)
	services.DB.Db.Find(&industry)
	services.DB.Db.Find(&country)
	services.DB.Db.Find(&province)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("company/update_company", fiber.Map{
		"Data":     data,
		"Industry": industry,
		"Country":  country,
		"Province": province,
	})
}

func CreateCompanyView(c *fiber.Ctx) error {
	industry := []model.Industry{}
	country := []model.Country{}
	province := []model.Province{}
	services.DB.Db.Find(&industry)
	services.DB.Db.Find(&country)
	services.DB.Db.Find(&province)
	return c.Render("company/create_company", fiber.Map{
		"Industry": industry,
		"Country":  country,
		"Province": province,
	})
}

func CreateCompany(c *fiber.Ctx) error {
	data := new(model.Company)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Name == "" {
		return c.Render("", fiber.Map{"Error": "Name is required"})
	}

	services.DB.Db.Create(&data)
	return c.Redirect("/company")
}

func UpdateCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.Company)

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

func DeleteCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Company{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func GetCompanyView(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Company{}
	result := services.DB.Db.Select("companies.id, companies.name, companies.company_description, industries.name as industry, companies.website, companies.phone_number, companies.email, companies.working_hour, companies.benefit, companies.language_use, countries.name as country, provinces.name as province, cities.name as city, companies.address").Joins("FULL JOIN industries ON companies.industry_id = industries.id").Joins("FULL JOIN countries ON companies.country_id = countries.id").Joins("FULL JOIN provinces ON companies.province_id = provinces.id").Joins("FULL JOIN cities ON companies.city_id = cities.id").Where("companies.id = ?", id).First(&data)
	// SELECT companies.id, companies.name, companies.company_description, industries.name as industry, companies.website, companies.phone_number, companies.email, companies.working_hour, companies.benefit, companies.language_use, countries.name as country, provinces.name as province, cities.name as city, companies.address
	// FROM "companies"
	// JOIN industries ON companies.industry_id = industries.id
	// JOIN countries ON companies.country_id = countries.id
	// JOIN provinces ON companies.province_id = provinces.id
	// JOIN cities ON companies.city_id = cities.id;
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("company/view_company", fiber.Map{
		"Data": data,
	})
}

func GetCompanyCity(c *fiber.Ctx) error {
	id := c.QueryInt("provinceid")
	city := []model.City{}
	services.DB.Db.Where("province_id = ?", id).Find(&city)
	return c.Status(200).JSON(fiber.Map{
		"Message": "Success",
		"City":    city,
	})
}
