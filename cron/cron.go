package cron

import (
	rxCron "web/cron/base"
)

func DoCron() error {
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
