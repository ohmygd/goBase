package router

import (
	"controller/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	//v1.Use(middleware.MiddleTest)
	{
		v1.GET("/test", api.Test)
		v1.GET("/test1", api.Test1)
		v1.GET("/test2", api.Test2)
		v1.GET("/test3", api.Test3)
		v1.GET("/test4", api.Test4)
		v1.GET("/test5", api.Test5)
		v1.GET("/test6", api.Test6)
		v1.GET("/test7", api.Test7)
		v1.GET("/test8", api.Test8)
		v1.GET("/test9", api.Test9)
		v1.GET("/test10", api.Test10)
		v1.GET("/test11", api.Test11)
		v1.GET("/test12", api.Test12)
		v1.GET("/test13", api.Test13)
		v1.GET("/test14", api.Test14)
		v1.GET("/test15", api.Test15)
		v1.GET("/test16", api.Test16)
		v1.GET("/test17", api.Test17)
		v1.GET("/test18", api.Test18)
		v1.GET("/test19", api.Test19)
		v1.GET("/test20", api.Test20)
		v1.GET("/test21", api.Test21)
		v1.GET("/test22", api.Test22)
		v1.GET("/test23", api.Test23)
		v1.GET("/test24", api.Test24)
		v1.GET("/test25", api.Test25)
		v1.GET("/test26", api.Test26)
		v1.GET("/test27", api.Test27)
		v1.GET("/test28", api.Test28)
		v1.GET("/test29", api.Test29)
	}

	return router
}