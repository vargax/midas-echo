module github.com/vargax/midas-echo

go 1.16

require (
	github.com/casbin/casbin/v2 v2.31.7
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // > Authentication (JWT Tokens support)
	github.com/go-playground/validator/v10 v10.6.1 // > Data-validation
	github.com/jackc/pgconn v1.8.1
	github.com/jackc/pgproto3/v2 v2.1.0 // indirect
	github.com/joho/godotenv v1.3.0 // > Load configurations from .env files
	github.com/labstack/echo-contrib v0.11.0 // > Authorization (echo-casbin)
	github.com/labstack/echo/v4 v4.3.0 // > Framework
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/text v0.3.8 // indirect
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6 // indirect
	gorm.io/driver/postgres v1.1.0 // > Gorm Postgres support
	gorm.io/gorm v1.21.11 // > Database
)
