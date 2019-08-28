package main

import (
	"router"
)

func main() {
	r := router.NewRouter()
	//gin.SetMode(gin.ReleaseMode)
	err := r.Run(":9000")
	if err != nil {
		return
	}
}

