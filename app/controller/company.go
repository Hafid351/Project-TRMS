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
		services.DB.Db.Model(&model.Company{}).Count(&total)
	}
	return c.Render("company/index_company", fiber.Map{
		"Data":       data,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
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
