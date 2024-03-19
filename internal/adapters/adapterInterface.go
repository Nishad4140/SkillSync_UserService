package adapters

import "github.com/Nishad4140/SkillSync_UserService/entities"

type AdapterInterface interface {
	UserSignup(entities.Client) (entities.Client, error)
}
