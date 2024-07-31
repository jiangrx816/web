package cron

import (
	"github.com/spf13/viper"
	rxCron "web/cron/base"
)

func DoCron() error {
	if !viper.GetBool("cron.switch") {
		return nil
	}
	cronSecondList := make([]rxCron.Cron, 0)
	cronSecondList = append(cronSecondList, NewTestSecond())
	if err := rxCron.InitFromSecond(cronSecondList); err != nil {
		return err
	}

	cronMinuteList := make([]rxCron.Cron, 0)
	cronMinuteList = append(cronMinuteList, NewTestMinute())
	if err := rxCron.InitFromMinute(cronMinuteList); err != nil {
		return err
	}
	return nil
}
