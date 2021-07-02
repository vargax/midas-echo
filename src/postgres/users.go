package postgres

import (
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo/src"
	"github.com/vargax/midas-echo/src/echo/validator"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func User(username string) (src.User, error) {
	user := src.User{
		Username: username,
	}
	result := db.Where(&user).First(&user)

	return user, result.Error
}

func NewUser(userToBe *validator.PostAppUsers) (*src.User, error) {
	pass, err := EncodePassword(userToBe.Password)
	if err != nil {
		return nil, err
	}

	newUser := src.User{
		Username: userToBe.Username,
		Password: pass,
		Role:     src.Role(userToBe.Role),
	}

	result := db.Create(&newUser)
	if e, ok := result.Error.(*pgconn.PgError); ok && e.Code == DuplicateKey {
		return nil, echo.NewHTTPError(http.StatusConflict, e.Detail)
	}

	return &newUser, err
}

func EncodePassword(in string) (out string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	out = string(hash)
	return out, nil
}

func PasswordMatch(h1, h2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h1), []byte(h2))
	return err == nil
}
