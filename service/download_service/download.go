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
	"web/utils"
)

var counter int = 10
var baseStep int = 0
var step int = 1000
var imgBaseUrl string = ""
var dirPathRoot string = "/Users/jiang/demo/book"

const maxConcurrent = 10

// Result 结果结构体，用于保存文件夹和文件的内容
type Result struct {
	Path string
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
				ds.execFfmpegCommand(fullPath, newPath)
			}
			results <- Result{Path: fullPath}
		}
	}
}

func (ds *DownloadService) ExecSh() {
	items := []string{"1"}
	for idx, _ := range items {

		rootDir := "/Users/jiang/demo/shell/items/" + items[idx]
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
				fileCount := countFilesInDir(dirPath, id)
				fmt.Printf("worker %d: directory %s, file count: %d\n", id, dirPath, fileCount/2)

				for ii := 1; ii <= fileCount/2; ii++ {
					if number, _ := extractNumber(dirPath); number > 0 {
						fileP := "/Users/jiang/demo/shell/" + strconv.Itoa(number) + ".sh"
						utils.CreateFile(fileP)
						ffm := `ffmpeg  -thread_queue_size 96   -loop 1   -t  2  -y -r 1 -i  ` + dirPath + `/` + strconv.Itoa(ii) + `.jpg   -i ` + dirPath + `/page_audio_` + strconv.Itoa(ii) + `.mp3  -x264-params keyint=1:scenecut=0  -vf "scale=2800:-2"   -absf aac_adtstoasc -s 1280x720 -c:v libx264 -pix_fmt yuv420p   ` + dirPath + `/` + strconv.Itoa(ii) + `.mp4  2>&1`
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
			if match, _ := regexp.MatchString(`/\d+\.jpg$`, path); match {
				matches := regexp.MustCompile(`/(\d+)/\d+\.jpg$`).FindStringSubmatch(path)
				if len(matches) > 1 {
					//matchList := regexp.MustCompile(`/\d+\.jpg$`).FindStringSubmatch(path)
					newPath := "/Users/jiang/demo/book1/86" + matches[1]
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
					newPath := "/Users/jiang/demo/book1/86" + matches[1]
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
