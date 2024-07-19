package picture_handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/utils/errs"
)

// ApiAlbumBookIndex 专辑绘本-首页
func (ph *PictureHandler) ApiAlbumBookIndex(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	response, err := ph.service.FindAlbumBookIndex(page)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiAlbumBookList 专辑绘本-列表
func (ph *PictureHandler) ApiAlbumBookList(ctx *gin.Context) {
	bookId := ctx.Query("book_id")
	response, err := ph.service.FindAlbumBookList(bookId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiAlbumBookInfo 专辑绘本-详情
func (ph *PictureHandler) ApiAlbumBookInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	response, err := ph.service.FindAlbumBookInfo(id)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
