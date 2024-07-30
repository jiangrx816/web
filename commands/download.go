package commands

import (
	"github.com/urfave/cli/v2"
	"web/service/download_service"
)

func Download() *cli.Command {
	return &cli.Command{
		Name:  "download",
		Usage: "./main download",
		Action: func(c *cli.Context) error {
			var downloadService download_service.DownloadService
			//将背景图与内容图进行合并成一张图片
			//downloadService.MoveFileToPath()

			//将图片与MP3合成为视频
			//downloadService.FFMpegImageToVideo()

			//执行顺序的sh
			//downloadService.ExecSh()

			//将需要合并的视频写入文件
			//downloadService.WriteFileVideo()

			//将多个视频合并成一个完整的视频
			//downloadService.MakeVideoShell()

			//删除对应的mp4文件
			downloadService.DownloadImageList()

			return nil
		},
	}
}
