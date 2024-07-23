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
			//downloadService.GetFileCount()
			//downloadService.DownloadImgList()

			downloadService.ReadDirList()
			return nil
		},
	}
}
