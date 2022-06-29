package response

import "github.com/gin-gonic/gin"

// UnifiedResponse 统一返回
type UnifiedResponse struct {
    Code    int    `json:"code"`
    Data    any    `json:"data"`
    Message string `json:"message"`
}

// HandleResponse 统一返回处理
func HandleResponse(c *gin.Context, data any, err error) {
    if err != nil {
        c.JSON(200, UnifiedResponse{
            Code:    400,
            Data:    nil,
            Message: err.Error(),
        })
        return
    }

    c.JSON(200, UnifiedResponse{
        Code:    100,
        Data:    data,
        Message: "成功",
    })
}

// HandleAbortResponse 统一 Abort 返回处理
func HandleAbortResponse(c *gin.Context, err string) {
    c.AbortWithStatusJSON(200, UnifiedResponse{
        Code:    400,
        Data:    nil,
        Message: err,
    })
}
