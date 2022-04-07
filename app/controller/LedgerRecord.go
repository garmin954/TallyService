package controller

import (
	"smg/app/middleware"
	"smg/app/service"
	"smg/utils"

	"github.com/gin-gonic/gin"
)

type LedgerRecord struct{}

func (ledgerRecord *LedgerRecord) Router(engine *gin.RouterGroup) {
	group := engine.Group("/ledger-records")
	route := group.Use(middleware.JwtToken())

	route.GET("", ledgerRecord.list)
}

// 获取记录
func (ledgerRecord *LedgerRecord) list(ctx *gin.Context) {
	var s service.LedgerRecord
	pr, err := s.Pagination()
	if err != nil {
		utils.Failed(ctx, err.Error)
	}
	utils.Success(ctx, pr)
}
