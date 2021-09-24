package system

import (
	"github.com/gin-gonic/gin"
	"amin/core/sdk/pkg"
	"amin/core/sdk/pkg/captcha"
	"amin/core/sdk/pkg/response"
)

func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	pkg.HasError(err, "验证码获取失败", 500)
	app.Custum(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
