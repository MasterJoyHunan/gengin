package cart

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/user/cart"
	cartType "github.com/MasterJoyHunan/gengin/test/example/types/user/cart"

	"github.com/gin-gonic/gin"
)

func AddCartHandle(c *gin.Context) {
	var req cartType.CartRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := cart.AddCart(&req)
	response.HandleResponse(c, resp, err)
}
