package adapters

import "github.com/Nishad4140/SkillSync_UserService/entities"

type AdapterInterface interface {
	ClientSignup(entities.Client) (entities.Client, error)
	GetClientByEmail(email string) (entities.Client, error)
	GetClientByPhone(phone string) (entities.Client, error)
	CreateProfile(userID string) error
	FreelancerSignup(entities.Freelancer) (entities.Freelancer, error)
	GetFreelancerByEmail(email string) (entities.Freelancer, error)
	GetFreelancerByPhone(phone string) (entities.Freelancer, error)
	GetCategoryById(categoryId int32) (entities.Category, error)
}
