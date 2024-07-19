package poetry_service

import (
	"github.com/jiangrx816/gopkg/server/api"
	"math"
	"web/common"
	poetryResp "web/common/response/poetry"
	"web/model"
	"web/model/poetry"
	"web/utils"
)

func (ps *PoetryService) FindDynastyList() (response poetryResp.DynastyResponse, apiErr api.Error) {
	var dynastyModelList []poetry.RXDynasty
	db := model.DefaultPoetry().Model(&poetry.RXDynasty{}).Debug().Where("status = 1")
	db = db.Order("s_sort desc").Order("id asc")
	db.Find(&dynastyModelList)

	var dynastyTemp poetryResp.ResponseDynasty
	for idx, _ := range dynastyModelList {
		dynastyTemp.DynastyName = dynastyModelList[idx].Name
		response.List = append(response.List, dynastyTemp)
	}
	return
}

func (ps *PoetryService) FindAuthorList(dynasty string, page int) (response poetryResp.AuthorResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	size := common.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	db := model.DefaultPoetry().Model(&poetry.RXAuthors{}).Debug()
	if dynasty != "" {
		db = db.Where("dynasty like ?", dynasty+"%")
	}
	db = db.Count(&response.Total)
	db = db.Order("id asc").Limit(size).Offset(offset)
	db.Find(&response.List)

	response.Page = page
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}

func (ps *PoetryService) FindAuthorInfo(name string, authorId int) (response poetryResp.AuthorInfoResponse, apiErr api.Error) {
	db := model.DefaultPoetry().Model(&poetry.RXAuthors{}).Debug()
	if authorId > 0 {
		db = db.Where("author_id = ?", authorId)
	} else {
		if name != "" {
			db = db.Where("name like ?", name+"%")
		} else {
			db = db.Where("author_id = 0")
		}
	}
	db.First(&response.Info)

	return
}

func (ps *PoetryService) FindKindList() (response poetryResp.KindResponse, apiErr api.Error) {
	db := model.DefaultPoetry().Model(&poetry.RXCollectionKinds{}).Debug()
	db = db.Order("sort asc").Order("id asc")
	db.Find(&response.List)
	return
}

func (ps *PoetryService) FindCollectionList(kindId int) (response poetryResp.CollectionResponse, apiErr api.Error) {
	db := model.DefaultPoetry().Model(&poetry.RXCollections{}).Debug().Where("kind_id = ?", kindId)
	db = db.Order("sort asc").Order("id asc")
	db.Find(&response.List)
	return
}

func (ps *PoetryService) FindCollectionQuotesList(collectionId, kindId, page int) (response poetryResp.CollectionQuotesResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	size := common.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	var collectionQuoteList []poetry.RXCollectionQuotes
	db := model.DefaultPoetry().Model(&poetry.RXCollectionQuotes{}).Debug()
	if collectionId > 0 {
		db = db.Where("collection_id = ?", collectionId)
	}
	if kindId > 0 {
		db = db.Where("collection_kind_id = ?", kindId)
	}
	db = db.Count(&response.Total)
	db = db.Order("sort asc").Order("id asc").Limit(size).Offset(offset)
	db.Find(&collectionQuoteList)

	var collectionIds, collectionKindIds []int
	for index, _ := range collectionQuoteList {
		collectionIds = append(collectionIds, collectionQuoteList[index].CollectionId)
		collectionKindIds = append(collectionKindIds, collectionQuoteList[index].CollectionKindId)
	}
	collectionIds = utils.RemoveDuplicateElement(collectionIds)
	collectionKindIds = utils.RemoveDuplicateElement(collectionKindIds)

	var collectionKindList []poetry.RXCollectionKinds
	db1 := model.DefaultPoetry().Model(&poetry.RXCollectionKinds{}).Debug()
	db1.Where("collection_kind_id in (?)", collectionKindIds).Find(&collectionKindList)

	var collectionList []poetry.RXCollections
	db2 := model.DefaultPoetry().Model(&poetry.RXCollections{}).Debug()
	db2.Where("collection_id in (?)", collectionIds).Find(&collectionList)

	var collectionQuotesItem poetryResp.CollectionQuotesItem
	for idx, _ := range collectionQuoteList {
		collectionQuotesItem.QuoteId = collectionQuoteList[idx].QuoteId
		collectionQuotesItem.Quote = collectionQuoteList[idx].Quote
		collectionQuotesItem.QuoteAuthor = collectionQuoteList[idx].QuoteAuthor
		collectionQuotesItem.QuoteWork = collectionQuoteList[idx].QuoteWork
		collectionQuotesItem.QuoteWorkId = collectionQuoteList[idx].QuoteWorkId
		collectionQuotesItem.CollectionId = collectionQuoteList[idx].CollectionId
		for index, _ := range collectionList {
			if collectionQuoteList[idx].CollectionId == collectionList[index].CollectionId {
				collectionQuotesItem.Collection = collectionList[index].Name
			}
		}
		collectionQuotesItem.CollectionKindId = collectionQuoteList[idx].CollectionKindId
		for index, _ := range collectionKindList {
			if collectionQuoteList[idx].CollectionKindId == collectionKindList[index].CollectionKindId {
				collectionQuotesItem.Kind = collectionKindList[index].Name
			}
		}
		response.List = append(response.List, collectionQuotesItem)
	}

	response.Page = page
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}

