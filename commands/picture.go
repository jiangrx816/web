package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"web/service/picture_service"
)

func Picture() *cli.Command {
	return &cli.Command{
		Name:  "picture",
		Usage: "./main picture",
		Action: func(c *cli.Context) error {
			var pictureService picture_service.PictureService
			fmt.Println(1111)
			pictureService.InsertChinesePictureCover()

			return nil
		},
	}
}
