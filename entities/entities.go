package entities

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID          uuid.UUID `gorm:"primaryKey;unique;not null"`
	Name        string
	Email       string
	Phone       string
	Password    string
	ReportCount int
	IsBlocked   bool
	CreatedAt   time.Time
}

type ClientProfile struct {
	ID        uuid.UUID `gorm:"primaryKey;unique;not null"`
	ClientID  uuid.UUID
	Client    Client `gorm:"foreignKey:ClientID"`
	AddressId uuid.UUID
	Address   Address `gorm:"foreignKey:AddressId"`
	Image     string
}

type Freelancer struct {
	ID         uuid.UUID `gorm:"primaryKey;unique;not null"`
	Name       string
	Email      string
	Phone      string
	CategoryId int32
	Category   Category `gorm:"foreignKey:CategoryId"`
	Rating     float64
	IsBlocked  bool
	Password   string
	CreatedAt  time.Time
}

type FreelancerProfile struct {
	ID                       uuid.UUID `gorm:"primaryKey;unique;not null"`
	FreelancerId             uuid.UUID
	Freelancer               Freelancer `gorm:"foreignKey:FreelancerId"`
	Title                    string
	AddressId                uuid.UUID
	Address                  Address `gorm:"foreignKey:AddressId"`
	ExperienceInCurrentField string
	Image                    string
}

type FreelancerSkill struct {
	ID                uuid.UUID
	ProfileId         uuid.UUID
	FreelancerProfile FreelancerProfile `gorm:"foreignKey:ProfileId"`
	SkillId           int
	Skill             Skill `gorm:"foreignKey:SkillId"`
}

type Admin struct {
	ID       uuid.UUID `gorm:"primaryKey;unique;not null"`
	Name     string
	Email    string
	Phone    string
	Password string
}

type Category struct {
	ID   int `gorm:"primaryKey; unique; not null"`
	Name string
}

type Skill struct {
	ID         int `gorm:"primaryKey; unique; not null"`
	CategoryId int
	Category   Category `gorm:"foreignKey:CategoryId"`
	Name       string
}

type Address struct {
	Id       uuid.UUID `gorm:"primaryKey;unique; not null"`
	Country  string
	State    string
	District string
	City     string
}

type Education struct {
	ID           uuid.UUID
	FreelancerID uuid.UUID
	Freelancer   Freelancer `gorm:"foreignKey:FreelancerID"`
	Degree       string
	Institution  string
	StartDate    time.Time `gorm:"type:date"`
	EndDate      time.Time `gorm:"type:date"`
}
