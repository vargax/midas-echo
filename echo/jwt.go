package echo

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/vargax/midas-echo"
	"github.com/vargax/midas-echo/env"
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

func authenticationConfig() emw.JWTConfig {
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

type jwtToken struct {
	Token string `json:"token"`
}

func jwtTokenFactory(tr *PostPublicToken) (*jwtToken, error) {
	user := midas.User{
		Username: tr.Username,
		Password: tr.Password,
	}

	err := ss.UserSrv.Authenticate(&user)
	if errors.Is(err, midas.UserNotFound) || errors.Is(err, midas.PasswordDontMatch) {
		return nil, echo.ErrUnauthorized
	}
	if err != nil {
		return nil, err
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

	response := jwtToken{
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
