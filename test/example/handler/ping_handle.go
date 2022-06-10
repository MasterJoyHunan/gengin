package handler

import (
	"github.com/MasterJoyHunan/gengin/test/example/logic"

	"github.com/gin-gonic/gin"
)

func PingHandle(c *gin.Context) {
	// 1.接受报文

	err := logic.Ping()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1000,
			"message": "失败",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "成功",
		})
	}
}
