package repository

import (
	"gorm.io/gorm"
	"jamal/api/api/models/domain"
)

type AuthRepository interface {
	Register(tx *gorm.DB, user domain.User) error
	Login(tx *gorm.DB, user domain.User) error
}
