package controller

import (
	"fmt"
	_ "log"
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
	_ "golang.org/x/crypto/bcrypt"
)

func GetAllProfile(c *fiber.Ctx) error {
	data := []model.Profile{}
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
		result = result.Where("fullname ILIKE ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)
	var total int64
	if search != "" {
		services.DB.Db.Where("fullname ILIKE ?", "%"+search+"%").Count(&total)
	} else {
		services.DB.Db.Model(&model.Profile{}).Count(&total)
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

	return c.Render("profile/index_profile", fiber.Map{
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

func GetProfileView(c *fiber.Ctx) error {
	id := c.Params("id")
	profile := model.Profile{}
	profileeducation := model.ProfileEducation{}
	profileworkexperience := model.ProfileWorkExperience{}
	profilelanguage := model.ProfileTraining{}
	profiletraining := model.ProfileTraining{}
	profilefiles := model.ProfileFile{}
	result := services.DB.Db.Where("profile.id = ?", id).First(&profile)
	services.DB.Db.Where("profile.id = ?", id).First(&profileeducation)
	services.DB.Db.Where("profile.id = ?", id).First(&profileworkexperience)
	services.DB.Db.Where("profile.id = ?", id).First(&profilelanguage)
	services.DB.Db.Where("profile.id = ?", id).First(&profiletraining)
	services.DB.Db.Where("profile.id = ?", id).First(&profilefiles)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("profile/view_profile", fiber.Map{
		"Data":                  profile,
		"ProfileEducation":      profileeducation,
		"ProfileWorkExperience": profileworkexperience,
		"ProfileLanguage":       profilelanguage,
		"ProfileTraining":       profiletraining,
		"ProfileFiles":          profilefiles,
	})
}

func GetProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Profile{}
	profile := []model.Profile{}
	country := []model.Country{}
	province := []model.Province{}
	religion := []model.Religion{}
	status := []model.MaritalStatus{}
	result := services.DB.Db.Select("skills.id, skills.name, skill_categories.name as category").Joins("JOIN skill_categories ON skills.category_id = skill_categories.id").Where("skills.id = ?", id).First(&data)
	services.DB.Db.Find(&profile)
	services.DB.Db.Find(&data)
	services.DB.Db.Find(&country)
	services.DB.Db.Find(&province)
	services.DB.Db.Find(&religion)
	services.DB.Db.Find(&status)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("profile/update_profile", fiber.Map{
		"Data":     data,
		"Profile":  profile,
		"Country":  country,
		"Province": province,
		"Religion": religion,
		"Status":   status,
	})
}

func ProfileWizardView(c *fiber.Ctx) error {
	// db := services.DB
	step := c.QueryInt("step", 0)
	switch step {
	case 0:
		data := []model.Profile{}
		country := []model.Country{}
		province := []model.Province{}
		religion := []model.Religion{}
		status := []model.MaritalStatus{}
		identification := []model.Identification{}
		position := []model.Position{}
		// db.Find(&data)
		services.DB.Db.Find(&data)
		services.DB.Db.Find(&country)
		services.DB.Db.Find(&province)
		services.DB.Db.Find(&religion)
		services.DB.Db.Find(&status)
		services.DB.Db.Find(&identification)
		services.DB.Db.Find(&position)
		return c.Render("profile/profile_wizard", fiber.Map{
			"Data":           data,
			"Country":        country,
			"Province":       province,
			"Religion":       religion,
			"Status":         status,
			"Identification": identification,
			"Position":       position,
		})
	case 1:
		data := []model.ProfileEducation{}
		qualification := []model.Qualification{}
		services.DB.Db.Find(&data)
		services.DB.Db.Order("id ASC").Find(&qualification)
		return c.Status(200).JSON(fiber.Map{
			"Message":       "Success",
			"Data":          data,
			"Qualification": qualification,
		})
	case 2:
		data := []model.ProfileWorkExperience{}
		company := []model.Company{}
		country := []model.Country{}
		province := []model.Province{}
		positionlevel := []model.JobPositionLevel{}
		position := []model.Position{}
		skillcategory := []model.SkillCategory{}
		skill := []model.Skill{}
		services.DB.Db.Find(&data)
		services.DB.Db.Select("id, name").Limit(100).Find(&company)
		services.DB.Db.Find(&country)
		services.DB.Db.Find(&province)
		services.DB.Db.Find(&positionlevel)
		services.DB.Db.Find(&position)
		services.DB.Db.Find(&skillcategory)
		services.DB.Db.Find(&skill)
		return c.Status(200).JSON(fiber.Map{
			"Message":       "Success",
			"Data":          data,
			"Company":       company,
			"Country":       country,
			"Province":      province,
			"PositionLevel": positionlevel,
			"Position":      position,
			"SkillCategory": skillcategory,
			"Skill":         skill,
		})
	case 3:
		data := []model.ProfileLanguage{}
		language := []model.Language{}
		languagelevel := []model.LanguageLevel{}
		services.DB.Db.Find(&data)
		services.DB.Db.Find(&language)
		services.DB.Db.Find(&languagelevel)
		return c.Status(200).JSON(fiber.Map{
			"Message":       "Success",
			"Data":          data,
			"Language":      language,
			"LanguageLevel": languagelevel,
		})
	case 4:
		data := []model.ProfileTraining{}
		services.DB.Db.Find(&data)
		return c.Status(200).JSON(fiber.Map{
			"Message": "Success",
			"Data":    data,
		})
	case 5:
		data := []model.ProfileFile{}
		services.DB.Db.Find(&data)
		return c.Status(200).JSON(fiber.Map{
			"Message": "Success",
			"Data":    data,
		})
	default:
		return c.Status(500).SendString("Error")
	}
}

func CreateProfileWizard(c *fiber.Ctx) error {
	step := c.QueryInt("step")
	switch step {
	case 1:
		data := new(model.Profile)
		if err := c.BodyParser(data); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": err.Error(),
			})
		}
		if data.Email == "" {
			return c.Status(500).JSON(fiber.Map{
				"Data": "Email Cannot be Empty!",
			})
		}
		var profile model.Profile
		if data.Email == profile.Email {
			return c.Status(500).JSON(fiber.Map{
				"Data": "Email Already Exists!",
			})
		}
		result := services.DB.Db.Create(&data)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": "Create Data Error!",
			})
		}
		fmt.Printf("User ID : %d\n", data.ID)
		return c.Status(200).JSON(fiber.Map{
			"Data":      "Ok",
			"Profileid": data.ID,
		})
	case 2:
		data := new(model.ProfileEducation)
		if err := c.BodyParser(data); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": err.Error(),
			})
		}
		if data.ProfileId == 0 {
			return c.Render("", fiber.Map{"Error": "Profile ID is required"})
		}
		if data.QualificationId == 0 {
			return c.Render("", fiber.Map{"Error": "Qualification is required"})
		}
		result := services.DB.Db.Create(&data)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": "Create Data Error!",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"Data": "Ok",
		})
	case 3:
		data := new(model.ProfileWorkExperience)
		if err := c.BodyParser(data); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": err.Error(),
			})
		}
		if data.ProfileId == 0 {
			return c.Render("", fiber.Map{"Error": "Profile ID is required"})
		}
		if data.CompanyId == 0 {
			return c.Render("", fiber.Map{"Error": "Company is required"})
		}
		result := services.DB.Db.Create(&data)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": "Create Data Error!",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"Data": "Ok",
		})
	case 4:
		data := new(model.ProfileLanguage)
		if err := c.BodyParser(data); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": err.Error(),
			})
		}
		if data.ProfileId == 0 {
			return c.Render("", fiber.Map{"Error": "Profile ID is required"})
		}
		if data.LanguageCode == 0 {
			return c.Render("", fiber.Map{"Error": "Language is required"})
		}
		result := services.DB.Db.Create(&data)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": "Create Data Error!",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"Data": "Ok",
		})
	case 5:
		data := new(model.ProfileTraining)
		if err := c.BodyParser(data); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": err.Error(),
			})
		}
		if data.ProfileId == 0 {
			return c.Render("", fiber.Map{"Error": "Profile ID is required"})
		}
		if data.TrainingTittle == "" {
			return c.Render("", fiber.Map{"Error": "Training Tittle is required"})
		}
		result := services.DB.Db.Create(&data)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": "Create Data Error!",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"Data": "Ok",
		})
	case 6:
		data := new(model.ProfileFile)
		if err := c.BodyParser(data); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": err.Error(),
			})
		}
		if data.ProfileId == 0 {
			return c.Render("", fiber.Map{"Error": "Profile ID is required"})
		}
		if data.Files == "" {
			return c.Render("", fiber.Map{"Error": "Files is required"})
		}
		result := services.DB.Db.Create(&data)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"Data": "Create Data Error!",
			})
		}
		// return c.Status(200).JSON(fiber.Map{
		// 	"Data":      "Ok",
		// 	"Profileid": data.ID,
		// })
		return c.Redirect("/profile/profile_wizard")
	default:
		return c.Status(500).SendString("Error")
	}
}

