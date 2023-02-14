package book

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/book"
	"github.com/MasterJoyHunan/gengin/test/example/svc"
	bookType "github.com/MasterJoyHunan/gengin/test/example/types/book"

	"github.com/gin-gonic/gin"
)

// AddBookHandle 添加书本信息
func AddBookHandle(c *gin.Context) {
	var req bookType.BookRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := book.AddBook(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
