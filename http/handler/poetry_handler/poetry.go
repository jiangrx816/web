package poetry_handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/utils/errs"
)

// ApiDynastyList 朝代-列表
func (ph *PoetryHandler) ApiDynastyList(ctx *gin.Context) {
	response, err := ph.service.FindDynastyList()
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiAuthorList 作者列表
func (ph *PoetryHandler) ApiAuthorList(ctx *gin.Context) {
	dynasty := ctx.Query("dynasty")
	page, _ := strconv.Atoi(ctx.Query("page"))
	response, err := ph.service.FindAuthorList(dynasty, page)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiAuthorInfo 作者详情
func (ph *PoetryHandler) ApiAuthorInfo(ctx *gin.Context) {
	name := ctx.Query("name")
	authorId, _ := strconv.Atoi(ctx.Query("author_id"))
	response, err := ph.service.FindAuthorInfo(name, authorId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiKindList 集合类别列表
func (ph *PoetryHandler) ApiKindList(ctx *gin.Context) {
	response, err := ph.service.FindKindList()
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiCollectionList 类别集合列表
func (ph *PoetryHandler) ApiCollectionList(ctx *gin.Context) {
	kindId, _ := strconv.Atoi(ctx.Query("kind_id"))
	response, err := ph.service.FindCollectionList(kindId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiQuotesList 名言警句列表
func (ph *PoetryHandler) ApiQuotesList(ctx *gin.Context) {
	collectionId, _ := strconv.Atoi(ctx.Query("collection_id"))
	kindId, _ := strconv.Atoi(ctx.Query("kind_id"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	response, err := ph.service.FindCollectionQuotesList(collectionId, kindId, page)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiQuotesInfo 名言警句详情
func (ph *PoetryHandler) ApiQuotesInfo(ctx *gin.Context) {
	quoteId, _ := strconv.Atoi(ctx.Query("quote_id"))
	response, err := ph.service.FindQuotesInfo(quoteId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiWorksList 集合作品列表
func (ph *PoetryHandler) ApiWorksList(ctx *gin.Context) {
	collectionId, _ := strconv.Atoi(ctx.Query("collection_id"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	workKind := ctx.Query("work_kind")
	response, err := ph.service.FindCollectionWorksList(collectionId, page, workKind)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiWorksInfo 集合作品详情
func (ph *PoetryHandler) ApiWorksInfo(ctx *gin.Context) {
	workId, _ := strconv.Atoi(ctx.Query("work_id"))
	response, err := ph.service.FindWorksInfo(workId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiSearchQuotes 名句搜索
func (ph *PoetryHandler) ApiSearchQuotes(ctx *gin.Context) {
	quote := ctx.Query("quote")
	author := ctx.Query("author")
	page, _ := strconv.Atoi(ctx.Query("page"))
	response, err := ph.service.FindSearchQuotes(quote, author, page)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiSearchWorks 古诗搜索
func (ph *PoetryHandler) ApiSearchWorks(ctx *gin.Context) {
	title := ctx.Query("title")
	author := ctx.Query("author")
	page, _ := strconv.Atoi(ctx.Query("page"))
	response, err := ph.service.FindSearchWorks(title, author, page)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
