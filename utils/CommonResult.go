package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 1
	FAILED  int = 0
)

func Success(cxt *gin.Context, v interface{}) {
	cxt.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}

func Failed(cxt *gin.Context, v interface{}) {
	cxt.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  v,
	})
}
