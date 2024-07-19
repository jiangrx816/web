package cron

import (
	rxCron "web/cron/base"
)

func DoCron() error {
	cronList := make([]rxCron.Cron, 0)
	cronList = append(cronList, NewUpdateStudent())
	if err := rxCron.InitFromSecond(cronList); err != nil {
		return err
	}
	return nil
}