func ProfileWizardProfileView(c *fiber.Ctx) error {
	data := []model.Profile{}
	country := []model.Country{}
	province := []model.Province{}
	religion := []model.Religion{}
	status := []model.MaritalStatus{}
	identification := []model.Identification{}
	position := []model.Position{}
	// db.Find(&data)
	services.DB.Db.Find(&data)
	services.DB.Db.Find(&country)
	services.DB.Db.Find(&province)
	services.DB.Db.Find(&religion)
	services.DB.Db.Find(&status)
	services.DB.Db.Find(&identification)
	services.DB.Db.Find(&position)
	return c.Render("profile/profile-wizard", fiber.Map{
		"Data":           data,
		"Country":        country,
		"Province":       province,
		"Religion":       religion,
		"Status":         status,
		"Identification": identification,
		"Position":       position,
	})
}

func CreateProfileWizardProfile(c *fiber.Ctx) error {
	// Ambil data gambar dari body request
	var requestData struct {
		Image string `json:"image"`
	}
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Data": "Failed to parse request body!",
		})
	}

	// Simpan informasi gambar ke dalam basis data
	data := model.Profile{
		// Isi dengan data lain yang diperlukan
		Photo: requestData.Image,
	}
	result := services.DB.Db.Create(&data)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Data": "Create Data Error!",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"msg":    "Image uploaded successfully",
	})
}

