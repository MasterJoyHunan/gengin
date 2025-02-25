package cart

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/user/cart"
	"github.com/MasterJoyHunan/gengin/test/example/svc"
	"github.com/MasterJoyHunan/gengin/test/example/types"

	"github.com/gin-gonic/gin"
)

// AddCartHandle 添加购物车信息
func AddCartHandle(c *gin.Context) {
	var req types.CartRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := cart.AddCart(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
