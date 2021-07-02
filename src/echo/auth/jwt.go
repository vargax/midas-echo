package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/vargax/midas-echo/env"
	"github.com/vargax/midas-echo/src/echo/validator"
	"github.com/vargax/midas-echo/src/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

// Authentication **********
// https://echo.labstack.com/cookbook/jwt/
// *************************
const (
	jwtclaimsUsername = "Username"
	jwtclaimsRole     = "Role"
	jwtclaimsExp      = "exp"
)

var (
	secret    string
	tokenLive time.Duration
)

func AuthenticationConfig() emw.JWTConfig {
	secret = os.Getenv(env.JwtSecret)

	var err error
	tokenLive, err = time.ParseDuration(os.Getenv(env.JwtTokenExp))
	if err != nil {
		panic(err)
	}

	return emw.JWTConfig{
		SigningKey: []byte(secret),
		Skipper: func(c echo.Context) bool {
			// We will skip authentication (i.e JWT token validation) if there is no Authorization Header
			// AuthorizationMiddleware will treat unauthenticated requests as RoleGuest
			return c.Request().Header.Get(echo.HeaderAuthorization) == ""
		},
	}
}

type JwtToken struct {
	Token string `json:"token"`
}

func JwtTokenFactory(tr *validator.PostPublicToken) (*JwtToken, error) {
	user, err := postgres.User(tr.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, echo.ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}

	if !postgres.PasswordMatch(user.Password, tr.Password) {
		return nil, echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[jwtclaimsUsername] = user.Username
	claims[jwtclaimsRole] = user.Role
	claims[jwtclaimsExp] = time.Now().Add(tokenLive).Unix()

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	response := JwtToken{
		Token: signedToken,
	}
	return &response, nil
}

func jwtExtractClaim(c echo.Context, claim string) (string, error) {
	ctxUser, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return "", errors.New("JWT token not found")
	}

	jwtClaims, ok := ctxUser.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("MapClaims not found")
	}

	claimValue, ok := jwtClaims[claim].(string)
	if !ok {
		return "", errors.New(claim + " claim not found")
	}

	return claimValue, nil
}