func (ps *PoetryService) FindQuotesInfo(quoteId int) (response poetryResp.QuotesResponse, apiErr api.Error) {
	db := model.DefaultPoetry().Model(&poetry.RXQuotes{}).Debug().Where("quote_id = ?", quoteId)
	db.First(&response.Info)
	return
}

func (ps *PoetryService) FindCollectionWorksList(collectionId, page int, workKind string) (response poetryResp.CollectionWorksResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	size := common.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	db := model.DefaultPoetry().Model(&poetry.RXCollectionWorks{}).Debug()
	if collectionId > 0 {
		db = db.Where("collection_id = ?", collectionId)
	}
	if workKind != "" {
		db = db.Where("work_kind = ?", workKind)
	}
	db = db.Count(&response.Total)
	db = db.Order("sort asc").Order("id asc").Limit(size).Offset(offset)
	db.Find(&response.List)

	response.Page = page
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}

func (ps *PoetryService) FindWorksInfo(workId int) (response poetryResp.WorksResponse, apiErr api.Error) {
	db := model.DefaultPoetry().Model(&poetry.RXWorks{}).Debug().Where("work_id = ?", workId)
	db.First(&response.Info)
	return
}

func (ps *PoetryService) FindSearchQuotes(quote, author string, page int) (response poetryResp.CollectionQuotesResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	size := common.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	var collectionQuoteList []poetry.RXCollectionQuotes
	db := model.DefaultPoetry().Model(&poetry.RXCollectionQuotes{}).Debug()

	if quote != "" {
		db = db.Where("quote like ?", quote+"%")
	}
	if author != "" {
		db = db.Where("quote_author like ?", author+"%")
	}
	db = db.Group("quote_id")
	db = db.Count(&response.Total)
	db = db.Order("sort asc").Order("id asc").Limit(size).Offset(offset)
	db.Find(&collectionQuoteList)

	var collectionIds, collectionKindIds []int
	for index, _ := range collectionQuoteList {
		collectionIds = append(collectionIds, collectionQuoteList[index].CollectionId)
		collectionKindIds = append(collectionKindIds, collectionQuoteList[index].CollectionKindId)
	}
	collectionIds = utils.RemoveDuplicateElement(collectionIds)
	collectionKindIds = utils.RemoveDuplicateElement(collectionKindIds)

	var collectionKindList []poetry.RXCollectionKinds
	db1 := model.DefaultPoetry().Model(&poetry.RXCollectionKinds{}).Debug()
	db1.Where("collection_kind_id in (?)", collectionKindIds).Find(&collectionKindList)

	var collectionList []poetry.RXCollections
	db2 := model.DefaultPoetry().Model(&poetry.RXCollections{}).Debug()
	db2.Where("collection_id in (?)", collectionIds).Find(&collectionList)

	var collectionQuotesItem poetryResp.CollectionQuotesItem
	for idx, _ := range collectionQuoteList {
		collectionQuotesItem.QuoteId = collectionQuoteList[idx].QuoteId
		collectionQuotesItem.Quote = collectionQuoteList[idx].Quote
		collectionQuotesItem.QuoteAuthor = collectionQuoteList[idx].QuoteAuthor
		collectionQuotesItem.QuoteWork = collectionQuoteList[idx].QuoteWork
		collectionQuotesItem.QuoteWorkId = collectionQuoteList[idx].QuoteWorkId
		collectionQuotesItem.CollectionId = collectionQuoteList[idx].CollectionId
		for index, _ := range collectionList {
			if collectionQuoteList[idx].CollectionId == collectionList[index].CollectionId {
				collectionQuotesItem.Collection = collectionList[index].Name
			}
		}
		collectionQuotesItem.CollectionKindId = collectionQuoteList[idx].CollectionKindId
		for index, _ := range collectionKindList {
			if collectionQuoteList[idx].CollectionKindId == collectionKindList[index].CollectionKindId {
				collectionQuotesItem.Kind = collectionKindList[index].Name
			}
		}
		response.List = append(response.List, collectionQuotesItem)
	}

	response.Page = page
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}

func (ps *PoetryService) FindSearchWorks(title, author string, page int) (response poetryResp.CollectionWorksResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	size := common.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	db := model.DefaultPoetry().Model(&poetry.RXCollectionWorks{}).Debug()
	if title != "" {
		db = db.Where("work_title like ?", title+"%")
	}
	if author != "" {
		db = db.Where("work_author like ?", author+"%")
	}
	db = db.Group("work_id")
	db = db.Count(&response.Total)
	db = db.Order("sort asc").Order("id asc").Limit(size).Offset(offset)
	db.Find(&response.List)

	response.Page = page
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}
