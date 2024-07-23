package download_service

import (
	"fmt"
	rxLog "github.com/jiangrx816/gopkg/log"
	"strconv"
	"strings"
	"sync"
	"web/utils"
)

var counter int = 10
var baseStep int = 10000
var step int = 1000
var imgBaseUrl string = ""

func (ds *DownloadService) DownloadImgList() {
	var wg sync.WaitGroup

	wg.Add(counter)
	for i := 0; i < counter; i++ {
		go ds.downloadFile(i, baseStep, step, &wg)
	}
	wg.Wait()

	rxLog.SugarContext(utils.BuildRequestIdCtx()).Infow("执行完毕")
}

func (ds *DownloadService) downloadFile(id, baseStep, step int, wg *sync.WaitGroup) {
	defer wg.Done()
	startIndex := baseStep + id*step
	endIndex := startIndex + (id+1)*step
	ds.fileUrlSaveToLog(startIndex, endIndex)
	fmt.Printf("goroutine:%#v startIndex:%d,endIndex:%d\n", id, startIndex, endIndex)
}

func (ds *DownloadService) fileUrlSaveToLog(startIndex, endIndex int) {
	fileLog := "/Users/jiang/demo/file.log"
	for bookId := startIndex; bookId < endIndex; bookId++ {
		bookIdStr := strconv.Itoa(bookId)
		imgPath := imgBaseUrl + bookIdStr
		for i := 1; i <= 30; i++ {
			filePathBg := imgPath + "/extract/page_bg_" + strconv.Itoa(i) + ".jpg"
			if exists := utils.UrlExists(filePathBg); !exists {
				continue
			}
			filePath := imgPath + "/extract/page_sub_" + strconv.Itoa(i) + ".png"
			utils.AppendToFile(fileLog, filePathBg)
			utils.AppendToFile(fileLog, filePath)

			mp3Path := imgPath + "/extract/page_audio_" + strconv.Itoa(i) + ".mp3"
			if exists := utils.UrlExists(mp3Path); exists {
				utils.AppendToFile(fileLog, mp3Path)
			}
		}
	}
}

func (ds *DownloadService) downloadFileExec(startIndex, endIndex int) {
	for bookId := startIndex; bookId < endIndex; bookId++ {
		bookIdStr := strconv.Itoa(bookId)
		imgPath := imgBaseUrl + bookIdStr
		for i := 1; i <= 30; i++ {
			filePathBg := imgPath + "/extract/page_bg_" + strconv.Itoa(i) + ".jpg"
			if exists := utils.UrlExists(filePathBg); !exists {
				continue
			}
			filePath := imgPath + "/extract/page_sub_" + strconv.Itoa(i) + ".png"

			bookIdInt := ds.findFileBookId(filePathBg)
			path := "/Users/jiang/demo/book/" + bookIdInt

			fileBgName := ds.findFileName(filePathBg)
			filePathBgSave := path + "/" + fileBgName
			err2 := utils.DownloadFile(filePathBg, path, filePathBgSave)
			if err2 != nil {
				fmt.Println("download err:", err2)
			}
			fmt.Println("download success:", filePathBg)

			fileName := ds.findFileName(filePath)
			filePathSave := path + "/" + fileName
			err := utils.DownloadFile(filePath, path, filePathSave)
			if err != nil {
				fmt.Println("download err:", err)
			}
			fmt.Println("download success:", filePath)

			mp3Path := imgPath + "/extract/page_audio_" + strconv.Itoa(i) + ".mp3"
			if exists := utils.UrlExists(mp3Path); exists {
				fileMp3Name := ds.findFileName(mp3Path)
				fileMp3 := path + "/" + fileMp3Name
				err3 := utils.DownloadFile(mp3Path, path, fileMp3)
				if err3 != nil {
					fmt.Println("download err:", err3)
				}
				fmt.Println("download success:", mp3Path)
			}
		}
	}
}

func (ds *DownloadService) findFileBookId(url string) string {
	// 找到book
	bookIDStart := strings.Index(url, "/cbook/") + len("/cbook/")
	bookIDEnd := strings.Index(url[bookIDStart:], "/") + bookIDStart
	bookID := url[bookIDStart:bookIDEnd]
	return bookID
}

func (ds *DownloadService) findFileName(url string) string {
	// 找到文件名
	fileNameStart := strings.LastIndex(url, "/") + 1
	fileName := url[fileNameStart:]
	return fileName
}
