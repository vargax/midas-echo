package src

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

type UserService interface {
	CreateUser(u *User) error
}