func ProfileWizardEducationView(c *fiber.Ctx) error {
	data := []model.ProfileEducation{}
	id := c.QueryInt("profileid")
	services.DB.Db.Select("profile_educations.profile_id AS profileid, qualifications.name AS qualification, universities.name AS university, departements.name AS departement, profile_educations.origin_school, profile_educations.major_sma, profile_educations.gpa, profile_educations.year_start, profile_educations.year_end").
		Joins("JOIN qualifications ON qualifications.id = profile_educations.qualification_id").
		Joins("LEFT JOIN universities ON universities.id = profile_educations.university_id").
		Joins("LEFT JOIN departements ON departements.id = profile_educations.departement_id").
		Where("profile_id", id).Order("qualification_id ASC").Find(&data)
	return c.Status(200).JSON(fiber.Map{
		"Data": data,
	})
}

func CreateProfileWizardEducation(c *fiber.Ctx) error {
	data := new(model.ProfileEducation)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.QualificationId == 0 {
		return c.Render("", fiber.Map{"Error": "Qualification is required"})
	}
	result := services.DB.Db.Create(&data)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"Data": "Ok",
	})
}

func ProfileWizardWorkExperienceView(c *fiber.Ctx) error {
	data := []model.ProfileWorkExperience{}
	id := c.QueryInt("profileid")
	services.DB.Db.Select("profile_work_experiences.profile_id AS profileid, positions.name AS startpositionjobtittle, companies.name AS company, countries.name AS country, provinces.name AS province, job_position_levels.name AS positionlevel, profile_work_experiences.salary, profile_work_experiences.experience_desc, profile_work_experiences.start_date, profile_work_experiences.end_date, positions.name AS lastpositionjobtittle, profile_work_experiences.reason_leaving").
		Joins("LEFT JOIN positions ON positions.id = profile_work_experiences.jobtittle").
		Joins("LEFT JOIN companies ON companies.id = profile_work_experiences.company_id").
		Joins("LEFT JOIN countries ON countries.id = profile_work_experiences.country_id").
		Joins("LEFT JOIN provinces ON provinces.id = profile_work_experiences.province_id").
		Joins("LEFT JOIN job_position_levels ON job_position_levels.id = profile_work_experiences.positionlevel_id").
		Joins("LEFT JOIN positions AS last_position ON last_position.id = profile_work_experiences.last_position_jobtittle").
		Where("profile_id", id).Order("company_id ASC").Find(&data)
	return c.Status(200).JSON(fiber.Map{
		"Data": data,
	})
}

