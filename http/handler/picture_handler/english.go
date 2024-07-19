package picture_handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/utils/errs"
)

// ApiEnglishBookList 英语绘本-列表
func (ph *PictureHandler) ApiEnglishBookList(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	typeId, _ := strconv.Atoi(ctx.Query("type_id"))
	response, err := ph.service.FindEnglishBookList(typeId, offset)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiEnglishBookInfo 英语绘本-详情
func (ph *PictureHandler) ApiEnglishBookInfo(ctx *gin.Context) {
	bookId, _ := strconv.Atoi(ctx.Query("book_id"))
	response, err := ph.service.FindEnglishBookInfo(bookId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
