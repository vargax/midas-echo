package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/vargax/midas-echo/env"
	"github.com/vargax/midas-echo/src/models"
	"github.com/vargax/midas-echo/src/repository"
	"github.com/vargax/midas-echo/src/utils"
	"gorm.io/gorm"
	"os"
	"time"
)

func Env() {
	var err error

	secret = os.Getenv(env.JwtSecret)

	tokenLive, err = time.ParseDuration(os.Getenv(env.JwtTokenExp))
	if err != nil {
		panic(err)
	}
}

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
	return emw.JWTConfig{
		SigningKey: []byte(secret),
		Skipper: func(c echo.Context) bool {
			// We will skip authentication (i.e JWT token validation) if there is no Authorization Header
			// AuthorizationMiddleware will treat unauthenticated requests as RoleGuest
			return c.Request().Header.Get(echo.HeaderAuthorization) == ""
		},
	}
}

func JwtTokenFactory(tr *models.PostPublicToken) (*models.JwtToken, error) {
	user, err := repository.ReadUser(tr.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, echo.ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}

	if !utils.PasswordMatch(user.Password, tr.Password) {
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

	response := models.JwtToken{
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

func ctxGetUser(c echo.Context) (*models.User, error) {
	username, err := jwtExtractClaim(c, jwtclaimsUsername)
	if err != nil {
		return nil, err
	}

	user, err := repository.ReadUser(username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
