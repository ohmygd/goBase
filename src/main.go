package main

import (
	"router"
	"time"
)

func main() {
	r := router.NewRouter()

	//go crontab.Cronttest()
	//gin.SetMode(gin.ReleaseMode)
	err := r.Run(":9000")
	if err != nil {
		return
	}
	time.Sleep(time.Minute)
}

