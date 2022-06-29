package cart

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/user/cart"
	cartType "github.com/MasterJoyHunan/gengin/test/example/types/user/cart"

	"github.com/gin-gonic/gin"
)

// EditCartHandle 编辑购物车信息
func EditCartHandle(c *gin.Context) {
	var req cartType.CartRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := cart.EditCart(&req)
	response.HandleResponse(c, resp, err)
}
