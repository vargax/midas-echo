module gitlab.activarsas.net/cvargasc/midas-echo

go 1.16

require (

	github.com/labstack/echo/v4 v4.2.1 	// > Framework
	gorm.io/gorm v1.21.5 				// > Database
	gorm.io/driver/postgres v1.0.8

	github.com/joho/godotenv v1.3.0 	// > Load configurations from .env files

	github.com/jackc/pgproto3/v2 v2.0.7 					// indirect
	github.com/jackc/pgx/v4 v4.11.0 						// indirect
	github.com/mattn/go-colorable v0.1.8 					// indirect
	golang.org/x/net v0.0.0-20210324205630-d1beb07c2056 	// indirect
	golang.org/x/text v0.3.5 								// indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
