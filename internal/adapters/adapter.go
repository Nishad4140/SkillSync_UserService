package adapters

import (
	"fmt"

	"github.com/Nishad4140/SkillSync_UserService/entities"
	helperstruct "github.com/Nishad4140/SkillSync_UserService/entities/helperStruct"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAdapter struct {
	DB *gorm.DB
}

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{
		DB: db,
	}
}

func (user *UserAdapter) ClientSignup(userData entities.Client) (entities.Client, error) {
	var res entities.Client
	id := uuid.New()
	query := "INSERT INTO clients (id, name, email, phone, password, is_blocked) VALUES ($1, $2, $3, $4, $5, false) RETURNING *"
	if err := user.DB.Raw(query, id, userData.Name, userData.Email, userData.Phone, userData.Password).Scan(&res).Error; err != nil {
		return entities.Client{}, fmt.Errorf("error in inserting the values")
	}
	return res, nil
}

func (user *UserAdapter) GetClientByEmail(email string) (entities.Client, error) {
	var res entities.Client
	query := "SELECT * FROM clients WHERE email = ?"
	if err := user.DB.Raw(query, email).Scan(&res).Error; err != nil {
		return entities.Client{}, err
	}
	return res, nil
}

func (user *UserAdapter) GetClientByPhone(phone string) (entities.Client, error) {
	var res entities.Client
	query := "SELECT * FROM clients WHERE phone = ?"
	if err := user.DB.Raw(query, phone).Scan(&res).Error; err != nil {
		return entities.Client{}, err
	}
	return res, nil
}

