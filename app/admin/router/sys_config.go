package router

import (
	"amin/app/admin/apis/system/sys_config"
	middleware2 "amin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "amin/core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysConfigRouter)
}

// 需认证的路由代码
func registerSysConfigRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r := v1.Group("/config").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{
		api := &sys_config.SysConfig{}
		r.GET("", api.GetSysConfigList)
		r.GET("/:id", api.GetSysConfig)
		r.POST("", api.InsertSysConfig)
		r.PUT("/:id", api.UpdateSysConfig)
		r.DELETE("/:id", api.DeleteSysConfig)
	}

	r1 := v1.Group("/configKey").Use(authMiddleware.MiddlewareFunc())
	{
		api := &sys_config.SysConfig{}
		r1.GET("/:configKey", api.GetSysConfigByKEYForService)
	}
}
