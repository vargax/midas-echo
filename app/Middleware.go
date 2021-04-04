package app

import (
	"encoding/json"
	"errors"
	auth0 "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"os"
)

var (
	iss      = "https://" + os.Getenv(jwtIss) + "/"
	audience = os.Getenv(jwtAudience)
)

type (
	Jwks struct {
		Keys []JSONWebKeys `json:"keys"`
	}

	Response struct {
		Message string `json:"message"`
	}

	JSONWebKeys struct {
		Kty string   `json:"kty"`
		Kid string   `json:"kid"`
		Use string   `json:"use"`
		N   string   `json:"n"`
		E   string   `json:"e"`
		X5c []string `json:"x5c"`
	}
)

var authMiddleware *auth0.JWTMiddleware

func AuthMiddlewareInit() {
	debug, _ := strconv.ParseBool(os.Getenv(debugJwt))

	options := auth0.Options{
		ValidationKeyGetter: validationKeyGetter,
		Debug:               debug,
		SigningMethod:       jwt.SigningMethodHS256,
	}

	authMiddleware = auth0.New(options)
}

func validationKeyGetter(token *jwt.Token) (interface{}, error) {
	// Verify 'aud' claim
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
	if !checkAud {
		return token, echo.ErrUnauthorized
	}
	// Verify 'iss' claim
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
	if !checkIss {
		return token, echo.ErrUnauthorized
	}

	cert, err := getPemCert(token)
	if err != nil {
		return nil, err
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(iss + ".well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
