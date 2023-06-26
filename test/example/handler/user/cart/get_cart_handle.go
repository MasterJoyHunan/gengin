package cart

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/user/cart"
	"github.com/MasterJoyHunan/gengin/test/example/svc"
	cartType "github.com/MasterJoyHunan/gengin/test/example/types/user/cart"

	"github.com/gin-gonic/gin"
)

// GetCartHandle 获取购物车信息
func GetCartHandle(c *gin.Context) {
	var req cartType.CartRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := cart.GetCart(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
