package picture_service

import (
	"github.com/jiangrx816/gopkg/server/api"
	"math"
	"web/common"
	pictureResp "web/common/response/picture"
	"web/model"
	"web/model/picture"
	"web/utils"
)

func (ps *PictureService) FindAlbumBookIndex(page int) (response pictureResp.AlbumBookIndexResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	size := common.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	db := model.DefaultPicture().Model(&picture.AlbumPicture{}).Debug()
	db = db.Where("status = 1").Count(&response.Total)
	db = db.Order("position desc").Limit(size).Offset(offset)
	db.Find(&response.List)

	response.Page = page
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}

func (ps *PictureService) FindAlbumBookList(bookId string) (response pictureResp.AlbumBookListResponse, apiErr api.Error) {
	size := common.DEFAULT_PAGE_SIZE

	db := model.DefaultPicture().Model(&picture.AlbumPictureInfo{}).Debug()
	db = db.Where("book_id = ?", bookId).Count(&response.Total)
	db = db.Order("position desc")
	db.Find(&response.List)

	response.Page = 1
	response.TotalPage = int(math.Ceil(float64(response.Total) / float64(size)))
	return
}

func (ps *PictureService) FindAlbumBookInfo(id int) (response pictureResp.AlbumPictureInfoResponse, apiErr api.Error) {
	utils.DefaultIntOne(&id)
	db := model.DefaultPicture().Model(&picture.AlbumPictureInfo{}).Debug()
	db = db.Where("id = ?", id).First(&response.Info)

	return
}
