package env

const (
	EchoPort = "ECHO_PORT"

	DebugEcho = "DEBUG_FRAMEWORK"
	DebugRepo = "DEBUG_REPOSITORY"
	DebugJwt  = "DEBUG_JWT"

	JwtIss      = "JWT_ISS"
	JwtAudience = "JWT_AUDIENCE"

	JwtSecret      = "JWT_SECRET"
	JwtTokenSecret = "JWT_TOKEN_SECRET"
	JwtTokenExp    = "JWT_TOKEN_EXP"

	CorsOrigin = "CORS_ORIGIN"

	DbHost = "POSTGRES_HOST"
	DbPort = "POSTGRES_PORT"
	DbUser = "POSTGRES_USER"
	DbPass = "POSTGRES_PASSWORD"
	DbName = "POSTGRES_DB"
)
