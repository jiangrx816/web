package commands

import (
	"github.com/urfave/cli/v2"
	"web/service/download_service"
)

func DownloadCover() *cli.Command {
	return &cli.Command{
		Name:  "download-cover",
		Usage: "./main download-cover",
		Action: func(c *cli.Context) error {
			var downloadService download_service.DownloadService
			downloadService.DownloadImgCover()
			return nil
		},
	}
}
