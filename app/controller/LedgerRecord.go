package controller

import (
<<<<<<< HEAD
	"fmt"
	"smg/app/middleware"
	"smg/app/model"
=======
	"smg/app/middleware"
>>>>>>> f5f42371f1c38e0ec7b08a45c1148b681c15329c
	"smg/app/service"
	"smg/utils"

	"github.com/gin-gonic/gin"
)

type LedgerRecord struct{}

func (ledgerRecord *LedgerRecord) Router(engine *gin.RouterGroup) {
	group := engine.Group("/ledger-records")
	route := group.Use(middleware.JwtToken())

	route.GET("", ledgerRecord.list)
<<<<<<< HEAD
	route.POST("", ledgerRecord.create)
	route.PUT("", ledgerRecord.update)
	route.DELETE("", ledgerRecord.delete)
=======
>>>>>>> f5f42371f1c38e0ec7b08a45c1148b681c15329c
}

// 获取记录
func (ledgerRecord *LedgerRecord) list(ctx *gin.Context) {
	var s service.LedgerRecord
<<<<<<< HEAD
	pr, err := s.Pagination(ctx)
=======
	pr, err := s.Pagination()
>>>>>>> f5f42371f1c38e0ec7b08a45c1148b681c15329c
	if err != nil {
		utils.Failed(ctx, err.Error)
	}
	utils.Success(ctx, pr)
}
<<<<<<< HEAD

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
=======
>>>>>>> f5f42371f1c38e0ec7b08a45c1148b681c15329c
