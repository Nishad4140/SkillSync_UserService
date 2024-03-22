package entities

import "github.com/google/uuid"

type Client struct {
	ID       uuid.UUID `gorm:"primaryKey;unique;not null"`
	Name     string
	Email    string
	Phone    string
	Password string
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
	Password   string
}

type Admin struct {
	ID       uuid.UUID `gorm:"primaryKey;unique;not null"`
	Name     string
	Email    string
	Phone    string
	Password string
}

type Category struct {
	ID   int
	Name string
}

type Address struct {
	Id       uuid.UUID `gorm:"primaryKey;unique; not null"`
	Country  string
	State    string
	District string
	City     string
}
