package {{.pkgName}}

import (
	{{.importPackages}}

    "github.com/gin-gonic/gin"
)

// {{.handlerName}} {{.comment}}
func {{.handlerName}}(c *gin.Context) {
{{- if .hasRequest -}}
    var req {{.requestType}}
    {{.parseRequest}}
{{- end -}}
    {{if .hasResp}}resp, {{end}}err := {{.logicCall}}(svc.NewServiceContext(c), {{if .hasRequest}}&req{{end}})
    response.HandleResponse(c, {{if .hasResp}}resp{{else}}nil{{end}}, err)
}
