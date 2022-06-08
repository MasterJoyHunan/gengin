package pkg

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

const req = "req"

type ParseRequestBody struct{}

func (p *ParseRequestBody) BuildParseRequestStr(requestName string, types []spec.Type) string {
	if requestName == "" {
		return ""
	}
	var sb strings.Builder
	var structType spec.DefineStruct
	for _, t := range types {
		if t.Name() == requestName {
			structType = t.(spec.DefineStruct)
			break
		}
	}
	sb.WriteString(p.header(structType))
	sb.WriteString(p.uri(structType))
	sb.WriteString(p.from(structType))
	return sb.String()
}

func (p *ParseRequestBody) header(i spec.DefineStruct) string {
	if p.hasTag(i, "header") {
		return p.returnCode("ShouldBindHeader", 1001)
	}
	return ""
}

func (p *ParseRequestBody) uri(i spec.DefineStruct) string {
	if p.hasTag(i, "path") {
		return p.returnCode("ShouldBindUri", 1001)
	}
	return ""
}

func (p *ParseRequestBody) from(i spec.DefineStruct) string {
	if p.hasTag(i, "from") || p.hasTag(i, "json") {
		return p.returnCode("ShouldBind", 1001)
	}
	return ""
}

func (p *ParseRequestBody) hasTag(i spec.DefineStruct, needTag string) bool {
	for _, m := range i.Members {
		before, _, found := strings.Cut(m.Tag, ":")
		if found && strings.HasSuffix(before, needTag) {
			return true
		}
	}
	return false
}

func (p *ParseRequestBody) returnCode(method string, code int) string {
	return fmt.Sprintf(`    if err := c.%s(&%s); err != nil {
		// TODO 处理异常
		c.JSON(200, gin.H{
            "code":    %d,
            "message": "失败",
        })
		return
	}`, method, req, code) + "\n"
}
