package router

import (
	"amin/app/admin/apis/system/sys_dept"
	middleware2 "amin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "amin/core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeptRouter)
}

// 需认证的路由代码
func registerSysDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &sys_dept.SysDept{}
	r := v1.Group("/dept").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{
		r.GET("", api.GetSysDeptList)
		r.GET("/:id", api.GetSysDept)
		r.POST("", api.InsertSysDept)
		r.PUT("/:id", api.UpdateSysDept)
		r.DELETE("/:id", api.DeleteSysDept)
	}

	r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("/deptTree", api.GetDeptTree)
	}

}
