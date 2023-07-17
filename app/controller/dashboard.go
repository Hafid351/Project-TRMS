package controller

import (
	"math"
	"strconv"
	"trms/app/model"
	"trms/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetData(c *fiber.Ctx) error {
	data := []model.Dashboard{}
	services.DB.Db.Model(&model.Skill{}).Select("skills.name AS skill, skills.id AS id, COUNT(*) AS total").Joins("FULL JOIN profile_skills AS Profile_Skills ON skills.id = Profile_Skills.skill_id").Group("skill, skills.id").Order("total DESC").Find(&data)
	// SELECT skills.name AS skill, skills.id AS id, COUNT(*) AS total
	// FROM "skills" FULL JOIN profile_skills AS Profile_Skills ON skills.id = Profile_Skills.skill_id
	// WHERE "skills"."deleted_at" IS NULL
	// GROUP BY skill, skills.id ORDER BY total DESC
	return c.JSON(fiber.Map{
		"Data": data,
	})
}

func GetTable(c *fiber.Ctx) error {
	table := []model.Skill1{}
	search := c.Query("search")
	perPage, err := strconv.Atoi(c.Query("perPage", "5"))
	if err != nil {
		return c.SendString("perPage harus angka")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendString("page harus angka")
	}
	offset := (page - 1) * perPage
	result := services.DB.Db.Table("skills").Select("skills.name", "COUNT(*) AS total").Joins("JOIN profile_skills AS Profile_Skills ON skills.id = Profile_Skills.skill_id")
	if search != "" {
		result = result.Where("name like ?", "%"+search+"%")
	}
	result.Group("skills.name, skills.id").Order("total DESC").Offset(offset).Limit(perPage).Find(&table)

	var total int64

	if search != "" {
		services.DB.Db.Joins("JOIN profile_skills AS Profile_Skills ON skills.id = Profile_Skills.skill_id").Where("skills.name like ?", "%"+search+"%").Group("skills.name, skills.id").Order("total DESC").Count(&total)
	} else {
		services.DB.Db.Table("skills").Model(&model.Skill1{}).Joins("JOIN profile_skills AS Profile_Skills ON skills.id = Profile_Skills.skill_id").Group("skills.name, skills.id").Count(&total)
	}
	return c.JSON(fiber.Map{
		"Table":      table,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

// TRYING DASHBOARD KEAHLIAN
// func GetData(c *fiber.Ctx) error {
// 	data := []model.Dashboard{}

// 	// Menambahkan indeks pada kolom skills.id dan Profile_Skills.skill_id jika belum ada
// 	if err := services.DB.Db.Exec("CREATE INDEX IF NOT EXISTS idx_skills_id ON skills(id)").Error; err != nil {
// 		return err
// 	}
// 	if err := services.DB.Db.Exec("CREATE INDEX IF NOT EXISTS idx_profile_skills_skill_id ON profile_skills(skill_id)").Error; err != nil {
// 		return err
// 	}

// 	// Menggunakan RAW query untuk mengoptimasi performa
// 	query := `SELECT skills.name AS skill, skills.id AS id, COUNT(*) AS total
// 		FROM skills
// 		FULL JOIN profile_skills AS Profile_Skills ON skills.id = Profile_Skills.skill_id
// 		GROUP BY skill, skills.id
// 		ORDER BY total DESC`

// 	if err := services.DB.Db.Raw(query).Scan(&data).Error; err != nil {
// 		return err
// 	}

// 	return c.JSON(fiber.Map{
// 		"Data": data,
// 	})
// }

// func GetTable(c *fiber.Ctx) error {
// 	search := c.Query("search")
// 	perPage, err := strconv.Atoi(c.Query("perPage", "5"))
// 	if err != nil {
// 		return c.SendString("perPage harus angka")
// 	}
// 	page, err := strconv.Atoi(c.Query("page", "1"))
// 	if err != nil {
// 		return c.SendString("page harus angka")
// 	}
// 	offset := (page - 1) * perPage

// 	// Menambahkan indeks pada kolom skills.id jika belum ada
// 	services.DB.Db.Exec("CREATE INDEX IF NOT EXISTS idx_skills_id ON skills(id)")

// 	var table []model.Skill1
// 	query := services.DB.Db.Table("skills").
// 		Select("skills.name", "COUNT(*) AS total").
// 		Joins("JOIN profile_skills AS Profile_Skills ON skills.id = Profile_Skills.skill_id").
// 		Preload("ProfileSkills").
// 		Group("skills.name, skills.id").
// 		Order("total DESC").
// 		Offset(offset).
// 		Limit(perPage)

// 	if search != "" {
// 		query = query.Where("skills.name ILIKE ?", "%"+search+"%")
// 	}

// 	query.Find(&table)

// 	var total int64
// 	countQuery := services.DB.Db.Joins("JOIN profile_skills AS Profile_Skills ON skills.id = Profile_Skills.skill_id").
// 		Group("skills.name, skills.id").
// 		Order("total DESC").
// 		Model(&model.Skill1{})

// 	if search != "" {
// 		countQuery = countQuery.Where("skills.name ILIKE ?", "%"+search+"%")
// 	}

// 	countQuery.Count(&total)

// 	return c.JSON(fiber.Map{
// 		"Table":      table,
// 		"TotalData":  total,
// 		"Page":       page,
// 		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
// 		"PerPage":    perPage,
// 	})
// }

//END OF TRYING DASHBOARD KEAHLIAN

func GetData1(c *fiber.Ctx) error {
	data := []model.Dashboard1{}
	services.DB.Db.Raw("SELECT source.name AS name, source.id AS id, source.total FROM (SELECT departements.name AS name, departements.id AS id, COUNT(*) AS total FROM departements FULL JOIN profile_educations AS profile_educations ON departements.id = profile_educations.departement_id GROUP BY departements.name, departements.id ORDER BY departements.name ASC, departements.id ASC) AS source WHERE source.total BETWEEN 100 AND 700 AND ((source.total <> 127) OR (source.total IS NULL))").Find(&data)
	// SELECT departements.name AS name, departements.id AS id, COUNT(*) AS count
	// FROM departements FULL JOIN profile_pendidikan AS Profile_Pendidikan ON departements.id = Profile_Pendidikan.departements_id
	// GROUP BY departements.name, departements.id
	// ORDER BY departements.name ASC, departements.id ASC AS source
	// WHERE source.count BETWEEN 100
	//    AND 700 AND (("source"."count" <> 127)
	//     OR ("source"."count" IS NULL))
	return c.JSON(fiber.Map{
		"Data": data,
		// "Tabel":tabel
	})
}

// SELECT departements.name AS name, departements.id AS id, COUNT(*) AS count
// FROM departements FULL JOIN profile_educations AS Profile_Pendidikan ON departements.id = Profile_Pendidikan.departements_id
// GROUP BY departements.name, departements.id
// ORDER BY count DESC

func GetTable1(c *fiber.Ctx) error {
	table := []model.Departement1{}
	search := c.Query("search")
	perPage, err := strconv.Atoi(c.Query("perPage", "5"))
	if err != nil {
		return c.SendString("perPage harus angka")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendString("page harus angka")
	}
	offset := (page - 1) * perPage
	result := services.DB.Db.Table("departements").Select("departements.name", "COUNT(*) AS total").Joins("JOIN profile_educations AS profile_educations ON departements.id = profile_educations.departement_id")
	if search != "" {
		result = result.Where("name like ?", "%"+search+"%")
	}
	result.Group("departements.name, departements.id").Order("total DESC").Offset(offset).Limit(perPage).Find(&table)

	var total int64

	if search != "" {
		services.DB.Db.Joins("JOIN profile_educations AS profile_educations ON departements.id = profile_educations.departement_id").Where("departements.name like ?", "%"+search+"%").Group("departements.name, departements.id").Order("total DESC").Count(&total)
	} else {
		services.DB.Db.Table("departements").Model(&model.Departement1{}).Joins("JOIN profile_educations AS profile_educations ON departements.id = profile_educations.departement_id").Group("departements.name, departements.id").Count(&total)
	}
	return c.JSON(fiber.Map{
		"Table":      table,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func GetData2(c *fiber.Ctx) error {
	data := []model.Dashboard2{}
	services.DB.Db.Raw("SELECT source.name AS name, source.total AS total FROM (SELECT positions.name AS name, COUNT(*) AS total FROM profiles INNER JOIN positions AS Positions ON profiles.job_title = Positions.id GROUP BY positions.name ORDER BY COUNT(*) DESC) AS source").Find(&data)
	// 	SELECT "source"."name" AS "name", "source"."Profile Work Experience_positionlevel_id" AS "Profile Work Experience_positionlevel_id", "source"."count" AS "count"
	// FROM (SELECT "public"."job_position_levels"."name" AS "name", "Profile Work Experience"."positionlevel_id" AS "Profile Work Experience_positionlevel_id", COUNT(*) AS "count" FROM "public"."job_position_levels" FULL JOIN "public"."profile_work_experience" AS "Profile Work Experience" ON "public"."job_position_levels"."id" = "Profile Work Experience"."positionlevel_id"
	// GROUP BY "public"."job_position_levels"."name", "Profile Work Experience"."positionlevel_id"
	// ORDER BY "public"."job_position_levels"."name" ASC, "Profile Work Experience"."positionlevel_id" ASC) AS "source"
	// WHERE ("source"."count" <> 48)
	//     OR ("source"."count" IS NULL)
	return c.JSON(fiber.Map{
		"Data": data,
	})
}

// SELECT job_position_levels.name AS name, Profile_Work_Experiences.positionlevel_id AS Profile_Work_Experience_positionlevel_id, COUNT(*) AS count
// FROM job_position_levels FULL JOIN profile_work_experiences AS Profile_Work_Experiences ON job_position_levels.id = Profile_Work_Experiences.positionlevel_id
// GROUP BY job_position_levels.name, Profile_Work_Experiences.positionlevel_id
// ORDER BY count DESC

func GetTable2(c *fiber.Ctx) error {
	table := []model.Position1{}
	search := c.Query("search")
	perPage, err := strconv.Atoi(c.Query("perPage", "5"))
	if err != nil {
		return c.SendString("perPage harus angka")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendString("page harus angka")
	}
	offset := (page - 1) * perPage
	result := services.DB.Db.Table("profiles").Select("positions.name", "COUNT(*) AS total").Joins("JOIN positions AS Positions ON profiles.job_title = Positions.id")
	if search != "" {
		result = result.Where("name like ?", "%"+search+"%")
	}
	result.Group("positions.name").Order("total DESC").Offset(offset).Limit(perPage).Find(&table)

	var total int64

	if search != "" {
		services.DB.Db.Joins("JOIN positions AS Positions ON profiles.job_title = Positions.id").Where("positions.name like ?", "%"+search+"%").Group("positions.name").Order("total DESC").Count(&total)
	} else {
		services.DB.Db.Table("profiles").Model(&model.Position1{}).Joins("JOIN positions AS Positions ON profiles.job_title = Positions.id").Group("positions.name").Count(&total)
	}
	return c.JSON(fiber.Map{
		"Table":      table,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func GetData3(c *fiber.Ctx) error {
	data := []model.Dashboard3{}
	services.DB.Db.Raw("SELECT job_position_levels.name AS name, Profile_Work_Experiences.salary AS Profile_Work_Experiences_salary, COUNT(*) AS total FROM job_position_levels LEFT JOIN profile_work_experiences AS Profile_Work_Experiences ON job_position_levels.id = Profile_Work_Experiences.positionlevel_id WHERE Profile_Work_Experiences.salary BETWEEN 2000000 AND 5000000 GROUP BY job_position_levels.name, Profile_Work_Experiences.salary ORDER BY job_position_levels.name ASC, Profile_Work_Experiences.salary ASC").Find(&data)
	// 	SELECT "public"."job_position_levels"."name" AS "name", "Profile Work Experience"."salary" AS "Profile Work Experience__salary", COUNT(*) AS "count"
	// FROM "public"."job_position_levels"
	// LEFT JOIN "public"."profile_work_experience" AS "Profile Work Experience" ON "public"."job_position_levels"."id" = "Profile Work Experience"."positionlevel_id"
	// WHERE "Profile Work Experience"."salary" BETWEEN 2000000
	//    AND 5000000
	// GROUP BY "public"."job_position_levels"."name", "Profile Work Experience"."salary"
	// ORDER BY "public"."job_position_levels"."name" ASC, "Profile Work Experience"."salary" ASC
	return c.JSON(fiber.Map{
		"Data": data,
	})
}

// SELECT job_position_levels.name AS name, Profile_Work_Experiences.salary AS Profile_Work_Experiences_salary, COUNT(*) AS total
// FROM job_position_levels LEFT JOIN profile_work_experiences AS Profile_Work_Experiences ON job_position_levels.id = Profile_Work_Experiences.positionlevel_id
// WHERE Profile_Work_Experiences.salary BETWEEN 2000000 AND 5000000
// GROUP BY job_position_levels.name, Profile_Work_Experiences.salary
// ORDER BY job_position_levels.name ASC, Profile_Work_Experiences.salary ASC

func GetTable3(c *fiber.Ctx) error {
	table := []model.ProfileWorkExperience1{}
	search := c.Query("search")
	perPage, err := strconv.Atoi(c.Query("perPage", "5"))
	if err != nil {
		return c.SendString("perPage harus angka")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendString("page harus angka")
	}
	offset := (page - 1) * perPage
	result := services.DB.Db.Table("job_position_levels").Select("job_position_levels.name", "COUNT(*) AS total").Joins("JOIN profile_work_experiences AS Profile_Work_Experiences ON job_position_levels.id = Profile_Work_Experiences.positionlevel_id")
	if search != "" {
		result = result.Where("name like ?", "%"+search+"%")
	}
	result.Group("job_position_levels.name, Profile_Work_Experiences.salary").Order("total DESC").Offset(offset).Limit(perPage).Find(&table)

	var total int64

	if search != "" {
		services.DB.Db.Joins("JOIN profile_work_experiences AS Profile_Work_Experiences ON job_position_levels.id = Profile_Work_Experiences.positionlevel_id").Where("job_position_levels.name like ?", "%"+search+"%").Group("job_position_levels.name, Profile_Work_Experiences.salary").Order("total DESC").Count(&total)
	} else {
		services.DB.Db.Table("job_position_levels").Model(&model.ProfileWorkExperience1{}).Joins("JOIN profile_work_experiences AS Profile_Work_Experiences ON job_position_levels.id = Profile_Work_Experiences.positionlevel_id").Group("job_position_levels.name, Profile_Work_Experiences.salary").Count(&total)
	}
	return c.JSON(fiber.Map{
		"Table":      table,
		"TotalData":  total,
		"Page":       int(page),
		"TotalPages": int(math.Ceil(float64(total) / float64(perPage))),
		"PerPage":    perPage,
	})
}

func DashboardView(c *fiber.Ctx) error {
	return c.Render("dashboard/dist/index", fiber.Map{
		"Data": "",
	})
}

func DashboardTable(c *fiber.Ctx) error {
	return c.Render("dashboard/dist/index", fiber.Map{
		"Table": "",
	})
}
