package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"amin/core/sdk/pkg"
	"amin/core/sdk/pkg/jwtauth/user"
	"amin/core/sdk/pkg/response"

	"amin/common/apis"
	"amin/common/dto"
	"amin/common/models"
)

// CreateAction 通用新增动作
func CreateAction(control dto.Control) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := apis.GetRequestLogger(c)
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		//新增操作
		req := control.Generate()
		err = req.Bind(c)
		if err != nil {
			app.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
			return
		}
		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			app.Error(c, http.StatusInternalServerError, err, "模型生成失败")
			return
		}
		object.SetCreateBy(user.GetUserId(c))
		err = db.WithContext(c).Create(object).Error
		if err != nil {
			log.Errorf("Create error: %s", err)
			app.Error(c, http.StatusInternalServerError, err, "创建失败")
			return
		}
		app.OK(c, object.GetId(), "创建成功")
		c.Next()
	}
}
