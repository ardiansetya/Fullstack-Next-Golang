package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"jamal/api/api/models/domain"
	"log"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (repository AuthRepositoryImpl) Register(tx *gorm.DB, user domain.User) error {
	// <-- check acc already ?
	var existUser domain.User
	if err := tx.Where("user_name = ?", user.UserName).First(&existUser).Error; err == nil {
		return errors.New("username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("failed to query database")
	}

	// <-- create
	result := tx.Create(&user).Error
	if result != nil {
		return errors.New("Error Create")
	}
	return nil
}

func (repository AuthRepositoryImpl) Login(tx *gorm.DB, user domain.User) error {
	var dataDB domain.User
	if err := tx.Take(&dataDB, "user_name=?", user.UserName).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("username atau password salah")
		}
	}

	/*log.Print("Username: ", user.UserName)
	log.Print("Stored Hashed Password: ", dataDB.Password)
	log.Print("Input Password: ", user.Password)*/

	// <-- data password di database yang sudah di-hash dibandingkan dengan password plaintext yang dikirim oleh user
	err := bcrypt.CompareHashAndPassword([]byte(dataDB.Password), []byte(user.Password))
	if err != nil {
		return errors.New("username atau password salah")
	}
	log.Print("passworld", user.Password)
	return nil
}
