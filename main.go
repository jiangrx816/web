package main

import (
	rxMysql "github.com/jiangrx816/gopkg/db"
	rxLog "github.com/jiangrx816/gopkg/log"
	rxRedis "github.com/jiangrx816/gopkg/redis"
	"github.com/jiangrx816/gopkg/utils"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"os"
	"web/cache"
	"web/commands"
	"web/cron"
)

var configFile string

func main() {
	app := cli.NewApp()
	app.Action = commands.Serve
	app.Before = initConfig
	app.Commands = commands.Commands
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Value:       "", // 默认从config目录读取
			Usage:       "specify the location of the configuration file",
			Required:    false,
			Destination: &configFile,
		},
	}
	if err := app.Run(os.Args); err != nil {
		rxLog.Sugar().Fatal(err)
	}
}

func initConfig(*cli.Context) error {
	viper.SetDefault("app", "jiang")
	if err := utils.LoadConfigInFile(configFile); err != nil {
		return err
	}
	if err := rxLog.InitFromViper(); err != nil {
		return err
	}
	if err := rxMysql.InitMysqlDB(); err != nil {
		return err
	}
	if err := rxRedis.InitFromViperDefault(); err != nil {
		return err
	}
	if err := cache.HttpCache(); err != nil {
		return err
	}
	if err := cron.DoCron(); err != nil {
		return err
	}
	return nil
}
