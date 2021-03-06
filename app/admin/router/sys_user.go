package router

import (
	"amin/app/admin/apis/sys_user"
	"amin/common/actions"
	middleware2 "amin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "amin/core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &sys_user.SysUser{}
	r := v1.Group("/sysUser").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("", api.GetSysUserList)
		r.GET("/:id", api.GetSysUser)
		r.POST("", api.InsertSysUser)
		r.PUT("", api.UpdateSysUser)
		r.DELETE("", api.DeleteSysUser)
	}

	user := v1.Group("/user").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole()).Use(actions.PermissionAction())
	{
		user.GET("/profile", api.GetSysUserProfile)
		user.POST("/avatar", api.InsetSysUserAvatar)
		user.PUT("/pwd", api.SysUserUpdatePwd)
	}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		v1auth.GET("/getinfo", api.GetInfo)
	}
}
