package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"primaryKey; varchar(300); not null"`
	Password string `gorm:"varchar(300); not null"`
	Name     string `gorm:"varchar(300); not null"`
}
