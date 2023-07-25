package controller

import (
	"log"
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GetAllUser(c *fiber.Ctx) error {
	data := []model.User{}
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
	result := services.DB.Db.Select("users.id, users.username, users.fullname, users.email, roles.name AS role").Joins("JOIN roles ON users.roles_id = roles.id")
	//SELECT users.id, users.username, users.fullname, users.email, roles.name AS role FROM "users" JOIN roles ON users.roles_id = roles.id
	if search != "" {
		result = result.Where("username ILIKE ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)
	var total int64
	if search != "" {
		services.DB.Db.Where("username ILIKE ?", "%"+search+"%").Select("users.id, users.username, users.fullname, users.email, roles.name AS role").Joins("JOIN roles ON users.roles_id = roles.id").Count(&total)
	} else {
		services.DB.Db.Model(&model.User{}).Select("users.id, users.username, users.fullname, users.email, roles.name AS role").Joins("JOIN roles ON users.roles_id = roles.id").Count(&total)
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

	return c.Render("user/index_user", fiber.Map{
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

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.User{}
	payload := []model.Role{}
	result := services.DB.Db.Select("users.id, users.username, users.fullname, users.email, roles.name AS role").Joins("JOIN roles ON users.roles_id = roles.id").Where("users.id = ?", id).First(&data)
	services.DB.Db.Find(&payload)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("user/update_user", fiber.Map{
		"Data":    data,
		"Payload": payload,
	})
}

func CreateUserView(c *fiber.Ctx) error {
	data := []model.Role{}
	services.DB.Db.Find(&data)
	return c.Render("user/create_user", fiber.Map{
		"Data": data,
	})
}

func CreateUser(c *fiber.Ctx) error {
	data := new(model.User)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Username == "" {
		return c.Render("", fiber.Map{"Error": "Username is Required!"})
	}
	data.PasswordHash = hashAndSalt([]byte(data.PasswordHash))
	services.DB.Db.Create(&data)
	return c.Redirect("/user")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.User)
	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Username == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Username is Required!"})
	}
	data.PasswordHash = hashAndSalt([]byte(data.PasswordHash))
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

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.User{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"Message": "Success",
	})
}
