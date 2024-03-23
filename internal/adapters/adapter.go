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
	query := "INSERT INTO clients (id, name, email, phone, password) VALUES ($1, $2, $3, $4, $5) RETURNING *"
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

func (user *UserAdapter) CreateClientProfile(userID string) error {
	profileID := uuid.New()
	query := "INSERT INTO client_profile (id, user_id) ($1, $2)"
	if err := user.DB.Raw(query, profileID, userID).Error; err != nil {
		return err
	}
	return nil
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

func (user *UserAdapter) GetAddressByUserId(userId string) (entities.Address, error) {
	var addressId string
	query := "SELECT address_id FROM client_profiles WHERE client_id = ?"
	if err := user.DB.Raw(query, userId).Scan(&addressId).Error; err != nil {
		return entities.Address{}, err
	}
	var res entities.Address
	selectQuery := "SELECT * FROM addresses WHERE id = ?"
	if err := user.DB.Raw(selectQuery, addressId).Scan(&res).Error; err != nil {
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
