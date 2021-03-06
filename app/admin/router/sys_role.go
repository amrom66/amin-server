package router

import (
	"amin/app/admin/apis/system/sys_role"
	middleware2 "amin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "amin/core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysRoleRouter)
}

// 需认证的路由代码
func registerSysRoleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &sys_role.SysRole{}
	r := v1.Group("/role").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{
		r.GET("", api.GetSysRoleList)
		r.GET("/:id", api.GetSysRole)
		r.POST("", api.InsertSysRole)
		r.PUT("/:id", api.UpdateSysRole)
		r.DELETE("", api.DeleteSysRole)
	}
	v1.PUT("/roledatascope", api.UpdateRoleDataScope)
}
