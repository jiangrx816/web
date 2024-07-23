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
var baseStep int = 0
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
	endIndex := baseStep + (id+1)*step
	ds.downloadFileExec(startIndex, endIndex)
	fmt.Printf("goroutine:%#v startIndex:%d,endIndex:%d\n", id, startIndex, endIndex)
}

func (ds *DownloadService) downloadFileExec(startIndex, endIndex int) {
	const maxConcurrentDownloads = 10
	semaphore := make(chan struct{}, maxConcurrentDownloads)
	var wg sync.WaitGroup

	for bookId := startIndex; bookId < endIndex; bookId++ {
		bookIdStr := strconv.Itoa(bookId)
		existDir := "/Users/jiang/demo/book/" + bookIdStr
		directory, err := utils.IsDirectory(existDir)
		if err != nil {
			fmt.Printf("Error checking directory: %v\n", err)
			continue
		}
		if directory {
			fmt.Printf("existDir: %v 存在\n", existDir)
			continue
		}

		imgPath := imgBaseUrl + bookIdStr
		for i := 1; i <= 30; i++ {
			filePathBg := imgPath + "/extract/page_bg_" + strconv.Itoa(i) + ".jpg"
			if exists := utils.UrlExists(filePathBg); !exists {
				continue
			}
			filePath := imgPath + "/extract/page_sub_" + strconv.Itoa(i) + ".png"
			mp3Path := imgPath + "/extract/page_audio_" + strconv.Itoa(i) + ".mp3"

			wg.Add(1)
			semaphore <- struct{}{} // Acquire semaphore
			go func(filePathBg, filePath, mp3Path string) {
				defer wg.Done()
				defer func() { <-semaphore }() // Release semaphore

				ds.downloadFileContent(filePathBg, filePath, mp3Path)
			}(filePathBg, filePath, mp3Path)
		}
	}

	wg.Wait()
}

func (ds *DownloadService) downloadFileContent(filePathBg, filePath, mp3Path string) {
	bookIdInt := ds.findFileBookId(filePathBg)
	path := "/Users/jiang/demo/book/" + bookIdInt

	if err := ds.downloadAndSaveFile(filePathBg, path); err != nil {
		fmt.Println("download err:", err)
	} else {
		fmt.Println("download success:", filePathBg)
	}

	if err := ds.downloadAndSaveFile(filePath, path); err != nil {
		fmt.Println("download err:", err)
	} else {
		fmt.Println("download success:", filePath)
	}

	if exists := utils.UrlExists(mp3Path); exists {
		if err := ds.downloadAndSaveFile(mp3Path, path); err != nil {
			fmt.Println("download err:", err)
		} else {
			fmt.Println("download success:", mp3Path)
		}
	}
}

func (ds *DownloadService) downloadAndSaveFile(url, path string) error {
	fileName := ds.findFileName(url)
	filePathSave := path + "/" + fileName
	return utils.DownloadFile(url, path, filePathSave)
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

func (ds *DownloadService) fileUrlSaveToLog(startIndex, endIndex int) {
	fileLog := "/Users/jiang/demo/file.log"
	var wg sync.WaitGroup
	for bookId := startIndex; bookId < endIndex; bookId++ {
		wg.Add(1)
		go func(bookId int) {
			defer wg.Done()
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
		}(bookId)
	}

	wg.Wait()
}
