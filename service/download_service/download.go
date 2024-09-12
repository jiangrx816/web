package download_service

import (
	"fmt"
	rxLog "github.com/jiangrx816/gopkg/log"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"web/model"
	"web/model/picture"
	"web/utils"
)

var counter int = 10
var baseStep int = 0
var step int = 1000
var imgBaseUrl string = ""
var dirPathRoot string = "/Users/jiang/demo/book1"

const maxConcurrent = 10

// Result 结果结构体，用于保存文件夹和文件的内容
type Result struct {
	Path string
}

func (ds *DownloadService) DownloadImgListByIds() {
	bookIds := []int{1004, 1005, 1006, 1007, 1010, 1013, 1014, 1015, 1016, 1017, 1530, 1535, 10081, 10221, 10229, 102, 10246, 10258, 10263, 10265, 10400, 10402, 10407, 10514, 10595, 10641, 10643, 10677, 10717, 10719, 10720, 10721, 10722, 10727, 10728, 10729, 10734, 10739, 10743, 10746, 10748, 10749, 10766, 10768, 10769, 10770, 10771, 11001, 11007, 11011, 11014, 11015, 11107, 11403, 11489, 11657, 12033, 13065, 20091, 20092, 20093, 20098, 20125, 20909, 20910, 20911, 20912, 20913, 20914, 20937, 20938, 20939, 20940, 20982, 20998, 21133, 21134, 21135, 21195, 21196, 21198, 21199, 10290, 11642, 13206, 20999, 20097, 20287, 20297, 20908, 20936, 21200, 10076, 10080, 10082, 10222, 10223, 10224, 10281, 10282, 10283, 10285, 10286, 10287, 10288, 10289, 10291, 10292, 10293, 10329, 10331, 10333, 10336, 10339, 10343, 10346, 10349, 10352, 10354, 10355, 10357, 10358, 10359, 10360, 10362, 10364, 10365, 10366, 10367, 10368, 10369, 10370, 10371, 10372, 10374, 10375, 10376, 10377, 10378, 10379, 10435, 10441, 10442, 10448, 10454, 10456, 10486, 10508, 10509, 10510, 10560, 10607, 10608, 10611, 10615, 10617, 10620, 10622, 10624, 10627, 10633, 10635, 10639, 10720, 10763, 10764, 10765, 11354, 11422, 11423, 11424, 11425, 11426, 11427, 11428, 11429, 11430, 11431, 11432, 11433, 11434, 11435, 11436, 11437, 11438, 11439, 11440, 11441, 11442, 11443, 11444, 11445, 11446, 11447, 11448, 11449, 11450, 11451, 11452, 11453, 11454, 11455, 11456, 11457, 11458, 11459, 11460, 11461, 11462, 11463, 11465, 20093, 10356, 10433, 10565, 11746, 20915, 13464, 20100}

	const workerCount = 20
	taskChan := make(chan int, len(bookIds))
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range taskChan {
				bookIdStr := strconv.Itoa(bookIds[idx])
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

					ds.downloadFileContent(filePathBg, filePath, mp3Path)
				}
			}
		}()
	}

	// Send tasks to the channel
	for idx := range bookIds {
		taskChan <- idx
	}

	// Close the channel to signal workers to stop
	close(taskChan)

	// Wait for all workers to finish
	wg.Wait()
}

