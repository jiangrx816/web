package cron

import (
	"github.com/jiangrx816/gopkg/log"
	rxCron "web/cron/base"
)

type UpdateStudent struct {
}

func NewUpdateStudent() rxCron.Cron {
	return &UpdateStudent{}
}

func (us *UpdateStudent) Spec() string {
	return "* * * * * *"
}

func (us *UpdateStudent) Run() {
	log.SugarContext(rxCron.BuildCtx()).Infow("1")
}
