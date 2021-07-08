package midas

import "errors"

// Role type for Authorization
type Role string

const (
	// RoleAdmin A super user
	RoleAdmin = Role("admin")
	// RoleUser A regular authenticated user
	RoleUser = Role("user")
	// RoleGuest Unauthenticated user / fallback
	RoleGuest = Role("guest")
)

type User struct {
	ID       uint
	Username string `gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`
	Role     Role   `gorm:"not null"`
}

type UserSrv interface {
	New(u *User) error
	User(username string) (*User, error)
	Authenticate(u *User) error
}

var (
	UserNotFound          = errors.New("user not found")
	PasswordDontMatch     = errors.New("password don't match")
	UserAlreadyRegistered = errors.New("user already registered")
)
