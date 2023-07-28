package routes

import (
	"trms/app/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_"github.com/gofiber/fiber/v2/middleware/session"
)

// Membuat manajer sesi baru
// var sess = session.New()

// // AuthMiddleware adalah fungsi middleware untuk melindungi rute-rute yang membutuhkan otentikasi
// func AuthMiddleware(c *fiber.Ctx) error {
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

// 	// Lanjut ke middleware berikutnya jika pengguna sudah login
// 	return c.Next()
// }

func Handlers(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/", controller.IndexLogin)
	app.Get("/index", controller.IndexTRMS)
	app.Post("/login", controller.Login)
	app.Get("/", controller.Logout)

	// app.Group("/", AuthMiddleware).Get("/protected", controller.ProtectedRoute)

	dashboard := app.Group("dashboard")
	dashboard.Get("/", controller.DashboardView)
	dashboard.Get("/", controller.DashboardTable)
	dashboard.Get("/skill", controller.GetData)
	dashboard.Get("/skill_table", controller.GetTable)
	dashboard.Get("/departement", controller.GetData1)
	dashboard.Get("/departement_table", controller.GetTable1)
	dashboard.Get("/position", controller.GetData2)
	dashboard.Get("/position_table", controller.GetTable2)
	dashboard.Get("/profileworkexperience", controller.GetData3)
	dashboard.Get("/profileworkexperience_table", controller.GetTable3)

	user := app.Group("user")
	user.Get("/", controller.GetAllUser)
	user.Get("/create_user", controller.CreateUserView)
	user.Post("/create_user", controller.CreateUser)
	user.Delete("/:id", controller.DeleteUser)
	user.Patch("/:id", controller.UpdateUser)
	user.Get("/:id", controller.GetUser)

	skill := app.Group("skill")
	skill.Get("/", controller.GetAllSkill)
	skill.Get("/create_skill", controller.CreateSkillView)
	skill.Post("/create_skill", controller.CreateSkill)
	skill.Delete("/:id", controller.DeleteSkill)
	skill.Patch("/:id", controller.UpdateSkill)
	skill.Get("/:id", controller.GetSkill)

	skillcategory := app.Group("skillcategory")
	skillcategory.Get("/", controller.GetAllSkillCategory)
	skillcategory.Get("/create_skillcategory", controller.CreateSkillCategoryView)
	skillcategory.Post("/create_skillcategory", controller.CreateSkillCategory)
	skillcategory.Delete("/:id", controller.DeleteSkillCategory)
	skillcategory.Patch("/:id", controller.UpdateSkillCategory)
	skillcategory.Get("/:id", controller.GetSkillCategory)

	position := app.Group("position")
	position.Get("/", controller.GetAllPosition)
	position.Get("/create_position", controller.CreatePositionView)
	position.Post("/create_position", controller.CreatePosition)
	position.Delete("/:id", controller.DeletePosition)
	position.Patch("/:id", controller.UpdatePosition)
	position.Get("/:id", controller.GetPosition)

	positioncategory := app.Group("positioncategory")
	positioncategory.Get("/", controller.GetAllPositionCategory)
	positioncategory.Get("/create_positioncategory", controller.CreatePositionCategoryView)
	positioncategory.Post("/create_positioncategory", controller.CreatePositionCategory)
	positioncategory.Delete("/:id", controller.DeletePositionCategory)
	positioncategory.Patch("/:id", controller.UpdatePositionCategory)
	positioncategory.Get("/:id", controller.GetPositionCategory)

	university := app.Group("university")
	university.Get("/", controller.GetAllUniversity)
	university.Get("/create_university", controller.CreateUniversityView)
	university.Post("/create_university", controller.CreateUniversity)
	university.Delete("/:id", controller.DeleteUniversity)
	university.Patch("/:id", controller.UpdateUniversity)
	university.Get("/:id", controller.GetUniversity)

	departement := app.Group("departement")
	departement.Get("/", controller.GetAllDepartement)
	departement.Get("/create_departement", controller.CreateDepartementView)
	departement.Post("/create_departement", controller.CreateDepartement)
	departement.Delete("/:id", controller.DeleteDepartement)
	departement.Patch("/:id", controller.UpdateDepartement)
	departement.Get("/:id", controller.GetDepartement)

	company := app.Group("company")
	company.Get("/", controller.GetAllCompany)
	company.Get("/create_company", controller.CreateCompanyView)
	company.Post("/create_company", controller.CreateCompany)
	company.Delete("/:id", controller.DeleteCompany)
	company.Patch("/:id", controller.UpdateCompany)
	company.Get("/:id", controller.GetCompany)
	company.Get("/view/:id", controller.GetCompanyView)
	company.Get("/create_company/country/city", controller.GetCompanyCity)
	company.Get("/:id/country/city", controller.GetCompanyCity)

	industry := app.Group("industry")
	industry.Get("/", controller.GetAllIndustry)
	industry.Get("/create_industry", controller.CreateIndustryView)
	industry.Post("/create_industry", controller.CreateIndustry)
	industry.Delete("/:id", controller.DeleteIndustry)
	industry.Patch("/:id", controller.UpdateIndustry)
	industry.Get("/:id", controller.GetIndustry)

	country := app.Group("country")
	country.Get("/", controller.GetAllCountry)
	country.Get("/create_country", controller.CreateCountryView)
	country.Post("/create_country", controller.CreateCountry)
	country.Delete("/:id", controller.DeleteCountry)
	country.Patch("/:id", controller.UpdateCountry)
	country.Get("/:id", controller.GetCountry)

	language := app.Group("language")
	language.Get("/", controller.GetAllLanguage)
	language.Get("/create_language", controller.CreateLanguageView)
	language.Post("/create_language", controller.CreateLanguage)
	language.Delete("/:id", controller.DeleteLanguage)
	language.Patch("/:id", controller.UpdateLanguage)
	language.Get("/:id", controller.GetLanguage)

	certificatecategory := app.Group("certificatecategory")
	certificatecategory.Get("/", controller.GetAllCertificateCategory)
	certificatecategory.Get("/create_certificatecategory", controller.CreateCertificateCategoryView)
	certificatecategory.Post("/create_certificatecategory", controller.CreateCertificateCategory)
	certificatecategory.Delete("/:id", controller.DeleteCertificateCategory)
	certificatecategory.Patch("/:id", controller.UpdateCertificateCategory)
	certificatecategory.Get("/:id", controller.GetCertificateCategory)

	profile := app.Group("profile")
	profile.Get("/", controller.GetAllProfile)
	profile.Get("/profile-wizard", controller.ProfileWizardView)
	profile.Post("/profile-wizard", controller.CreateProfileWizard)
	profile.Delete("/:id", controller.DeleteProfile)
	profile.Patch("/:id", controller.UpdateProfile)
	profile.Get("/profile_update/:id", controller.GetProfile)
	profile.Get("/profile_wizard/country/city", controller.GetProfileCity)
	profile.Get("/profile_wizard/skill", controller.GetProfileSkill)
	profile.Get("/view/:id", controller.GetProfileView)
	profile.Get("/profile-wizard/qualification", controller.GetQualification)
	//profile.Post("/profile-wizard/profile", controller.CreateProfileWizardProfile)
	profile.Get("/profile-wizard/profile", controller.ProfileWizardProfileView)
	profile.Post("/profile-wizard/education", controller.CreateProfileWizardEducation)
	profile.Get("/profile-wizard/education", controller.ProfileWizardEducationView)
	profile.Post("/profile-wizard/work", controller.CreateProfileWizardWorkExperience)
	profile.Get("/profile-wizard/work", controller.ProfileWizardWorkExperienceView)
	profile.Post("/profile-wizard/language", controller.CreateProfileWizardLanguage)
	profile.Get("/profile-wizard/language", controller.ProfileWizardLanguageView)
	profile.Post("/profile-wizard/training", controller.CreateProfileWizardTraining)
	profile.Get("/profile-wizard/training", controller.ProfileWizardTrainingView)
	profile.Post("/profile-wizard/files", controller.UploadFiles)
	profile.Get("/profile-wizard/country", controller.GetProfileCountry)
}
