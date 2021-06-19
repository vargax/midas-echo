module gitlab.activarsas.net/cvargasc/midas-echo

go 1.16

require (
	github.com/casbin/casbin/v2 v2.31.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // > Authentication (JWT Tokens support)
	github.com/go-playground/validator/v10 v10.6.1 // > Data-validation
	github.com/jackc/pgconn v1.8.1
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/joho/godotenv v1.3.0 // > Load configurations from .env files
	github.com/labstack/echo-contrib v0.11.0 // > Authorization (echo-casbin)
	github.com/labstack/echo/v4 v4.3.0 // > Framework
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
	golang.org/x/time v0.0.0-20210608053304-ed9ce3a009e4 // indirect
	gorm.io/driver/postgres v1.1.0 // > Gorm Postgres support
	gorm.io/gorm v1.21.10 // > Database
)
