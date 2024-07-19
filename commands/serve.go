package commands

import (
	"github.com/jiangrx816/gopkg/graceful"
	"github.com/jiangrx816/gopkg/server"
	"github.com/urfave/cli/v2"
	"web/http/router"
)

func Serve(c *cli.Context) error {
	// 运行HTTP服务
	graceful.Start(server.NewHttp(server.Addr(":8082"), server.Router(router.All())))

	graceful.Wait()
	return nil
}
