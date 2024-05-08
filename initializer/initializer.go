package initializer

import (
	"github.com/Nishad4140/SkillSync_UserService/concurrency"
	"github.com/Nishad4140/SkillSync_UserService/internal/adapters"
	"github.com/Nishad4140/SkillSync_UserService/internal/service"
	"github.com/Nishad4140/SkillSync_UserService/internal/usecase"
	"gorm.io/gorm"
)

func Initializer(db *gorm.DB) *service.UserService {
	adapter := adapters.NewUserAdapter(db)
	usecase := usecase.NewUserUsecase(adapter)
	service := service.NewUserService(adapter, usecase)
	c := concurrency.NewConcurrency(db, adapter, service)
	c.Concurrency()
	return service
}
