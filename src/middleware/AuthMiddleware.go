package middleware

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	ecb "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/vargax/midas-echo/env"
	"github.com/vargax/midas-echo/src/models"
	"github.com/vargax/midas-echo/src/repository"
	"github.com/vargax/midas-echo/src/utils"
	"gorm.io/gorm"
	"os"
	"path"
	"strings"
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
	jwtclaimsExp      = "exp"
)

var (
	secret    string
	tokenLive time.Duration
)

func AuthenticationConfig() emw.JWTConfig {
	return emw.JWTConfig{
		SigningKey: []byte(secret),
		Skipper:    skipper,
	}
}

func NewJwtToken(tokenRequest *models.PostPublicToken) (*models.JwtToken, error) {

	user, err := repository.ReadUser(tokenRequest.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, echo.ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}

	if !utils.PasswordMatch(user.Password, tokenRequest.Password) {
		return nil, echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[jwtclaimsUsername] = user.Username
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

func ctxGetUser(c echo.Context) (*models.User, error) {
	ctxUser := c.Get("user").(*jwt.Token)
	ctxClaims := ctxUser.Claims.(jwt.MapClaims)

	username := ctxClaims[jwtclaimsUsername].(string)

	user, err := repository.ReadUser(username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Public All routes under /public skip Authentication and Authorization
const Public = "/public"

func skipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), Public)
}

// Authorization ***********
// https://echo.labstack.com/middleware/casbin-auth/
// *************************
const (
	model  = "casbin/model.conf"
	policy = "casbin/policy.csv"
)

func AuthorizationConfig() ecb.Config {

	modelPath := path.Join(utils.GoFilePath(), model)
	policyPath := path.Join(utils.GoFilePath(), policy)

	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}

	return ecb.Config{
		Skipper:  skipper,
		Enforcer: enforcer,
		UserGetter: func(c echo.Context) (string, error) {
			user, err := ctxGetUser(c)
			return user.Role, err
		},
	}
}
