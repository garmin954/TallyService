package service

import (
	"smg/app/model"
	"smg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LedgerRecord struct{}

// 获取记录
func (ledgerRecord *LedgerRecord) Pagination(ctx *gin.Context) (pg Pagination, err error) {

	ps, _ := strconv.Atoi(ctx.DefaultQuery("page_num", "1"))
	pn, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	var row []model.LedgerRecord
	uid := int(ctx.Keys["uid"].(int))
	where := utils.DB.Table(model.LedgerRecordTable).Where("uid=?", uid)

	where.Count(&pg.Total)
	where.Offset((ps - 1) * pn).Limit(pn).Find(&row)
	pg.List = row

	pg.PageNum = pn
	pg.PageSize = ps

	return pg, nil
}
