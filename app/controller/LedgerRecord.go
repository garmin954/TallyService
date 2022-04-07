package controller

import (
	"fmt"
	"smg/app/middleware"
	"smg/app/model"
	"smg/app/service"
	"smg/utils"

	"github.com/gin-gonic/gin"
)

type LedgerRecord struct{}

func (ledgerRecord *LedgerRecord) Router(engine *gin.RouterGroup) {
	group := engine.Group("/ledger-records")
	route := group.Use(middleware.JwtToken())

	route.GET("", ledgerRecord.list)
	route.POST("", ledgerRecord.create)
	route.PUT("", ledgerRecord.update)
	route.DELETE("", ledgerRecord.delete)
}

// 获取记录
func (ledgerRecord *LedgerRecord) list(ctx *gin.Context) {
	var s service.LedgerRecord
	pr, err := s.Pagination(ctx)
	if err != nil {
		utils.Failed(ctx, err.Error)
	}
	utils.Success(ctx, pr)
}

func (ledgerRecord *LedgerRecord) create(ctx *gin.Context) {
	var params model.LedgerRecord
	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.Failed(ctx, err.Error())
		return
	}

	params.Uid = int(ctx.Keys["uid"].(int))

	r := utils.DB.Table(model.LedgerRecordTable).Create(&params)
	if r.Error != nil {
		utils.Failed(ctx, r.Error.Error())
		return
	}

	utils.Success(ctx, params)
}

func (ledgerRecord *LedgerRecord) update(ctx *gin.Context) {
	utils.Success(ctx, ctx.Param("id"))

}

// 删除
func (ledgerRecord *LedgerRecord) delete(ctx *gin.Context) {
	fmt.Println(ctx.Param("id"), ctx.PostForm("id"))
	id := ctx.PostForm("id")
	if id == "" {
		utils.Failed(ctx, "ID 不存在")
		return
	}

	var record model.LedgerRecord
	r := utils.DB.Table(model.LedgerRecordTable).Where("id=?", id).Delete(&record)
	if r.Error != nil {
		utils.Failed(ctx, r.Error.Error())
		return
	}

	utils.Success(ctx, "")
}
