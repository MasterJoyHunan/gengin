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
		return p.returnCode("ShouldBindHeader")
	}
	return ""
}

func (p *ParseRequestBody) uri(i spec.DefineStruct) string {
	if p.hasTag(i, "path") {
		return p.returnCode("ShouldBindUri")
	}
	return ""
}

func (p *ParseRequestBody) from(i spec.DefineStruct) string {
	if p.hasTag(i, "from") || p.hasTag(i, "json") {
		return p.returnCode("ShouldBind")
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

func (p *ParseRequestBody) returnCode(method string) string {
	return fmt.Sprintf(`    if err := c.%s(&%s); err != nil {
		// response.HandleResponse(c, nil, err)
		return
	}`, method, req) + "\n"
}
