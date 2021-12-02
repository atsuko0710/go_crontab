package job

import (
	"master/internal/master/params"

	"master/internal/pkg/code"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"master/internal/pkg/validation"
)

func Create(c *gin.Context) {
	var params params.CreateJobRequest

	if err := c.ShouldBindJSON(&params); err != nil {
		// 判断错误是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			code.SendResponse(c, code.InvalidParam, gin.H{})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": validation.RemoveTopStruct(errs.Translate(validation.Trans)),
		})
		return
	}
}
