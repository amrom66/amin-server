package router

import (
	"amin/app/admin/apis/system/sys_post"
	middleware2 "amin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "amin/core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSyPostRouter)
}

// 需认证的路由代码
func registerSyPostRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &sys_post.SysPost{}
	r := v1.Group("/post").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{
		r.GET("", api.GetSysPostList)
		r.GET("/:id", api.GetSysPost)
		r.POST("", api.InsertSysPost)
		r.PUT("/:id", api.UpdateSysPost)
		r.DELETE("/:id", api.DeleteSysPost)
	}
}