func (ds *DownloadService) DownloadImgListByIds1() {
	bookIds := []int{1004, 1005, 1006, 1007, 1010, 1013, 1014, 1015, 1016, 1017, 1530, 1535, 10081, 10221, 10229, 102, 10246, 10258, 10263, 10265, 10400, 10402, 10407, 10514, 10595, 10641, 10643, 10677, 10717, 10719, 10720, 10721, 10722, 10727, 10728, 10729, 10734, 10739, 10743, 10746, 10748, 10749, 10766, 10768, 10769, 10770, 10771, 11001, 11007, 11011, 11014, 11015, 11107, 11403, 11489, 11657, 12033, 13065, 20091, 20092, 20093, 20098, 20125, 20909, 20910, 20911, 20912, 20913, 20914, 20937, 20938, 20939, 20940, 20982, 20998, 21133, 21134, 21135, 21195, 21196, 21198, 21199, 10290, 11642, 13206, 20999, 20097, 20287, 20297, 20908, 20936, 21200, 10076, 10080, 10082, 10222, 10223, 10224, 10281, 10282, 10283, 10285, 10286, 10287, 10288, 10289, 10291, 10292, 10293, 10329, 10331, 10333, 10336, 10339, 10343, 10346, 10349, 10352, 10354, 10355, 10357, 10358, 10359, 10360, 10362, 10364, 10365, 10366, 10367, 10368, 10369, 10370, 10371, 10372, 10374, 10375, 10376, 10377, 10378, 10379, 10435, 10441, 10442, 10448, 10454, 10456, 10486, 10508, 10509, 10510, 10560, 10607, 10608, 10611, 10615, 10617, 10620, 10622, 10624, 10627, 10633, 10635, 10639, 10720, 10763, 10764, 10765, 11354, 11422, 11423, 11424, 11425, 11426, 11427, 11428, 11429, 11430, 11431, 11432, 11433, 11434, 11435, 11436, 11437, 11438, 11439, 11440, 11441, 11442, 11443, 11444, 11445, 11446, 11447, 11448, 11449, 11450, 11451, 11452, 11453, 11454, 11455, 11456, 11457, 11458, 11459, 11460, 11461, 11462, 11463, 11465, 20093, 10356, 10433, 10565, 11746, 20915, 13464, 20100}
	for idx, _ := range bookIds {
		bookIdStr := strconv.Itoa(bookIds[idx])
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

			ds.downloadFileContent(filePathBg, filePath, mp3Path)
		}
	}
}

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

// FFMpegImageMergeBGToSub 将背景图与内容图进行合并成一张图片
func (ds *DownloadService) FFMpegImageMergeBGToSub() {
	sem := make(chan struct{}, maxConcurrent)
	results := make(chan Result)
	var wg sync.WaitGroup

	wg.Add(1)
	go ds.processDir(dirPathRoot, &wg, sem, results)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("File: %s\n", result.Path)
	}
}

