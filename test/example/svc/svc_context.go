package svc

import (
	"github.com/gin-gonic/gin"
)

type ServiceContext struct {
	GinContext *gin.Context
}

func NewServiceContext(c *gin.Context) *ServiceContext {
	return &ServiceContext{
		GinContext: c,
	}
}
