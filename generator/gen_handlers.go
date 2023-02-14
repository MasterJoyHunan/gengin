package generator

import (
	"fmt"
	"strings"

	"github.com/MasterJoyHunan/gengin/pkg"
	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func GenHandlers() error {
	for _, group := range PluginInfo.Api.Service.Groups {
		for _, route := range group.Routes {
			if err := genHandler(group, route); err != nil {
				return err
			}
		}
	}
	return nil
}

func genHandler(group spec.Group, route spec.Route) error {
	handler := getHandlerName(route)
	logicGroupNameParse := parseGroupName(group.GetAnnotation(groupProperty), logicDir, logicPacket)
	handlerGroupNameParse := parseGroupName(group.GetAnnotation(groupProperty), handlerDir, handlerPacket)

	// 解析请求
	p := new(pkg.ParseRequestBody)
	parseRequest := p.BuildParseRequestStr(route.RequestTypeName(), PluginInfo.Api.Types)

	// 根据请求体找到对应的group
	alias := ""
	if len(route.RequestTypeName()) > 0 {
		typeGroupNameParse := parseGroupName(typeGroup[route.RequestTypeName()], typesDir, typesPacket)
		alias = getTypesUseAlias(typeGroupNameParse)
	}

	filename, err := format.FileNamingFormat(PluginInfo.Style, handler)
	if err != nil {
		return err
	}

	logic := getLogicName(route)
	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          handlerGroupNameParse.dirPath,
		filename:        filename + ".go",
		templateName:    "handlerTemplate",
		builtinTemplate: tpl.HandlerTemplate,
		data: map[string]interface{}{
			"comment":        parseComment(route),
			"pkgName":        handlerGroupNameParse.pkgName,
			"importPackages": genHandlerImports(group, route),
			"handlerName":    util.Title(handler),
			"requestType":    alias + util.Title(route.RequestTypeName()),
			"logicCall":      logicGroupNameParse.pkgName + "." + util.Title(strings.TrimSuffix(logic, "Logic")),
			"hasResp":        len(route.ResponseTypeName()) > 0,
			"hasRequest":     len(route.RequestTypeName()) > 0,
			"parseRequest":   parseRequest,
		},
	})
}

func genHandlerImports(group spec.Group, route spec.Route) string {
	var imports []string

	// handler 需要找 logic
	groupNameParse := parseGroupName(group.GetAnnotation(groupProperty), logicDir, logicPacket)
	imports = append(imports, fmt.Sprintf("\"%s\"",
		pathx.JoinPackages(RootPkg, groupNameParse.dirPath)))

	// handler 需要找 types, type 有可能需要 alias
	if len(route.RequestTypeName()) > 0 {
		groupNameParse = parseGroupName(typeGroup[route.RequestTypeName()], typesPacket, typesPacket)
		alias := getTypesImportAlias(groupNameParse)
		imports = append(imports, fmt.Sprintf("%s\"%s\"",
			alias, pathx.JoinPackages(RootPkg, groupNameParse.dirPath)))
	}

	// handler 需要统一返回处理
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(RootPkg, responseDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(RootPkg, "svc")))

	return strings.Join(imports, "\n\t")
}