// 遍历目录并读取文件内容
func (ds *DownloadService) processDir(dirPath string, wg *sync.WaitGroup, sem chan struct{}, results chan<- Result) {
	defer wg.Done()

	sem <- struct{}{}        // 获取一个许可
	defer func() { <-sem }() // 释放许可

	entries, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())
		if entry.IsDir() {
			wg.Add(1)
			go ds.processDir(fullPath, wg, sem, results) // 递归遍历子目录
		} else {
			if contains := ds.isPageBg(fullPath); contains {
				newPath := ds.makePageSub(fullPath)
				os.Remove(fullPath)
				os.Remove(newPath)
				ds.execFfmpegCommand(fullPath, newPath)
			}
			results <- Result{Path: fullPath}
		}
	}
}
func (ds *DownloadService) ExecSh() {
	// 指定目录
	dir := "/Users/jiang/demo/shell/items/"

	// 读取目录中的文件
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("读取目录失败:", err)
		return
	}

	var wg sync.WaitGroup
	fileChan := make(chan string, len(files))

	// 将所有 .sh 文件添加到通道
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sh" {
			fileChan <- filepath.Join(dir, file.Name())
		}
	}
	close(fileChan)

	// 开启10个goroutine并发执行
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range fileChan {
				fmt.Printf("正在执行 %s\n", file)
				cmd := exec.Command("sh", file)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Printf("执行 %s 失败: %v\n", file, err)
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println("所有脚本执行完毕")
}
func (ds *DownloadService) ExecSh1() {
	items := []int{241004,
		241005,
		241006,
		241007,
		241010,
		241013,
		241014,
		241015,
		241016,
		241017,
		241530,
		241535,
		2410081,
		2410221,
		2410229,
		2410242,
		2410246,
		2410258,
		2410263,
		2410265,
		2410400,
		2410402,
		2410407,
		2410514,
		2410595,
		2410641,
		2410643,
		2410677,
		2410717,
		2410719,
		2410720,
		2410721,
		2410722,
		2410727,
		2410728,
		2410729,
		2410734,
		2410739,
		2410743,
		2410746,
		2410748,
		2410749,
		2410766,
		2410768,
		2410769,
		2410770,
		2410771,
		2411001,
		2411007,
		2411011,
		2411014,
		2411015,
		2411107,
		2411403,
		2411489,
		2411657,
		2412033,
		2413065,
		2420091,
		2420092,
		2420093,
		2420098,
		2420125,
		2420909,
		2420910,
		2420911,
		2420912,
		2420913,
		2420914,
		2420937,
		2420938,
		2420939,
		2420940,
		2420982,
		2420998,
		2421133,
		2421134,
		2421135,
		2421195,
		2421196,
		2421198,
		2421199,
		2410290,
		2411642,
		2413206,
		2420999,
		2420097,
		2420287,
		2420297,
		2420908,
		2420936,
		2421200,
		2410076,
		2410080,
		2410082,
		2410222,
		2410223,
		2410224,
		2410281,
		2410282,
		2410283,
		2410285,
		2410286,
		2410287,
		2410288,
		2410289,
		2410291,
		2410292,
		2410293,
		2410329,
		2410331,
		2410333,
		2410336,
		2410339,
		2410343,
		2410346,
		2410349,
		2410352,
		2410354,
		2410355,
		2410357,
		2410358,
		2410359,
		2410360,
		2410362,
		2410364,
		2410365,
		2410366,
		2410367,
		2410368,
		2410369,
		2410370,
		2410371,
		2410372,
		2410374,
		2410375,
		2410376,
		2410377,
		2410378,
		2410379,
		2410435,
		2410441,
		2410442,
		2410448,
		2410454,
		2410456,
		2410486,
		2410508,
		2410509,
		2410510,
		2410560,
		2410607,
		2410608,
		2410611,
		2410615,
		2410617,
		2410620,
		2410622,
		2410624,
		2410627,
		2410633,
		2410635,
		2410639,
		2410720,
		2410763,
		2410764,
		2410765,
		2411354,
		2411422,
		2411423,
		2411424,
		2411425,
		2411426,
		2411427,
		2411428,
		2411429,
		2411430,
		2411431,
		2411432,
		2411433,
		2411434,
		2411435,
		2411436,
		2411437,
		2411438,
		2411439,
		2411440,
		2411441,
		2411442,
		2411443,
		2411444,
		2411445,
		2411446,
		2411447,
		2411448,
		2411449,
		2411450,
		2411451,
		2411452,
		2411453,
		2411454,
		2411455,
		2411456,
		2411457,
		2411458,
		2411459,
		2411460,
		2411461,
		2411462,
		2411463,
		2411465,
		2420093,
		2410356,
		2410433,
		2410565,
		2411746,
		2420915,
		2413464,
		2420100}
	for idx, _ := range items {
		rootDir := "/Users/jiang/demo/shell/items/" + strconv.Itoa(items[idx])
		numWorkers := 10

		// 创建一个 channel 用于传递 shell 脚本路径
		scriptChan := make(chan string, numWorkers)
		var wg sync.WaitGroup

		// 启动工作 goroutines
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for scriptPath := range scriptChan {
					fmt.Printf("worker %d: 执行脚本 %s\n", id, scriptPath)
					if err := executeShellScript(scriptPath); err != nil {
						fmt.Printf("worker %d: 执行脚本 %s 时出错: %v\n", id, scriptPath, err)
					}
				}
			}(i)
		}

		// 遍历根目录并发送脚本路径到 channel
		go func() {
			defer close(scriptChan)
			err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("访问文件 %s 时出错: %v\n", path, err)
					return err
				}
				if !info.IsDir() && filepath.Ext(path) == ".sh" {
					scriptChan <- path
				}
				return nil
			})
			if err != nil {
				fmt.Printf("遍历路径 %s 时出错: %v\n", rootDir, err)
			}
		}()

		// 等待所有 goroutines 完成
		wg.Wait()
		fmt.Println("所有脚本执行完成。")

		deleteFolder(rootDir)
	}
}

