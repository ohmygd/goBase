package api

import (
	"fmt"
	"github.com/robfig/cron"
	"testing"
	"time"
)

func TestM1(t *testing.T) {
	c := cron.New()
	err := c.AddFunc("*/1 * * * *", func(){
		fmt.Println("crontab test")
	})

	if err != nil {
		t.Log(err.Error())
	}
	c.Start()
	fmt.Println("=========")

	time.Sleep(time.Second * 10)
}
