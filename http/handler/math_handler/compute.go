package math_handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/utils/errs"
)

func (mh *MathHandler) ApiMakeCompute(ctx *gin.Context) {
	level, _ := strconv.Atoi(ctx.Query("level"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	response, err := mh.service.MakeComputeList(level, limit)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
