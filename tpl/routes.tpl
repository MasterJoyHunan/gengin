// Code generated by goctl. DO NOT EDIT.
package {{.pkgName}}

import (
    {{.importPackages}}

	"github.com/gin-gonic/gin"
)

func Register{{.function}}Route(e *gin.Engine) {
    g := e.Group("{{if .hasPrefix}}{{.prefix}}{{end}}"){{if .hasMiddleware}}
    g.Use({{.middleware}}){{end}}
    {{.routesAdditions}}
}
