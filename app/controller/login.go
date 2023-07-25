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
			"Message": "Nama Pengguna atau Kata Sandi salah",
		})
	}

	// Membandingkan kata sandi yang diberikan dengan kata sandi yang telah di-hash
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.PasswordHash))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Nama Pengguna atau Kata Sandi salah",
		})
	}

	// Membuat sesi untuk menandakan bahwa pengguna sudah login
	session, _ := sess.Get(c)
	session.Set("isLoggedIn", true)
	session.Set("username", data.Username)
	session.Save()

	return c.Render("dashboard/dist/index", fiber.Map{
		"Username": data.Username,
	})
}

// Logout mengelola proses logout pengguna dan menghapus sesi
func Logout(c *fiber.Ctx) error {
	session, _ := sess.Get(c)
	session.Delete("isLoggedIn")
	session.Delete("username")
	session.Save()

	return c.JSON(fiber.Map{
		"Message": "Logout berhasil",
	})
}

// AuthMiddleware adalah fungsi middleware untuk melindungi rute-rute yang membutuhkan otentikasi
func AuthMiddleware(c *fiber.Ctx) error {
	session, _ := sess.Get(c)

	// Memeriksa apakah pengguna sudah login
	isLoggedIn := session.Get("isLoggedIn")
	if isLoggedIn == nil || isLoggedIn.(bool) == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Anda harus login untuk mengakses rute ini",
		})
	}

	// Lanjut ke middleware berikutnya jika pengguna sudah login
	return c.Next()
}

// ProtectedRoute adalah contoh rute yang dilindungi dan memerlukan otentikasi
func ProtectedRoute(c *fiber.Ctx) error {
	// Logika rute yang dilindungi di sini
	return c.JSON(fiber.Map{
		"Message": "Ini adalah rute dilindungi, dan hanya pengguna yang sudah login yang dapat mengaksesnya.",
	})
}
