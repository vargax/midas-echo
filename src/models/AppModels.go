package models

import "gorm.io/gorm"

// Roles ******************
// Used for Authorization
// *************************
type Role string

const (
	// RoleAdmin A super user
	RoleAdmin = Role("admin")
	// RoleUser A regular authenticated user
	RoleUser = Role("user")
	// RoleGuest Unauthenticated user / fallback
	RoleGuest = Role("guest")
)

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
	Role     Role   `gorm:"not null"`
	gorm.Model
}

// DTOs ************************
// Data Transfer Objects
// *****************************

type JwtToken struct {
	Token string `json:"token"`
}
