package echo

import (
	"github.com/casbin/casbin/v2"
	ecb "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo"
	"path"
	"runtime"
)

// Authorization ***********
// https://echo.labstack.com/middleware/casbin-auth/
// *************************
const (
	model  = "casbin/model.conf"
	policy = "casbin/policy.csv"
)

func authorizationConfig() ecb.Config {
	modelPath := path.Join(filePath(), model)
	policyPath := path.Join(filePath(), policy)

	enf, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}

	return ecb.Config{
		Enforcer: enf,
		UserGetter: func(c echo.Context) (string, error) {
			role, err := jwtExtractClaim(c, jwtclaimsRole)
			if err != nil {
				// If there is any problem getting the role, we will default to Guest
				return string(midas.RoleGuest), nil
			}
			return role, nil
		},
	}
}

func filePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
