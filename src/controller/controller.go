package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ohmygd/mgo/merror"
	"log"
	"net/http"
	"pconst"
)

func Render(c *gin.Context, err error, data interface{}) {
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":pconst.ErrorOk,
			"msg":"success",
			"data":data,
		})

		return
	}

	var merr *merror.Merr
	var ok bool

	if merr, ok = err.(*merror.Merr); !ok {
		log.Println(err.Error())
		merr = merror.New(pconst.ErrorSystem)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":merr.GetCode(),
		"msg":merr.GetMsg(),
		"data":data,
	})

	return
}