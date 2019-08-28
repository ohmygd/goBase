package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MiddleTest(c *gin.Context) {
	fmt.Println("test in")
	c.Next()
	fmt.Println("test over")
}
