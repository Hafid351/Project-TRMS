package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username           string `json:"Username"`
	Fullname           string `json:"Fullname"`
	AuthKey            string `json:"AuthKey"`
	PasswordHash       string `json:"PasswordHash"`
	PasswordResetToken string `json:"PasswordResetToken"`
	Email              string `json:"Email"`
	VerificationToken  string `json:"VerificationToken"`
	RolesId            int    `json:"RolesId"`
	Role               string `gorm:"->" json:"Role"`
}

type Skill struct {
	gorm.Model
	Name       string `json:"Name"`
	CategoryId int    `json:"CategoryId"`
	Category   string `gorm:"->" json:"Category"`
}

type Skill1 struct {
	gorm.Model
	Name       string `json:"Name"`
	CategoryId int    `json:"CategoryId"`
	Category   string `gorm:"->" json:"Category"`
	Total      int    `json:"Total"`
}

type SkillCategory struct {
	gorm.Model
	Name string `json:"Name"`
}

type Position struct {
	gorm.Model
	Name       string `json:"Name"`
	CategoryId int    `json:"CategoryId"`
	Category   string `gorm:"->" json:"Category"`
}

type Position1 struct {
	Name  string `json:"Name"`
	Total int    `json:"Total"`
}

type PositionCategory struct {
	gorm.Model
	Name string `json:"Name"`
}

type University struct {
	gorm.Model
	Name string `json:"Name"`
	Url  string `json:"Url"`
}

type Departement struct {
	gorm.Model
	Name string `json:"Name"`
}

type Departement1 struct {
	gorm.Model
	Name  string `json:"Name"`
	Total int    `json:"Total"`
}

type Company struct {
	gorm.Model
	Logo               string `json:"Logo"`
	Name               string `json:"Name"`
	CompanyDescription string `json:"CompanyDescription"`
	IndustryId         int    `json:"IndustriId"`
	Industry           string `gorm:"->" json:"Industry"`
	Website            string `json:"Website"`
	PhoneNumber        string `json:"PhoneNumber"`
	Email              string `json:"Email"`
	WorkingHour        string `json:"WorkingHour"`
	Benefit            string `json:"Benefit"`
	LanguageUse        string `json:"LanguageUse"`
	CountryId          string `json:"CountryId"`
	Country            string `gorm:"->" json:"Country"`
	ProvinceId         int    `json:"ProvinceId"`
	Province           string `gorm:"->" json:"Province"`
	CityId             int    `json:"CityId"`
	City               string `gorm:"->" json:"City"`
	Address            string `json:"Address"`
}

type Industry struct {
	gorm.Model
	Name   string `json:"Name"`
	Status string `json:"Status"`
}

type Country struct {
	gorm.Model
	Code string `json:"Code"`
	Name string `json:"Name"`
}

type Language struct {
	gorm.Model
	Code string `json:"Code"`
	Name string `json:"Name"`
}

type CertificateCategory struct {
	gorm.Model
	Name     string `json:"Name"`
	Files    string `json:"Files"`
	SkillsId int    `json:"SkillsId"`
	Skill    string `gorm:"->" json:"Skill"`
}

type Role struct {
	gorm.Model
	Name string `json:"Name"`
}

type Dashboard struct {
	Skill string `json:"Skill"`
	Id    int    `json:"Id"`
	Total int    `json:"Total"`
}

type Dashboard1 struct {
	Name  string `json:"Name"`
	Id    int    `json:"Id"`
	Total int    `json:"Total"`
}

type Dashboard2 struct {
	Name  string `json:"Name"`
	Total int    `json:"Total"`
}

type Dashboard3 struct {
	Name                         string `json:"Name"`
	ProfileWorkExperiencesSalary int    `json:"ProfileWorkExperienceSalary"`
	Total                        int    `json:"Total"`
}

type Province struct {
	gorm.Model
	Name string `json:"Name"`
}

type City struct {
	gorm.Model
	Name       string `json:"Name"`
	ProvinceId int    `json:"ProvinceId"`
}

