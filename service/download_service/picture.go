package download_service

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"web/model"
	"web/model/picture"
	"web/utils"
)

func (ds *DownloadService) InsertChinesePictureInfo() {
	rootDir := "/Users/jiang/demo/book"
	numWorkers := 1

	// 创建一个 channel 用于传递目录路径
	dirChan := make(chan string, numWorkers)
	var wg sync.WaitGroup

	// 启动工作 goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			var temp picture.ChineseBookInfo
			var tempList []picture.ChineseBookInfo
			for dirPath := range dirChan {
				fileCount := countFilesInDir(dirPath, id)
				fmt.Printf("worker %d: directory %s, file count: %d\n", id, dirPath, fileCount/2)

				for ii := 1; ii <= fileCount/2; ii++ {
					if number, _ := extractNumber(dirPath); number > 0 {
						temp.BookId = utils.MD5String(strconv.Itoa(number))
						temp.BookIdOld = number
						temp.Pic = "https://oss.58haha.com/chinese_book/file_rx/" + strconv.Itoa(number) + "/" + strconv.Itoa(ii) + ".jpg"
						temp.Mp3 = "https://oss.58haha.com/chinese_book/file_rx/" + strconv.Itoa(number) + "/page_audio_" + strconv.Itoa(ii) + ".mp3"
						temp.Position = ii
						tempList = append(tempList, temp)
					}
				}
				model.DefaultPicture().Model(&picture.ChineseBookInfo{}).Debug().Create(&tempList)
			}
		}(i)
	}

	// 遍历根目录并发送子目录路径到 channel
	go func() {
		defer close(dirChan)
		err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("访问文件 %s 时出错: %v\n", path, err)
				return err
			}
			if info.IsDir() && path != rootDir {
				dirChan <- path
			}
			return nil
		})
		if err != nil {
			fmt.Printf("遍历路径 %s 时出错: %v\n", rootDir, err)
		}
	}()

	// 等待所有 goroutines 完成
	wg.Wait()
	fmt.Println("所有目录处理完成。")
}
