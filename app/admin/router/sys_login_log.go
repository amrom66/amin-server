package router

import (
	"amin/app/admin/apis/system/sys_login_log"
	middleware2 "amin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "amin/core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysLoginLogRouter)
}

// 需认证的路由代码
func registerSysLoginLogRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &sys_login_log.SysLoginLog{}
	r := v1.Group("/sys-login-log").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{
		r.GET("", api.GetSysLoginLogList)
		r.GET("/:id", api.GetSysLoginLog)
		r.POST("", api.InsertSysLoginLog)
		r.PUT("/:id", api.UpdateSysLoginLog)
		r.DELETE("", api.DeleteSysLoginLog)
	}
}