func (user *UserAdapter) ClientCreateProfile(userID string) error {
	profileID := uuid.New()
	query := "INSERT INTO client_profiles (id, client_id) VALUES ($1, $2)"
	if err := user.DB.Exec(query, profileID, userID).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetClientProfileIdByUserId(userId string) (string, error) {
	var profileId string
	query := "SELECT id FROM client_profiles WHERE client_id = ?"
	if err := user.DB.Raw(query, userId).Scan(&profileId).Error; err != nil {
		return "", err
	}
	return profileId, nil
}

func (user *UserAdapter) UploadClientProfileImage(image, profileId string) (string, error) {
	var imageUrl string
	query := "UPDATE client_profiles SET image = $1 WHERE id = $2 RETURNING image"
	if err := user.DB.Raw(query, image, profileId).Scan(&imageUrl).Error; err != nil {
		return "", err
	}
	return imageUrl, nil
}

func (user *UserAdapter) GetClientProfileImage(profileId string) (string, error) {
	var image string
	query := "SELECT image FROM client_profiles WHERE id = $1 AND image IS NOT NULL"
	if err := user.DB.Raw(query, profileId).Scan(&image).Error; err != nil {
		return "", err
	}
	return image, nil
}

func (user *UserAdapter) ClientEditName(req entities.Client) error {
	query := "UPDATE clients SET name = $1 WHERE id = $2"
	if err := user.DB.Exec(query, req.Name, req.ID).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) ClientEditPhone(req entities.Client) error {
	query := "UPDATE clients SET phone = $1 WHERE id = $2"
	if err := user.DB.Exec(query, req.Phone, req.ID).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetClientById(userId string) (entities.Client, error) {
	var res entities.Client
	query := "SELECT * FROM clients WHERE id = ?"
	if err := user.DB.Raw(query, userId).Scan(&res).Error; err != nil {
		return entities.Client{}, err
	}
	return res, nil
}

func (user *UserAdapter) FreelancerSignup(freelancerData entities.Freelancer) (entities.Freelancer, error) {
	freelancerId := uuid.New()
	var res entities.Freelancer
	query := "INSERT INTO freelancers (id, name, email, phone, category_id, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	if err := user.DB.Raw(query, freelancerId, freelancerData.Name, freelancerData.Email, freelancerData.Phone, freelancerData.CategoryId, freelancerData.Password).Scan(&res).Error; err != nil {
		return entities.Freelancer{}, fmt.Errorf("error in inserting values to the freelancer table")
	}
	return res, nil
}

func (user *UserAdapter) CreateFreelancerProfile(req entities.FreelancerProfile) error {
	id := uuid.New()
	query := "INSERT INTO freelancer_profiles (id, freelancer_id) VALUES ($1, $2)"
	if err := user.DB.Exec(query, id, req.FreelancerId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetFreelancerByEmail(email string) (entities.Freelancer, error) {
	var res entities.Freelancer
	query := "SELECT * FROM freelancers WHERE email = ?"
	if err := user.DB.Raw(query, email).Scan(&res).Error; err != nil {
		return entities.Freelancer{}, err
	}
	return res, nil
}

func (user *UserAdapter) GetFreelancerByPhone(phone string) (entities.Freelancer, error) {
	var res entities.Freelancer
	query := "SELECT * FROM freelancers WHERE phone = ?"
	if err := user.DB.Raw(query, phone).Scan(&res).Error; err != nil {
		return entities.Freelancer{}, err
	}
	return res, nil
}

func (user *UserAdapter) GetFreelancerProfileIdByUserId(userId string) (string, error) {
	var profileId string
	query := "SELECT id FROM freelancer_profiles WHERE freelancer_id = ?"
	if err := user.DB.Raw(query, userId).Scan(&profileId).Error; err != nil {
		return "", err
	}
	return profileId, nil
}

func (user *UserAdapter) UploadFreelancerProfileImage(image, profileId string) (string, error) {
	var imageUrl string
	query := "UPDATE freelancer_profiles SET image = $1 WHERE id = $2 RETURNING image"
	if err := user.DB.Raw(query, image, profileId).Scan(&imageUrl).Error; err != nil {
		return "", err
	}
	return imageUrl, nil
}

func (user *UserAdapter) GetFreelancerProfileImage(profileId string) (string, error) {
	var image string
	query := "SELECT image FROM freelancer_profiles WHERE id = $1 AND image IS NOT NULL"
	if err := user.DB.Raw(query, profileId).Scan(&image).Error; err != nil {
		return "", err
	}
	return image, nil
}

func (user *UserAdapter) FreelancerEditName(req entities.Freelancer) error {
	query := "UPDATE freelancers SET name = $1 WHERE id = $2"
	if err := user.DB.Exec(query, req.Name, req.ID).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerEditPhone(req entities.Freelancer) error {
	query := "UPDATE freelancers SET phone = $1 WHERE id = $2"
	if err := user.DB.Exec(query, req.Phone, req.ID).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetCategoryById(categoryId int32) (entities.Category, error) {
	var res entities.Category
	query := "SELECT * FROM categories WHERE id = ?"
	if err := user.DB.Raw(query, categoryId).Scan(&res).Error; err != nil {
		return entities.Category{}, err
	}
	return res, nil
}

func (user *UserAdapter) GetAdminByEmail(email string) (entities.Admin, error) {
	var res entities.Admin
	query := "SELECT * FROM admins WHERE email = ?"
	if err := user.DB.Raw(query, email).Scan(&res).Error; err != nil {
		return entities.Admin{}, err
	}
	return res, nil
}

func (user *UserAdapter) AdminAddCategory(category entities.Category) error {
	query := "INSERT INTO categories (name) VALUES ($1)"
	if err := user.DB.Exec(query, category.Name).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) AdminUpdateCategory(category entities.Category) error {
	query := "UPDATE categories SET name=$1 WHERE id=$2"
	if err := user.DB.Exec(query, category.Name, category.ID).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetAllCategories() ([]entities.Category, error) {
	var res []entities.Category
	query := "SELECT * FROM categories"
	if err := user.DB.Raw(query).Scan(&res).Error; err != nil {
		return []entities.Category{}, err
	}
	return res, nil
}

func (user *UserAdapter) AdminAddSkill(skill entities.Skill) error {
	query := "INSERT INTO skills (category_id, name) VALUES ($1, $2)"
	if err := user.DB.Exec(query, skill.CategoryId, skill.Name).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetSkillByName(skill string) (entities.Skill, error) {
	var res entities.Skill
	query := "SELECT * FROM skills WHERE name = ?"
	if err := user.DB.Raw(query, skill).Scan(&res).Error; err != nil {
		return entities.Skill{}, err
	}
	return res, nil
}

func (user *UserAdapter) AdminUpdateSkill(skill entities.Skill) error {
	query := "UPDATE skills SET name = $1, category_id = $2 WHERE id = $3"
	if err := user.DB.Exec(query, skill.Name, skill.CategoryId, skill.ID).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) AdminGetAllSkills() ([]helperstruct.SkillHelper, error) {
	var res []helperstruct.SkillHelper
	query := "SELECT s.id AS skill_id, s.name AS skill_name, c.id AS category_id, c.name AS category_name FROM skills s JOIN categories c ON c.id = s.category_id"
	if err := user.DB.Raw(query).Scan(&res).Error; err != nil {
		return []helperstruct.SkillHelper{}, err
	}
	return res, nil
}

func (user *UserAdapter) GetSkillById(id int) (helperstruct.SkillHelper, error) {
	var res helperstruct.SkillHelper
	query := "SELECT s.id AS skill_id, s.name AS skill_name, c.id AS category_id, c.name AS category_name FROM skills s JOIN categories c ON c.id = s.category_id WHERE s.id = ?"
	if err := user.DB.Raw(query, id).Scan(&res).Error; err != nil {
		return helperstruct.SkillHelper{}, err
	}
	return res, nil
}

func (user *UserAdapter) GetFreelancerSkillById(profileId string, skillId int) (entities.FreelancerSkill, error) {
	var res entities.FreelancerSkill
	query := "SELECT * FROM freelancer_skills WHERE profile_id = $1 AND skill_id = $2"
	if err := user.DB.Raw(query, profileId, skillId).Scan(&res).Error; err != nil {
		return entities.FreelancerSkill{}, err
	}
	return res, nil
}

func (user *UserAdapter) FreelancerAddSkill(req entities.FreelancerSkill) error {
	id := uuid.New()
	query := "INSERT INTO freelancer_skills (id, skill_id, profile_id) VALUES ($1, $2, $3)"
	if err := user.DB.Exec(query, id, req.SkillId, req.ProfileId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerDeleteSkill(req entities.FreelancerSkill) error {
	query := "DELETE FROM freelancer_skills WHERE skill_id = $1 AND profile_id = $2"
	if err := user.DB.Exec(query, req.SkillId, req.ProfileId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerGetAllSkill(profileId string) ([]helperstruct.SkillHelper, error) {
	var res []helperstruct.SkillHelper
	query := "SELECT s.id as skill_id, s.name AS skill_name, c.id AS category_id, c.name as category_name FROM skills s JOIN categories c ON c.id = s.category_id JOIN freelancer_skills f ON f.skill_id = s.id WHERE profile_id = $1"
	if err := user.DB.Raw(query, profileId).Scan(&res).Error; err != nil {
		return []helperstruct.SkillHelper{}, err
	}
	return res, nil
}

func (user *UserAdapter) FreelancerAddEducation(req entities.Education) error {
	id := uuid.New()
	query := "INSERT INTO educations (id, freelancer_id, degree, institution, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6)"
	if err := user.DB.Exec(query, id, req.FreelancerID, req.Degree, req.Institution, req.StartDate, req.EndDate).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerEditEducation(req entities.Education) error {
	query := "UPDATE educations SET degree = $1, institution = $2, start_date = $3, end_date = $4 WHERE id = $5"
	if err := user.DB.Exec(query, req.Degree, req.Institution, req.StartDate, req.EndDate, req.ID).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerGetEducation(userId string) ([]entities.Education, error) {
	var res []entities.Education
	query := "SELECT * FROM educations WHERE freelancer_id = ?"
	if err := user.DB.Raw(query, userId).Scan(&res).Error; err != nil {
		return []entities.Education{}, err
	}
	return res, nil
}

func (user *UserAdapter) FreelancerRemoveEducation(educationId string) error {
	query := "DELETE FROM educations WHERE id  = ?"
	if err := user.DB.Exec(query, educationId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerAddTitle(req entities.FreelancerProfile) error {
	query := "UPDATE freelancer_profiles SET title = $1 WHERE freelancer_id = $2"
	if err := user.DB.Exec(query, req.Title, req.FreelancerId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerGetProfile(freelancerId string) (entities.FreelancerProfile, error) {
	var res entities.FreelancerProfile
	query := "SELECT * FROM freelancer_profiles WHERE freelancer_id = ?"
	if err := user.DB.Raw(query, freelancerId).Scan(&res).Error; err != nil {
		return entities.FreelancerProfile{}, err
	}
	return res, nil
}

func (user *UserAdapter) GetCategoryByName(name string) (entities.Category, error) {
	var res entities.Category
	query := "SELECT * FROM categories WHERE name = ?"
	if err := user.DB.Raw(query, name).Scan(&res).Error; err != nil {
		return entities.Category{}, err
	}
	return res, nil
}

func (user *UserAdapter) ClientAddAddress(req entities.Address, userId string) error {
	id := uuid.New()
	insertQuery := "INSERT INTO addresses (id, country, state, district, city) VALUES ($1, $2, $3, $4, $5)"
	if err := user.DB.Exec(insertQuery, id, req.Country, req.State, req.District, req.City).Error; err != nil {
		return err
	}
	updateQuery := "UPDATE client_profiles SET address_id = $1 WHERE client_id = $2"
	if err := user.DB.Exec(updateQuery, id, userId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetAddressByClientId(clientId string) (entities.Address, error) {
	var clientProfile entities.ClientProfile
	query := "SELECT * FROM client_profiles WHERE client_id = ?"
	if err := user.DB.Raw(query, clientId).Scan(&clientProfile).Error; err != nil {
		return entities.Address{}, err
	}
	var res entities.Address
	selectQuery := "SELECT * FROM addresses WHERE id = ?"
	if err := user.DB.Raw(selectQuery, clientProfile.AddressId).Scan(&res).Error; err != nil {
		return entities.Address{}, err
	}
	return res, nil
}

func (user *UserAdapter) ClientUpdateAddress(req entities.Address) error {
	query := "UPDATE addresses SET country = $1, state = $2, district = $3, city = $4 WHERE id = $5"
	if err := user.DB.Exec(query, req.Country, req.State, req.District, req.City, req.Id).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerAddAddress(req entities.Address, freelancerId string) error {
	id := uuid.New()
	insertQuery := "INSERT INTO addresses (id, country, state, district, city) VALUES ($1, $2, $3, $4, $5)"
	if err := user.DB.Exec(insertQuery, id, req.Country, req.State, req.District, req.City).Error; err != nil {
		return err
	}
	updateQuery := "UPDATE freelancer_profiles SET address_id = $1 WHERE freelancer_id = $2"
	if err := user.DB.Exec(updateQuery, id, freelancerId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetAddressByFreelancerId(freelancerId string) (entities.Address, error) {
	var freelancerProfile entities.FreelancerProfile
	query := "SELECT * FROM freelancer_profiles WHERE freelancer_id = ?"
	if err := user.DB.Raw(query, freelancerId).Scan(&freelancerProfile).Error; err != nil {
		return entities.Address{}, err
	}
	var res entities.Address
	selectQuery := "SELECT * FROM addresses WHERE id = ?"
	if err := user.DB.Raw(selectQuery, freelancerProfile.AddressId).Scan(&res).Error; err != nil {
		return entities.Address{}, err
	}
	return res, nil
}

func (user *UserAdapter) FreelancerUpdateAddress(req entities.Address) error {
	query := "UPDATE addresses SET country = $1, state = $2, district = $3, city = $4 WHERE id = $5"
	if err := user.DB.Exec(query, req.Country, req.State, req.District, req.City, req.Id).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) GetFreelancerById(freelancerId string) (entities.Freelancer, error) {
	var res entities.Freelancer
	query := "SELECT * FROM freelancers WHERE id = ?"
	if err := user.DB.Raw(query, freelancerId).Scan(&res).Error; err != nil {
		return entities.Freelancer{}, err
	}
	return res, nil
}

func (user *UserAdapter) FreelancerAddExperience(freelancerId, experience string) error {
	query := "UPDATE freelancer_profiles SET experience_in_current_field = $1 WHERE freelancer_id = $2"
	if err := user.DB.Exec(query, experience, freelancerId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerGetExperience(freelancerId string) (string, error) {
	var res string
	query := "SELECT COALESCE(experience_in_current_field, '0 years') FROM profiles WHERE freelancer_id = ?"
	if err := user.DB.Raw(query, freelancerId).Scan(&res).Error; err != nil {
		return "", err
	}
	return res, nil
}

func (user *UserAdapter) ClientBlock(clientId string) error {
	query := "UPDATE clients SET is_blocked = true WHERE id = ?"
	if err := user.DB.Exec(query, clientId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) ClientUnBlock(clientId string) error {
	query := "UPDATE clients SET is_blocked = false WHERE id = ?"
	if err := user.DB.Exec(query, clientId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerBlock(freelancerId string) error {
	query := "UPDATE freelancers SET is_blocked = true WHERE id = ?"
	if err := user.DB.Exec(query, freelancerId).Error; err != nil {
		return err
	}
	return nil
}

func (user *UserAdapter) FreelancerUnblock(freelancerId string) error {
	query := "UPDATE freelancers SET is_blocked = false WHERE id = ?"
	if err := user.DB.Exec(query, freelancerId).Error; err != nil {
		return err
	}
	return nil
}
