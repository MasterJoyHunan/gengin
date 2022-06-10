package user

import (
	"github.com/MasterJoyHunan/gengin/test/example/logic/user"
	userType "github.com/MasterJoyHunan/gengin/test/example/types/user"

	"github.com/gin-gonic/gin"
)

func AddUserHandle(c *gin.Context) {
	// 1.接受报文
	var req userType.UserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		// TODO 处理异常
		c.JSON(200, gin.H{
			"code":    1001,
			"message": "失败",
		})
		return
	}
	if err := c.ShouldBind(&req); err != nil {
		// TODO 处理异常
		c.JSON(200, gin.H{
			"code":    1001,
			"message": "失败",
		})
		return
	}

	resp, err := user.AddUser(&req)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1000,
			"message": "失败",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    0,
			"data":    resp,
			"message": "成功",
		})
	}
}
