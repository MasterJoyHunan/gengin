package pkg

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

type ParseRequestBody struct{}

func (p *ParseRequestBody) BuildParseRequestStr(requestType spec.Type) string {
	if requestType == nil {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(p.from(requestType))
	sb.WriteString(p.header(requestType))
	sb.WriteString(p.uri(requestType))
	return sb.String()
}

func (p *ParseRequestBody) header(i spec.Type) string {
	if p.hasTag(i, "header") {
		return p.returnCode("ShouldBindHeader")
	}
	return ""
}

func (p *ParseRequestBody) uri(i spec.Type) string {
	if p.hasTag(i, "path") || p.hasTag(i, "uri") {
		return p.returnCode("ShouldBindUri")
	}
	return ""
}

func (p *ParseRequestBody) from(i spec.Type) string {
	if p.hasTag(i, "form") || p.hasTag(i, "json") {
		return p.returnCode("ShouldBind")
	}
	return ""
}

func (p *ParseRequestBody) hasTag(i spec.Type, needTag string) bool {
	switch v := i.(type) {
	case spec.DefineStruct:
		for _, vv := range v.Members {
			before, _, found := strings.Cut(vv.Tag, ":")
			if found && strings.HasSuffix(before, needTag) {
				return true
			}
			if p.hasTag(vv.Type, needTag) {
				return true
			}
		}

	//case spec.PrimitiveType: // 内置

	case spec.MapType:
		if p.hasTag(v.Value, needTag) {
			return true
		}
	case spec.ArrayType:
		if p.hasTag(v.Value, needTag) {
			return true
		}
	//case spec.InterfaceType:
	// 不允许
	case spec.PointerType:
		// 不允许
		if p.hasTag(v.Type, needTag) {
			return true
		}
	}

	return false
}

func (p *ParseRequestBody) returnCode(method string) string {
	return fmt.Sprintf(`    if err := c.%s(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}`, method) + "\n"
}
