package handler

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic"

	"github.com/gin-gonic/gin"
)

// PingHandle 无Group信息
func PingHandle(c *gin.Context) {
	err := logic.Ping()
	response.HandleResponse(c, nil, err)
}
