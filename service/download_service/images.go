package download_service

import (
	"strconv"
	"web/model"
	"web/model/picture"
	"web/utils"
)

func (ds *DownloadService) DownloadImageList() {
	var list []picture.ChineseBook
	db := model.DefaultPicture().Model(&picture.ChineseBook{}).Debug().Where("type = -1")
	db.Find(&list)

	for idx, _ := range list {
		if list[idx].Icon_1 != "" {
			path := "/Users/jiang/demo/cover/"
			filePathSave := path + strconv.Itoa(list[idx].Id) + ".jpg"

			utils.DownloadFile(list[idx].Icon_1, path, filePathSave)
		}
	}
}
