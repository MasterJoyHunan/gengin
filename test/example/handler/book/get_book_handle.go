package book

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/book"
	bookType "github.com/MasterJoyHunan/gengin/test/example/types/book"

	"github.com/gin-gonic/gin"
)

func GetBookHandle(c *gin.Context) {
	var req bookType.BookRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := book.GetBook(&req)
	response.HandleResponse(c, resp, err)
}
