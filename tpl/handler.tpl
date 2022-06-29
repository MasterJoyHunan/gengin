package {{.pkgName}}

import (
	{{.importPackages}}

    "github.com/gin-gonic/gin"
)

func {{.handlerName}}(c *gin.Context) {
{{if .hasRequest}}    var req {{.requestType}}
{{.parseRequest}}
{{end}}    {{if .hasResp}}resp, {{end}}err := {{.logicCall}}({{if .hasRequest}}&req{{end}})
    response.HandleResponse(c, {{if .hasResp}}resp{{else}}nil{{end}}, err)
}
