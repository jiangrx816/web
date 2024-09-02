package download_service

import (
	"web/model"
	"web/model/picture"
)

func (ds *DownloadService) DownloadImageList() {
	var list []picture.ChineseBook
	db := model.DefaultPicture().Model(&picture.ChineseBook{}).Debug().Where("icon_1 != ''")
	db.Find(&list)

	for idx, _ := range list {
		var temp picture.ChineseBook
		//path := "/Users/jiang/demo/cover/"
		//filePathSave := path + strconv.Itoa(list[idx].Id) + ".jpg"
		//utils.DownloadFile(list[idx].Icon_1, path, filePathSave)
		temp.Icon_1 = ""
		db1 := model.DefaultPicture().Model(&picture.ChineseBook{}).Debug().Where("id = ?", list[idx].Id)
		db1.Updates(&temp)
	}
}
