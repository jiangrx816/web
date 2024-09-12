package picture_service

import (
	"github.com/jiangrx816/gopkg/server/api"
	"math"
	"sort"
	"web/common"
	pictureResp "web/common/response/picture"
	"web/model"
	"web/model/picture"
	"web/utils"
)

func (ps *PictureService) FindBookNavList(typeId int) (response pictureResp.ChineseBookNavNameResponse, apiErr api.Error) {
	utils.DefaultIntOne(&typeId)
	db := model.DefaultPicture().Model(&picture.ChineseBookNavName{}).Debug()
	db = db.Where("status = 1 and s_type = ?", typeId)
	db = db.Order("s_sort asc").Order("id asc")
	db.Find(&response.List)
	return
}

func (ps *PictureService) FindChineseBookList(page, level int) (response pictureResp.ChineseBookResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	utils.DefaultIntOne(&level)
	size := common.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	var bookList []picture.ChineseBook
	db := model.DefaultPicture().Model(&picture.ChineseBook{}).Debug()
	db = db.Where("type = ? and status = 1", level).Count(&response.Total)
	db = db.Order("position desc").Limit(size).Offset(offset)
	db.Find(&bookList)

	var bookInfoCountList []pictureResp.ResponseBookInfoCount
	db1 := model.DefaultPicture().Model(&picture.ChineseBookInfo{}).Debug()
	db1.Raw("SELECT book_id,count(id) as book_count FROM s_chinese_picture_info where status = 1 GROUP BY book_id").Scan(&bookInfoCountList)

	var temp pictureResp.ResponseChineseBook
	for _, item := range bookList {
		temp.Id = item.Id
		temp.BookId = item.BookId
		temp.Title = item.Title
		temp.Icon = item.Icon
		temp.Level = item.Type
		temp.Position = item.Position
		response.List = append(response.List, temp)
	}
	for index, item := range response.List {
		for _, it := range bookInfoCountList {
			if item.BookId == it.BookId {
				response.List[index].BookCount = it.BookCount
			}
		}
	}
	sort.Slice(response.List, func(i, j int) bool {
		if response.List[i].Position > response.List[j].Position {
			return true
		}
		return response.List[i].Position == response.List[j].Position && response.List[i].Id < response.List[j].Id
	})
	response.Page = page
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}

func (ps *PictureService) FindChineseBookInfo(bookId string) (response pictureResp.ChineseBookInfoResponse, apiErr api.Error) {
	db := model.DefaultPicture().Model(&picture.ChineseBookInfo{}).Debug()
	db = db.Where("book_id = ? and status = 1", bookId).Order("position asc").Find(&response.Info)
	return
}

func (ps *PictureService) FindSearchChineseBookList(page int, value string) (response pictureResp.ChineseBookResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	size := 100
	offset := size * (page - 1)

	var bookList []picture.ChineseBook
	db := model.DefaultPicture().Model(&picture.ChineseBook{}).Debug()
	db = db.Where("status = 1 and title like ?", "%"+value+"%")
	db = db.Count(&response.Total)
	db = db.Order("position desc").Limit(size).Offset(offset)
	db.Find(&bookList)

	var bookInfoCountList []pictureResp.ResponseBookInfoCount
	db1 := model.DefaultPicture().Model(&picture.ChineseBookInfo{}).Debug()
	db1.Raw("SELECT book_id,count(id) as book_count FROM s_chinese_picture_info GROUP BY book_id").Scan(&bookInfoCountList)

	var temp pictureResp.ResponseChineseBook
	for _, item := range bookList {
		temp.Id = item.Id
		temp.BookId = item.BookId
		temp.Title = item.Title
		temp.Icon = item.Icon
		temp.Level = item.Type
		temp.Position = item.Position
		response.List = append(response.List, temp)
	}
	for index, item := range response.List {
		for _, it := range bookInfoCountList {
			if item.BookId == it.BookId {
				response.List[index].BookCount = it.BookCount
			}
		}
	}
	sort.Slice(response.List, func(i, j int) bool {
		if response.List[i].Position > response.List[j].Position {
			return true
		}
		return response.List[i].Position == response.List[j].Position && response.List[i].Id < response.List[j].Id
	})
	response.Page = page
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}
