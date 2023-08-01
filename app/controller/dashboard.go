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

func GetData1(c *fiber.Ctx) error {
	data := []model.Dashboard1{}
	services.DB.Db.Raw("SELECT source.name AS name, source.id AS id, source.total FROM (SELECT departements.name AS name, departements.id AS id, COUNT(*) AS total FROM departements FULL JOIN profile_educations AS Profile_Educations ON departements.id = Profile_Educations.departement_id GROUP BY departements.name, departements.id ORDER BY departements.name ASC, departements.id ASC) AS source WHERE source.total BETWEEN 100 AND 700 AND ((source.total <> 127) OR (source.total IS NULL)) AND ((source.total <> 132) OR (source.total IS NULL))").Find(&data)
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
// FROM departements FULL JOIN profile_pendidikans AS Profile_Pendidikan ON departements.id = Profile_Pendidikan.departements_id
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
	result := services.DB.Db.Table("departements").Select("departements.name", "COUNT(*) AS total").Joins("JOIN profile_educations AS Profile_Educations ON departements.id = Profile_Educations.departement_id")
	if search != "" {
		result = result.Where("name like ?", "%"+search+"%")
	}
	result.Group("departements.name, departements.id").Order("total DESC").Offset(offset).Limit(perPage).Find(&table)

	var total int64

	if search != "" {
		services.DB.Db.Joins("JOIN profile_educations AS Profile_Educations ON departements.id = Profile_Pendidikans.departement_id").Where("departements.name like ?", "%"+search+"%").Group("departements.name, departements.id").Order("total DESC").Count(&total)
	} else {
		services.DB.Db.Table("departements").Model(&model.Departement1{}).Joins("JOIN profile_educations AS Profile_Educations ON departements.id = Profile_Educations.departement_id").Group("departements.name, departements.id").Count(&total)
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

// WITH cte AS (
//     SELECT positions.name AS name,
//            profiles.salary_expectation AS Profile_Work_Experiences_salary_expectation,
//            COUNT(*) AS total
//     FROM positions
//     LEFT JOIN profiles
//         ON positions.id = profiles.job_title
//     WHERE profiles.salary_expectation BETWEEN 2000000 AND 5000000
//     GROUP BY positions.name, profiles.salary_expectation
// )
// SELECT cte.name,
//        CASE
//            WHEN cte.Profile_Work_Experiences_salary_expectation >= 2000000 AND cte.Profile_Work_Experiences_salary_expectation < 3000000 THEN '2-3 Juta'
//            WHEN cte.Profile_Work_Experiences_salary_expectation >= 3000000 AND cte.Profile_Work_Experiences_salary_expectation < 4000000 THEN '3-4 Juta'
//            WHEN cte.Profile_Work_Experiences_salary_expectation >= 4000000 AND cte.Profile_Work_Experiences_salary_expectation <= 5000000 THEN '4-5 Juta'
//            ELSE 'Tidak Diketahui'
//        END AS salary_category,
//        MAX(cte.total) AS total
// FROM cte
// GROUP BY cte.name, cte.Profile_Work_Experiences_salary_expectation
// ORDER BY cte.name ASC;

func GetData3(c *fiber.Ctx) error {
	data := []model.Dashboard3{}
	services.DB.Db.Raw("WITH cte AS (SELECT positions.name AS name, profiles.salary_expectation AS Profiles_salary_expectation, COUNT(*) AS total FROM positions LEFT JOIN profiles ON positions.id = profiles.job_title WHERE profiles.salary_expectation BETWEEN 2000000 AND 5000000 GROUP BY positions.name, profiles.salary_expectation) SELECT cte.name, CASE WHEN cte.Profiles_salary_expectation >= 2000000 AND cte.Profiles_salary_expectation < 3000000 THEN '2-3 Juta' WHEN cte.Profiles_salary_expectation >= 3000000 AND cte.Profiles_salary_expectation < 4000000 THEN '3-4 Juta' WHEN cte.Profiles_salary_expectation >= 4000000 AND cte.Profiles_salary_expectation <= 5000000 THEN '4-5 Juta' ELSE 'Tidak Diketahui' END AS salary_category, MAX(cte.total) AS total FROM cte GROUP BY cte.name, cte.Profiles_salary_expectation ORDER BY cte.name ASC;").Find(&data)
	// SELECT positions.name AS name, profiles.salary_expectation AS Profile_Work_Experiences_salary_expectation, COUNT(*) AS total 
	// FROM positions LEFT JOIN profiles AS Profiles ON positions.id = Profiles.job_title 
	// WHERE Profiles.salary_expectation BETWEEN 2000000 AND 5000000 
	// GROUP BY positions.name, Profiles.salary_expectation 	
	// ORDER BY positions.name ASC, Profiles.salary_expectation ASC
	return c.JSON(fiber.Map{
		"Data": data,
	})
}

func GetTable3(c *fiber.Ctx) error {
	table := []model.Position2{}
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
	result := services.DB.Db.Table("positions").Select("positions.name", "COUNT(*) AS total").Joins("JOIN profiles AS Profiles ON positions.id = profiles.job_title")
	if search != "" {
		result = result.Where("name like ?", "%"+search+"%")
	}
	result.Group("positions.name, profiles.salary_expectation").Order("total DESC").Offset(offset).Limit(perPage).Find(&table)

	var total int64

	if search != "" {
		services.DB.Db.Joins("JOIN profiles AS Profiles ON positions.id = profiles.job_title").Where("positions.name like ?", "%"+search+"%").Group("positions.name, profiles.salary_expectation").Order("total DESC").Count(&total)
	} else {
		services.DB.Db.Table("positions").Model(&model.Position2{}).Joins("JOIN profiles AS Profiles ON positions.id = profiles.job_title").Group("positions.name, profiles.salary_expectation").Count(&total)
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
