package src

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

// DTOs ************************
// Data Transfer Objects
// *****************************
