package controller

import (
	"net/http"
	"smg/app/model"
	"smg/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 路由
func (user *UserController) Router(engine *gin.Engine) {
	engine.GET("/api/query_users", user.queryUsers)
}

func (user *UserController) queryUsers(ctx *gin.Context) {
	var results []map[string]interface{}
	utils.DB.Table(model.UserTable).Select("nickname,id").Find(&results)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "成功！先把",
		"data": results,
	})
}
