package user

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/user"
	"github.com/MasterJoyHunan/gengin/test/example/svc"
	userType "github.com/MasterJoyHunan/gengin/test/example/types/user"

	"github.com/gin-gonic/gin"
)

// GetUserHandle 获取用户信息
func GetUserHandle(c *gin.Context) {
	var req userType.UserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := user.GetUser(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
