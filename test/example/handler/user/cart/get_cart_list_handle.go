package cart

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/user/cart"
	cartType "github.com/MasterJoyHunan/gengin/test/example/types/user/cart"

	"github.com/gin-gonic/gin"
)

// GetCartListHandle 获取所有购物车信息
func GetCartListHandle(c *gin.Context) {
	var req cartType.CartRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := cart.GetCartList(&req)
	response.HandleResponse(c, resp, err)
}