func deleteFolder(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to delete folder %s: %w", path, err)
	}
	return nil
}

// executeShellScript 执行 shell 脚本文件
func executeShellScript(filePath string) error {
	cmd := exec.Command("sh", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (ds *DownloadService) FFMpegImageToVideo() {
	//ffmpeg  -thread_queue_size 96   -loop 1   -t  2  -y -r 1 -i  /Users/jiang/demo/book1/8610011/2.jpg   -i   /Users/jiang/demo/book1/8610011/page_audio_2.mp3  -x264-params keyint=1:scenecut=0  -vf "scale=2800:-2"   -absf aac_adtstoasc -s 1280x720 -c:v libx264 -pix_fmt yuv420p   /Users/jiang/demo/mp4/8610011-4.mp4  2>&1
	rootDir := "/Users/jiang/demo/book"
	numWorkers := 10

	// 创建一个 channel 用于传递目录路径
	dirChan := make(chan string, numWorkers)
	var wg sync.WaitGroup

	// 启动工作 goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for dirPath := range dirChan {
				fileCount := countFilesInDir(dirPath, id)
				fmt.Printf("worker %d: directory %s, file count: %d\n", id, dirPath, fileCount/2)

				for ii := 1; ii <= fileCount/4; ii++ {
					if number, _ := extractNumber(dirPath); number > 0 {
						fileP := "/Users/jiang/demo/shell/" + strconv.Itoa(number) + ".sh"
						utils.CreateFile(fileP)
						ffm := `ffmpeg  -thread_queue_size 96   -loop 1   -t  2  -y -r 1 -i  ` + dirPath + `/` + strconv.Itoa(ii) + `.jpg   -i ` + dirPath + `/` + strconv.Itoa(ii) + `.mp3  -x264-params keyint=1:scenecut=0  -vf "scale=2800:-2"   -absf aac_adtstoasc -s 1280x720 -c:v libx264 -pix_fmt yuv420p   ` + dirPath + `/` + strconv.Itoa(ii) + `.mp4  2>&1`
						utils.AppendToFile(fileP, ffm)
					}
				}
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

// countFilesInDir 计算目录中的文件数量
func countFilesInDir(dirPath string, workerID int) int {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("worker %d: 读取目录 %s 时出错: %v\n", workerID, dirPath, err)
		return 0
	}
	return len(files)
}

// extractNumber 从给定的字符串中提取数字
func extractNumber(input string) (int, error) {
	re := regexp.MustCompile(`(\d+)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return strconv.Atoi(matches[1])
	}
	return 0, fmt.Errorf("未找到数字")
}

func (ds *DownloadService) execFfmpegCommand(pageBg, pageSub string) {
	outFile := ds.ffmpegCommandOutFile(pageBg)
	// 构造FFmpeg命令
	cmd := exec.Command("ffmpeg", "-i", pageBg, "-i", pageSub, "-filter_complex", "[0][1]overlay=0:0", outFile)
	// 运行命令
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing FFmpeg command:", err)
		return
	}
	// 输出命令执行结果
	fmt.Println(string(output))
}

func (ds *DownloadService) ffmpegCommandOutFile(originalPath string) (outFile string) {

	// 提取目录和文件名
	dir := filepath.Dir(originalPath)
	filename := filepath.Base(originalPath)

	// 分割文件名和扩展名
	nameParts := strings.Split(filename, "_")
	if len(nameParts) < 2 {
		fmt.Println("Unexpected file name format")
		return
	}

	// 获取最后一个部分（即“4.jpg”）
	newFilename := nameParts[len(nameParts)-1]

	// 构造新的路径
	newPath := filepath.Join(dir, newFilename)

	// 输出新的路径
	fmt.Println(newPath)
	outFile = newPath
	return
}

func (ds *DownloadService) isPageBg(imgPath string) (contains bool) {
	// 要查找的子字符串
	substr := "page_bg_"
	// 判断字符串中是否包含子字符串
	if strings.Contains(imgPath, substr) {
		contains = true
		return
	}
	return
}

func (ds *DownloadService) makePageSub(imgPath string) (newPath string) {
	// 定义正则表达式，匹配 "page_bg_X.jpg"
	re := regexp.MustCompile(`page_bg_(\d+)\.jpg`)

	newPath = re.ReplaceAllString(imgPath, "page_sub_${1}.png")
	return
}

func (ds *DownloadService) GetFileCount() {
	countFilesInSubdirectories(dirPathRoot)
}

func countFilesInSubdirectories(root string) {
	// 打开根目录
	dirEntries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// 遍历根目录中的条目
	for _, entry := range dirEntries {
		if entry.IsDir() {
			// 获取子目录路径
			subdir := filepath.Join(root, entry.Name())

			// 统计子目录中的文件数量
			fileCount, err := countFiles(subdir)
			if err != nil {
				fmt.Printf("Error counting files in directory %s: %v\n", subdir, err)
				continue
			}

			// 输出子目录文件数量
			if fileCount < 15 {
				//fmt.Printf("Directory: %s, File count: %d\n", subdir, fileCount)
				// 定义正则表达式，用于匹配数字
				re := regexp.MustCompile(`\d+`)
				// 提取字符串中的数字
				numbers := re.FindAllString(subdir, -1)
				if len(numbers) > 0 {
					//fmt.Println(numbers[0])
					utils.AppendToFile("/Users/jiang/demo/book1.log", numbers[0]+",")
				}
			}
		}
	}
}

func countFiles(dir string) (int, error) {
	// 初始化文件计数器
	fileCount := 0

	// 遍历目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是文件，计数器加1
		if !info.IsDir() {
			fileCount++
		}

		return nil
	})

	return fileCount, err
}

func (ds *DownloadService) MoveFileToPath() {
	var wg sync.WaitGroup

	wg.Add(1)
	go walkDir(dirPathRoot, copyFile, &wg)

	wg.Wait()
	fmt.Println("All files have been moved.")
}

func walkDir(dir string, copyFile func(string, string) error, wg *sync.WaitGroup) {
	defer wg.Done()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by returning filepath.WalkErr: %v\n", err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".jpg" {
			//if match, _ := regexp.MatchString(`/\d+\.jpg$`, path); match {
			if match, _ := regexp.MatchString(`page_bg_\d+\.jpg$`, path); match {
				matches := regexp.MustCompile(`/(\d+)/page_bg_\d+\.jpg$`).FindStringSubmatch(path)
				if len(matches) > 1 {
					//matchList := regexp.MustCompile(`/\d+\.jpg$`).FindStringSubmatch(path)
					newPath := "/Users/jiang/demo/book1/24" + matches[1]
					destinationFile := filepath.Join(newPath, filepath.Base(path))
					// 创建目标目录（如果不存在）
					err := os.MkdirAll(newPath, os.ModePerm)
					if err != nil {
						fmt.Println("创建目录失败:", err)
					}

					if err := copyFile(path, destinationFile); err != nil {
						fmt.Printf("failed to move file: %v\n", err)
					}
				}
			}
		}

		if !info.IsDir() && filepath.Ext(path) == ".png" {
			//if match, _ := regexp.MatchString(`/\d+\.jpg$`, path); match {
			if match, _ := regexp.MatchString(`page_sub_\d+\.png$`, path); match {
				matches := regexp.MustCompile(`/(\d+)/page_sub_\d+\.png$`).FindStringSubmatch(path)
				if len(matches) > 1 {
					//matchList := regexp.MustCompile(`/\d+\.jpg$`).FindStringSubmatch(path)
					newPath := "/Users/jiang/demo/book1/24" + matches[1]
					destinationFile := filepath.Join(newPath, filepath.Base(path))
					// 创建目标目录（如果不存在）
					err := os.MkdirAll(newPath, os.ModePerm)
					if err != nil {
						fmt.Println("创建目录失败:", err)
					}

					if err := copyFile(path, destinationFile); err != nil {
						fmt.Printf("failed to move file: %v\n", err)
					}
				}
			}
		}

		if !info.IsDir() && filepath.Ext(path) == ".mp3" {
			if match, _ := regexp.MatchString(`page_audio_\d+\.mp3$`, path); match {
				matches := regexp.MustCompile(`/(\d+)/page_audio_\d+\.mp3$`).FindStringSubmatch(path)
				if len(matches) > 1 {
					//matchList := regexp.MustCompile(`/page_audio_\d+\.mp3$`).FindStringSubmatch(path)
					newPath := "/Users/jiang/demo/book1/24" + matches[1]
					destinationFile := filepath.Join(newPath, filepath.Base(path))
					// 创建目标目录（如果不存在）
					err := os.MkdirAll(newPath, os.ModePerm)
					if err != nil {
						fmt.Println("创建目录失败:", err)
					}

					if err := copyFile(path, destinationFile); err != nil {
						fmt.Printf("failed to move file: %v\n", err)
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
	}
}

// copyFile 复制文件从 src 到 dst
func copyFile(src, dst string) error {

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// 确保文件复制成功
	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (ds *DownloadService) WriteFileVideo() {
	rootDir := "/Users/jiang/demo/book"
	numWorkers := 10

	// 创建一个 channel 用于传递目录路径
	dirChan := make(chan string, numWorkers)
	var wg sync.WaitGroup

	// 启动工作 goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for dirPath := range dirChan {
				fileCount := countFilesInDir(dirPath, id)
				//fmt.Printf("worker %d: directory %s, file count: %d\n", id, dirPath, fileCount/3)
				for ii := 1; ii <= fileCount/3; ii++ {
					if number, _ := extractNumber(dirPath); number > 0 {
						fileP := "/Users/jiang/demo/shell/items/" + strconv.Itoa(number) + ".txt"
						utils.CreateFile(fileP)
						ffm := `file '` + dirPath + `/` + strconv.Itoa(ii) + `.mp4'`
						utils.AppendToFile(fileP, ffm)
					}
				}
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

func (ds *DownloadService) MakeVideoShell() {
	var wg sync.WaitGroup

	wg.Add(1)
	go walkDirVideoShell("/Users/jiang/demo/shell/items", &wg)

	wg.Wait()
	fmt.Println("All files have been finished.")
}

func walkDirVideoShell(dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by returning filepath.WalkErr: %v\n", err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			// 定义正则表达式，匹配数字
			re := regexp.MustCompile(`\d+`)
			// 查找字符串中的第一个数字
			number := re.FindString(path)
			fileP := "/Users/jiang/demo/shell/shell.sh"
			ffm := `ffmpeg -f concat -safe 0 -i ` + path + ` -c copy /Users/jiang/demo/mp4/` + number + `.mp4`
			utils.AppendToFile(fileP, ffm)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
	}
}

func (ds *DownloadService) InsertData() {
	rootDir := "/Users/jiang/demo/book1"
	numWorkers := 10

	// 创建一个 channel 用于传递目录路径
	dirChan := make(chan string, numWorkers)
	var wg sync.WaitGroup

	// 启动工作 goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for dirPath := range dirChan {
				number, _ := extractNumber(dirPath)
				bookId := utils.MD5String(strconv.Itoa(number))
				var temp picture.ChineseBookTemp
				temp.BookId = bookId
				temp.BookIdOld = number
				temp.Type = -1
				model.DefaultPicture().Model(&picture.ChineseBookTemp{}).Debug().Create(&temp)
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

func (ds *DownloadService) DeleteMp4Data() {
	rootDir := "/Users/jiang/demo/cover"
	numWorkers := 10

	// 创建一个 channel 用于传递目录路径
	dirChan := make(chan string, numWorkers)
	var wg sync.WaitGroup

	// 启动工作 goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for dirPath := range dirChan {
				filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						fmt.Printf("prevent panic by returning filepath.WalkErr: %v\n", err)
						return err
					}
					if !info.IsDir() && filepath.Ext(path) == ".mp4" {
						fmt.Println("Deleting:", path)
						err = os.Remove(path)
						if err != nil {
							return err
						}
					}
					return nil
				})
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
