package download_service

import (
	"fmt"
	"strconv"
	"web/model"
	"web/model/picture"
	"web/utils"
)

func (ds *DownloadService) DownloadImgCover() {
	var list []picture.ChineseBook
	db := model.DefaultPicture().Model(&picture.ChineseBook{}).Debug().Where("type = 1 AND `status` = 1")
	db.Find(&list)

	for idx, _ := range list {
		path := "/Users/jiang/demo/cover/"
		filePathSave := path + strconv.Itoa(list[idx].Id) + ".jpg"
		fmt.Printf("开始下载：%#v\n", list[idx].Icon)
		utils.DownloadFile(list[idx].Icon, path, filePathSave)
	}
}