func CreateProfileWizardWorkExperience(c *fiber.Ctx) error {
	data := new(model.ProfileWorkExperience)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.CompanyId == 0 {
		return c.Render("", fiber.Map{"Error": "Company is required"})
	}
	result := services.DB.Db.Create(&data)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"Data": "Ok",
	})
}

func ProfileWizardLanguageView(c *fiber.Ctx) error {
	data := []model.ProfileLanguage{}
	id := c.QueryInt("profileid")
	services.DB.Db.Select("profile_languages.profile_id AS profileid, languages.name AS language, spoken_levels.name AS spokenlevel, written_levels.name AS writtenlevel, listening_levels.name AS listeninglevel").
		Joins("LEFT JOIN languages ON languages.id = profile_languages.language_code").
		Joins("LEFT JOIN language_levels AS spoken_levels ON spoken_levels.id = profile_languages.spoken_level").
		Joins("LEFT JOIN language_levels AS written_levels ON written_levels.id = profile_languages.written_level").
		Joins("LEFT JOIN language_levels AS listening_levels ON listening_levels.id = profile_languages.listening_level").
		Where("profile_id", id).Order("language_code ASC").Find(&data)
	return c.Status(200).JSON(fiber.Map{
		"Data": data,
	})
}

func CreateProfileWizardLanguage(c *fiber.Ctx) error {
	data := new(model.ProfileLanguage)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.LanguageCode == 0 {
		return c.Render("", fiber.Map{"Error": "Language is required"})
	}
	result := services.DB.Db.Create(&data)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"Data": "Ok",
	})
}

func ProfileWizardTrainingView(c *fiber.Ctx) error {
	data := []model.ProfileTraining{}
	id := c.QueryInt("profileid")
	services.DB.Db.Select("profile_trainings.profile_id AS profileid, profile_trainings.training_tittle, profile_trainings.vendor, profile_trainings.training_year, profile_trainings.duration_day, profile_trainings.financed_by").
		Where("profile_id", id).Order("training_tittle ASC").Find(&data)
	return c.Status(200).JSON(fiber.Map{
		"Data": data,
	})
}

func CreateProfileWizardTraining(c *fiber.Ctx) error {
	data := new(model.ProfileTraining)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.TrainingTittle == "" {
		return c.Render("", fiber.Map{"Error": "Training Tittle is required"})
	}
	result := services.DB.Db.Create(&data)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"Data": "Ok",
	})
}

func ProfileWizardFilesView(c *fiber.Ctx) error {
	data := []model.ProfileFile{}
	services.DB.Db.Find(&data)
	return c.Render("profile/profile_wizard_step-6", fiber.Map{
		"Data": data,
	})
}

func CreateProfileWizardFiles(c *fiber.Ctx) error {
	data := new(model.ProfileFile)
	if err := c.BodyParser(data); err != nil {
		return c.Render("", fiber.Map{"Error": err.Error()})
	}
	if data.Files == "" {
		return c.Render("", fiber.Map{"Error": "Files is required"})
	}
	result := services.DB.Db.Create(&data)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	// fmt.Printf("User ID : %d\n", data.ID)
	// return c.Status(200).JSON(fiber.Map{
	// 	"Data":      "Ok",
	// 	"Profileid": data.ID,
	// })
	return c.Redirect("/profile/profile_wizard")
}

func UpdateProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(model.Profile)
	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Fullname == "" {
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

func DeleteProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	data := model.Profile{}
	services.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func GetProfileCity(c *fiber.Ctx) error {
	id := c.QueryInt("provinceid")
	city := []model.City{}
	services.DB.Db.Where("province_id = ?", id).Find(&city)
	return c.Status(200).JSON(fiber.Map{
		"Message": "Success",
		"City":    city,
	})
}

func GetProfileSkill(c *fiber.Ctx) error {
	id := c.QueryInt("categoryid")
	skill := []model.Skill{}
	services.DB.Db.Where("category_id = ?", id).Find(&skill)
	return c.Status(200).JSON(fiber.Map{
		"Message": "Success",
		"Skill":   skill,
	})
}

func GetQualification(c *fiber.Ctx) error {
	university := []model.University{}
	departement := []model.Departement{}
	services.DB.Db.Find(&university)
	services.DB.Db.Find(&departement)
	return c.Status(200).JSON(fiber.Map{
		"Message":     "Success",
		"University":  university,
		"Departement": departement,
	})
}
