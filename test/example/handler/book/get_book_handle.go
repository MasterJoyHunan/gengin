package book

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/book"
	"github.com/MasterJoyHunan/gengin/test/example/svc"
	"github.com/MasterJoyHunan/gengin/test/example/types"

	"github.com/gin-gonic/gin"
)

// GetBookHandle 获取书本信息
func GetBookHandle(c *gin.Context) {
	var req types.BookRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := book.GetBook(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
