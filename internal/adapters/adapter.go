package adapters

import (
	"github.com/Nishad4140/SkillSync_UserService/entities"
	"gorm.io/gorm"
)

type UserAdapter struct {
	DB *gorm.DB
}

// UserSignup implements AdapterInterface.
func (*UserAdapter) UserSignup(entities.Client) (entities.Client, error) {
	panic("unimplemented")
}

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{
		DB: db,
	}
}
