package usecase

import "github.com/Nishad4140/SkillSync_UserService/internal/adapters"

type UserUsecase struct {
	userAdapter adapters.AdapterInterface
}

func NewUserUsecase(userAdapter adapters.AdapterInterface) *UserUsecase {
	return &UserUsecase{
		userAdapter: userAdapter,
	}
}
