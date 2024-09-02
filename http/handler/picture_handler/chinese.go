package picture_handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/utils/errs"
)

// ApiChineseBookList 汉语绘本-列表
func (ph *PictureHandler) ApiChineseBookList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	level, _ := strconv.Atoi(ctx.Query("level"))
	response, err := ph.service.FindChineseBookList(page, level)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiChineseBookInfo 汉语绘本-详情
func (ph *PictureHandler) ApiChineseBookInfo(ctx *gin.Context) {
	bookId := ctx.Query("book_id")
	response, err := ph.service.FindChineseBookInfo(bookId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiChineseSearchList 汉语绘本-搜索列表
func (ph *PictureHandler) ApiChineseSearchList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	value := ctx.Query("value")
	response, err := ph.service.FindSearchChineseBookList(page, value)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
