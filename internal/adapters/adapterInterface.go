package adapters

import "github.com/Nishad4140/SkillSync_UserService/entities"

type AdapterInterface interface {
	// Client
	ClientSignup(entities.Client) (entities.Client, error)
	GetClientByEmail(email string) (entities.Client, error)
	GetClientByPhone(phone string) (entities.Client, error)
	CreateClientProfile(userID string) error
	ClientAddAddress(address entities.Address, userId string) error
	ClientUpdateAddress(address entities.Address) error
	GetAddressByUserId(userId string) (entities.Address, error)

	// Freelancer
	FreelancerSignup(entities.Freelancer) (entities.Freelancer, error)
	GetFreelancerByEmail(email string) (entities.Freelancer, error)
	GetFreelancerByPhone(phone string) (entities.Freelancer, error)

	// Admin
	GetAdminByEmail(email string) (entities.Admin, error)
	GetCategoryById(categoryId int32) (entities.Category, error)
	AdminAddCategory(entities.Category) error
	AdminUpdateCategory(category entities.Category) error
	GetAllCategories() ([]entities.Category, error)
	GetCategoryByName(name string) (entities.Category, error)
}
