package swagger

import (
	"github.com/MasterJoyHunan/gengin/test/example/logic/swagger"

	"github.com/gin-gonic/gin"
)

func SwaggerHandle(c *gin.Context) {
	// 1.接受报文

	err := swagger.Swagger()
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
