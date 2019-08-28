package controller

import "time"

// binding 提供validate功能
type Test3 struct {
	Name string `form:"name" binding:"required"`
	Age int `form:"age"`
}

type Test7 struct {
	Name string `form:"name" binding:"required,contains=mc"`
	Age int `form:"age" binding:"required,min=1,max=10"`
	Born time.Time `form:"born" time_format:"2006-01-02"`
}