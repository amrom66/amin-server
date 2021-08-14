package router

import (
	"amin/app/admin/apis/system/sys_menu"
	middleware2 "amin/common/middleware"
	"mime"

	"amin/app/admin/apis/monitor"
	"amin/app/admin/apis/public"
	"amin/app/admin/apis/system"
	"amin/app/admin/apis/system/dict"
	. "amin/app/admin/apis/tools"
	"amin/app/admin/middleware/handler"
	_ "amin/docs"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/ws"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	sysStaticFileRouter(g)
	// swagger；注意：生产环境可以注释掉
	sysSwaggerRouter(g)
	// 无需认证
	sysNoCheckRoleRouter(g)
	// 需要认证
	sysCheckRoleRouterInit(g, authMiddleware)
	return g
}

func sysBaseRouter(r *gin.RouterGroup) {

	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendAllService()

	r.GET("/", system.HelloWorld)
	r.GET("/info", handler.Ping)
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	mime.AddExtensionType(".js", "application/javascript")
	r.Static("/static", "./static")
	r.Static("/form-generator", "./static/form-generator")
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.GET("/monitor/server", monitor.ServerInfo)
	v1.GET("/getCaptcha", system.GenerateCaptchaHandler)
	v1.GET("/gen/preview/:tableId", Preview)
	v1.GET("/gen/toproject/:tableId", GenCodeV3)
	v1.GET("/gen/todb/:tableId", GenMenuAndApi)
	v1.GET("/gen/tabletree", GetSysTablesTree)

	registerDBRouter(v1)
	registerSysTableRouter(v1)
	registerPublicRouter(v1)
	registerSysSettingRouter(v1)
}

func registerDBRouter(api *gin.RouterGroup) {
	db := api.Group("/db")
	{
		db.GET("/tables/page", GetDBTableList)
		db.GET("/columns/page", GetDBColumnList)
	}
}

func registerSysTableRouter(v1 *gin.RouterGroup) {
	systables := v1.Group("/sys/tables")
	{
		systables.GET("/page", GetSysTableList)
		tablesinfo := systables.Group("/info")
		{
			tablesinfo.POST("", InsertSysTable)
			tablesinfo.PUT("", UpdateSysTable)
			tablesinfo.DELETE("/:tableId", DeleteSysTables)
			tablesinfo.GET("/:tableId", GetSysTables)
			tablesinfo.GET("", GetSysTablesInfo)
		}
	}
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/login", authMiddleware.LoginHandler)
	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	r.Group("").Use(authMiddleware.MiddlewareFunc()).GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
	r.Group("").Use(authMiddleware.MiddlewareFunc()).GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	v1 := r.Group("/api/v1")

	//registerPageRouter(v1, authMiddleware)
	registerBaseRouter(v1, authMiddleware)
	registerDictRouter(v1, authMiddleware)
	//registerSysUserRouter(v1, authMiddleware)
	//registerUserCenterRouter(v1, authMiddleware)
}

func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := sys_menu.SysMenu{}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{
		//v1auth.GET("/getinfo", system.GetInfo)
		v1auth.GET("/roleMenuTreeselect/:roleId", api.GetMenuTreeSelect)
		v1.GET("/menuTreeselect", api.GetMenuTreeSelect)
		//v1auth.GET("/roleDeptTreeselect/:roleId", system.GetDeptTreeRoleselect)
		v1auth.POST("/logout", handler.LogOut)
	}
}

//func registerPageRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
//	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
//	{
//		v1auth.GET("/sysUserList", system.GetSysUserList)
//	}
//}

//func registerUserCenterRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
//	user := v1.Group("/user").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
//	{
//		user.GET("/profile", system.GetSysUserProfile)
//		user.POST("/avatar", system.InsetSysUserAvatar)
//		user.PUT("/pwd", system.SysUserUpdatePwd)
//	}
//}

//func registerPostRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
//	post := v1.Group("/post").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
//	{
//		post.GET("/:postId", system.GetPost)
//		post.POST("", system.InsertPost)
//		post.PUT("", system.UpdatePost)
//		post.DELETE("/:postId", system.DeletePost)
//	}
//}

//func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
//	sysuser := v1.Group("/sysUser").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
//	{
//		sysuser.GET("/:userId", system.GetSysUser)
//		sysuser.GET("/", system.GetSysUserInit)
//		sysuser.POST("", system.InsertSysUser)
//		sysuser.PUT("", system.UpdateSysUser)
//		sysuser.DELETE("/:userId", system.DeleteSysUser)
//	}
//}

func registerDictRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	dictApi := &dict.SysDictType{}
	dataApi := &dict.SysDictData{}
	dicts := v1.Group("/dict").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{

		dicts.GET("/data", dataApi.GetSysDictDataList)
		dicts.GET("/data/:dictCode", dataApi.GetSysDictData)
		dicts.POST("/data", dataApi.InsertSysDictData)
		dicts.PUT("/data/:dictCode", dataApi.UpdateSysDictData)
		dicts.DELETE("/data", dataApi.DeleteSysDictData)

		dicts.GET("/type-option-select", dictApi.GetSysDictTypeAll)
		dicts.GET("/type", dictApi.GetSysDictTypeList)
		dicts.GET("/type/:id", dictApi.GetSysDictType)
		dicts.POST("/type", dictApi.InsertSysDictType)
		dicts.PUT("/type/:id", dictApi.UpdateSysDictType)
		dicts.DELETE("/type", dictApi.DeleteSysDictType)
	}
	v1.Group("/dict").Use(authMiddleware.MiddlewareFunc()).GET("/data-all", dataApi.GetSysDictDataAll)
}

//func registerDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
//	dept := v1.Group("/dept").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
//	{
//		dept.GET("/:deptId", system.GetDept)
//		dept.POST("", system.InsertDept)
//		dept.PUT("", system.UpdateDept)
//		dept.DELETE("/:id", system.DeleteDept)
//	}
//}
func registerSysSettingRouter(v1 *gin.RouterGroup) {
	api := system.SysSetting{}
	setting := v1.Group("/setting")
	{
		setting.GET("", api.GetSetting)
		setting.POST("", api.CreateOrUpdateSetting)
		setting.GET("/serverInfo", monitor.ServerInfo)
	}
}

func registerPublicRouter(v1 *gin.RouterGroup) {
	p := v1.Group("/public")
	{
		p.POST("/uploadFile", public.UploadFile)
	}
}
