package adapters

import (
	"fmt"

	"github.com/Nishad4140/SkillSync_UserService/entities"
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

func (user *UserAdapter) CreateProfile(userID string) error {
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
