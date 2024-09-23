package download_service

import (
	"fmt"
	"path"
	"strconv"
	"sync"
	"web/model"
	"web/model/picture"
	"web/utils"
)

func (ds *DownloadService) DownloadImgInfo() {
	var list []picture.ChineseBook
	db := model.DefaultPicture().Model(&picture.ChineseBook{}).Debug().Where("type = 1 AND `status` = 1")
	db.Find(&list)

	var bookIds []string
	for _, item := range list {
		bookIds = append(bookIds, item.BookId)
	}

	var listInfo []picture.ChineseBookInfo
	dbInfo := model.DefaultPicture().Model(&picture.ChineseBookInfo{}).Debug().Where("book_id in(?)", bookIds)
	dbInfo.Find(&listInfo)

	var wg sync.WaitGroup
	for _, item := range listInfo {
		wg.Add(1)
		go func(info picture.ChineseBookInfo) {
			defer wg.Done()
			dir := "/Users/jiang/demo/file/"

			pic := path.Base(info.Pic)
			////mp3 := path.Base(info.Mp3)
			bookOldId := info.BookIdOld
			fileDir := dir + strconv.Itoa(bookOldId) + "/"

			picPath := fileDir + pic
			//mp3Path := realDir + mp3

			fmt.Printf("开始下载：pic->%#v mp3:->%#v \n", info.Pic, info.Mp3)

			utils.DownloadFile(info.Pic, fileDir, picPath)
			//utils.DownloadFile(info.Mp3, realDir, mp3Path)
		}(item)
	}
	wg.Wait()
}
