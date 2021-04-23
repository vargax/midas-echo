module gitlab.activarsas.net/cvargasc/midas-echo

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // > JWT Tokens support
	github.com/go-playground/validator/v10 v10.4.1 // > Data-validation
	github.com/joho/godotenv v1.3.0 // > Load configurations from .env files
	github.com/labstack/echo/v4 v4.2.1 // > Framework
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b
	gorm.io/driver/postgres v1.0.8 // > Gorm Postgres support
	gorm.io/gorm v1.21.5 // > Database
)
