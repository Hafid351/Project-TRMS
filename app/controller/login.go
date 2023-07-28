package controller

import (
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

// Membuat manajer sesi baru
var sess = session.New()

func IndexLogin(c *fiber.Ctx) error {
	return c.Render("login/login", fiber.Map{
		"Title": "Login TRMS",
	})
}

func IndexTRMS(c *fiber.Ctx) error {
	return c.Render("login/login", fiber.Map{
		"Title": "TOG Recruitment",
	})
}

// Login mengelola otentikasi pengguna dan membuat sesi setelah login berhasil
func Login(c *fiber.Ctx) error {
	data := new(model.User)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Format permintaan tidak valid",
		})
	}

	// Memeriksa apakah nama pengguna yang diberikan ada di dalam database
	var user model.User
	result := services.DB.Db.Where("username = ?", data.Username).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Username Doesn't Exist!",
		})
	}

	// Membandingkan password yang diberikan dengan password yang telah di-hash di database
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Password Incorrect!",
		})
	}

	// Membuat sesi untuk menandakan bahwa pengguna sudah login
	session, err := sess.Get(c)
	if err != nil {
		// Tangani kesalahan saat mendapatkan sesi
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Failed to create session.",
		})
	}

	session.Set("isLoggedIn", true)
	session.Set("username", data.Username)
	if err := session.Save(); err != nil {
		// Tangani kesalahan saat menyimpan sesi
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Failed to save session.",
		})
	}

	return c.Render("dashboard/dist/index", fiber.Map{
		"Message":  "Sucess",
		"Username": data.Username,
	})
}

// Logout mengelola proses logout pengguna dan menghapus sesi
func Logout(c *fiber.Ctx) error {
	session, _ := sess.Get(c)
	session.Delete("isLoggedIn")
	session.Delete("username")
	session.Save()
	if err := session.Save(); err != nil {
		// Tangani kesalahan saat menyimpan sesi
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Failed to save session.",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Logout berhasil",
	})
}

// ProtectedRoute adalah contoh rute yang dilindungi dan memerlukan otentikasi
// func ProtectedRoute(c *fiber.Ctx) error {
// 	session, err := sess.Get(c)
// 	if err != nil {
// 		// Tangani kesalahan saat mendapatkan sesi
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"Message": "Failed to get session.",
// 		})
// 	}

// 	// Memeriksa apakah pengguna sudah login
// 	isLoggedIn := session.Get("isLoggedIn")
// 	if isLoggedIn == nil || isLoggedIn.(bool) == false {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"Message": "Anda harus login untuk mengakses rute ini",
// 		})
// 	}

// 	// Jika pengguna sudah login, Anda dapat memberikan respons dengan data yang relevan.
// 	// Misalnya, Anda dapat menampilkan pesan "Selamat datang" atau data lain yang diperlukan.
// 	username := session.Get("username")
// 	return c.JSON(fiber.Map{
// 		"Message":  "Ini adalah rute dilindungi, dan hanya pengguna yang sudah login yang dapat mengaksesnya.",
// 		"Username": username,
// 	})
// }