type Profile struct {
	gorm.Model
	Fullname          string    `json:"Fullname"`
	Gender            string    `json:"Gender"`
	Photo             string    `json:"Photo"`
	Filename          string    `json:"Filename"`
	Avatar            string    `json:"Avatar"`
	CountryId         int       `json:"CountryId"`
	ProvinceId        int       `json:"ProvinceId"`
	CityId            int       `json:"CityId"`
	Address           string    `json:"Address"`
	NationalityId     string    `json:"NationalityId"`
	Height            int       `json:"Height"`
	Weight            int       `json:"Weight"`
	ReligionId        int       `json:"ReligionId"`
	MaritalStatus     int       `json:"Maritalstatus"`
	IdentificationId  int       `json:"IdentificationId"`
	IdCardNo          string    `json:"IdCardNo"`
	PhoneNumber       string    `json:"PhoneNumber"`
	OtherPhoneNumber  string    `json:"OtherPhoneNumber"`
	Email             string    `json:"Email"`
	SalaryExpectation int       `json:"SalaryExpectation"`
	JobTitle          int       `json:"JobTitle"`
	About             string    `json:"About"`
	InstagramProfile  string    `json:"InstagramProfile"`
	FacebookProfile   string    `json:"FacebookProfile"`
	LinkedinProfile   string    `json:"LinkedinProfile"`
	Dob               time.Time `json:"Dob"`
	Pob               string    `json:"Pob"`
}

type ProfileEducation struct {
	gorm.Model
	ProfileId       int    `json:"ProfileId"`
	QualificationId int    `json:"QualificationId"`
	Qualification   string `gorm:"->" json:"Qualification"`
	UniversityId    int    `json:"UniversityId"`
	University      string `gorm:"->" json:"University"`
	DepartementId   int    `json:"DepartementId"`
	Departement     string `gorm:"->" json:"Departement"`
	OriginSchool    string `json:"OriginSchool"`
	MajorSMA        string `json:"MajorSMA"`
	Gpa             string `json:"Gpa"`
	YearStart       int    `json:"YearStart"`
	YearEnd         int    `json:"YearEnd"`
}

type ProfileWorkExperience struct {
	gorm.Model
	ProfileId             int    `json:"ProfileId"`
	JobTittle             int    `json:"JobTittle"`
	CompanyId             int    `json:"CompanyId"`
	CountryId             int    `json:"CountryId"`
	ProvinceId            int    `json:"ProvinceId"`
	PositionlevelId       int    `json:"PositionlevelId"`
	Salary                int    `json:"Salary"`
	ExperienceDesc        string `json:"ExperienceDesc"`
	StartDate             string `json:"StartDate"`
	EndDate               string `json:"EndDate"`
	LastPositionJobtittle int    `json:"LastPositionJobtittle"`
	ReasonLeaving         string `json:"ReasonLeaving"`
}

type ProfileLanguage struct {
	gorm.Model
	ProfileId      int `json:"ProfileId"`
	LanguageCode   int `json:"LanguageCode"`
	SpokenLevel    int `json:"SpokenLevel"`
	WritenLevel    int `json:"WritenLevel"`
	ListeningLevel int `json:"ListeningLevel"`
}

type ProfileTraining struct {
	gorm.Model
	ProfileId      int    `json:"ProfileId"`
	TrainingTittle string `json:"TrainingTittle"`
	Vendor         string `json:"Vendor"`
	TrainingYear   int    `json:"TrainingYear"`
	DurationDay    int    `json:"DurationDay"`
	FinancedBy     string `json:"FinancedBy"`
	CompanyName    string `json:"CompanyName"`
}

type ProfileFile struct {
	gorm.Model
	ProfileId    int    `json:"ProfileId"`
	ProfileFiles int    `json:"ProfileFiles"`
	Files        string `json:"Files"`
	Filename     string `json:"Filename"`
	Avatar       string `json:"Avatar"`
	Name         string `json:"Name"`
}

type ProfileWorkExperience1 struct {
	Name   string `json:"Name"`
	Salary int    `json:"Salary"`
	Total  int    `json:"Total"`
}

type Religion struct {
	gorm.Model
	Name string `json:"Name"`
}

type MaritalStatus struct {
	gorm.Model
	Name string `json:"Name"`
}

type Identification struct {
	gorm.Model
	Name string `json:"Name"`
}

type Qualification struct {
	gorm.Model
	Name string `json:"Name"`
}

type JobPositionLevel struct {
	gorm.Model
	Name   string `json:"Name"`
	Status string `json:"Status"`
}

type LanguageLevel struct {
	gorm.Model
	Name string `json:"Name"`
}
