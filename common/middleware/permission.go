package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"amin/core/sdk"
	"amin/core/sdk/pkg/jwtauth"
	"amin/core/sdk/pkg/response"

	"amin/common/apis"
)

//权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := apis.GetRequestLogger(c)
		data, _ := c.Get(jwtauth.JwtPayloadKey)
		v := data.(jwtauth.MapClaims)
		e := sdk.Runtime.GetCasbinKey(c.Request.Host)
		var res bool
		var err error
		//检查权限
		if v["rolekey"] == "admin" {
			res = true
			log.Infof("info:%s method:%s path:%s", v["rolekey"], c.Request.Method, c.Request.URL.Path)
		} else {
			res, err = e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
			if err != nil {
				log.Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
				app.Error(c, 500, err, "")
				return
			}
		}

		if res {
			log.Infof("isTrue: %v role: %s method: %s path: %s", res, v["rolekey"], c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			log.Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, v["rolekey"], c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}
