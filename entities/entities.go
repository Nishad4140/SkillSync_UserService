package entities

import "github.com/google/uuid"

type Client struct {
	ID        uuid.UUID `gorm:"primaryKey;unique;not null"`
	Name      string
	Email     string
	Phone     string
	Password  string
	IsBlocked bool
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
	IsBlocked  bool
	Password   string
}

type FreelancerProfile struct {
	ID                       uuid.UUID `gorm:"primaryKey;unique;not null"`
	FreelancerId             uuid.UUID
	Freelancer               Freelancer `gorm:"foreignKey:FreelancerId"`
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
