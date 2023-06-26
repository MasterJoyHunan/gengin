package book

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/book"
	"github.com/MasterJoyHunan/gengin/test/example/svc"
	bookType "github.com/MasterJoyHunan/gengin/test/example/types/book"

	"github.com/gin-gonic/gin"
)

// EditBookHandle 编辑书本信息
func EditBookHandle(c *gin.Context) {
	var req bookType.BookRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := book.EditBook(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
