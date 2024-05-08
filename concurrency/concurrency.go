package concurrency

import (
	"log"
	"sync"
	"time"

	"github.com/Nishad4140/SkillSync_UserService/internal/adapters"
	"github.com/Nishad4140/SkillSync_UserService/internal/service"
	"gorm.io/gorm"
)

type Concurrency struct {
	DB       *gorm.DB
	adapters adapters.AdapterInterface
	mu       sync.Mutex
	service  *service.UserService
}

func NewConcurrency(DB *gorm.DB, adapters adapters.AdapterInterface, service *service.UserService) *Concurrency {
	return &Concurrency{
		DB:       DB,
		adapters: adapters,
		service:  service,
	}
}

func (c *Concurrency) Concurrency() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			c.mu.Lock()
			if err := c.DB.Exec(`
			UPDATE clients SET is_blocked = true WHERE report_count > 50
			`).Error; err != nil {
				log.Print("error while performing concurrency", err)
				break
			}
		}
		c.mu.Unlock()
	}()
}
