package middleware

import (
	"time"

	"amin/app/admin/middleware/handler"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

// AuthInit jwt验证new
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	if config.ApplicationConfig.Mode == "dev" {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if config.JwtConfig.Timeout != 0 {
			timeout = time.Duration(config.JwtConfig.Timeout) * time.Second
		}
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(config.ApplicationConfig.JwtSecret),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler,
		Authenticator:   handler.Authenticator,
		Authorizator:    handler.Authorizator,
		Unauthorized:    handler.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

}
