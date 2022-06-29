package swagger

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/swagger"

	"github.com/gin-gonic/gin"
)

// SwaggerHandle
func SwaggerHandle(c *gin.Context) {
	err := swagger.Swagger()
	response.HandleResponse(c, nil, err)
}
