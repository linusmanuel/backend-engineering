package database

import (
	"github.com/linusmanuel/backend-engineering/goexpert/9-APIS/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
