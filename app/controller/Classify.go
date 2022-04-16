package controller

import (
	"smg/app/model"
	"smg/utils"

	"github.com/gin-gonic/gin"
)

type Classify struct{}

func (classify *Classify) Router(engine *gin.RouterGroup) {
	group := engine.Group("/classify")
	group.GET("", classify.classifies)
}

func (classify *Classify) classifies(ctx *gin.Context) {
	uid := ctx.Keys["uid"]
	var row []model.Classify

	r := utils.DB.Table(model.ClassifyTable).Where("uid in (?,?)", uid, 0).Find(&row)
	if r.Error != nil {
		utils.Failed(ctx, r.Error)
		return
	}

	utils.Success(ctx, row)
}
