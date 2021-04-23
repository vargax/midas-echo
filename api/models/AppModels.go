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
	ID     uint
	Email  string `gorm:"unique;not null"`
	Passwd string `json:"-" gorm:"not null"`
	gorm.Model
}

// DTOs ************************
// Data Transfer Objects
// *****************************

type JwtToken struct {
	Token string `json:"token"`
}
