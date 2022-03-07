package v1

import (
	"BLOGplus/middleware"
	"BLOGplus/model"
	"BLOGplus/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int
	c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, _ = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
