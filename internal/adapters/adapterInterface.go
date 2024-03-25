package adapters

import (
	"github.com/Nishad4140/SkillSync_UserService/entities"
	helperstruct "github.com/Nishad4140/SkillSync_UserService/entities/helperStruct"
)

type AdapterInterface interface {
	// Client
	ClientSignup(entities.Client) (entities.Client, error)
	GetClientByEmail(email string) (entities.Client, error)
	GetClientByPhone(phone string) (entities.Client, error)
	CreateClientProfile(userID string) error
	ClientAddAddress(address entities.Address, userId string) error
	ClientUpdateAddress(req entities.Address) error
	GetAddressByClientId(clientId string) (entities.Address, error)
	GetClientProfileIdByUserId(userId string) (string, error)
	UploadClientProfileImage(image, profileId string) (string, error)
	GetClientProfileImage(profileId string) (string, error)
	ClientEditName(req entities.Client) error
	ClientEditPhone(req entities.Client) error

	// Freelancer
	FreelancerSignup(entities.Freelancer) (entities.Freelancer, error)
	GetFreelancerByEmail(email string) (entities.Freelancer, error)
	GetFreelancerByPhone(phone string) (entities.Freelancer, error)
	CreateFreelancerProfile(req entities.FreelancerProfile) error
	FreelancerAddAddress(address entities.Address, userId string) error
	FreelancerUpdateAddress(req entities.Address) error
	GetAddressByFreelancerId(freelancerId string) (entities.Address, error)
	GetFreelancerProfileIdByUserId(userId string) (string, error)
	UploadFreelancerProfileImage(image, profileId string) (string, error)
	GetFreelancerProfileImage(profileId string) (string, error)
	FreelancerEditName(req entities.Freelancer) error
	FreelancerEditPhone(req entities.Freelancer) error
	GetSkillById(id int) (helperstruct.SkillHelper, error)
	GetFreelancerSkillById(profileId string, skillId int) (entities.FreelancerSkill, error)
	FreelancerAddSkill(req entities.FreelancerSkill) error
	FreelancerDeleteSkill(req entities.FreelancerSkill) error
	FreelancerGetAllSkill(profileId string) ([]helperstruct.SkillHelper, error)

	// Admin
	GetAdminByEmail(email string) (entities.Admin, error)
	GetCategoryById(categoryId int32) (entities.Category, error)
	AdminAddCategory(category entities.Category) error
	AdminUpdateCategory(category entities.Category) error
	GetAllCategories() ([]entities.Category, error)
	GetCategoryByName(name string) (entities.Category, error)
	AdminAddSkill(skill entities.Skill) error
	AdminUpdateSkill(skill entities.Skill) error
	AdminGetAllSkills() ([]helperstruct.SkillHelper, error)
	GetSkillByName(skill string) (entities.Skill, error)
	ClientBlock(clientId string) error
	ClientUnBlock(clientId string) error
	FreelancerBlock(freelancerId string) error
	FreelancerUnblock(freelancerId string) error
}
