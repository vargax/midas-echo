package models

import "gorm.io/gorm"

// Entities ****************
// Persistent data maintained in the database
// *************************

type File struct {
	ID            int
	Name          string
	Path          string
	PublicacionID int
	gorm.Model    `json:"-"`
}

type User struct {
	ID     int
	email  string
	passwd string
	gorm.Model
}

// DTOs ************************
// Data Transfer Objects
// *****************************

type JwtToken struct {
	Token string `json:"token"`
}
