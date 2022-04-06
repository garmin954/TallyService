package main

import (
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

	// 初始化路由
	routes.Init(engine)

	engine.Run(":8080")
}
