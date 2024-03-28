package {{.pkgName}}

import (
    {{if .hasResp}}"{{.rootPkg}}/types"{{end}}
    "{{.rootPkg}}/{{.logicPkg}}"
    "{{.rootPkg}}/svc"
    "{{.rootPkg}}/internal/response"

    "github.com/gin-gonic/gin"
)

// {{.handlerName}}Handle {{.comment}}
func {{.handlerName}}Handle(c *gin.Context) {
    {{- if .hasReq -}}
    var req types.{{.requestType}}
    {{.parseRequest}}
    {{- end -}}
    {{if .hasResp}}resp, {{end}}err := {{.logicBase}}.{{.handlerName}}(svc.NewServiceContext(c), {{if .hasReq}}&req{{end}})
    response.HandleResponse(c, {{if .hasResp}}resp{{else}}nil{{end}}, err)
}
