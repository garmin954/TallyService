package controller

import (
	"smg/app/middleware"
	"smg/app/model"
	"smg/utils"

	"github.com/gin-gonic/gin"
)

type Ledger struct{}

type CreateParams struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Desc  string `form:"desc"`
	Cover string `form:"cover"`
	Uid   int    `form:"uid" json:"uid"`
}

// 路由
func (ledger *Ledger) Router(engine *gin.RouterGroup) {
	group := engine.Group("/ledger")
	route := group.Use(middleware.JwtToken())
	route.GET("", ledger.info)
	route.POST("", ledger.create)
	route.GET("/user", ledger.ledgers)
}

// 创建
func (ledger *Ledger) create(ctx *gin.Context) {
	params := CreateParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Failed(ctx, err.Error())
		return
	}

	uid := int(ctx.Keys["uid"].(int))
	params.Uid = uid

	r := utils.DB.Table(model.LedgerTable).Create(&params)
	if r.Error != nil {
		utils.Failed(ctx, r.Error.Error())
		return
	}

	utils.Success(ctx, params)
}

// 获取信息
func (ledger *Ledger) info(ctx *gin.Context) {
	utils.Failed(ctx, "success")
}

// 获取我的账本
func (ledger *Ledger) ledgers(ctx *gin.Context) {
	uid := int(ctx.Keys["uid"].(int))
	var result []model.Ledger
	r := utils.DB.Table(model.LedgerTable).Where("uid=?", uid).Find(&result)

	if r.Error != nil {
		utils.Failed(ctx, r.Error.Error())
		return
	}

	utils.Success(ctx, result)
}
