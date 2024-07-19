package picture_handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/utils/errs"
)

// ApiPoetryBookList 古诗绘本-列表
func (ph *PictureHandler) ApiPoetryBookList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	typeId, _ := strconv.Atoi(ctx.Query("type_id"))
	response, err := ph.service.FindPoetryBookList(page, limit, typeId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiPoetryBookInfo 古诗绘本-详情
func (ph *PictureHandler) ApiPoetryBookInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	response, err := ph.service.FindPoetryBookInfo(id)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
