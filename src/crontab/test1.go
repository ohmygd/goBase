package crontab

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func Cronttest() {
	cc := cron.New()
	err := cc.AddFunc("*/1 * * * *", func(){
		fmt.Println("crontab test")
	})
	if err != nil {
		log.Println(err.Error(), "==========")
	}

	cc.Start()
}
