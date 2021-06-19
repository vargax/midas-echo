package models

import "gorm.io/gorm"

// Entities ****************
// Persistent data maintained in the database
// *************************

type File struct {
	ID            uint
	Name          string
	Path          string
	PublicacionID int
	gorm.Model    `json:"-"`
}

type User struct {
	ID       uint
	Username string `gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`
	Role     string `gorm:"not null"`
	gorm.Model
}

// DTOs ************************
// Data Transfer Objects
// *****************************

type JwtToken struct {
	Token string `json:"token"`
}
