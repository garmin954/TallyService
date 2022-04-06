package main

import (
	"net/http"
	"smg/app/model"
	"smg/routes"
	"smg/utils"

	"github.com/gin-gonic/gin"
)

type InsertParam struct {
	Name string `form:"name"`
	Pass string `form:"pass"`
}

func main() {

	utils.ConnectDB()
	engine := gin.Default()

	routes.Init()
	engine.GET("/", func(ctx *gin.Context) {
		var results []map[string]interface{}
		utils.DB.Table(model.UserTable).Select("nickname,id").Find(&results)

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "成功！先把",
			"data": results,
		})
	})

	engine.POST("/insert", func(ctx *gin.Context) {
		var params InsertParam
		ctx.ShouldBind(&params)

		user := model.User{Nickname: params.Name, Unionid: params.Pass}
		utils.DB.Create(&user)

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "成功！",
			"data": params,
		})
	})
	engine.Run()

}
