package base

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"web/utils"
)

type Cron interface {
	Spec() string
	Run()
}

// InitFromSecond 秒级
func InitFromSecond(cronList []Cron) error {
	if !viper.GetBool("cron.switch") {
		return nil
	}
	if err := initFromViper(cron.New(cron.WithSeconds()), cronList); err != nil {
		return err
	}
	return nil
}

// InitFromMinute 分级
func InitFromMinute(cronList []Cron) error {
	if !viper.GetBool("cron.switch") {
		return nil
	}
	if err := initFromViper(cron.New(), cronList); err != nil {
		return err
	}
	return nil
}

func initFromViper(c *cron.Cron, cronList []Cron) error {
	for _, task := range cronList {
		if _, err := c.AddFunc(task.Spec(), task.Run); err != nil {
			return err
		}
	}
	c.Start()
	return nil
}

func BuildCtx() context.Context {
	return context.WithValue(context.Background(), "x-request-id", utils.GenUUID())
}
