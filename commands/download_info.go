package commands

import (
	"github.com/urfave/cli/v2"
	"web/service/download_service"
)

func DownloadInfo() *cli.Command {
	return &cli.Command{
		Name:  "download-info",
		Usage: "./main download-info",
		Action: func(c *cli.Context) error {
			var downloadService download_service.DownloadService
			downloadService.DownloadImgInfo()
			return nil
		},
	}
}
