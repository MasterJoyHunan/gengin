package {{.pkgName}}

import (
	{{.importPackages}}

    "github.com/gin-gonic/gin"
)

func {{.handlerName}}(c *gin.Context) {
    // 1.接受报文
    {{if .hasRequest}}var req {{.requestType}}
{{.parseParam}}{{end}}
    {{if .hasResp}}resp, {{end}}err := {{.logicCall}}({{if .hasRequest}}&req{{end}})
    if err != nil {
        c.JSON(200, gin.H{
            "code": 1000,
            "message": "失败",
        })
    } else {
        c.JSON(200, gin.H{
            "code": 0,{{if .hasResp}}
            "data": resp,{{end}}
            "message": "成功",
        })
    }
}
