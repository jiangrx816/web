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
			downloadService.FFMpegImageMergeBGToSub()

			//将图片与MP3合成为视频
			//downloadService.FFMpegImageToVideo()

			return nil
		},
	}
}
