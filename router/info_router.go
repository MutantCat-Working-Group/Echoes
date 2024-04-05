package router

import (
	"com.mutantcat.echoes/status"
	"github.com/gin-gonic/gin"
)

type InfoRouter struct {
	ServerName string
}

func (r *InfoRouter) PrepareRouter() error {
	return nil
}

func (r *InfoRouter) InitRouter(c *gin.Engine) error {
	c.Any("/ping", ping)
	c.Any("/info", getAllInfo)
	return nil
}

func (r *InfoRouter) DestroyRouter() error {
	return nil
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "pong",
	})
}

func getAllInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		// data里面有两项
		"data": gin.H{
			"system": status.GetSysInfo(),
			"disk":   status.GetDiskInfo(),
		},
	})
}
