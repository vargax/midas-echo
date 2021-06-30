package middleware

import (
	"github.com/casbin/casbin/v2"
	ecb "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo/src/models"
	"github.com/vargax/midas-echo/src/utils"
	"path"
)

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

	e, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}

	return ecb.Config{
		Enforcer: e,
		UserGetter: func(c echo.Context) (string, error) {
			role, err := jwtExtractClaim(c, jwtclaimsRole)
			if err != nil {
				// If there is any problem getting the role, we will default to Guest
				return models.RoleGuest, nil
			}
			return role, nil
		},
	}
}
