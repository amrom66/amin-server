package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"amin/core/sdk/config"
	"amin/core/sdk/pkg"
	jwt "amin/core/sdk/pkg/jwtauth"
	"amin/core/sdk/pkg/jwtauth/user"
	"amin/core/sdk/pkg/response"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"

	"amin/app/admin/models/system"
	"amin/app/admin/service"
	"amin/common/apis"
)

var store = base64Captcha.DefaultMemStore

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)
		return jwt.MapClaims{
			jwt.IdentityKey:  u.UserId,
			jwt.RoleIdKey:    r.RoleId,
			jwt.RoleKey:      r.RoleKey,
			jwt.NiceKey:      u.Username,
			jwt.DataScopeKey: r.DataScope,
			jwt.RoleNameKey:  r.RoleName,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"RoleKey":     claims["rolekey"],
		"UserId":      claims["identity"],
		"RoleIds":     claims["roleid"],
		"DataScope":   claims["datascope"],
	}
}

// @Summary 登陆
// @Description 获取token
// @Description LoginHandler can be used by clients to get a jwt token.
// @Description Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// @Description Reply will be of the form {"token": "TOKEN"}.
// @Description dev mode：It should be noted that all fields cannot be empty, and a value of 0 can be passed in addition to the account password
// @Description 注意：开发模式：需要注意全部字段不能为空，账号密码外可以传入0值
// @Accept  application/json
// @Product application/json
// @Param account body models.Login  true "account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	log := apis.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		app.Error(c, http.StatusInternalServerError, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}

	var loginVals system.Login
	var status = "2"
	var msg = "登录成功"
	var username = ""
	defer func() {
		LoginLogToDB(c, status, msg, username)
	}()

	if err = c.ShouldBind(&loginVals); err != nil {
		username = loginVals.Username
		msg = "数据解析失败"
		status = "1"

		return nil, jwt.ErrMissingLoginValues
	}
	if config.ApplicationConfig.Mode != "dev" {
		if !store.Verify(loginVals.UUID, loginVals.Code, true) {
			username = loginVals.Username
			msg = "验证码错误"
			status = "1"

			return nil, jwt.ErrInvalidVerificationode
		}
	}
	user, role, e := loginVals.GetUser(db)
	if e == nil {
		username = loginVals.Username

		return map[string]interface{}{"user": user, "role": role}, nil
	} else {
		msg = "登录失败"
		status = "1"
		log.Warnf("%s login failed!", loginVals.Username)
	}
	return nil, jwt.ErrFailedAuthentication
}

// LoginLogToDB Write log to database
func LoginLogToDB(c *gin.Context, status string, msg string, username string) {
	log := apis.GetRequestLogger(c)
	if config.LoggerConfig.EnabledDB {
		var loginLog system.SysLoginLog
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Errorf("获取Orm失败, error:%s", err)
		}
		ua := user_agent.New(c.Request.UserAgent())
		loginLog.Ipaddr = c.ClientIP()
		loginLog.Username = username
		location := pkg.GetLocation(c.ClientIP())
		loginLog.LoginLocation = location
		loginLog.LoginTime = pkg.GetCurrentTime()
		loginLog.Status = status
		loginLog.Remark = c.Request.UserAgent()
		browserName, browserVersion := ua.Browser()
		loginLog.Browser = browserName + " " + browserVersion
		loginLog.Os = ua.OS()
		loginLog.Msg = msg
		loginLog.Platform = ua.Platform()
		serviceLoginLog := service.SysLoginLog{}
		serviceLoginLog.Orm = db
		_ = serviceLoginLog.InsertSysLoginLog(loginLog.Generate())
	}
}

// @Summary 退出登录
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "msg": "成功退出系统" }"
// @Router /logout [post]
// @Security Bearer
func LogOut(c *gin.Context) {
	log := apis.GetRequestLogger(c)
	var loginLog system.SysLoginLog
	ua := user_agent.New(c.Request.UserAgent())
	loginLog.Ipaddr = c.ClientIP()
	location := pkg.GetLocation(c.ClientIP())
	loginLog.LoginLocation = location
	loginLog.LoginTime = pkg.GetCurrentTime()
	loginLog.Status = "2"
	loginLog.Remark = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	loginLog.Browser = browserName + " " + browserVersion
	loginLog.Os = ua.OS()
	loginLog.Platform = ua.Platform()
	loginLog.Username = user.GetUserName(c)
	loginLog.Msg = "退出成功"
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("获取Orm失败, error:%s", err)
	}
	serviceLoginLog := service.SysLoginLog{}
	serviceLoginLog.Orm = db
	_ = serviceLoginLog.InsertSysLoginLog(loginLog.Generate())

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}

func Authorizator(data interface{}, c *gin.Context) bool {

	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)
		c.Set("role", r.RoleName)
		c.Set("roleIds", r.RoleId)
		c.Set("userId", u.UserId)
		c.Set("userName", u.Username)
		c.Set("dataScope", r.DataScope)

		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
