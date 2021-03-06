package router

import (
	"amin/app/admin/apis/system/sys_menu"
	middleware2 "amin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "amin/core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysMenuRouter)
}

// 需认证的路由代码
func registerSysMenuRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &sys_menu.SysMenu{}
	//menu := v1.Group("/menu").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	//{
	//	menu.GET("/:id", system.GetMenu)
	//	menu.POST("", system.InsertMenu)
	//	menu.PUT("", system.UpdateMenu)
	//	menu.DELETE("/:id", system.DeleteMenu)
	//}
	r := v1.Group("/menu").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{
		r.GET("", api.GetSysMenuList)
		r.GET("/:id", api.GetSysMenu)
		r.POST("", api.InsertSysMenu)
		r.PUT("/:id", api.UpdateSysMenu)
		r.DELETE("/:id", api.DeleteSysMenu)
	}

	r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("/menurole", api.GetMenuRole)
		r1.GET("/menuids", api.GetMenuIDS)
	}

}
